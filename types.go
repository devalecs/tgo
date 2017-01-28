package tgo

import "encoding/json"

// Telegram Bot API Types
// https://core.telegram.org/bots/api

// https://core.telegram.org/bots/api#making-requests
type Response struct {
	Ok          bool            `json:"ok"`
	Description string          `json:"description,omitempty"`
	Result      json.RawMessage `json:"result,omitempty"`
	ErrorCode   int             `json:"error_code,omitempty"`
}

func (r *Response) encodeResult(v interface{}) error {
	return json.Unmarshal(r.Result, v)
}

// https://core.telegram.org/bots/api#update
type Update struct {
	ID                 int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
}

// https://core.telegram.org/bots/api#user
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
}

// https://core.telegram.org/bots/api#chat
type Chat struct {
	ID                          int    `json:"id"`
	Type                        string `json:"type"`
	Title                       string `json:"title,omitempty"`
	Username                    string `json:"username,omitempty"`
	FirstName                   string `json:"first_name,omitempty"`
	LastName                    string `json:"last_name,omitempty"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators,omitempty"`
}

// https://core.telegram.org/bots/api#message
type Message struct {
	ID                    int              `json:"message_id"`
	From                  User             `json:"from,omitempty"`
	Date                  int              `json:"date"`
	Chat                  *Chat            `json:"chat"`
	ForwardFrom           *User            `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat            `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID  int              `json:"forward_from_message_id,omitempty"`
	ForwardDate           int              `json:"forward_date,omitempty"`
	ReplyToMessage        *Message         `json:"reply_to_message,omitempty"`
	EditDate              int              `json:"edit_date,omitempty"`
	Text                  string           `json:"text,omitempty"`
	Entities              *[]MessageEntity `json:"entities,omitempty"`
	Audio                 *Audio           `json:"audio,omitempty"`
	Document              *Document        `json:"document,omitempty"`
	Game                  *Game            `json:"game,omitempty"`
	Photo                 *[]PhotoSize     `json:"photo,omitempty"`
	Sticker               *Sticker         `json:"sticker,omitempty"`
	Video                 *Video           `json:"video,omitempty"`
	Voice                 *Voice           `json:"voice,omitempty"`
	Caption               string           `json:"caption,omitempty"`
	Contact               *Contact         `json:"contact,omitempty"`
	Location              *Location        `json:"location,omitempty"`
	Venue                 *Venue           `json:"venue,omitempty"`
	NewChatMember         *User            `json:"new_chat_member,omitempty"`
	LeftChatMember        *User            `json:"left_chat_member,omitempty"`
	NewChatTitle          string           `json:"new_chat_title,omitempty"`
	NewChatPhoto          *[]PhotoSize     `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool             `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool             `json:"group_chat_created,omitempty"`
	SupergroupChatCreated bool             `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool             `json:"channel_chat_created,omitempty"`
	MigrateToChatID       int              `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID     int              `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message         `json:"pinned_message,omitempty"`
}

// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url,omitempty"`
	User   *User  `json:"user,omitempty"`
}

// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer,omitempty"`
	Title     string `json:"title,omitempty"`
	MimeType  string `json:"mime_type,omitempty"`
	FileSize  int    `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#document
type Document struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	Emoji    string     `json:"emoji,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#video
type Video struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize int    `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
}

// https://core.telegram.org/bots/api#location
type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

// https://core.telegram.org/bots/api#venue
type Venue struct {
	Location     *Location `json:"location"`
	Title        string    `json:"title"`
	Address      string    `json:"address"`
	FoursquareID string    `json:"foursquare_id"`
}

// https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	TotalCount int          `json:"total_count"`
	Photos     *[]PhotoSize `json:"photos"`
}

// https://core.telegram.org/bots/api#file
type File struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size,omitempty"`
	FilePath string `json:"file_path,omitempty"`
}

// https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard        *[][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                `json:"resize_keyboard"`
	OneTimeKeyboard bool                `json:"one_time_keyboard"`
	Selective       bool                `json:"selective"`
}

// https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact,omitempty"`
	RequestLocation bool   `json:"request_location,omitempty"`
}

// https://core.telegram.org/bots/api#replykeyboardremove
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard *[][]InlineKeyboardButton `json:"inline_keyboardt"`
}

// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url,omitempty"`
	Callbackdata                 string        `json:"callback_data,omitempty"`
	SwitchInlinequery            string        `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
}

// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance,omitempty"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

// https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"`
}

