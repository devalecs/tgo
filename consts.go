package tgo

const (
	TelegramAPIEndpoint = "https://api.telegram.org"

	// https://core.telegram.org/bots/api#chatmember
	ChatMemberStatusCreator       = "creator"
	ChatMemberStatusAdministrator = "administrator"
	ChatMemberStatusMember        = "member"
	ChatMemberStatusLeft          = "left"
	ChatMemberStatusKicked        = "kicked"

	// https://core.telegram.org/bots/api#markdown-style
	ParseModeMarkdown = "Markdown"
	// https://core.telegram.org/bots/api#html-style
	ParseModeHTML = "HTML"

	// https://core.telegram.org/bots/api#inlinequeryresult
	InlineQueryResultTypeArticle  = "article"
	InlineQueryResultTypePhoto    = "photo"
	InlineQueryResultTypeGif      = "gif"
	InlineQueryResultTypeMpeg4Gif = "mpeg4_gif"
	InlineQueryResultTypeVideo    = "video"
	InlineQueryResultTypeAudio    = "audio"
	InlineQueryResultTypeVoice    = "voice"
	InlineQueryResultTypeDocument = "document"
	InlineQueryResultTypeLocation = "location"
	InlineQueryResultTypeVenue    = "venue"
	InlineQueryResultTypeContact  = "contact"
	InlineQueryResultTypeGame     = "game"
	InlineQueryResultTypeSticker  = "sticker"

	MimeTypeTextHTML       = "text/html"
	MimeTypeVideoMP4       = "video/mp4"
	MimeTypeApplicationPDF = "application/pdf"
	MimeTypeApplicationZIP = "application/zip"
	MimeTypeAudioMPEG      = "audio/mpeg"
	MimeTypeAudioOgg       = "audio/ogg"
)
