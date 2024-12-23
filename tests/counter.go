package tests

import (
	"context"
	"fmt"
	"log"
	"test-mayflower/pages"
	"time"

	"github.com/chromedp/chromedp"
)

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
