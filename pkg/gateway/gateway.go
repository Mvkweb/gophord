// Package gateway provides a WebSocket client for the Discord Gateway.
//
// This package handles the WebSocket connection to Discord's gateway,
// including heartbeating, reconnection, and event dispatching.
package gateway

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Mvkweb/gophord/pkg/json"
	"github.com/Mvkweb/gophord/pkg/types"
	"github.com/lxzan/gws"
)

// Opcodes for gateway payloads.
const (
	OpcodeDispatch            = 0
	OpcodeHeartbeat           = 1
	OpcodeIdentify            = 2
	OpcodePresenceUpdate      = 3
	OpcodeVoiceStateUpdate    = 4
	OpcodeResume              = 6
	OpcodeReconnect           = 7
	OpcodeRequestGuildMembers = 8
	OpcodeInvalidSession      = 9
	OpcodeHello               = 10
	OpcodeHeartbeatACK        = 11
)

// GatewayPayload represents a gateway WebSocket payload.
type GatewayPayload struct {
	// Op is the opcode.
	Op int `json:"op"`
	// D is the event data.
	D interface{} `json:"d,omitempty"`
	// S is the sequence number (for Op 0).
	S *int64 `json:"s,omitempty"`
	// T is the event name (for Op 0).
	T string `json:"t,omitempty"`
}

// HelloData represents the data in a Hello payload.
type HelloData struct {
	// HeartbeatInterval is the interval for heartbeats in milliseconds.
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// IdentifyData represents the data in an Identify payload.
type IdentifyData struct {
	// Token is the authentication token.
	Token string `json:"token"`
	// Intents is the gateway intents.
	Intents types.IntentFlags `json:"intents"`
	// Properties contains client properties.
	Properties IdentifyProperties `json:"properties"`
	// Compress indicates whether to use zlib compression.
	Compress bool `json:"compress,omitempty"`
	// LargeThreshold is the threshold for offline member handling.
	LargeThreshold int `json:"large_threshold,omitempty"`
	// Shard is the shard information [shard_id, num_shards].
	Shard *[2]int `json:"shard,omitempty"`
	// Presence is the initial presence.
	Presence *PresenceUpdate `json:"presence,omitempty"`
}

// IdentifyProperties contains client identification properties.
type IdentifyProperties struct {
	// OS is the operating system.
	OS string `json:"os"`
	// Browser is the library name.
	Browser string `json:"browser"`
	// Device is the library name.
	Device string `json:"device"`
}

// PresenceUpdate represents a presence update payload.
type PresenceUpdate struct {
	// Since is the Unix time when the client went idle.
	Since *int64 `json:"since"`
	// Activities are the user's activities.
	Activities []Activity `json:"activities"`
	// Status is the user's status.
	Status string `json:"status"`
	// AFK indicates whether the client is AFK.
	AFK bool `json:"afk"`
}

// Activity represents a user activity.
type Activity struct {
	// Name is the activity name.
	Name string `json:"name"`
	// Type is the activity type (0: Playing, 1: Streaming, 2: Listening, 3: Watching, 4: Custom Status, 5: Competing).
	Type int `json:"type"`
	// State is the custom status text (used with type 4 for custom status).
	State string `json:"state,omitempty"`
	// URL is the stream URL (for streaming activities).
	URL string `json:"url,omitempty"`
}

// ResumeData represents the data in a Resume payload.
type ResumeData struct {
	// Token is the authentication token.
	Token string `json:"token"`
	// SessionID is the session ID to resume.
	SessionID string `json:"session_id"`
	// Seq is the last sequence number received.
	Seq int64 `json:"seq"`
}

// Event represents a dispatched gateway event.
type Event struct {
	// Type is the event type name.
	Type string
	// Data is the raw event data.
	Data []byte
	// Sequence is the event sequence number.
	Sequence int64
}

// Client is a gateway WebSocket client.
type Client struct {
	token        string
	intents      types.IntentFlags
	conn         *gws.Conn
	sessionID    string
	resumeURL    string
	sequence     atomic.Int64
	heartbeatACK atomic.Bool
	events       chan *Event
	errors       chan error
	done         chan struct{}
	mu           sync.RWMutex
	closed       bool
	mobileStatus bool
}

// ClientOption is a function that configures a Client.
type ClientOption func(*Client)

// WithIntents sets the gateway intents.
func WithIntents(intents types.IntentFlags) ClientOption {
	return func(c *Client) {
		c.intents = intents
	}
}

// WithMobileStatus sets whether to identify as a mobile client (Discord Android).
func WithMobileStatus(enabled bool) ClientOption {
	return func(c *Client) {
		c.mobileStatus = enabled
	}
}

// New creates a new gateway client.
func New(token string, opts ...ClientOption) *Client {
	c := &Client{
		token:   token,
		intents: types.IntentsDefault,
		events:  make(chan *Event, 100),
		errors:  make(chan error, 10),
		done:    make(chan struct{}),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Connect establishes a connection to the Discord gateway.
func (c *Client) Connect(ctx context.Context) error {
	url := types.GatewayURL
	if c.resumeURL != "" {
		url = c.resumeURL
	}

	conn, _, err := gws.NewClient(c, &gws.ClientOption{
		Addr: url,
		PermessageDeflate: gws.PermessageDeflate{
			Enabled:               true,
			ServerContextTakeover: true,
			ClientContextTakeover: true,
		},
	})
	if err != nil {
		return fmt.Errorf("dial gateway: %w", err)
	}

	c.mu.Lock()
	c.conn = conn
	c.closed = false
	c.mu.Unlock()

	// Start reading messages
	go conn.ReadLoop()

	return nil
}

// Events returns the channel for receiving events.
func (c *Client) Events() <-chan *Event {
	return c.events
}

// Errors returns the channel for receiving errors.
func (c *Client) Errors() <-chan error {
	return c.errors
}

// Close closes the gateway connection.
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil
	}
	c.closed = true

	close(c.done)

	if c.conn != nil {
		// NetConn().Close() is the underlying connection
		return c.conn.NetConn().Close()
	}

	return nil
}

// OnOpen implements gws.EventHandler.
func (c *Client) OnOpen(socket *gws.Conn) {
	// No-op
}

// OnClose implements gws.EventHandler.
func (c *Client) OnClose(socket *gws.Conn, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return
	}
	c.closed = true

	select {
	case <-c.done:
	default:
		close(c.done)
	}
}

