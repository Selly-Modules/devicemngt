package devicemngt

import (
	"context"

	"github.com/Selly-Modules/logger"
)

// findByDeviceID ...
func (s Service) findByDeviceID(ctx context.Context, id string) (result Device) {
	stm, args, _ := s.Builder.Select("*").From(TableDeviceMngt).Where("device_id = ?", id).ToSql()

	if err := s.DB.SelectContext(ctx, &result, stm, args...); err != nil {
		logger.Error("devicemngt - findByDeviceID", logger.LogData{
			"device_id": id,
			"error":     err.Error(),
		})
	}

	return
}
