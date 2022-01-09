// AUTOGENERATED - DO NOT EDIT

package tdlib

// WebPageInstantView Describes an instant view page for a web page
type WebPageInstantView struct {
	tdCommon
	PageBlocks []PageBlock `json:"page_blocks"` // Content of the web page
	ViewCount  int32       `json:"view_count"`  // Number of the instant view views; 0 if unknown
	Version    int32       `json:"version"`     // Version of the instant view, currently can be 1 or 2
	IsRtl      bool        `json:"is_rtl"`      // True, if the instant view must be shown from right to left
	IsFull     bool        `json:"is_full"`     // True, if the instant view contains the full page. A network request might be needed to get the full web page instant view
}

// MessageType return the string telegram-type of WebPageInstantView
func (webPageInstantView *WebPageInstantView) MessageType() string {
	return "webPageInstantView"
}

// NewWebPageInstantView creates a new WebPageInstantView
//
// @param pageBlocks Content of the web page
// @param viewCount Number of the instant view views; 0 if unknown
// @param version Version of the instant view, currently can be 1 or 2
// @param isRtl True, if the instant view must be shown from right to left
// @param isFull True, if the instant view contains the full page. A network request might be needed to get the full web page instant view
func NewWebPageInstantView(pageBlocks []PageBlock, viewCount int32, version int32, isRtl bool, isFull bool) *WebPageInstantView {
	webPageInstantViewTemp := WebPageInstantView{
		tdCommon:   tdCommon{Type: "webPageInstantView"},
		PageBlocks: pageBlocks,
		ViewCount:  viewCount,
		Version:    version,
		IsRtl:      isRtl,
		IsFull:     isFull,
	}

	return &webPageInstantViewTemp
}