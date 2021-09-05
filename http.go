package gokoreanbots

import (
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

func (c *HTTPClient) PostServers(token, botID string, servers, shards int) error {
	var err error

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	body, err := BotStatRequest{
		Servers: servers,
		Shards:  shards,
	}.MarshalJSON()
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

	err = rawData.UnmarshalJSON(response.Body())
	voteData.raw = &rawData
	if err != nil {
		return nil, err
	}
	err = voteData.UnmarshalJSON(rawData.Data)

	return &voteData, getStatusError(response.StatusCode())
}

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

	err = rawData.UnmarshalJSON(response.Body())
	bots.raw = &rawData
	if err != nil {
		return nil, err
	}
	err = bots.UnmarshalJSON(rawData.Data)

	return &bots, getStatusError(response.StatusCode())
}

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

	err = rawData.UnmarshalJSON(response.Body())
	bots.raw = &rawData
	if err != nil {
		return nil, err
	}
	err = bots.UnmarshalJSON(rawData.Data)

	return &bots, getStatusError(response.StatusCode())
}

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

	err = rawData.UnmarshalJSON(response.Body())
	bots.raw = &rawData
	if err != nil {
		return nil, err
	}
	err = bots.UnmarshalJSON(rawData.Data)

	return &bots, getStatusError(response.StatusCode())
}
