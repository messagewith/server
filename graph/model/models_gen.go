// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Chat struct {
	ID                string             `json:"id"`
	User              *User              `json:"user"`
	Messages          []*Message         `json:"messages"`
	MessagesCount     int                `json:"messagesCount"`
	LastViewedMessage *LastViewedMessage `json:"lastViewedMessage"`
}

type ChatFilter struct {
	ID     *string `json:"id"`
	UserID *string `json:"userId"`
}

type LastViewedMessage struct {
	Message *Message `json:"message"`
	User    *User    `json:"user"`
	Time    int      `json:"time"`
}

type Message struct {
	ID          string      `json:"id"`
	User        *User       `json:"user"`
	MessageType MessageType `json:"messageType"`
	Content     string      `json:"content"`
	SendTime    int         `json:"sendTime"`
	Reactions   []*Reaction `json:"reactions"`
	IsForwarded bool        `json:"isForwarded"`
}

type Reaction struct {
	Emoji string `json:"emoji"`
	User  *User  `json:"user"`
}

type User struct {
	ID         string  `json:"id"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	MiddleName *string `json:"middleName"`
	FullName   string  `json:"fullName"`
	Nickname   string  `json:"nickname"`
	Email      string  `json:"email"`
}

type UserFilter struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email"`
	FullName  *string `json:"fullName"`
}

type UserInput struct {
	FirstName  string  `json:"firstName"`
	MiddleName *string `json:"middleName"`
	LastName   string  `json:"lastName"`
	Password   string  `json:"password"`
	Nickname   *string `json:"nickname"`
	Email      string  `json:"email"`
}

type MessageType string

const (
	MessageTypeText      MessageType = "Text"
	MessageTypeReply     MessageType = "Reply"
	MessageTypeGif       MessageType = "GIF"
	MessageTypeImage     MessageType = "Image"
	MessageTypeImages    MessageType = "Images"
	MessageTypeSticker   MessageType = "Sticker"
	MessageTypeFile      MessageType = "File"
	MessageTypeFiles     MessageType = "Files"
	MessageTypeWithdrawn MessageType = "Withdrawn"
)

var AllMessageType = []MessageType{
	MessageTypeText,
	MessageTypeReply,
	MessageTypeGif,
	MessageTypeImage,
	MessageTypeImages,
	MessageTypeSticker,
	MessageTypeFile,
	MessageTypeFiles,
	MessageTypeWithdrawn,
}

func (e MessageType) IsValid() bool {
	switch e {
	case MessageTypeText, MessageTypeReply, MessageTypeGif, MessageTypeImage, MessageTypeImages, MessageTypeSticker, MessageTypeFile, MessageTypeFiles, MessageTypeWithdrawn:
		return true
	}
	return false
}

func (e MessageType) String() string {
	return string(e)
}

func (e *MessageType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MessageType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MessageType", str)
	}
	return nil
}

func (e MessageType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
