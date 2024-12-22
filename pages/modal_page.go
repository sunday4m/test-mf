package pages

import (
	"context"

	"github.com/chromedp/chromedp"
)

const (
	LoginModalCloseButton = `//div[@qa-element="modal-title"]//button`
)

type ModalPage struct {
	ctx context.Context
}

func NewModalPage(ctx context.Context) *ModalPage {
	return &ModalPage{ctx: ctx}
}

func (mp *ModalPage) CloseModal() error {
	return chromedp.Run(mp.ctx,
		chromedp.WaitVisible(LoginModalCloseButton),
		chromedp.Click(LoginModalCloseButton),
	)
}
