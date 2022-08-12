package models

import "errors"

// TextBlockObjectAlign is a type to define align text
type TextBlockObjectAlign string

// TextBlockObjectType is a type to define type text
type TextBlockObjectType string

// TextBlockObjectColor is a type to define color text
type TextBlockObjectColor string

const (

	// center align is center position text
	TextBlockObjectAlignCenter TextBlockObjectAlign = "center"

	// left align is left position text
	TextBlockObjectAlignLeft TextBlockObjectAlign = "left"

	// right align is right position text
	TextBlockObjectAlignRight TextBlockObjectAlign = "right"

	// span type is 12px medium
	TextBlockObjectTypeSpan TextBlockObjectType = "span"

	// label type is 12px medium (display: block)
	TextBlockObjectTypeLabel TextBlockObjectType = "label"

	// paragraph type is 14px medium
	TextBlockObjectTypeParagraph TextBlockObjectType = "paragraph"

	// heading type is 16px bold
	TextBlockObjectTypeHeading TextBlockObjectType = "heading"

	// subheading type is 14px bold
	TextBlockObjectTypeSubheading TextBlockObjectType = "subheading"

	// subheading2 type is 12px bold
	TextBlockObjectTypeSubheading2 TextBlockObjectType = "subheading2"

	// figure type is 16px bold (color: primary (#269CD9))
	TextBlockObjectTypeFigure TextBlockObjectType = "figure"

	// text color is text (#3D4F5C)
	TextBlockObjectColorText TextBlockObjectColor = "text"

	// primary color is primary (#269CD9)
	TextBlockObjectColorPrimary TextBlockObjectColor = "primary"

	// secondary color is secondary (#7590A3)
	TextBlockObjectColorSecondary TextBlockObjectColor = "secondary"

	// danger color is red (#E05D52)
	TextBlockObjectColorDanger TextBlockObjectColor = "danger"

	// disabled color is silver or text disabled (#D0D9E0)
	TextBlockObjectColorDisabled TextBlockObjectColor = "disabled"
)

// TextBlock is a subtype of block. It represents a text block.
type TextBlock struct {
	block
	Text *TextBlockObject `json:"text"`
}

// TextBlockObject holds the detail of the TextBlock.
type TextBlockObject struct {
	appendable
	// body is content of text block
	// this field is required
	Body string `json:"body"`

	// align is positioning of the text
	// align has the default value left
	Align TextBlockObjectAlign `json:"align"`

	// type is the size of the text (px)
	// type has the default value is not field type (*default: 16px medium)
	Type TextBlockObjectType `json:"type,omitempty"`

	// color is the color of the text
	// color has the default value text
	Color TextBlockObjectColor `json:"color"`
}

// Validate performs validation to the TextBlock. Field `Body`
// should not be empty.
func (s *TextBlock) Validate() (bool, []error) {
	// TextBlock validation implementation
	var errs []error
	if s.Text.Body == "" {
		errs = append(errs, errors.New("body is missing"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewTextBlock returns a new instance of a section block to be rendered
func NewTextBlock(textObj *TextBlockObject) *TextBlock {
	var (
		alignText TextBlockObjectAlign
		typeText  TextBlockObjectType
		colorText TextBlockObjectColor
	)

	alignText = TextBlockObjectAlignLeft
	colorText = TextBlockObjectColorText
	typeText = ""

	if string(textObj.Align) != "" {
		alignText = textObj.Align
	}

	if string(textObj.Type) != "" {
		typeText = textObj.Type
	}

	if string(textObj.Color) != "" {
		colorText = textObj.Color
	}

	block := TextBlock{
		Text: &TextBlockObject{
			Body:  textObj.Body,
			Align: alignText,
			Type:  typeText,
			Color: colorText,
		},
	}
	block.Type = MBTText

	return &block
}
