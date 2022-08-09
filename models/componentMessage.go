package models

import (
	"errors"
)

type MessageComponent struct {
	component
	Message MessageObject `json:"message"`
}

type MessageObject struct {
	TenantID string              `json:"tenant_id"`
	BrandID  string              `json:"brand_id"`
	RoomID   string              `json:"room_id"`
	Channel  string              `json:"channel"`
	Texts    []MessageTextObject `json:"texts"`
}

type MessageTextObject struct {
	Body string `json:"body"`
}

func (ths *MessageComponent) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MCTMessage {
		errs = append(errs, errors.New("invalid message component type"))
	}

	if ths.Message.TenantID == "" {
		errs = append(errs, errors.New("tenant ID in message component can't be empty"))
	}

	if ths.Message.BrandID == "" {
		errs = append(errs, errors.New("brand ID in message component can't be empty"))
	}

	if ths.Message.RoomID == "" {
		errs = append(errs, errors.New("room ID in message component can't be empty"))
	}

	if ths.Message.Channel == "" {
		errs = append(errs, errors.New("channel in message component can't be empty"))
	}

	if len(ths.Message.Texts) == 0 {
		errs = append(errs, errors.New("texts in message component can't be empty"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

func NewMessage() *MessageComponent {
	var c MessageComponent
	c.Type = MCTMessage
	c.component.IComponent = &c
	return &c
}
