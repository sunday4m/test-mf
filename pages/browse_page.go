package pages

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
	"golang.design/x/clipboard"
)

// Локаторы для страницы статистики
const (
	CardButtonXpath = `(//div[@qa-element="card-body"])[7]//button`
)

type BrowsePage struct {
	ctx context.Context
}

func NewBrowsePage(ctx context.Context) *BrowsePage {
	return &BrowsePage{ctx: ctx}
}

func (bp *BrowsePage) CopyLink() (string, error) {
	err := chromedp.Run(bp.ctx,
		chromedp.WaitVisible(CardButtonXpath),
		chromedp.Click(CardButtonXpath),
	)

	if err = clipboard.Init(); err != nil {
		log.Fatalf("Ошибка при инициализации буфера обмена: %v", err)
	}

	link := clipboard.Read(clipboard.FmtText)
	if link == nil {
		log.Fatalf("Буфер обмена пуст или не содержит текст. %v", err)
	}
	copiedURL := string(link)

	if !strings.HasPrefix(copiedURL, "http://") && !strings.HasPrefix(copiedURL, "https://") {
		return "", fmt.Errorf("некорректная ссылка: %s", link)
	}

	return copiedURL, err
}
