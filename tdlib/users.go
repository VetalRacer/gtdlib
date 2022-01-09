// AUTOGENERATED - DO NOT EDIT

package tdlib

// Users Represents a list of users
type Users struct {
	tdCommon
	TotalCount int32   `json:"total_count"` // Approximate total count of users found
	UserIDs    []int32 `json:"user_ids"`    // A list of user identifiers
}

// MessageType return the string telegram-type of Users
func (users *Users) MessageType() string {
	return "users"
}

// NewUsers creates a new Users
//
// @param totalCount Approximate total count of users found
// @param userIDs A list of user identifiers
func NewUsers(totalCount int32, userIDs []int32) *Users {
	usersTemp := Users{
		tdCommon:   tdCommon{Type: "users"},
		TotalCount: totalCount,
		UserIDs:    userIDs,
	}

	return &usersTemp
}