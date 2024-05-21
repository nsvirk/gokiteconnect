package kiteconnect

import (
	"fmt"
	"net/http"

	"github.com/nsvirk/gokiteconnect/v4/models"
)

// UserSession represents the response after a successful authentication.
type UserSession struct {
	UserProfile
	UserSessionTokens

	UserID      string      `json:"user_id"`
	APIKey      string      `json:"api_key"`
	PublicToken string      `json:"public_token"`
	LoginTime   models.Time `json:"login_time"`
}

// UserSessionTokens represents response after renew access token.
type UserSessionTokens struct {
	UserID       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Bank represents the details of a single bank account entry on a user's file.
type Bank struct {
	Name    string `json:"name"`
	Branch  string `json:"branch"`
	Account string `json:"account"`
}

// UserMeta contains meta data of the user.
type UserMeta struct {
	DematConsent string `json:"demat_consent"`
}

// FullUserMeta contains full meta data of the user.
type FullUserMeta struct {
	DematConsent  string   `json:"poa"`
	Silo          string   `json:"silo"`
	AccountBlocks []string `json:"account_blocks"`
}

// UserProfile represents a user's personal and financial profile.
type UserProfile struct {
	UserID        string   `json:"user_id"`
	UserName      string   `json:"user_name"`
	UserShortName string   `json:"user_shortname"`
	AvatarURL     string   `json:"avatar_url"`
	UserType      string   `json:"user_type"`
	Email         string   `json:"email"`
	Broker        string   `json:"broker"`
	Meta          UserMeta `json:"meta"`
	Products      []string `json:"products"`
	OrderTypes    []string `json:"order_types"`
	Exchanges     []string `json:"exchanges"`
}

type FullUserProfile struct {
	UserID            string       `json:"user_id"`
	UserName          string       `json:"user_name"`
	AvatarURL         string       `json:"avatar_url"`
	UserType          string       `json:"user_type"`
	Email             string       `json:"email"`
	Phone             string       `json:"phone"`
	Broker            string       `json:"broker"`
	TwoFAType         string       `json:"twofa_type"`
	Banks             []Bank       `json:"bank_accounts"`
	DPIDs             []string     `json:"dp_ids"`
	Products          []string     `json:"products"`
	OrderTypes        []string     `json:"order_types"`
	Exchanges         []string     `json:"exchanges"`
	PAN               string       `json:"pan"`
	UserShortName     string       `json:"user_shortname"`
	Tags              []string     `json:"tags"`
	PasswordTimestamp models.Time  `json:"password_timestamp"`
	TwoFATimestamp    models.Time  `json:"twofa_timestamp"`
	Meta              FullUserMeta `json:"meta"`
}

// Margins represents the user margins for a segment.
type Margins struct {
	Category  string           `json:"-"`
	Enabled   bool             `json:"enabled"`
	Net       float64          `json:"net"`
	Available AvailableMargins `json:"available"`
	Used      UsedMargins      `json:"utilised"`
}

// AvailableMargins represents the available margins from the margins response for a single segment.
type AvailableMargins struct {
	AdHocMargin    float64 `json:"adhoc_margin"`
	Cash           float64 `json:"cash"`
	Collateral     float64 `json:"collateral"`
	IntradayPayin  float64 `json:"intraday_payin"`
	LiveBalance    float64 `json:"live_balance"`
	OpeningBalance float64 `json:"opening_balance"`
}

// UsedMargins represents the used margins from the margins response for a single segment.
type UsedMargins struct {
	Debits           float64 `json:"debits"`
	Exposure         float64 `json:"exposure"`
	M2MRealised      float64 `json:"m2m_realised"`
	M2MUnrealised    float64 `json:"m2m_unrealised"`
	OptionPremium    float64 `json:"option_premium"`
	Payout           float64 `json:"payout"`
	Span             float64 `json:"span"`
	HoldingSales     float64 `json:"holding_sales"`
	Turnover         float64 `json:"turnover"`
	LiquidCollateral float64 `json:"liquid_collateral"`
	StockCollateral  float64 `json:"stock_collateral"`
	Delivery         float64 `json:"delivery"`
}

// AllMargins contains both equity and commodity margins.
type AllMargins struct {
	Equity    Margins `json:"equity"`
	Commodity Margins `json:"commodity"`
}

// GetUserProfile gets user profile.
func (c *Client) GetUserProfile() (UserProfile, error) {
	var userProfile UserProfile
	err := c.doEnvelope(http.MethodGet, URIUserProfile, nil, nil, &userProfile)
	return userProfile, err
}

// GetFullUserProfile gets full user profile.
func (c *Client) GetFullUserProfile() (FullUserProfile, error) {
	var fUserProfile FullUserProfile
	err := c.doEnvelope(http.MethodGet, URIFullUserProfile, nil, nil, &fUserProfile)
	return fUserProfile, err
}

// GetUserMargins gets all user margins.
func (c *Client) GetUserMargins() (AllMargins, error) {
	var allUserMargins AllMargins
	err := c.doEnvelope(http.MethodGet, URIUserMargins, nil, nil, &allUserMargins)
	return allUserMargins, err
}

// GetUserSegmentMargins gets segmentwise user margins.
func (c *Client) GetUserSegmentMargins(segment string) (Margins, error) {
	var margins Margins
	err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIUserMarginsSegment, segment), nil, nil, &margins)
	return margins, err
}
