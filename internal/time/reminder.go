package time

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"GriBotTimeLev/internal/config"
)

// Функция для отправки сообщения
func sendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка отправки сообщения: %v", err)
	}
}

// Функция напоминания о времени
func StartTimeReminder(bot *tgbotapi.BotAPI) {
	for {
		currentTime := time.Now().Format("15:04")
		message := fmt.Sprintf("Текущее время: %s", currentTime)
		sendMessage(bot, config.ChatID, message)
		log.Printf("Отправлено сообщение с временем: %s", currentTime)
		time.Sleep(config.TimeInterval)
	}
}

