package tgo

import "time"
import "log"

type requestMaker interface {
	makeRequest(method string, params interface{}, val interface{}) error
	makeMultipartRequest(method string, params interface{}, val interface{}) error
}

type Client struct {
	requestMaker requestMaker
}

func NewClient(token string) *Client {
	return &Client{
		requestMaker: &apiRequestMaker{token: token},
	}
}

func (c *Client) GetUpdatesChan(params GetUpdatesParams) chan *Update {
	updatesChan := make(chan *Update)

	go func() {
		for {
			updates, err := c.GetUpdates(params)
			if err != nil {
				log.Println(err)
				time.Sleep(10 * time.Second)
				continue
			}

			if len(*updates) == 0 {
				continue
			}

			for _, update := range *updates {
				updatesChan <- &update
			}

			params.Offset = (*updates)[len(*updates)-1].ID + 1
		}
	}()

	return updatesChan
}

func (c *Client) GetUpdates(params GetUpdatesParams) (*[]Update, error) {
	u := new([]Update)

	err := c.requestMaker.makeRequest("getUpdates", params, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// TODO:
// - setWebhook
// - deleteWebhook
// - getWebhookInfo

func (c *Client) GetMe() (*User, error) {
	u := new(User)

	err := c.requestMaker.makeRequest("getMe", nil, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (c *Client) SendMessage(params SendMessageParams) (*Message, error) {
	m := new(Message)

	err := c.requestMaker.makeRequest("sendMessage", params, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) ForwardMessage(params ForwardMessageParams) (*Message, error) {
	m := new(Message)

	err := c.requestMaker.makeRequest("forwardMessage", params, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) SendPhoto(params SendPhotoParams) (*Message, error) {
	m := new(Message)

	switch params.Photo.(type) {
	case InputFilePath:
		err := c.requestMaker.makeMultipartRequest("sendPhoto", params, m)
		if err != nil {
			return nil, err
		}
	default:
		err := c.requestMaker.makeRequest("sendPhoto", params, m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (c *Client) SendAudio(params SendAudioParams) (*Message, error) {
	m := new(Message)

	switch params.Audio.(type) {
	case InputFilePath:
		err := c.requestMaker.makeMultipartRequest("sendAudio", params, m)
		if err != nil {
			return nil, err
		}
	default:
		err := c.requestMaker.makeRequest("sendAudio", params, m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (c *Client) SendDocument(params SendDocumentParams) (*Message, error) {
	m := new(Message)

	switch params.Document.(type) {
	case InputFilePath:
		err := c.requestMaker.makeMultipartRequest("sendDocument", params, m)
		if err != nil {
			return nil, err
		}
	default:
		err := c.requestMaker.makeRequest("sendDocument", params, m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (c *Client) SendSticker(params SendStickerParams) (*Message, error) {
	m := new(Message)

	switch params.Sticker.(type) {
	case InputFilePath:
		err := c.requestMaker.makeMultipartRequest("sendSticker", params, m)
		if err != nil {
			return nil, err
		}
	default:
		err := c.requestMaker.makeRequest("sendSticker", params, m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (c *Client) SendVideo(params SendVideoParams) (*Message, error) {
	m := new(Message)

	switch params.Video.(type) {
	case InputFilePath:
		err := c.requestMaker.makeMultipartRequest("sendVideo", params, m)
		if err != nil {
			return nil, err
		}
	default:
		err := c.requestMaker.makeRequest("sendVideo", params, m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (c *Client) SendVoice(params SendVoiceParams) (*Message, error) {
	m := new(Message)

	switch params.Voice.(type) {
	case InputFilePath:
		err := c.requestMaker.makeMultipartRequest("sendVoice", params, m)
		if err != nil {
			return nil, err
		}
	default:
		err := c.requestMaker.makeRequest("sendVoice", params, m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (c *Client) SendLocation(params SendLocationParams) (*Message, error) {
	m := new(Message)

	err := c.requestMaker.makeRequest("sendLocation", params, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) SendVenue(params SendVenueParams) (*Message, error) {
	m := new(Message)

	err := c.requestMaker.makeRequest("sendVenue", params, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) SendContact(params SendContactParams) (*Message, error) {
	m := new(Message)

	err := c.requestMaker.makeRequest("sendContact", params, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) SendChatAction(params SendChatActionParams) (*bool, error) {
	b := new(bool)

	err := c.requestMaker.makeRequest("sendChatAction", params, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) GetUserProfilePhotos(params GetUserProfilePhotosParams) (*UserProfilePhotos, error) {
	upp := new(UserProfilePhotos)

	err := c.requestMaker.makeRequest("sendChatAction", params, upp)
	if err != nil {
		return nil, err
	}

	return upp, nil
}

// https://core.telegram.org/bots/api#getfile
// TODO

func (c *Client) KickChatMember(params KickChatMemberParams) (*bool, error) {
	b := new(bool)

	err := c.requestMaker.makeRequest("kickChatMember", params, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) LeaveChat(params LeaveChatParams) (*bool, error) {
	b := new(bool)

	err := c.requestMaker.makeRequest("leaveChat", params, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) UnbanChatMember(params UnbanChatMemberParams) (*bool, error) {
	b := new(bool)

	err := c.requestMaker.makeRequest("unbanChatMember", params, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) GetChat(params GetChatParams) (*Chat, error) {
	chat := new(Chat)

	err := c.requestMaker.makeRequest("getChat", params, chat)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (c *Client) GetChatAdministrators(params GetChatParams) (*[]ChatMember, error) {
	cm := new([]ChatMember)

	err := c.requestMaker.makeRequest("getChatAdministrators", params, cm)
	if err != nil {
		return nil, err
	}

	return cm, nil
}

func (c *Client) GetChatMembersCount(params GetChatMembersCountParams) (*uint, error) {
	u := new(uint)

	err := c.requestMaker.makeRequest("getChatMembersCount", params, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (c *Client) GetChatMember(params GetChatMemberParams) (*ChatMember, error) {
	cm := new(ChatMember)

	err := c.requestMaker.makeRequest("getChatMember", params, cm)
	if err != nil {
		return nil, err
	}

	return cm, nil
}

func (c *Client) AnswerCallbackQuery(params AnswerCallbackQueryParams) (*bool, error) {
	b := new(bool)

	err := c.requestMaker.makeRequest("answerCallbackQuery", params, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) AnswerInlineQuery(params AnswerInlineQueryParams) (*bool, error) {
	b := new(bool)

	err := c.requestMaker.makeRequest("answerInlineQuery", params, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
