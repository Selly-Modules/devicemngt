package devicemngt

import (
	"context"
	"fmt"

	"github.com/Selly-Modules/logger"
)

// GetDevicePermissionsByToken ...
func (s Service) GetDevicePermissionsByToken(token string) (doc StaffDevicePermissions) {
	ctx := context.Background()

	stm, args, _ := s.Builder.
		Select("dm.id AS device_id, s.id, s.name, s.account_type, sr.permissions").
		From(fmt.Sprintf("%s AS dm", TableDeviceMngt)).
		Join(fmt.Sprintf("LEFT JOIN %s s ON s.id = dm.owner_id", TableStaff)).
		Join(fmt.Sprintf("LEFT JOIN %s sr ON s.role_id = sr.id", TableStaffRole)).
		Where("dm.auth_token = ?", token).
		ToSql()
	if err := s.DB.GetContext(ctx, &doc, stm, args...); err != nil {
		logger.Error("devicemngt - GetDevicePermissionsByToken", logger.LogData{
			"token": token,
			"error": err.Error(),
		})
	}
	return
}
