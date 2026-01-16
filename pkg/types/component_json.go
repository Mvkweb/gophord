// Package types provides Discord API type definitions.
package types

import (
	"encoding/json"
	"fmt"

	"github.com/bytedance/sonic"
)

// componentJSON is used for JSON marshaling/unmarshaling of components.
type componentJSON struct {
	Type ComponentType `json:"type"`
}

// MarshalJSON implements json.Marshaler for ActionRow.
func (a *ActionRow) MarshalJSON() ([]byte, error) {
	type actionRowAlias ActionRow
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*actionRowAlias
	}{
		Type:           ComponentTypeActionRow,
		actionRowAlias: (*actionRowAlias)(a),
	})
}

// MarshalJSON implements json.Marshaler for Button.
func (b *Button) MarshalJSON() ([]byte, error) {
	type buttonAlias Button
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*buttonAlias
	}{
		Type:        ComponentTypeButton,
		buttonAlias: (*buttonAlias)(b),
	})
}

// MarshalJSON implements json.Marshaler for StringSelect.
func (s *StringSelect) MarshalJSON() ([]byte, error) {
	type stringSelectAlias StringSelect
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*stringSelectAlias
	}{
		Type:              ComponentTypeStringSelect,
		stringSelectAlias: (*stringSelectAlias)(s),
	})
}

// MarshalJSON implements json.Marshaler for UserSelect.
func (u *UserSelect) MarshalJSON() ([]byte, error) {
	type userSelectAlias UserSelect
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*userSelectAlias
	}{
		Type:            ComponentTypeUserSelect,
		userSelectAlias: (*userSelectAlias)(u),
	})
}

// MarshalJSON implements json.Marshaler for RoleSelect.
func (r *RoleSelect) MarshalJSON() ([]byte, error) {
	type roleSelectAlias RoleSelect
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*roleSelectAlias
	}{
		Type:            ComponentTypeRoleSelect,
		roleSelectAlias: (*roleSelectAlias)(r),
	})
}

// MarshalJSON implements json.Marshaler for MentionableSelect.
func (m *MentionableSelect) MarshalJSON() ([]byte, error) {
	type mentionableSelectAlias MentionableSelect
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*mentionableSelectAlias
	}{
		Type:                   ComponentTypeMentionableSelect,
		mentionableSelectAlias: (*mentionableSelectAlias)(m),
	})
}

// MarshalJSON implements json.Marshaler for ChannelSelect.
func (c *ChannelSelect) MarshalJSON() ([]byte, error) {
	type channelSelectAlias ChannelSelect
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*channelSelectAlias
	}{
		Type:               ComponentTypeChannelSelect,
		channelSelectAlias: (*channelSelectAlias)(c),
	})
}

// MarshalJSON implements json.Marshaler for Section.
func (s *Section) MarshalJSON() ([]byte, error) {
	type sectionAlias Section
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*sectionAlias
	}{
		Type:         ComponentTypeSection,
		sectionAlias: (*sectionAlias)(s),
	})
}

// MarshalJSON implements json.Marshaler for TextDisplay.
func (t *TextDisplay) MarshalJSON() ([]byte, error) {
	type textDisplayAlias TextDisplay
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*textDisplayAlias
	}{
		Type:             ComponentTypeTextDisplay,
		textDisplayAlias: (*textDisplayAlias)(t),
	})
}

// MarshalJSON implements json.Marshaler for Thumbnail.
func (t *Thumbnail) MarshalJSON() ([]byte, error) {
	type thumbnailAlias Thumbnail
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*thumbnailAlias
	}{
		Type:           ComponentTypeThumbnail,
		thumbnailAlias: (*thumbnailAlias)(t),
	})
}

// MarshalJSON implements json.Marshaler for MediaGallery.
func (m *MediaGallery) MarshalJSON() ([]byte, error) {
	type mediaGalleryAlias MediaGallery
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*mediaGalleryAlias
	}{
		Type:              ComponentTypeMediaGallery,
		mediaGalleryAlias: (*mediaGalleryAlias)(m),
	})
}

// MarshalJSON implements json.Marshaler for File.
func (f *File) MarshalJSON() ([]byte, error) {
	type fileAlias File
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*fileAlias
	}{
		Type:      ComponentTypeFile,
		fileAlias: (*fileAlias)(f),
	})
}

// MarshalJSON implements json.Marshaler for Separator.
func (s *Separator) MarshalJSON() ([]byte, error) {
	type separatorAlias Separator
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*separatorAlias
	}{
		Type:           ComponentTypeSeparator,
		separatorAlias: (*separatorAlias)(s),
	})
}

// MarshalJSON implements json.Marshaler for Container.
func (c *Container) MarshalJSON() ([]byte, error) {
	type containerAlias Container
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*containerAlias
	}{
		Type:           ComponentTypeContainer,
		containerAlias: (*containerAlias)(c),
	})
}