// https://core.telegram.org/bots/api#chatmember
type ChatMember struct {
	User   *User  `json:"user"`
	Status string `json:"status"`
}

// https://core.telegram.org/bots/api#responseparameters
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}

// https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

// https://core.telegram.org/bots/api#game
type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        *[]PhotoSize     `json:"photo"`
	Text         string           `json:"text,omitempty"`
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation       `json:"migrate_to_chat_id,omitempty"`
}

// https://core.telegram.org/bots/api#animation
type Animation struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#gamehighscore
type GameHighScore struct {
	Position int   `json:"position"`
	User     *User `json:"user"`
	Score    int   `json:"score"`
}

// https://core.telegram.org/bots/api#callbackgame
type CallbackGame struct{}

// Inline mode

// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// https://core.telegram.org/bots/api#inlinequeryresult
type InlineQueryResult interface{}

// https://core.telegram.org/bots/api#inlinequeryresultarticle
type InlineQueryResultArticle struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 string                `json:"url,omitempty"`
	HideURL             bool                  `json:"hide_url,omitempty"`
	Description         string                `json:"description,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultphoto
type InlineQueryResultPhoto struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhotoURL            string                `json:"photo_url"`
	ThumbURL            string                `json:"thumb_url"`
	PhotoWidth          int                   `json:"photo_width,omitempty"`
	PhotoHeight         int                   `json:"photo_height,omitempty"`
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultgif
type InlineQueryResultGif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	GifURL              string                `json:"gif_url"`
	GifWidth            int                   `json:"gif_width,omitempty"`
	GifHeight           int                   `json:"gif_height,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
type InlineQueryResultMpeg4Gif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Mpeg4URL            string                `json:"mpeg4_url"`
	Mpeg4Width          int                   `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int                   `json:"mpeg4_height,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultvideo
type InlineQueryResultVideo struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VideoURL            string                `json:"video_url"`
	MimeType            string                `json:"mime_type"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	VideoWidth          int                   `json:"video_width,omitempty"`
	VideoHeight         int                   `json:"video_height,omitempty"`
	VideoDuration       int                   `json:"video_duration,omitempty"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultaudio
type InlineQueryResultAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	Performer           string                `json:"performer,omitempty"`
	AudioDuration       int                   `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultvoice
type InlineQueryResultVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceURL            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	VoiceDuration       string                `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultdocument
type InlineQueryResultDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	DocumentURL         string                `json:"document_url"`
	MimeType            string                `json:"mime_type"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultlocation
type InlineQueryResultLocation struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Latitude            float32               `json:"latitude"`
	Longitude           float32               `json:"longitude"`
	Title               string                `json:"title"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultvenue
type InlineQueryResultVenue struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Latitude            string                `json:"latitude"`
	Longitude           string                `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareID        string                `json:"foursquare_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          string                `json:"thumb_width,omitempty"`
	ThumbHeight         string                `json:"thumb_height,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcontact
type InlineQueryResultContact struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            string                `json:"last_name,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultgame
type InlineQueryResultGame struct {
	Type          string               `json:"type"`
	ID            string               `json:"id"`
	GameShortName string               `json:"game_short_name"`
	ReplyMarkup   InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
type InlineQueryResultCachedPhoto struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhotoFileID         string                `json:"photo_file_id"`
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
type InlineQueryResultCachedGif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	GifFileID           string                `json:"gif_file_id"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
type InlineQueryResultCachedMpeg4Gif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Mpeg4FileID         string                `json:"mpeg4_file_id"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	StickerFileID       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	DocumentFileID      string                `json:"document_file_id"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
type InlineQueryResultCachedVideo struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VideoFileID         string                `json:"video_file_id"`
	Title               string                `json:"title"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioFileID         string                `json:"audio_file_id"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inputmessagecontent
type InputMessageContent interface{}

// https://core.telegram.org/bots/api#inputtextmessagecontent
type InputTextMessageContent struct {
	MessageText           string `json:"message_text"`
	ParseMode             string `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
}

// https://core.telegram.org/bots/api#inputlocationmessagecontent
type InputLocationMessageContent struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// https://core.telegram.org/bots/api#inputvenuemessagecontent
type InputVenueMessageContent struct {
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareID string  `json:"foursquare_id,omitempty"`
}

// https://core.telegram.org/bots/api#inputcontactmessagecontent
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
}

// https://core.telegram.org/bots/api#inputfile
type InputFilePath string
