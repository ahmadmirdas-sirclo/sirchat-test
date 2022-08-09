package models

import (
	"errors"
)

type DrawerComponent struct {
	component
	appendable
	Title      Title      `json:"title"`
	Action     Action     `json:"action"`
	Subheading Subheading `json:"subheading,omitempty"`
}

func (ths *DrawerComponent) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MCTDrawer {
		errs = append(errs, errors.New("invalid drawer component type"))
	}

	for _, v := range ths.Blocks {
		if valid, err := v.Validate(); !valid {
			errs = append(errs, err...)
		}
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

func NewDrawer() *DrawerComponent {
	var c DrawerComponent
	c.Type = MCTDrawer
	c.component.IComponent = &c
	return &c
}
