package model

type User struct {
	APIKey      string       `json:"apikey" bson:"apikey"`
	APISecret   string       `json:"apiSecrete" bson:"apiSecrete"`
	UserName    string       `json:"userName" bson:"userName"`
	Password    string       `json:"password" bson:"password"`
	TOTPToken   string       `json:"totpToken" bson:"totpToken"`
	Name        string       `json:"name" bson:"name"`
	AccessToken string       `json:"acc_token" bson:"acc_token"`
	Telegram    TelegramInfo `json:"telegramInfo" bson:"telegramInfo"`
}

type TelegramInfo struct {
	BotToken  string `json:"bot_token" bson:"bot_token"`
	BotChatID string `json:"botchatId" bson:"botchatId"`
}
