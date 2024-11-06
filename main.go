package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"GriBotTimeLev/internal/config"
	"GriBotTimeLev/internal/time"
)

func main() {
	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %v", err)
	}

	bot.Debug = true // включить режим отладки, если нужно
	log.Println("Запуск GriBotTimeLev...")

	// Запуск функции напоминания о времени в отдельной горутине
	go time.StartTimeReminder(bot)

	// Основной цикл для обработки команд бота, если это понадобится в будущем
	select {} // Используем select{} чтобы оставить бота в запущенном состоянии
}

