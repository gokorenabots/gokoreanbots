package gokoreanbots

import (
	"fmt"
	"testing"
)

func TestHTTPClient_GetBot(t *testing.T) {
	client := NewHTTPClient()
	bot, err := client.GetBot("770246143652397069")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s, (%d Servers/%d Votes)\n", bot.Name, bot.Servers, bot.Votes)
}

func TestHTTPClient_GetUser(t *testing.T) {
	client := NewHTTPClient()
	user, err := client.GetUser("441202161481809922")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s, (%d Bots)\n", user.Username, len(user.Bots))
}

func TestHTTPClient_SearchBots(t *testing.T) {
	client := NewHTTPClient()
	bots, err := client.SearchBots("도박", 1)
	if err != nil {
		t.Error(err)
	}
	for idx, bot := range bots.Data {
		fmt.Printf("%d. %s, (%d Servers/%d Votes)\n", idx+1, bot.Name, bot.Servers, bot.Votes)
	}
}

func TestHTTPClient_GetBotsByVote(t *testing.T) {
	client := NewHTTPClient()
	bots, err := client.GetBotsByVote(1)
	if err != nil {
		t.Error(err)
	}
	for idx, bot := range bots.Data {
		fmt.Printf("%d. %s, (%d Servers/%d Votes)\n", idx+1, bot.Name, bot.Servers, bot.Votes)
	}
}

func TestHTTPClient_GetNewBots(t *testing.T) {
	client := NewHTTPClient()
	bots, err := client.GetNewBots()
	if err != nil {
		t.Error(err)
	}
	for idx, bot := range bots.Data {
		fmt.Printf("%d. %s, (%d Servers/%d Votes)\n", idx+1, bot.Name, bot.Servers, bot.Votes)
	}
}
