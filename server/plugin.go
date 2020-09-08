package main

import (
	"encoding/json"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

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

	triggerSentences []string
	// BotId of the created bot account.
	botID string
	// Probability for the trigger message to be sent.
	probabilityFactor int
}

func (p *Plugin) containsTriggerSentence(message string) bool {
	// Check if the message contains any of the trigger sentence
	for _, triggerSentence := range p.triggerSentences {
		matched, _ := regexp.MatchString(triggerSentence, message)
		if matched {
			return true
		}
	}
	return false
}

func (p *Plugin) modifyMessage(post *model.Post) (*model.Post, string) {
	message := post.Message
	if p.containsTriggerSentence(strings.ToLower(message)) {
		// Send the trigger message based on probability here
		p.IncrementTrackingCount("trigger_sentence_detect_count")
		if rand.Intn(10-1)+1 < (p.probabilityFactor / 10) {
			// Sending the Trigger message when chr trigger word has been detected and satsifies the probability factor
			ephemeralPost := &model.Post{
				UserId:    p.botID,
				ChannelId: post.ChannelId,
				Props: model.StringInterface{
					"attachments": []*model.SlackAttachment{
						{
							Text:       "Now, you can ask doubts for any concept and get them answered by getting on a ***1-1 live call of ~ 15 minutes with a helper/TA***.",
							Pretext:    "##### Seems like you have a Doubt! :thinking:. You can click on Raise Concept Help Request link below to get on call :telephone_receiver: with a TA ",
							AuthorName: "Scaler",
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
			p.IncrementTrackingCount("chr_message_trigger_count")
			// Tracking for the user the the trigger message have been sent to
			postedUser, _ := p.API.GetUser(post.UserId)
			p.IncrementTrackingCount("chr_trigr_usr_" + postedUser.Username)
			// Tracking for the channel the trigger message have been sent to
			p.IncrementTrackingCount("chr_trigr_chnl_" + post.ChannelId)
			postedChannel, err := p.API.GetChannel(post.ChannelId)
			if postedChannel.TeamId != "" && err == nil {
				// Tracking for the team the trigger message have been sent to
				p.IncrementTrackingCount("chr_trigr_team_" + postedChannel.TeamId)
			}
			// Date wise Tracking for the post
			currentTime := time.Now().Local()
			postList, err := p.API.KVGet("chr_trgr_date_" + currentTime.Format("2006-01-02"))
			if postList == nil {
				newPostList := []model.StringInterface{{
					"post":         post,
					"channel_name": postedChannel.Name,
					"user":         postedUser.Username,
				}}
				newPostListJSON, _ := json.Marshal(newPostList)
				p.API.KVSet("chr_trgr_date_"+currentTime.Format("2006-01-02"), newPostListJSON)
			} else {
				newPostList := []model.StringInterface{}
				json.Unmarshal(postList, &newPostList)
				newPostList = append(newPostList, model.StringInterface{
					"post":         post,
					"channel_name": postedChannel.Name,
					"user":         postedUser.Username,
				})
				newPostListJSON, _ := json.Marshal(newPostList)
				p.API.KVSet("chr_trgr_date_"+currentTime.Format("2006-01-02"), newPostListJSON)
			}

			return post, ""
		}
		return post, ""
	}
	return post, ""
}

// IncrementTrackingCount : A function to implement tracking for the plugin
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

// MessageWillBePosted : The function which runs right before the message is posted
func (p *Plugin) MessageWillBePosted(_ *plugin.Context, post *model.Post) (*model.Post, string) {
	return p.modifyMessage(post)
}

// MessageWillBeUpdated : The function which runs right before the message is updated
func (p *Plugin) MessageWillBeUpdated(_ *plugin.Context, post *model.Post) (*model.Post, string) {
	return p.modifyMessage(post)
}