// OnPing implements gws.EventHandler.
func (c *Client) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.WritePong(payload)
}

// OnPong implements gws.EventHandler.
func (c *Client) OnPong(socket *gws.Conn, payload []byte) {
	// No-op
}

// OnMessage implements gws.EventHandler.
func (c *Client) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	if err := c.handleMessage(message.Bytes()); err != nil {
		select {
		case c.errors <- err:
		default:
		}
	}
}

// handleMessage processes a gateway message.
func (c *Client) handleMessage(data []byte) error {
	var payload GatewayPayload
	if err := json.Unmarshal(data, &payload); err != nil {
		return fmt.Errorf("unmarshal payload: %w", err)
	}

	switch payload.Op {
	case OpcodeDispatch:
		return c.handleDispatch(&payload, data)
	case OpcodeHeartbeat:
		return c.sendHeartbeat()
	case OpcodeReconnect:
		return c.handleReconnect()
	case OpcodeInvalidSession:
		return c.handleInvalidSession(data)
	case OpcodeHello:
		return c.handleHello(data)
	case OpcodeHeartbeatACK:
		c.heartbeatACK.Store(true)
	}

	return nil
}

// handleDispatch handles a dispatch event.
func (c *Client) handleDispatch(payload *GatewayPayload, raw []byte) error {
	if payload.S != nil {
		c.sequence.Store(*payload.S)
	}

	// Extract the event data using json.Get (references buffer)
	// We need to use MarshalJSON immediately to get a safe copy
	eventData, err := json.Get(raw, "d")
	if err != nil {
		return fmt.Errorf("extract event data: %w", err)
	}

	dataBytes, err := eventData.MarshalJSON()
	if err != nil {
		return fmt.Errorf("marshal event data: %w", err)
	}

	// Handle READY event specially
	if payload.T == "READY" {
		var ready struct {
			SessionID string `json:"session_id"`
			ResumeURL string `json:"resume_gateway_url"`
		}
		if err := json.Unmarshal(dataBytes, &ready); err == nil {
			c.mu.Lock()
			c.sessionID = ready.SessionID
			c.resumeURL = ready.ResumeURL
			c.mu.Unlock()
		}
	}

	// Send event to channel
	event := &Event{
		Type:     payload.T,
		Data:     dataBytes, // Safe copy from MarshalJSON
		Sequence: c.sequence.Load(),
	}

	select {
	case c.events <- event:
	default:
		// Event channel full, drop oldest event
		select {
		case <-c.events:
		default:
		}
		c.events <- event
	}

	return nil
}

