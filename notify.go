package main

import (
	"fmt"
	"os"

	hook "github.com/versai-pro/discord-go/webhook"
)

func notify(q Question) {
	webhook := os.Getenv("DISCORD_WEBHOOK")

	if webhook != "" {
		fmt.Println(webhook)
		client := hook.NewClient(webhook)

		embed := hook.NewEmbed().SetTitle(q.Title).SetDescription(q.Body).SetAuthor(q.Author, "https://duckduckgo.com/?t=ffab&q=%22"+q.Author+"%22", "https://files.catbox.moe/7hz92d.jpeg")

		if err := client.SendEmbed(embed); err != nil {
			fmt.Println("Invalid webhook url supplied")
		}
	}
}
