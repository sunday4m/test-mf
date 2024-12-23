package main

import (
	"context"
	"log"
	"os"

	"test-mayflower/tests"

	"github.com/chromedp/chromedp"
)

func main() {
	login := os.Getenv("LOGIN_USERNAME")
	password := os.Getenv("LOGIN_PASSWORD")

	if login == "" || password == "" {
		log.Fatal("LOGIN_USERNAME и LOGIN_PASSWORD должны быть установлены")
	}

	// Настраиваем параметры для браузера
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("start-maximized", true),
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()

	ctx, ctxCancel := chromedp.NewContext(allocCtx)
	defer ctxCancel()

	// Выполняем тест
	err := tests.TestCounterIncrease(ctx, login, password)
	if err != nil {
		log.Fatalf("Тест завершился с ошибкой: %v", err)
	}

	log.Println("Тест успешно завершен!")
}
