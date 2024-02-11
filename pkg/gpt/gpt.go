package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IGPT interface {
	SendRequest() (Response, error)
	GetID() string
	MessagesCount(totalMessages int) int
}

type IGPTService interface {
	Get(id string) (IGPT, error)
	Create(message IGPT) (IGPT, error)
}

type GPTReader interface {
	GetMessage(id string) (IGPT, error)
}
type GPTWriter interface {
	Save(message IGPT) (IGPT, error)
}
type IGPTPersistence interface {
	GPTReader
	GPTWriter
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Choice struct {
	Input   string  `json:"input"`
	Message Message `json:"message"`
}
type Request struct {
	ID        string    `json:"id"`
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens,omitempty"`
}
type Response struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
}

type GPT struct {
	OpenAIKey string
	Request   Request
	Response  Response
}

func NewGPT(openAIKey string) *GPT {
	return &GPT{
		OpenAIKey: openAIKey,
	}
}

func (g *GPT) SendRequest() (Response, error) {
	reqJson, err := json.Marshal(g.Request)
	if err != nil {
		return Response{}, err
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqJson))
	if err != nil {
		return Response{}, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+g.OpenAIKey)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Response{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println()
		}
	}(response.Body)

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}

	var resp Response
	err = json.Unmarshal(responseData, &resp)
	if err != nil {
		return Response{}, err
	}

	return resp, nil
}

func (g *GPT) GetID() string {
	return g.Request.ID
}

func (g *GPT) MessagesCount(totalMessages int) int {
	status := len(g.Request.Messages) - totalMessages
	return status
}
