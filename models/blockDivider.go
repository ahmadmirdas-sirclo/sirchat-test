package models

import "errors"

type DividerBlock struct {
	block
}

func (s DividerBlock) Validate() (bool, []error) {
	// DividerBlock validation implementation
	var errs []error
	if s.Type != MBTDivider {
		errs = append(errs, errors.New("invalid divider block type"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewButtonBlock returns a new instance of a section block to be rendered
func NewDividerBlock() *DividerBlock {
	block := DividerBlock{}
	block.Type = MBTDivider

	return &block
}
