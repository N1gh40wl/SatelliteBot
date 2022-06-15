package main

import (
	"fmt"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
)

func telegramBot() {

	//Создаем бота
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		//Проверяем что от пользователья пришло именно текстовое сообщение
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

			switch update.Message.Text {
			case "/start":

				//Отправлем сообщение
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm a satellite photo bot, i can show you photo from NASA satellites, send me your city")
				bot.Send(msg)
			default:
				//Отправлем сообщение
				//https://api.nasa.gov/planetary/earth/imagery?lon=100.75&lat=1.5&date=2014-02-01&api_key=DEMO_KEY

				lon, lat, _ := runGetYandex(update.Message.Text)

				// 55.75 37.61
				//lon := "55.75"
				//lat := "37.61"

				rss, _ := runGetNASA(lon, lat)

				resp, err := http.Get(rss)
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()

				// Create the file
				photoName := string(update.Message.Chat.ID) + ".png"
				out, err := os.Create(photoName)
				if err != nil {
					panic(err)
				}
				defer out.Close()

				// Write the body to file
				_, err = io.Copy(out, resp.Body)

				photoBytes, err := ioutil.ReadFile(photoName)
				if err != nil {
					panic(err)
				}
				photoFileBytes := tgbotapi.FileBytes{
					Name:  "picture",
					Bytes: photoBytes,
				}
				chatID := update.Message.Chat.ID
				message, err := bot.Send(tgbotapi.NewPhotoUpload(int64(chatID), photoFileBytes))

				fmt.Printf(message.Text)
				err = os.Remove(photoName)
				if err != nil {
					panic(err)
				}
				//msg := tgbotapi.NewMessage(update.Message.Chat.ID, rss)

				//bot.Send(msg)

			}
		}
	}
}
