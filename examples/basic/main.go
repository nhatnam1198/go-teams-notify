// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/go-teams-notify
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

/*

This is an example of a simple client application which uses this library.

Of note:

- default timeout
- package-level logging is disabled by default
- validation of known webhook URL prefixes is *enabled*
- simple message submitted to Microsoft Teams consisting of formatted body and
  title

*/

package main

import (
	goteamsnotify "github.com/phuc1998/go-teams-notify"
)

func main() {
	_ = sendTheMessage()
}

func sendTheMessage() error {
	// init the client
	mstClient := goteamsnotify.NewClient()

	// setup webhook url
	webhookUrl := "https://outlook.office.com/webhook/YOUR_WEBHOOK_URL_OF_TEAMS_CHANNEL"

	// setup message card
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = "Hello world"
	msgCard.Text = "Here are some examples of formatted stuff like " +
		"<br> * this list itself  <br> * **bold** <br> * *italic* <br> * ***bolditalic***"
	msgCard.ThemeColor = "#DF813D"

	// send
	return mstClient.Send(webhookUrl, msgCard)
}
