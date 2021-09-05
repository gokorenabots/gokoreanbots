package gokoreanbots

import (
	"fmt"
	"testing"
)

func TestHTTPClient_SearchBots(t *testing.T) {
	client := NewHTTPClient()
	bots, err := client.SearchBots("도박", 1)
	if err != nil {
		t.Error(err)
	}
	for idx, bot := range bots.Data {
		fmt.Printf("%d. %s, (%d Servers/%d Votes)\n", idx+1, bot.Name, bot.Servers, bot.Servers)
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
