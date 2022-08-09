package models

import "errors"

type InputBlockObjectType string

const (
	InputBlockObjectTypeText           InputBlockObjectType = "text"
	InputBlockObjectTypeRadio          InputBlockObjectType = "radio"
	InputBlockObjectTypeCounter        InputBlockObjectType = "counter"
	InputBlockObjectTypeNumber         InputBlockObjectType = "number"
	InputBlockObjectTypeSelect         InputBlockObjectType = "select"
	InputBlockObjectTypeDistrictSelect InputBlockObjectType = "district-select"
)

// InputBlock defines a new block of type section
type InputBlock struct {
	block
	Input *InputBlockObject `json:"input,omitempty"`
}

type InputBlockObject struct {
	appendable
	Type        string                    `json:"type"`
	Value       string                    `json:"value"`
	Name        string                    `json:"name"`
	Placeholder string                    `json:"placeholder,omitempty"`
	Options     []InputBlockOptionsObject `json:"options,omitempty"`
	Label       string                    `json:"label,omitempty"`
	Tooltip     string                    `json:"tooltip,omitempty"`
	Required    bool                      `json:"required,omitempty"`
}

type InputBlockOptionsObject struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

func (s InputBlock) Validate() (bool, []error) {
	// InputBlock validation implementation
	var errs []error
	if s.Type != MBTInput {
		errs = append(errs, errors.New("invalid input block type"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewInputBlock returns a new instance of a section block to be rendered
func NewInputBlock(inputObj *InputBlockObject) *InputBlock {
	block := InputBlock{
		Input: inputObj,
	}
	block.Type = MBTInput

	return &block
}
