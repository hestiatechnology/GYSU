package user

import (
	"context"
	"hestia/api/utils/db"

	"github.com/google/uuid"
)

func IsEmployeeIdUsed(ctx context.Context, employeeId uuid.UUID) bool {
	db, err := db.GetDBPoolConn()
	if err != nil {
		return false
	}

	var count int
	err = db.QueryRow(ctx, "SELECT COUNT(employee_id) FROM users.user_company WHERE employee_id = $1", employeeId).Scan(&count)
	if err != nil {
		return false
	}

	if count > 0 {
		return true
	}

	return false
}
