package http

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/jesusEstaba/calculator/internal"
	"github.com/jesusEstaba/calculator/pkg/domain"
	"io"
	"net/http"
	"strings"
)

type RandomStringRepositoryImplementation struct{}

func NewRandomStringRepositoryImplementation() domain.RandomStringRepository {
	return &RandomStringRepositoryImplementation{}
}

func (r *RandomStringRepositoryImplementation) Generate() (string, error) {
	url := "https://api.random.org/json-rpc/2/invoke"

	resp, err := http.Post(url, "application/json", strings.NewReader(getBody()))
	if err != nil {
		return "", errors.New("random string generation: failed call")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("random string generation: unavailable service")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("random string generation: can not read body")
	}

	var response randomStringResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", errors.New("random string generation: can not parse body")
	}

	return response.Result.Random.Data[0], nil
}

func getBody() string {
	apiKey := internal.Config.RandomStringAPIKey
	id := uuid.New()

	return `
	{
		"jsonrpc": "2.0",
		"method": "generateUUIDs",
		"params": {
			"apiKey": "` + apiKey + `",
			"n": 1,
			"length": 32,
			"characters": "abcdefghijklmnopqrstuvwxyz",
		},
		"id": "` + id.String() + `"
	}
	`
}

type randomStringResponse struct {
	Result resultRandomStringResponse `json:"result"`
}

type resultRandomStringResponse struct {
	Random dataRandomResultRandomStringResponse `json:"random"`
}

type dataRandomResultRandomStringResponse struct {
	Data []string `json:"data"`
}
