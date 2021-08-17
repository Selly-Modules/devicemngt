package devicemngt

import (
	"context"
	"time"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	ua "github.com/mssola/user_agent"
)

// UpsertPayload ...
type UpsertPayload struct {
	DeviceID      string
	IP            string
	UserAgent     string
	AuthToken     string
	FCMToken      string
	OwnerID       string
	OwnerType     string
	FirstSignInAt time.Time
}

// Upsert ...
func (s Service) Upsert(payload UpsertPayload) {
	ctx := context.Background()

	// Find device id existed or not
	device := s.findByDeviceID(ctx, payload.DeviceID)

	// Read UA
	var (
		uaData                      = ua.New(payload.UserAgent)
		platform                    = uaData.Platform()
		osInfo                      = uaData.OSInfo()
		browserName, browserVersion = uaData.Browser()
	)

	if !mongodb.IsValidID(device.ID) {
		// If not exist, create new
		stm, args, _ := s.Builder.Insert(TableDeviceMngt).
			Columns(
				"id", "device_id", "ip", "platform",
				"os_name", "os_version", "browser_name", "browser_version",
				"auth_token", "fcm_token", "owner_id", "owner_type",
				"first_sign_in_at", "last_activity_at",
			).Values(
			mongodb.NewStringID(), payload.DeviceID, payload.IP, platform,
			osInfo.Name, osInfo.Version, browserName, browserVersion,
			payload.AuthToken, payload.FCMToken, payload.OwnerID, payload.OwnerType,
			payload.FirstSignInAt, now(),
		).ToSql()

		if _, err := s.DB.ExecContext(ctx, stm, args); err != nil {
			logger.Error("devicemngt - Upsert - Create new", logger.LogData{
				"payload": payload,
				"error":   err.Error(),
			})
		}
	} else {
		// Else update
		stm, args, _ := s.Builder.Update(TableDeviceMngt).
			Set("ip", payload.IP).
			Set("platform", platform).
			Set("os_name", osInfo.Name).
			Set("os_version", osInfo.Version).
			Set("browser_name", browserName).
			Set("browser_version", browserVersion).
			Set("auth_token", payload.AuthToken).
			Set("fcm_token", payload.FCMToken).
			Set("owner_id", payload.OwnerID).
			Set("owner_type", payload.OwnerType).
			Set("last_activity_at", now()).
			Where("device_id = ?", payload.DeviceID).
			ToSql()

		if _, err := s.DB.ExecContext(ctx, stm, args); err != nil {
			logger.Error("devicemngt - Upsert - Update", logger.LogData{
				"payload": payload,
				"error":   err.Error(),
			})
		}
	}
}
