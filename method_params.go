package tgo

// GetUpdatesParams https://core.telegram.org/bots/api#getupdates
type GetUpdatesParams struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          uint     `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// https://core.telegram.org/bots/api#sendmessage
type SendMessageParams struct {
	ChatID                interface{} `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID      int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#forwardmessage
type ForwardMessageParams struct {
	ChatID              interface{} `json:"chat_id"`
	FromChatID          interface{} `json:"from_chat_id"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	MessageID           int         `json:"message_id"`
}

// https://core.telegram.org/bots/api#sendphoto
type SendPhotoParams struct {
	ChatID              interface{} `json:"chat_id"                        form:"chat_id"`
	Photo               interface{} `json:"photo"                          form:"photo,inputFile"`
	Caption             string      `json:"caption,omitempty"              form:"caption,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"  form:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"         form:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendaudio
type SendAudioParams struct {
	ChatID              interface{} `json:"chat_id"                        form:"chat_id"`
	Audio               interface{} `json:"audio"                          form:"audio,inputFile"`
	Caption             string      `json:"caption,omitempty"              form:"caption,omitempty"`
	Duration            uint        `json:"duration,omitempty"             form:"duration,omitempty"`
	Performer           string      `json:"performer,omitempty"            form:"performer,omitempty"`
	Title               string      `json:"title,omitempty"                form:"title,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"  form:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"         form:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#senddocument
type SendDocumentParams struct {
	ChatID              interface{} `json:"chat_id"                        form:"chat_id"`
	Document            interface{} `json:"document"                       form:"document,inputFile"`
	Caption             string      `json:"caption,omitempty"              form:"caption,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"  form:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"         form:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendsticker
type SendStickerParams struct {
	ChatID              interface{} `json:"chat_id"                        form:"chat_id"`
	Sticker             interface{} `json:"sticker"                        form:"sticker,inputFile"`
	Caption             string      `json:"caption,omitempty"              form:"caption,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"  form:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"         form:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvideo
type SendVideoParams struct {
	ChatID              interface{} `json:"chat_id"                        form:"chat_id"`
	Video               interface{} `json:"video"                          form:"video,inputFile"`
	Duration            uint        `json:"duration,omitempty"             form:"duration,omitempty"`
	Width               uint        `json:"width,omitempty"                form:"width,omitempty"`
	Height              uint        `json:"height,omitempty"               form:"height,omitempty"`
	Caption             string      `json:"caption,omitempty"              form:"caption,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"  form:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"         form:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvoice
type SendVoiceParams struct {
	ChatID              interface{} `json:"chat_id"                        form:"chat_id"`
	Voice               interface{} `json:"voice"                          form:"voice,inputFile"`
	Caption             string      `json:"caption,omitempty"              form:"caption,omitempty"`
	Duration            uint        `json:"duration,omitempty"             form:"duration,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"  form:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"         form:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendlocation
type SendLocationParams struct {
	ChatID              interface{} `json:"chat_id"`
	Latitude            float32     `json:"latitude"`
	Longitude           float32     `json:"longitude"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvenue
type SendVenueParams struct {
	ChatID              interface{} `json:"chat_id"`
	Latitude            float32     `json:"latitude"`
	Longitude           float32     `json:"longitude"`
	Title               string      `json:"title"`
	Address             string      `json:"address"`
	FoursquareID        string      `json:"foursquare_id,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendcontact
type SendContactParams struct {
	ChatID              interface{} `json:"chat_id"`
	PhoneNumber         string      `json:"phone_number"`
	FirstName           string      `json:"first_name"`
	LastName            string      `json:"last_name,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendchataction
type SendChatActionParams struct {
	ChatID interface{} `json:"chat_id"`
	Action string      `json:"action"`
}

// https://core.telegram.org/bots/api#getuserprofilephotos
type GetUserProfilePhotosParams struct {
	UserID int  `json:"user_id"`
	Offset int  `json:"offset,omitempty"`
	Limit  uint `json:"limit,omitempty"`
}

// https://core.telegram.org/bots/api#kickchatmember
type KickChatMemberParams struct {
	ChatID interface{} `json:"chat_id"`
	UserID int         `json:"user_id"`
}

// https://core.telegram.org/bots/api#leavechat
type LeaveChatParams struct {
	ChatID interface{} `json:"chat_id"`
}

// https://core.telegram.org/bots/api#unbanchatmember
type UnbanChatMemberParams struct {
	ChatID interface{} `json:"chat_id"`
	UserID int         `json:"user_id"`
}

// https://core.telegram.org/bots/api#getchat
type GetChatParams struct {
	ChatID interface{} `json:"chat_id"`
}

// https://core.telegram.org/bots/api#getchatadministrators
type GetChatAdministratorsParams struct {
	ChatID interface{} `json:"chat_id"`
}

// https://core.telegram.org/bots/api#getchatmemberscount
type GetChatMembersCountParams struct {
	ChatID interface{} `json:"chat_id"`
}

// https://core.telegram.org/bots/api#getchatmember
type GetChatMemberParams struct {
	ChatID interface{} `json:"chat_id"`
	UserID int         `json:"user_id"`
}

// https://core.telegram.org/bots/api#answercallbackquery
type AnswerCallbackQueryParams struct {
	ID        string `json:"callback_query_id"`
	Text      string `json:"text,omitempty"`
	ShowAlert bool   `json:"show_alert,omitempty"`
	URL       string `json:"url,omitempty"`
	CacheTime uint   `json:"cache_time,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequery
type AnswerInlineQueryParams struct {
	InlineQueryID     string              `json:"inline_query_id"`
	Results           []InlineQueryResult `json:"results"`
	CacheTime         uint                `json:"cache_time,omitempty"`
	IsPersonal        bool                `json:"is_personal,omitempty"`
	NextOffset        string              `json:"next_offset,omitempty"`
	SwitchPmText      string              `json:"switch_pm_text,omitempty"`
	SwitchPmParameter string              `json:"switch_pm_parameter,omitempty"`
}