// MarshalJSON implements json.Marshaler for TextInput.
func (t *TextInput) MarshalJSON() ([]byte, error) {
	type textInputAlias TextInput
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*textInputAlias
	}{
		Type:           ComponentTypeTextInput,
		textInputAlias: (*textInputAlias)(t),
	})
}

// MarshalJSON implements json.Marshaler for Label.
func (i *Label) MarshalJSON() ([]byte, error) {
	type labelAlias Label
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*labelAlias
	}{
		Type:       ComponentTypeLabel,
		labelAlias: (*labelAlias)(i),
	})
}

// UnmarshalJSON implements json.Unmarshaler for Label.
// This handles the nested Component field which is an interface.
func (l *Label) UnmarshalJSON(data []byte) error {
	type labelAlias Label
	aux := &struct {
		*labelAlias
		Component json.RawMessage `json:"component"`
	}{
		labelAlias: (*labelAlias)(l),
	}

	if err := sonic.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal label: %w", err)
	}

	if len(aux.Component) > 0 {
		comp, err := UnmarshalComponent(aux.Component)
		if err != nil {
			return fmt.Errorf("failed to unmarshal label component: %w", err)
		}
		l.Component = comp
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler for Section.
// This handles the Accessory field which is an interface.
func (s *Section) UnmarshalJSON(data []byte) error {
	type sectionAlias Section
	aux := &struct {
		*sectionAlias
		Accessory json.RawMessage `json:"accessory"`
	}{
		sectionAlias: (*sectionAlias)(s),
	}

	if err := sonic.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal section: %w", err)
	}

	if len(aux.Accessory) > 0 {
		acc, err := UnmarshalComponent(aux.Accessory)
		if err != nil {
			return fmt.Errorf("failed to unmarshal section accessory: %w", err)
		}
		s.Accessory = acc
	}

	return nil
}

// MarshalJSON implements json.Marshaler for FileUpload.
func (f *FileUpload) MarshalJSON() ([]byte, error) {
	type fileUploadAlias FileUpload
	return sonic.Marshal(&struct {
		Type ComponentType `json:"type"`
		*fileUploadAlias
	}{
		Type:            ComponentTypeFileUpload,
		fileUploadAlias: (*fileUploadAlias)(f),
	})
}

// UnmarshalJSON implements json.Unmarshaler for ComponentList.
func (cl *ComponentList) UnmarshalJSON(data []byte) error {
	components, err := UnmarshalComponents(data)
	if err != nil {
		return err
	}
	*cl = components
	return nil
}

// UnmarshalComponent unmarshals a component from JSON data.
// It inspects the "type" field to determine which concrete type to use.
func UnmarshalComponent(data []byte) (Component, error) {
	var cj componentJSON
	if err := sonic.Unmarshal(data, &cj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal component type: %w", err)
	}

	var component Component
	switch cj.Type {
	case ComponentTypeActionRow:
		component = &ActionRow{}
	case ComponentTypeButton:
		component = &Button{}
	case ComponentTypeStringSelect:
		component = &StringSelect{}
	case ComponentTypeTextInput:
		component = &TextInput{}
	case ComponentTypeUserSelect:
		component = &UserSelect{}
	case ComponentTypeRoleSelect:
		component = &RoleSelect{}
	case ComponentTypeMentionableSelect:
		component = &MentionableSelect{}
	case ComponentTypeChannelSelect:
		component = &ChannelSelect{}
	case ComponentTypeSection:
		component = &Section{}
	case ComponentTypeTextDisplay:
		component = &TextDisplay{}
	case ComponentTypeThumbnail:
		component = &Thumbnail{}
	case ComponentTypeMediaGallery:
		component = &MediaGallery{}
	case ComponentTypeFile:
		component = &File{}
	case ComponentTypeSeparator:
		component = &Separator{}
	case ComponentTypeContainer:
		component = &Container{}
	case ComponentTypeLabel:
		component = &Label{}
	case ComponentTypeFileUpload:
		component = &FileUpload{}
	default:
		return nil, fmt.Errorf("unknown component type: %d", cj.Type)
	}

	if err := sonic.Unmarshal(data, component); err != nil {
		return nil, fmt.Errorf("failed to unmarshal component: %w", err)
	}

	return component, nil
}

// UnmarshalComponents unmarshals a slice of components from JSON data.
func UnmarshalComponents(data []byte) ([]Component, error) {
	var rawComponents []json.RawMessage
	if err := sonic.Unmarshal(data, &rawComponents); err != nil {
		return nil, fmt.Errorf("failed to unmarshal components array: %w", err)
	}

	components := make([]Component, 0, len(rawComponents))
	for _, raw := range rawComponents {
		component, err := UnmarshalComponent(raw)
		if err != nil {
			return nil, err
		}
		components = append(components, component)
	}

	return components, nil
}
