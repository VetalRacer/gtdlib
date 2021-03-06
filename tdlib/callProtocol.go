// AUTOGENERATED - DO NOT EDIT

package tdlib

// CallProtocol Specifies the supported call protocols
type CallProtocol struct {
	tdCommon
	UDPP2p          bool     `json:"udp_p2p"`          // True, if UDP peer-to-peer connections are supported
	UDPReflector    bool     `json:"udp_reflector"`    // True, if connection through UDP reflectors is supported
	MinLayer        int32    `json:"min_layer"`        // The minimum supported API layer; use 65
	MaxLayer        int32    `json:"max_layer"`        // The maximum supported API layer; use 65
	LibraryVersions []string `json:"library_versions"` // List of supported libtgvoip versions
}

// MessageType return the string telegram-type of CallProtocol
func (callProtocol *CallProtocol) MessageType() string {
	return "callProtocol"
}

// NewCallProtocol creates a new CallProtocol
//
// @param uDPP2p True, if UDP peer-to-peer connections are supported
// @param uDPReflector True, if connection through UDP reflectors is supported
// @param minLayer The minimum supported API layer; use 65
// @param maxLayer The maximum supported API layer; use 65
// @param libraryVersions List of supported libtgvoip versions
func NewCallProtocol(uDPP2p bool, uDPReflector bool, minLayer int32, maxLayer int32, libraryVersions []string) *CallProtocol {
	callProtocolTemp := CallProtocol{
		tdCommon:        tdCommon{Type: "callProtocol"},
		UDPP2p:          uDPP2p,
		UDPReflector:    uDPReflector,
		MinLayer:        minLayer,
		MaxLayer:        maxLayer,
		LibraryVersions: libraryVersions,
	}

	return &callProtocolTemp
}
