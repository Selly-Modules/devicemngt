package devicemngt

import (
	"context"

	"github.com/Selly-Modules/logger"
)

// AllQuery ...
type AllQuery struct {
	OwnerID string
}

// FindAllDevicesByOwnerID ...
func (s Service) FindAllDevicesByOwnerID(ownerID string) []ResponseDevice {
	ctx := context.Background()

	var (
		docs   = make([]Device, 0)
		result = make([]ResponseDevice, 0)
	)

	stm, args, _ := s.Builder.Select("*").
		From(TableDeviceMngt).
		Where("owner_id = ?", ownerID).
		ToSql()
	if err := s.DB.SelectContext(ctx, &docs, stm, args...); err != nil {
		logger.Error("devicemngt - FindAllDevicesByOwnerID", logger.LogData{
			"ownerID": ownerID,
			"error":   err.Error(),
		})
		return result
	}

	// Get data
	for _, doc := range docs {
		result = append(result, ResponseDevice{
			ID:       doc.ID,
			IP:       doc.IP,
			Platform: doc.Platform,
			OS: ResponseOS{
				Name:    doc.OSName,
				Version: doc.OSVersion,
			},
			Browser: ResponseBrowser{
				Name:    doc.BrowserName,
				Version: doc.BrowserVersion,
			},
			FCMToken:      doc.FCMToken,
			FirstSignInAt: doc.FirstSignInAt.Format(dateLayoutFull),
		})
	}
	return result
}
