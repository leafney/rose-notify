/**
 * @Author:      leafney
 * @Date:        2022-10-09 09:44
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package dingtalk

// Roboter is the interface implemented by Robot that can send multiple types of messages.
type Roboter interface {
	SetHost(url string)
	SetSecret(secret string)

	SendText(content string) error
	SendTextAt(content string, atMobiles []string, isAtAll bool) error
	SendLink(title, text, messageURL, picURL string) error
	SendMarkdown(title, text string) error
	SendMarkdownAt(title, text string, atMobiles []string, isAtAll bool) error
	//SendActionCard(title, text, singleTitle, singleURL, btnOrientation, hideAvatar string) error
	//SendActionCards(title, text string)error
	//SendFeedCard(title string)
}

// Robot represents a dingtalk custom robot that can send messages to groups.
type Robot struct {
	host   string
	token  string
	secret string
	debug  bool
}

func (r *Robot) SetSecret(secret string) {
	r.secret = secret
}

func (r *Robot) SetHost(host string) {
	r.host = host
}

func (r *Robot) SendText(content string) error {
	return r.send(&textMessage{
		MsgType: msgTypeText,
		Text: textParams{
			Content: content,
		},
	})
}

func (r *Robot) SendTextAt(content string, atMobiles []string, isAtAll bool) error {
	return r.send(&textMessage{
		MsgType: msgTypeText,
		Text: textParams{
			Content: content,
		},
		At: atParams{
			AtMobiles: atMobiles,
			IsAtAll:   isAtAll,
		},
	})
}

func (r *Robot) SendLink(title, text, messageURL, picURL string) error {
	return r.send(&linkMessage{
		MsgType: msgTypeLink,
		Link: linkParams{
			Title:      title,
			Text:       text,
			MessageURL: messageURL,
			PicURL:     picURL,
		},
	})
}

func (r *Robot) SendMarkdown(title, text string) error {
	return r.send(&markdownMessage{
		MsgType: msgTypeMarkdown,
		Markdown: markdownParams{
			Title: title,
			Text:  text,
		},
	})
}

func (r *Robot) SendMarkdownAt(title, text string, atMobiles []string, isAtAll bool) error {
	return r.send(&markdownMessage{
		MsgType: msgTypeMarkdown,
		Markdown: markdownParams{
			Title: title,
			Text:  text,
		},
		At: atParams{
			AtMobiles: atMobiles,
			IsAtAll:   isAtAll,
		},
	})
}

// NewRobot returns a roboter that can send messages.
func NewRobot(token string) *Robot {
	return &Robot{
		host:  "https://oapi.dingtalk.com/robot/send",
		token: token,
	}
}

func (r *Robot) DebugMode() *Robot {
	r.debug = true
	return r
}
