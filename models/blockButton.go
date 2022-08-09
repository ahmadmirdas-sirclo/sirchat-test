package models

import "errors"

// ButtonBlock defines a new block of type section
type ButtonBlock struct {
	block
	Button *ButtonBlockObject `json:"button,omitempty"`
}

func (ths *ButtonBlock) Validate() (bool, []error) {
	// ButtonBlock validation implementation
	var errs []error
	if ths.Type != MBTButton {
		errs = append(errs, errors.New("invalid button block type"))
	}

	if len(ths.Button.Action.Buttons) == 0 {
		errs = append(errs, errors.New("there are no action buttons in the component"))
	}

	var submitCount, cancelCount int
	for i := 1; i < len(ths.Button.Action.Buttons); i++ {
		v := ths.Button.Action.Buttons[i]

		switch v.GetType() {
		case MBTTSubmit:
			obj := v.(*SubmitButton)
			if obj.ButtonBlockObject.Query == nil {
				errs = append(errs, errors.New("submit button must have query object"))
			}
			submitCount++
		case MBTTCancel:
			cancelCount++
		}

		if submitCount > 1 || cancelCount > 1 {
			errs = append(errs, errors.New("there should be only one submit button and one cancel button in the action buttons"))
		}

		if i == len(ths.Button.Action.Buttons)-1 && (submitCount == 0 || cancelCount == 0) {
			errs = append(errs, errors.New("there should be one submit button and one cancel button in the action buttons"))
		}
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewButtonBlock returns a new instance of a section block to be rendered
func NewButtonBlock(buttonObj *ButtonBlockObject) *ButtonBlock {
	block := ButtonBlock{
		Button: buttonObj,
	}
	block.Type = MBTButton

	return &block
}
