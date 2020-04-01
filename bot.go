package main

import (
	"JestAndMemeBot/pkg"
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func prediction(client *pkg.Client, args string) string {
	var pageLine []int
	arg := strings.Split(args, " ")
	for _, val := range arg {
		i, err := strconv.Atoi(val)
		if err != nil {
			return "дай координаты"
		}
		pageLine = append(pageLine, i)
	}
	return client.GetPrediction(pageLine[0], pageLine[1])
}

func startJestBot(ctx context.Context) error {
	var (
		BotToken   = os.Getenv("BotToken")
		Port       = os.Getenv("PORT")
		WebhookURL = os.Getenv("WebhookURL")
	)
	log.Printf("%+v", tgbotapi.NewWebhook(WebhookURL))

	client := &pkg.Client{
		InternalClient:        &fasthttp.Client{},
		VKToken:               os.Getenv("VKToken"),
		YandexToken:           os.Getenv("YandexToken"),
		RandomToken:           os.Getenv("RandomToken"),
		TextContentURL:        "https://archive.org/stream/InfiniteJest/Infinite-Jest_djvu.txt",
		PictureContentGroupID: "-51322788",
	}

	bot, _ := tgbotapi.NewBotAPI(BotToken)
	bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	updates := bot.ListenForWebhook("/")

	go func() {
		http.ListenAndServe(":"+Port, nil)
	}()
	log.Println("jestbot started listening")

	for update := range updates {
		upd := update.Message

		if upd == nil { // ignore any non-Message updates
			continue
		}

		if !upd.IsCommand() { // ignore any non-command Messages
			continue
		}

		msg := tgbotapi.NewMessage(upd.Chat.ID, "")

		switch upd.Command() {
		case "prediction":
			msg.Text = prediction(client, upd.CommandArguments())
		case "meme":
			s := client.GetPicture()
			u, _ := url.Parse(s)
			m := tgbotapi.NewPhotoUpload(upd.Chat.ID, u)
			log.Printf("%+v", m)
			bot.Send(m)
			msg.Text = "tried to upload"
		case "full":
			msg.Text = prediction(client, upd.CommandArguments()) + "\n" + client.GetPicture()
		default:
			msg.Text = "я так не умею"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

	return nil
}

func main() {
	startJestBot(context.Background())
}
