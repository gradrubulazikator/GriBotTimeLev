package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
    botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
    if botToken == "" {
        log.Fatal("TELEGRAM_BOT_TOKEN environment variable required")
    }

    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)

    reminderChan := make(chan string)

    go func() {
        for {
            time.Sleep(4 * time.Hour)
            currentTime, err := getCurrentTime()
            if err != nil {
                log.Println("Failed to get current time:", err)
                reminderChan <- "Напоминание: проверьте свое расписание!"
            } else {
                reminderChan <- fmt.Sprintf("Напоминание: текущее время %s. Проверьте свое расписание!", currentTime)
            }
        }
    }()

    go func() {
        for reminder := range reminderChan {
            msg := tgbotapi.NewMessage(123456789, reminder) // Замените 123456789 на ваш Telegram user ID
            if _, err := bot.Send(msg); err != nil {
                log.Println("Failed to send reminder:", err)
            }
        }
    }()

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message != nil && update.Message.Text == "/start" {
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать в Рычаг времени! Напоминания будут отправляться каждые 4 часа.")
            if _, err := bot.Send(msg); err != nil {
                log.Println("Failed to send message:", err)
            }
        }
    }
}

// getCurrentTime запрашивает текущее время через API и возвращает его в виде строки
func getCurrentTime() (string, error) {
    resp, err := http.Get("http://worldtimeapi.org/api/timezone/Etc/UTC")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var data struct {
        Datetime string `json:"datetime"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return "", err
    }

    return data.Datetime, nil
}

