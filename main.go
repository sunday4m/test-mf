package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"test-mayflower/pages"

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
	err := TestCounterIncrease(ctx, login, password)
	if err != nil {
		log.Fatalf("Тест завершился с ошибкой: %v", err)
	}

	log.Println("Тест успешно завершен!")
}

func TestCounterIncrease(parentCtx context.Context, login string, password string) error {
	ctx, ctxCancel := context.WithTimeout(parentCtx, 180*time.Second)
	defer ctxCancel()

	// Инициализация страниц
	loginPage := pages.NewLoginPage(ctx)
	modalPage := pages.NewModalPage(ctx)
	browsePage := pages.NewBrowsePage(ctx)
	statisticsPage := pages.NewStatisticsPage(ctx)

	// Авторизация
	if err := loginPage.Login(login, password); err != nil {
		return err
	}

	// Закрытие всплывающей модалки
	if err := modalPage.CloseModal(); err != nil {
		return err
	}

	// Копирование ссылки
	copiedLink, err := browsePage.CopyLink()
	if err != nil {
		return err
	}

	// Переход на страницу статистики и создание отчета
	if err := statisticsPage.OpenPage(); err != nil {
		return err
	}
	if err := statisticsPage.RunReport(); err != nil {
		return err
	}

	// Получаем счетчик до перехода по ссылке
	counterBefore, err := statisticsPage.GetCounter()
	if err != nil {
		return err
	}
	log.Printf("Счетчик до клика: %s", counterBefore)

	// Открываем новую вкладку по скопированной ссылке
	if err := chromedp.Run(ctx,
		chromedp.Navigate(copiedLink),
		chromedp.Sleep(time.Second*40)); err != nil { // Слип, чтобы подождать пока счетчик обновится
		return err
	}

	// Переход на страницу статистики и создание отчета
	if err := statisticsPage.OpenPage(); err != nil {
		return err
	}
	if err := statisticsPage.RunReport(); err != nil {
		return err
	}

	// Получаем счетчик после перехода по ссылке
	counterAfter, err := statisticsPage.GetCounter()
	if err != nil {
		return err
	}
	log.Printf("Счетчик после клика: %s", counterAfter)

	if counterAfter <= counterBefore {
		return fmt.Errorf("счетчик не увеличился: before=%s, after=%s", counterBefore, counterAfter)
	}

	return nil
}
