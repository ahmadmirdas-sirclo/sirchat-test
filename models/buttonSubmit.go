package models

import "errors"

type SubmitButton struct {
	button
	Label  string  `json:"label,omitempty"`
	Action *Action `json:"action,omitempty"`
}

func (ths *SubmitButton) Validate() (bool, error) {
	// SubmitButton validation implementation
	var errs []error
	if ths.ButtonBlockObject.Type != MBTTSubmit {
		errs = append(errs, errors.New("invalid submit button block object type"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewSubmitButton returns a new instance of a section block to be rendered
func NewSubmitButton(label string) *SubmitButton {
	var button SubmitButton
	button.Label = label
	button.Type = MBTTSubmit

	return &button
}
