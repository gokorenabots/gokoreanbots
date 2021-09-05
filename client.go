package gokoreanbots

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

type Koreanbots struct {
	http       *HTTPClient
	dgoSession *discordgo.Session

	token string
}

// NewKoreanbots creates new Koreanbots client
func NewKoreanbots(session *discordgo.Session, token string, postServers bool) *Koreanbots {
	koreanbots := &Koreanbots{
		http:       NewHTTPClient(),
		dgoSession: session,
		token:      token,
	}

	if postServers {
		go func() {
			for {
				_ = koreanbots.PostServers()
				time.Sleep(time.Minute * 3)
			}
		}()
	}
	return koreanbots
}

// GetVote gets vote of user from koreanbots
func (k Koreanbots) GetVote(userID string) (*Vote, error) {
	return k.http.GetVote(k.token, k.dgoSession.State.User.ID, userID)
}

// PostServers post servers to koreanbots
func (k Koreanbots) PostServers() error {
	return k.http.PostServers(
		k.token, k.dgoSession.State.User.ID, len(k.dgoSession.State.Guilds), k.dgoSession.ShardCount)
}
