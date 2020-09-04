package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration

	triggerWords map[string]bool
	// BotId of the created bot account.
	botID string
}

func (p *Plugin) isChrTriggerWord(word string) bool {
	_, ok := p.triggerWords[strings.ToLower(word)]
	return ok
}

func (p *Plugin) modifyMessage(post *model.Post) (*model.Post, string) {
	message := post.Message
	words := strings.Split(message, " ")
	for _, word := range words {
		if p.isChrTriggerWord(word) {
			ephemeralPost := &model.Post{
				UserId:    p.botID,
				ChannelId: post.ChannelId,
				Message:   "",
			}
			var props map[string]interface{}
			data := `
			{
				"attachments": [
					{
						"fallback": "test",
						"color": "#FF8000",
						"pretext": "Seems like you have a Doubt! :thinking:. You can click on **Raise Concept Help Request** link below to raise a Concept help Request",
						"text": "Now, you can ask doubts for any concept and get them answered by getting on a **1-1 live call of ~ 15 minutes with the helper/TA**.",
						"author_name": "Scaler",
						"author_icon": "https://assets.scaler.com/assets/academy/scalar-chat-icon-7dc8c6cce5bc388bd2ce9de1d347df05e5999d50d1b3a50ed910c93a97d97eca.png",
						"author_link": "http://www.scaler.com",
						"title": "Raise Concept Help Request",
						"title_link": "http://www.scaler.com/academy/mentee-dashboard/mentee_help_request_dashboard/?ref=open-chr-modal",
						"image_url": "https://assets.scaler.com/assets/academy/help_requests/bulb_question-57932b17f7273b95ad9b9bc23e8880437e35ad616927f09a1ad9c613372c5e18.png"
					}
				]
			}`
			err := json.Unmarshal([]byte(data), &props)
			if err != nil {
				fmt.Println("Failed to create the props")
			}
			ephemeralPost.SetProps(props)

			p.API.SendEphemeralPost(post.UserId, ephemeralPost)
			return post, ""
		}
	}
	return post, ""
}

func (p *Plugin) MessageWillBePosted(_ *plugin.Context, post *model.Post) (*model.Post, string) {
	return p.modifyMessage(post)
}

func (p *Plugin) MessageWillBeUpdated(_ *plugin.Context, post *model.Post) (*model.Post, string) {
	return p.modifyMessage(post)
}

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
