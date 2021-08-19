package devicemngt

import (
	"context"
	"net/http"
	"time"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	"github.com/kr/pretty"
	ua "github.com/mssola/user_agent"
)

// UpsertPayload ...
type UpsertPayload struct {
	IP            string
	Headers       http.Header
	AuthToken     string
	FCMToken      string
	OwnerID       string
	OwnerType     string
	FirstSignInAt time.Time
}

// Upsert ...
func (s Service) Upsert(payload UpsertPayload) {
	ctx := context.Background()

	// Read UA
	var (
		headerData = getHeaderData(payload.Headers)
		uaData     = ua.New(headerData.UserAgent)
	)

	// DB data
	var (
		deviceID       = ""
		platform       = ""
		osName         = ""
		osVersion      = ""
		appVersion     = ""
		appVersionCode = ""
		browserName    = ""
		browserVersion = ""
	)

	// Set deviceID
	deviceID = headerData.DeviceID
	if deviceID == "" {
		logger.Error("devicemngt - Upsert: no device_id data", logger.LogData{
			"payload": payload,
		})
		return
	}

	// OS, if there is os name, means mobile app, else browser
	if headerData.OSName != "" {
		platform = headerData.OSName
		osName = headerData.OSName
		osVersion = headerData.OSVersion
	} else {
		platform = uaData.Platform()
		osName = uaData.OSInfo().Name
		osVersion = uaData.OSInfo().Version
		browserName, browserVersion = uaData.Browser()
	}

	// App version
	if headerData.AppVersion != "" {
		appVersion = headerData.AppVersion
		appVersionCode = headerData.AppVersionCode
	}

	pretty.Println("- platform", platform)
	pretty.Println("- osName", osName)
	pretty.Println("- osVersion", osVersion)
	pretty.Println("- appVersion", appVersion)
	pretty.Println("- appVersionCode", appVersionCode)
	pretty.Println("- browserName", browserName)
	pretty.Println("- browserVersion", browserVersion)
	pretty.Println("----------------")

	// Find device id existed or not
	device := s.findByDeviceID(ctx, deviceID)

	if !mongodb.IsValidID(device.ID) {
		// If not exist, create new
		stm, args, _ := s.Builder.Insert(TableDeviceMngt).
			Columns(
				"id", "device_id", "ip", "platform", "app_version",
				"app_version_code", "os_name", "os_version", "browser_name", "browser_version",
				"auth_token", "fcm_token", "owner_id", "owner_type",
				"first_sign_in_at", "last_activity_at",
			).Values(
			mongodb.NewStringID(), deviceID, payload.IP, platform, appVersion,
			appVersionCode, osName, osVersion, browserName, browserVersion,
			payload.AuthToken, payload.FCMToken, payload.OwnerID, payload.OwnerType,
			payload.FirstSignInAt, now(),
		).ToSql()

		pretty.Println("Create new")
		pretty.Println("stm -", stm)
		pretty.Println("args -", args)

		if _, err := s.DB.ExecContext(ctx, stm, args); err != nil {
			logger.Error("devicemngt - Upsert: Create new", logger.LogData{
				"payload": payload,
				"error":   err.Error(),
			})
		}
	} else {
		// Else update
		stm, args, _ := s.Builder.Update(TableDeviceMngt).
			Set("ip", payload.IP).
			Set("platform", platform).
			Set("app_version", appVersion).
			Set("app_version_code", appVersionCode).
			Set("os_name", osName).
			Set("os_version", osVersion).
			Set("browser_name", browserName).
			Set("browser_version", browserVersion).
			Set("auth_token", payload.AuthToken).
			Set("fcm_token", payload.FCMToken).
			Set("owner_id", payload.OwnerID).
			Set("owner_type", payload.OwnerType).
			Set("last_activity_at", now()).
			Where("device_id = ?", deviceID).
			ToSql()

		pretty.Println("Update")
		pretty.Println("stm -", stm)
		pretty.Println("args -", args)

		if _, err := s.DB.ExecContext(ctx, stm, args); err != nil {
			logger.Error("devicemngt - Upsert: Update", logger.LogData{
				"payload": payload,
				"error":   err.Error(),
			})
		}
	}
}
