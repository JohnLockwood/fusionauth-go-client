package auth

// TwoFactorDelivery enum for User
type TwoFactorDelivery string

// Valid values.  If TextMessage, User.MobilePhone must also be valid
const (
	None        TwoFactorDelivery = "None"
	TextMessage TwoFactorDelivery = "TextMessage"
)
