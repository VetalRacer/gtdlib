// AUTOGENERATED - DO NOT EDIT

package tdlib

// TextEntities Contains a list of text entities
type TextEntities struct {
	tdCommon
	Entities []TextEntity `json:"entities"` // List of text entities
}

// MessageType return the string telegram-type of TextEntities
func (textEntities *TextEntities) MessageType() string {
	return "textEntities"
}

// NewTextEntities creates a new TextEntities
//
// @param entities List of text entities
func NewTextEntities(entities []TextEntity) *TextEntities {
	textEntitiesTemp := TextEntities{
		tdCommon: tdCommon{Type: "textEntities"},
		Entities: entities,
	}

	return &textEntitiesTemp
}
