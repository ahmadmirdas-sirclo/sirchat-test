package models

type MessageBlockType string

const (
	MBTText      MessageBlockType = "text"
	MBTTextList  MessageBlockType = "text_list"
	MBTTable     MessageBlockType = "table"
	MBTContainer MessageBlockType = "container"
	MBTImage     MessageBlockType = "image"
	MBTInput     MessageBlockType = "input"
	MBTButton    MessageBlockType = "button"
	MBTDivider   MessageBlockType = "divider"
	MBTCarousel  MessageBlockType = "carousel"
)

// IBlock is the interface for blocks. All type of blocks that embed `block`
// struct should satisfy this interface.
type IBlock interface {
	GetType() MessageBlockType
	Validate() (bool, []error)
}

// `block` is the base struct for every other block type. It is meant
// to be embedded to a block subtype. `block` provides the embedding
// structs with field `Type` and the basic methods for a block.
type block struct {
	// Type is the block type of a block.
	Type MessageBlockType `json:"type"`
}

// GetType returns the type of the block. Use this method as the
// alternative for getting the value of field `Type` conventionally,
// such as when handling structs as IBlock.
func (ths *block) GetType() MessageBlockType {
	return ths.Type
}
