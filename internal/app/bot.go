package app

type Bot struct {
	token string
}

func NewBot() *Bot {
	b := &Bot{
		token: "1436318012:AAF6mfxveYh213kd5Ge9ce_uydO3IqDjtuU",
	}
	return b
}
