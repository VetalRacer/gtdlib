// AUTOGENERATED - DO NOT EDIT

package tdlib

// Emojis Represents a list of emoji
type Emojis struct {
	tdCommon
	Emojis []string `json:"emojis"` // List of emojis
}

// MessageType return the string telegram-type of Emojis
func (emojis *Emojis) MessageType() string {
	return "emojis"
}

// NewEmojis creates a new Emojis
//
// @param emojis List of emojis
func NewEmojis(emojis []string) *Emojis {
	emojisTemp := Emojis{
		tdCommon: tdCommon{Type: "emojis"},
		Emojis:   emojis,
	}

	return &emojisTemp
}