// handleHello handles a Hello payload.
func (c *Client) handleHello(data []byte) error {
	// Extract heartbeat interval
	intervalNode, err := json.Get(data, "d", "heartbeat_interval")
	if err != nil {
		return fmt.Errorf("extract heartbeat interval: %w", err)
	}

	interval, err := intervalNode.Int64()
	if err != nil {
		return fmt.Errorf("parse heartbeat interval: %w", err)
	}

	// Start heartbeat
	go c.heartbeatLoop(time.Duration(interval) * time.Millisecond)

	// Send identify or resume
	c.mu.RLock()
	sessionID := c.sessionID
	c.mu.RUnlock()

	if sessionID != "" {
		return c.sendResume()
	}
	return c.sendIdentify()
}

// handleReconnect handles a Reconnect opcode.
func (c *Client) handleReconnect() error {
	// gws handles reconnection logic mostly, but we trigger a reconnect manually here
	c.Close()
	return c.Connect(context.Background())
}

// handleInvalidSession handles an Invalid Session opcode.
func (c *Client) handleInvalidSession(data []byte) error {
	// Check if resumable
	resumable, err := json.Get(data, "d")
	if err == nil {
		if b, err := resumable.Bool(); err == nil && b {
			time.Sleep(time.Second * time.Duration(1+time.Now().UnixNano()%5))
			return c.sendResume()
		}
	}

	// Reset session and re-identify
	c.mu.Lock()
	c.sessionID = ""
	c.resumeURL = ""
	c.mu.Unlock()

	time.Sleep(time.Second * time.Duration(1+time.Now().UnixNano()%5))
	return c.sendIdentify()
}

// heartbeatLoop sends heartbeats at the specified interval.
func (c *Client) heartbeatLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Send first heartbeat after jitter
	jitter := time.Duration(float64(interval) * 0.5)
	time.Sleep(jitter)
	c.sendHeartbeat()
	c.heartbeatACK.Store(true)

	for {
		select {
		case <-c.done:
			return
		case <-ticker.C:
			if !c.heartbeatACK.Load() {
				// No ACK received, reconnect
				c.errors <- fmt.Errorf("heartbeat ACK not received")
				c.handleReconnect()
				return
			}
			c.heartbeatACK.Store(false)
			c.sendHeartbeat()
		}
	}
}

// sendHeartbeat sends a heartbeat payload.
func (c *Client) sendHeartbeat() error {
	seq := c.sequence.Load()
	var d interface{} = nil
	if seq > 0 {
		d = seq
	}

	return c.send(GatewayPayload{
		Op: OpcodeHeartbeat,
		D:  d,
	})
}

// sendIdentify sends an Identify payload.
func (c *Client) sendIdentify() error {
	browser := "gophord"
	if c.mobileStatus {
		browser = "Discord Android"
	}

	return c.send(GatewayPayload{
		Op: OpcodeIdentify,
		D: IdentifyData{
			Token:   c.token,
			Intents: c.intents,
			Properties: IdentifyProperties{
				OS:      "linux",
				Browser: browser,
				Device:  "gophord",
			},
			LargeThreshold: 250,
		},
	})
}

// sendResume sends a Resume payload.
func (c *Client) sendResume() error {
	c.mu.RLock()
	sessionID := c.sessionID
	c.mu.RUnlock()

	return c.send(GatewayPayload{
		Op: OpcodeResume,
		D: ResumeData{
			Token:     c.token,
			SessionID: sessionID,
			Seq:       c.sequence.Load(),
		},
	})
}

// send sends a payload over the WebSocket connection.
func (c *Client) send(payload GatewayPayload) error {
	c.mu.RLock()
	conn := c.conn
	c.mu.RUnlock()

	if conn == nil {
		return fmt.Errorf("not connected")
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	return conn.WriteMessage(gws.OpcodeText, data)
}

// UpdatePresence updates the client's presence.
func (c *Client) UpdatePresence(presence *PresenceUpdate) error {
	return c.send(GatewayPayload{
		Op: OpcodePresenceUpdate,
		D:  presence,
	})
}
