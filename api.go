package pubg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Api struct {
	Client  *http.Client
	Key     string
	BaseUrl string
	/**
	TODO:
		GetMatches(shard Shard) *[]Match
		GetMatch(shard Shard, matchID string) Match
		GetPlayers(shard Shard, playerIDs []string, playerNames []string) []Player
		GetTelemetry(shard Shard,assetID string)
	*/
}

func (c *Api) Status() (*Status, error) {
	status := &Status{}
	u, err := url.Parse(fmt.Sprintf("%s/status", c.BaseUrl))
	if err != nil {
		return status, err
	}
	err = c.get(u, status)
	return status, err
}

func (c *Api) GetPlayers(shard Shard, playerIDs []string, playerNames []string) (*[]Player, error) {
	players := &[]Player{}
	u, err := url.Parse(fmt.Sprintf("%s/shards/%s/players", c.BaseUrl, shard))
	if err != nil {
		return players, err
	}
	fmt.Println(playerIDs)
	fmt.Println(playerNames)
	query := u.Query()
	playerIDvalue := strings.Join(playerIDs, ",")
	if len(playerIDvalue) > 0 {
		query.Set("filter[playerIds]", playerIDvalue)
	}

	playerNameValue := strings.Join(playerNames, ",")
	if len(playerNameValue) > 0 {
		query.Set("filter[playerNames]", playerNameValue)
	}
	u.RawQuery = query.Encode()
	err = c.get(u, players)
	return players, err
}

func (c *Api) GetMatches(shard Shard, matchID string) (*[]Match, error) {
	matches := &[]Match{}
	u, err := url.Parse(fmt.Sprintf("%s/shards/%s/matches/%s", c.BaseUrl, shard, matchID))
	if err != nil {
		return matches, err
	}
	err = c.get(u, matches)
	return matches, err
}

func (c *Api) get(url *url.URL, target interface{}) error {

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Running %s",req.URL.String()))

	//req.Header.Set("User-Agent", "pubg-go")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Key))
	req.Header.Set("Accept", "application/vnd.api+json")

	r, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	//TODO: Rate limit API requests.
	//This can be done by analyzing the X-Ratelimit-Remaining response header and if it is 0 then setting a rate limit flag to wait for X-Ratelimit-Reset milliseconds
	rateLimit := r.Header.Get("X-Ratelimit-Limit")
	remainingRequests := r.Header.Get("X-Ratelimit-Remaining")
	resetTime := r.Header.Get("X-Ratelimit-Reset")
	fmt.Println(fmt.Sprintf("Currently %s requests remain out of your %s request limit.", remainingRequests, rateLimit))
	fmt.Println(fmt.Sprintf("Rate limit resets in %s milliseconds.", resetTime))

	defer r.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	fmt.Println(r.StatusCode)
	fmt.Println(string(bodyBytes))

	if r.StatusCode == http.StatusOK {
		r.Body.Close()
		return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(bodyBytes))).Decode(target)
	}
	if r.StatusCode == http.StatusNotFound {
		//This _could_ be ok. If a response finds nothing, the API returns a 404 with a body
		r.Body.Close()
		//try to decode to EmptyResponse
		emptyResponse := &EmptyResponse{}
		err := json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(bodyBytes))).Decode(emptyResponse)
		if err == nil && emptyResponse != nil {
			return emptyResponse
		}
		fmt.Println(fmt.Sprintf("error decoding empty response: %v", err))
		//ok, try to decode to error response
		errorResponse := &ErrorResponse{}
		err = json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(bodyBytes))).Decode(errorResponse)
		if err == nil && errorResponse != nil {
			return errorResponse
		}
		fmt.Println(fmt.Sprintf("error decoding error response: %v", err))
		//ok this is an unknown error
		return fmt.Errorf("failed to process %s", string(bodyBytes))
	}
	return fmt.Errorf("status Code: %d", r.StatusCode)
}

type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

type Errors struct {
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func (e *Errors) String() string {
	return fmt.Sprintf("%s: %s", e.Title, e.Detail)
}

type EmptyResponse struct {
	Errs []*Errors `json:"errors,omitempty"`
}

func (e *EmptyResponse) Error() string {
	var str = bytes.Buffer{}
	for _, err := range e.Errs {
		str.WriteString(err.String())
		str.WriteString("; ")
	}
	return str.String()
}
