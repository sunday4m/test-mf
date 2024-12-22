package pages

import (
	"context"

	"github.com/chromedp/chromedp"
)

const (
	LoginURL           = "https://stripcash.com/login"
	LoginUsernameField = `input[name="username"]`
	LoginPasswordField = `input[name="password"]`
	LoginSubmitButton  = `button[type="submit"]`
)

type LoginPage struct {
	ctx context.Context
}

func NewLoginPage(ctx context.Context) *LoginPage {
	return &LoginPage{ctx: ctx}
}

func (lp *LoginPage) Login(username, password string) error {
	return chromedp.Run(lp.ctx,
		chromedp.Navigate(LoginURL),
		chromedp.SendKeys(LoginUsernameField, username),
		chromedp.SendKeys(LoginPasswordField, password),
		chromedp.Click(LoginSubmitButton),
	)
}
