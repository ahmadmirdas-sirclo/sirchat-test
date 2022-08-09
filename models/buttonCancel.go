package models

import "errors"

type CancelButton struct {
	button
}

func (ths *CancelButton) Validate() (bool, []error) {
	// CancelButton validation implementation
	var errs []error
	if ths.ButtonBlockObject.Type != MBTTCancel {
		errs = append(errs, errors.New("invalid cancel button block object type"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewCancelButton returns a new instance of a section block to be rendered
func NewCancelButton(label string) *CancelButton {
	var button CancelButton
	button.Label = label
	button.Type = MBTTCancel

	return &button
}
