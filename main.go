package main

import (
	"log"

	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"net/http"
	"os"
)

type Meme struct {
	Image  string
	Reddit string
	Title  string
}

type Neko struct {
   Url string
}


func main() {
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {

			case "meme":
				resp, err := http.Get("https://nksamamemeapi.pythonanywhere.com")
				if err != nil {
					fmt.Println(err)
				} else {
					data, _ := ioutil.ReadAll(resp.Body)
					fmt.Println(string(data))
					var meme Meme
					json.Unmarshal([]byte(data), &meme)
					fmt.Printf(meme.Title)
					file := meme.Image
					msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, nil)
					msg.FileID = file
					msg.Caption = meme.Title
					msg.UseExisting = true
					bot.Send(msg)
				}
			case "start":
				msg.Text = "Hello"
			case "help":
				msg.Text = "/meme"

			case "neko":
				res, err := http.Get("https://api.waifu.pics/sfw/neko")
				if err != nil {
					fmt.Println(err)
				} else {
					data, _ := ioutil.ReadAll(res.Body)
					// fmt.Println(string(data))
					var cat Neko
					json.Unmarshal([]byte(data), &cat)
					fmt.Println(cat.Url)
					nekochan := cat.Url
					msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, nil)
					msg.FileID = nekochan
					msg.UseExisting = true
					bot.Send(msg)

				}
				
			case "hentai":
				res, err := http.Get("https://api.waifu.pics/nsfw/neko")
				if err != nil {
					fmt.Println(err)
				} else {
					data, _ := ioutil.ReadAll(res.Body)
					// fmt.Println(string(data))
					var cat Neko
					json.Unmarshal([]byte(data), &cat)
					fmt.Println(cat.Url)
					nekochan := cat.Url
					msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, nil)
					msg.FileID = nekochan
					msg.UseExisting = true
					bot.Send(msg)

				}

			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}

	}
}
