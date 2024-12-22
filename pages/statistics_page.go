package pages

import (
	"context"

	"github.com/chromedp/chromedp"
)

// Локаторы для страницы статистики
const (
	StatisticsURL          = "https://stripcash.com/analytics/statistics"
	StatisticsCounterXpath = `//div[@qa-element="summary-table"]/div[1]/div[1]/div[1]/div[1]/div[2]/div[2]/div[1]/div[position()=last()-1]`
	RunReportButton        = `button[type="submit"]`
)

type StatisticsPage struct {
	ctx context.Context
}

func NewStatisticsPage(ctx context.Context) *StatisticsPage {
	return &StatisticsPage{ctx: ctx}
}

func (sp *StatisticsPage) OpenPage() error {
	err := chromedp.Run(sp.ctx,
		chromedp.Navigate(StatisticsURL),
	)
	return err
}

func (sp *StatisticsPage) GetCounter() (string, error) {
	var counter string
	err := chromedp.Run(sp.ctx,
		chromedp.WaitVisible(StatisticsCounterXpath),
		chromedp.Text(StatisticsCounterXpath, &counter, chromedp.NodeVisible),
	)
	return counter, err
}

func (sp *StatisticsPage) RunReport() error {
	return chromedp.Run(sp.ctx,
		chromedp.WaitVisible(RunReportButton),
		chromedp.Click(RunReportButton),
		chromedp.WaitVisible(StatisticsCounterXpath),
	)
}
