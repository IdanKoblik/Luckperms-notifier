package main

import (
	"log"
	"luckperms-notifier/utils"
	"luckperms-notifier/config"
)

func main() {
	username, err := config.GetName()
	if err != nil {
		log.Fatalf("Error while getting webhook username: %v\n", err)
	}

	url, err := config.GetURL()
	if (err != nil) {
		log.Fatalf("Error while getting webhook url: %v\n", err)
	}

	test := "test"

	embed := utils.Embed{
		Title: &test,
	}

	message := utils.Message{
		Username: &username,
		Embeds:   &[]utils.Embed{embed},
	}

	err = utils.SendMessage(url, &message)
	if err != nil {
		log.Fatal(err)
	}
}