package notifier

import (
	"bytes"
	"net/http"

	"github.com/goccy/go-json"
)

type DiscordAdapter struct {
	WebHookURL string
}

type DiscordNotification struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func (n DiscordNotification) Query() map[string]any {
	return make(map[string]any)
}

func (n DiscordNotification) Body() []byte {
	payload, _ := json.Marshal(n)
	return payload
}

func (d *DiscordAdapter) Send(n DiscordNotification) error {
	payload := n.Body()
	resp, err := http.Post(d.WebHookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	return nil
}

func NewDiscordAdapter(webhookURL string) *DiscordAdapter {
	return &DiscordAdapter{
		WebHookURL: webhookURL,
	}
}
