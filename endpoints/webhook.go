package endpoints

import (
	"encoding/json"
	"net/http"
)

type Webhook struct {
	Avatar string `json:"avatar"`
	ChannelID string `json:"channel_id"`
	GuildID string `json:"guild_id"`
	Id string `json:"id"`
	Name string `json:"name"`
}

func Fetch(url string) (*Webhook, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var webhook Webhook
	err = json.NewDecoder(resp.Body).Decode(&webhook)
	if err != nil {
		return nil, err
	}

	return &webhook, nil
}