package hoi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/nlopes/slack"
)

type Notifier interface {
	Notify(to, message string) error
}

func NewNotifier(conf Notification) Notifier {
	switch conf.To {
	case "slack":
		return &SlackNotifier{
			From:   conf.From,
			Client: slack.New(conf.Token),
		}
	case "takosan":
		return &TakosanNotifier{
			From: conf.From,
			Host: conf.Host,
			Port: conf.Port,
		}
	default:
		return nil
	}
}

type SlackNotifier struct {
	From   string
	Client *slack.Slack
}

func (s SlackNotifier) Notify(to, message string) error {
	_, _, err := s.Client.PostMessage(
		to,
		message,
		slack.PostMessageParameters{
			Username: s.From,
		},
	)
	if err != nil {
		return fmt.Errorf("Failed to send message to %s: %s\n", to, err)
	}
	return nil
}

type TakosanNotifier struct {
	From string
	Host string
	Port int
}

func (t TakosanNotifier) Notify(to, message string) error {
	res, err := http.PostForm(
		fmt.Sprintf("http://%s:%d/privmsg", t.Host, t.Port),
		url.Values{"channel": {to}, "message": {message}},
	)
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if res.Status == "400" {
		return fmt.Errorf("%s\n", body)
	}
	return nil
}
