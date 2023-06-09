package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", test)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	log.Println("New req!!11!11")
	_, err := w.Write([]byte("ТУТ МОГЛО БЫТЬ ЧТО-ТО БЫТЬ, НО ТУТ ТОЛЬКО НИХУЯ"))
	if err != nil {
		return
	}

	bot, err := tgbotapi.NewBotAPI("6045862064:AAFes5NsbKdIWm9BQ1qPwIDeHoUmTRnUr0E")
	if err != nil {
		log.Fatal("Error creating bot: ", err)
	}

	msg := tgbotapi.NewMessage(-906041044, "И ТУТ НИХУЯ\n https://img.armtek.ru/img/article/979/9795946/500x500/9795946_0.png")

	if _, err := bot.Send(msg); err != nil {
		log.Fatal("Error sending message: ", err)
	}

	//photo := tgbotapi.NewPhoto(-906041044, "https://img.armtek.ru/img/article/979/9795946/500x500/9795946_0.png")
	//_, err = bot.Send(photo)
	//if err != nil {
	//	panic(err)
	//}
}
