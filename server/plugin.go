package main

import (
	"math/rand"
	"strconv"
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
	// Probability for the trigger message to be sent.
	probabilityFactor int
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
			// Tracking for no of times the trigger word has been detected
			p.IncrementTrackingCount("trigger_word_detect_count")
			if rand.Intn(10-1)+1 < (p.probabilityFactor / 10) {
				// Sending the Trigger message when chr trigger word has been detected and satsifies the probability factor
				ephemeralPost := &model.Post{
					UserId:    p.botID,
					ChannelId: post.ChannelId,
					Props: model.StringInterface{
						"attachments": []*model.SlackAttachment{
							{
								Text:       "Now, you can ask doubts for any concept and get them answered by getting on a **1-1 live call of ~ 15 minutes with the helper/TA**.",
								Pretext:    "Seems like you have a Doubt! :thinking:. You can click on **Raise Concept Help Request** link below to raise a Concept help Request",
								AuthorName: "Scaler",
								AuthorLink: "https://www.scaler.com",
								AuthorIcon: "https://assets.scaler.com/assets/academy/scalar-chat-icon-7dc8c6cce5bc388bd2ce9de1d347df05e5999d50d1b3a50ed910c93a97d97eca.png",
								Title:      "Raise Concept Help Request",
								TitleLink:  "http://www.scaler.com/academy/mentee-dashboard/mentee_help_request_dashboard/?ref=open-chr-modal",
								ImageURL:   "https://assets.scaler.com/assets/academy/help_requests/bulb_question-57932b17f7273b95ad9b9bc23e8880437e35ad616927f09a1ad9c613372c5e18.png",
							},
						},
					},
				}
				p.API.SendEphemeralPost(post.UserId, ephemeralPost)
				// Tracking for no of times we've sent the trigger message
				p.IncrementTrackingCount("chr_trigger_count")
				return post, ""
			}
			return post, ""
		}
	}
	return post, ""
}

func (p *Plugin) IncrementTrackingCount(trackingCount string) {
	data, error := p.API.KVGet(trackingCount)
	if error != nil {
		p.API.KVSet(trackingCount, []byte("0"))
	} else {
		triggerCount, _ := strconv.Atoi(string(data))
		triggerCount++
		p.API.KVSet(trackingCount, []byte(strconv.Itoa(triggerCount)))
	}
}

func (p *Plugin) MessageWillBePosted(_ *plugin.Context, post *model.Post) (*model.Post, string) {
	return p.modifyMessage(post)
}

func (p *Plugin) MessageWillBeUpdated(_ *plugin.Context, post *model.Post) (*model.Post, string) {
	return p.modifyMessage(post)
}
