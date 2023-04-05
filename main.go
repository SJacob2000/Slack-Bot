package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5066001063762-5059436515334-kUvSQjr5iTKQif5NHeaNyLYD")
	os.Setenv("CHANNEL_ID", "C052ML2BT4Y")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Shubham's Resume.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}

		fmt.Printf("Name %s, URL: %s\n", file.Name, file.URL)
	}
}
