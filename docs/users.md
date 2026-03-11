# Users

> Managing user, including the profiles current bot user.

## Overview

The Users endpoints allow you to get and modify the current bot user's profile information. This includes setting the bot's bio ("about me"), username, and avatar.

## Detailed Usage

### Get Current User

Retrieve the current bot user's profile.

```go
package main

import (
    "context"
    "log"

    "github.com/Mvkweb/gophord/pkg/rest"
)

func getCurrentUser(restClient *rest.Client) {
    user, err := restClient.GetCurrentUser(context.Background())
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    log.Printf("Bot: %s#%s", user.Username, user.Discriminator)
}
```

### Modify Current User Profile

Modify the current user's profile settings. This is commonly used to set a bot's bio/profile description.

```go
package main

import (
    "context"
    "log"

    "github.com/Mvkweb/gophord/pkg/rest"
    "github.com/Mvkweb/gophord/pkg/types"
)

func modifyBotProfile(restClient *rest.Client) {
    updatedUser, err := restClient.ModifyCurrentUser(context.Background(), &types.ModifyCurrentUserParams{
        Bio: "evelith.net",
    })
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    log.Printf("Updated bio: %s", updatedUser.Bio)
}
```

### Set Bot Bio and Username

You can set multiple profile fields at once:

```go
func updateBotProfile(restClient *rest.Client) {
    _, err := restClient.ModifyCurrentUser(context.Background(), &types.ModifyCurrentUserParams{
        Username: "Evelith",
        Bio:      "evelith.net",
    })
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
}
```

## API Reference

### GET /users/@me

Returns the user object for the authenticated user.

### PATCH /users/@me

Modifies the authenticated user's profile settings.

#### Parameters

- **Username** (string, optional) - User's username. For bots, 2-32 characters.
- **Avatar** (string, optional) - Avatar image encoded as base64.
- **Bio** (string, optional) - User's "about me" field. Maximum 190 characters for bots.

#### Response

Returns the updated [User](#) object.
