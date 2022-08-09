package models

import "errors"

type ActionButton struct {
	button
	Action *Action `json:"action"`
}

func (ths *ActionButton) Validate() (bool, []error) {
	// ActionButton validation implementation
	var errs []error
	if ths.ButtonBlockObject.Type != MBTTAction {
		errs = append(errs, errors.New("invalid action button block object type"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewActionButton returns a new instance of a section block to be rendered
func NewActionButton(label, actionID string, query interface{}) *ActionButton {
	var button ActionButton
	button.Label = label
	button.Action = &Action{
		ID: actionID,
	}
	button.Type = MBTTAction
	button.Query = query

	return &button
}
