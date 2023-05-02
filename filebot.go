package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("Slack bot token", "xoxb-5177666699623-5197699722739-p3rzqIDyOadJAGZM1WITKlJ5")
	os.Setenv("Channel_ID", "C0557KLMC5D")
	api := slack.New(os.Getenv("Slack bot token"))
	channelArr := []string{os.Getenv("Channel_ID")}
	fileArr := []string{"first.txt"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s \n", file.Name, file.URL)
	}

}
