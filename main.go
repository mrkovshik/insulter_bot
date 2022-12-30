package main

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	Insulter "github.com/mrkovshik/Insulter_lib"
	years "github.com/mrkovshik/deathClocker_lib"

)


func main() {
	var msg tgbotapi.MessageConfig
  // подключаемся к боту с помощью токена
  bot, err := tgbotapi.NewBotAPI("5851301531:AAEwg_vNy7wEGt-PffSQE5XDc9XHX7gvXBE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
	
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			switch{
			case update.Message.Text == "/start": 
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, "+update.Message.From.FirstName+"! Добро пожаловать к Оскорблятелю! Выбери из списка команд, что ты хочешь узнать!")
			case update.Message.Text == "/insult": 
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, Insulter.Insult(update.Message.From.FirstName))
			case update.Message.Text == "/deathclock": 
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Хочешь узнать дату своей смерти? Тогда напиши свой возраст или слово 'Нет' для выхода")
			bot.Send(msg) 
			FirstLoop:
			for update := range updates {
				if update.Message != nil{
					age, _:=strconv.Atoi(update.Message.Text)
					switch{
						case update.Message.Text == "нет" ||  update.Message.Text == "Нет":
							msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Тогда возвращаемся в предыдущее меню")
							break FirstLoop
						case age > 0&& age <100:	
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, years.Calculate(age))
						bot.Send(msg)
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Выбери из списка команд, что ты хочешь узнать!")
						break FirstLoop
						case age <= 0|| age >100: 
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Чота сомнительна! Давай вводи нормальный возраст!")
						default:
							msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Ну же, "+update.Message.From.FirstName+" не тупи - введи свой возраст циферками или просто напиши слово 'нет', если передумал")
					}
					bot.Send(msg) 
				}
			}
		
		
		default:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Ну же, "+update.Message.From.FirstName+" не тупи - выбери команду из списка: хочешь обзывательство или дату смерти?")
		}
		bot.Send(msg)
	}
	
	}
}