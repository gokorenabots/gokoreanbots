package gokoreanbots

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/url"
	"strconv"
)

const baseURL = "https://koreanbots.dev/api/v2"

type HTTPClient struct {
	fasthttpClient fasthttp.Client
}

// NewHTTPClient create new http client
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		fasthttpClient: fasthttp.Client{},
	}
}

// PostServers post servers to koreanbots
func (c *HTTPClient) PostServers(token, botID string, servers, shards int) error {
	var err error

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	body, err := json.Marshal(&BotStatRequest{
		Servers: servers,
		Shards:  shards,
	})
	if err != nil {
		return err
	}

	request.Header.Add("Authorization", token)
	request.Header.SetMethod("POST")
	request.SetBody(body)
	request.SetRequestURI(baseURL + "/bots/" + botID + "/stats")

	err = c.fasthttpClient.Do(request, response)
	fasthttp.ReleaseRequest(request)
	if err != nil {
		return err
	}

	return getStatusError(response.StatusCode())
}

// GetVote gets vote of user from koreanbots
func (c *HTTPClient) GetVote(token, botID, userID string) (*Vote, error) {
	var (
		err      error
		rawData  RawResponse
		voteData Vote
	)

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.Header.Add("Authorization", token)
	request.Header.Add("Content-Type", "application/json")
	request.Header.SetMethod("GET")
	request.SetRequestURI(baseURL + "/v2/bots/" + botID + "/vote?userID=" + userID)

	err = c.fasthttpClient.Do(request, response)
	fasthttp.ReleaseRequest(request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response.Body(), &rawData)
	voteData.raw = &rawData
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawData.Data, &voteData)
	if err != nil {
		return nil, err
	}

	return &voteData, getStatusError(response.StatusCode())
}

func (c *HTTPClient) GetBot(id string) (*Bot, error) {
	var (
		err     error
		rawData RawResponse
		bot     Bot
	)

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.Header.Add("Content-Type", "application/json")
	request.Header.SetMethod("GET")
	request.SetRequestURI(baseURL + "/bots/" + id)

	err = c.fasthttpClient.Do(request, response)
	fasthttp.ReleaseRequest(request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response.Body(), &rawData)
	bot.raw = &rawData
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawData.Data, &bot)

	return &bot, getStatusError(response.StatusCode())
}

// SearchBots searches bots from koreanbots
func (c *HTTPClient) SearchBots(query string, page int) (*Bots, error) {
	var (
		err     error
		rawData RawResponse
		bots    Bots
	)

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.Header.Add("Content-Type", "application/json")
	request.Header.SetMethod("GET")
	request.SetRequestURI(baseURL + "/search/bots?query=" + url.QueryEscape(query) + "&page=" + strconv.Itoa(page))

	err = c.fasthttpClient.Do(request, response)
	fasthttp.ReleaseRequest(request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response.Body(), &rawData)
	bots.raw = &rawData
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawData.Data, &bots)

	return &bots, getStatusError(response.StatusCode())
}

// GetBotsByVote gets bots by vote in koreanbots
func (c *HTTPClient) GetBotsByVote(page int) (*Bots, error) {
	var (
		err     error
		rawData RawResponse
		bots    Bots
	)

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.Header.Add("Content-Type", "application/json")
	request.Header.SetMethod("GET")
	request.SetRequestURI(baseURL + "/list/bots/votes?page=" + strconv.Itoa(page))

	err = c.fasthttpClient.Do(request, response)
	fasthttp.ReleaseRequest(request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response.Body(), &rawData)
	bots.raw = &rawData
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawData.Data, &bots)

	return &bots, getStatusError(response.StatusCode())
}

// GetNewBots gets new bots from koreanbots
func (c *HTTPClient) GetNewBots() (*Bots, error) {
	var (
		err     error
		rawData RawResponse
		bots    Bots
	)

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.Header.Add("Content-Type", "application/json")
	request.Header.SetMethod("GET")
	request.SetRequestURI(baseURL + "/list/bots/new")

	err = c.fasthttpClient.Do(request, response)
	fasthttp.ReleaseRequest(request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response.Body(), &rawData)
	bots.raw = &rawData
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawData.Data, &bots)

	return &bots, getStatusError(response.StatusCode())
}

// GetUser get koreanbots user
func (c *HTTPClient) GetUser(id string) (*User, error) {
	var (
		err     error
		rawData RawResponse
		user    User
	)

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.Header.Add("Content-Type", "application/json")
	request.Header.SetMethod("GET")
	request.SetRequestURI(baseURL + "/users/" + id)

	err = c.fasthttpClient.Do(request, response)
	fasthttp.ReleaseRequest(request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response.Body(), &rawData)
	user.raw = &rawData
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawData.Data, &user)

	return &user, getStatusError(response.StatusCode())
}
