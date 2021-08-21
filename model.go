package devicemngt

import (
	"time"

	"github.com/lib/pq"
)

// Device ...
type Device struct {
	ID             string    `db:"id"`
	DeviceID       string    `db:"device_id"`
	IP             string    `db:"ip"`
	Platform       string    `db:"platform"`
	OSName         string    `db:"os_name"`
	OSVersion      string    `db:"os_version"`
	AppVersion     string    `db:"app_version"`
	AppVersionCode string    `db:"app_version_code"`
	BrowserName    string    `db:"browser_name"`
	BrowserVersion string    `db:"browser_version"`
	AuthToken      string    `db:"auth_token"`
	FCMToken       string    `db:"fcm_token"`
	OwnerID        string    `db:"owner_id"`
	OwnerType      string    `db:"owner_type"`
	FirstSignInAt  time.Time `db:"first_sign_in_at"`
	LastActivityAt time.Time `db:"last_activity_at"`
}

// ResponseDevice ...
type ResponseDevice struct {
	ID            string          `json:"id"`
	IP            string          `json:"ip"`
	Platform      string          `json:"platform"`
	OS            ResponseOS      `json:"os"`
	Browser       ResponseBrowser `json:"browser"`
	FCMToken      string          `json:"fcmToken"`
	FirstSignInAt string          `json:"firstSignInAt"`
}

// ResponseOS ...
type ResponseOS struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ResponseBrowser ...
type ResponseBrowser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// HeaderData ...
type HeaderData struct {
	UserAgent      string
	DeviceID       string
	AppVersion     string
	AppVersionCode string
	OSName         string
	OSVersion      string
	AuthToken      string
}

// StaffPermissions ...
type StaffPermissions struct {
	ID          string         `db:"id"`
	Name        string         `db:"name"`
	Permissions pq.StringArray `db:"permissions"`
	AccountType string         `db:"account_type"`
	Active      bool           `db:"active"`
	DeviceID    string         `db:"device_id"`
}
