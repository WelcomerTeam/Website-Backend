package backend

import "github.com/WelcomerTeam/Discord/discord"

// User represents the structure sent when doing a /users/@me request.
type User struct {
	*discord.User
}
