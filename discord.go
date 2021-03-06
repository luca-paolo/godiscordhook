package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func SendWebhook(url string, webhook Webhook) error {
	requestBody, err := json.Marshal(webhook)

	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode > 299 || response.StatusCode < 200 {
		return errors.New(fmt.Sprintf("Request failed with status code %d", response.StatusCode))
	}

	return nil
}

type Webhook struct {
	Content   string  `json:"content,omitempty"`
	Username  string  `json:"username,omitempty"`
	AvatarURL string  `json:"avatar_url,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
}

type Embed struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url,omitempty"`
	Color       int       `json:"color,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	Footer      Footer    `json:"footer,omitempty"`
	Thumbnail   Thumbnail `json:"thumbnail,omitempty"`
	Image       Image     `json:"image,omitempty"`
	Author      Author    `json:"author,omitempty"`
	Fields      []Field   `json:"fields,omitempty"`
}

type Footer struct {
	IconUrl string `json:"icon_url,omitempty"`
	Text    string `json:"text"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

type Image struct {
	URL string `json:"url"`
}

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
