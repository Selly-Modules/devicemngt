package devicemngt

import (
	"context"
	"net/http"

	"github.com/Selly-Modules/logger"
)

// findByDeviceID ...
func (s Service) findByDeviceID(ctx context.Context, id string) (result Device) {
	stm, args, _ := s.Builder.Select("*").From(TableDeviceMngt).Where("device_id = ?", id).ToSql()

	if err := s.DB.GetContext(ctx, &result, stm, args...); err != nil {
		logger.Error("devicemngt - findByDeviceID", logger.LogData{
			"device_id": id,
			"error":     err.Error(),
		})
	}

	return
}

// getHeaderData ...
func getHeaderData(headers http.Header) HeaderData {
	return HeaderData{
		UserAgent:      headers.Get("User-Agent"),
		DeviceID:       headers.Get("Deviceid"),
		AppVersion:     headers.Get("App-Version"),
		AppVersionCode: headers.Get("App-Version-Code"),
		OSName:         headers.Get("Os-Name"),
		OSVersion:      headers.Get("Os-Version"),
		AuthToken:      headers.Get("Authorization"),
	}
}
