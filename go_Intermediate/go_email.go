package main

import (
	"github.com/mailgun/mailgun-go"
)

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun("danielgongying@gmail.com","key-danielgG12345")
	m := mg.NewMessage(
		"Excited User <elliot@tutorialedge.net>",
		"Hello",
		"Testing some Mailgun!",
		"elliot@tutorialedge.net",
	)
	_, id, err := mg.Send(nil,m)
	return id, err
}

func main(){
	SendSimpleMessage("danielgongying@gmail.com", "key-danielgG12345")

}