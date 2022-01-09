// AUTOGENERATED - DO NOT EDIT

package tdlib

// RecommendedChatFilters Contains a list of recommended chat filters
type RecommendedChatFilters struct {
	tdCommon
	ChatFilters []RecommendedChatFilter `json:"chat_filters"` // List of recommended chat filters
}

// MessageType return the string telegram-type of RecommendedChatFilters
func (recommendedChatFilters *RecommendedChatFilters) MessageType() string {
	return "recommendedChatFilters"
}

// NewRecommendedChatFilters creates a new RecommendedChatFilters
//
// @param chatFilters List of recommended chat filters
func NewRecommendedChatFilters(chatFilters []RecommendedChatFilter) *RecommendedChatFilters {
	recommendedChatFiltersTemp := RecommendedChatFilters{
		tdCommon:    tdCommon{Type: "recommendedChatFilters"},
		ChatFilters: chatFilters,
	}

	return &recommendedChatFiltersTemp
}
