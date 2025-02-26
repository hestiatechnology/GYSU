package auth

import (
	"context"
	"errors"
	"hestia/jobfair/api/utils/db"
	"time"

	"github.com/google/uuid"
)

type CtxKey string

const (
	UserIdKey   CtxKey = "user_id"
	RoleKey     CtxKey = "role"
	StudentRole        = "student"
	CompanyRole        = "company"
	StaffRole          = "staff"
	AdminRole          = "admin"
)

var (
	ErrExpiredToken = errors.New("TokenExpired")
	ErrInvalidToken = errors.New("TokenInvalid")
	ErrNoCompanyId  = errors.New("NoCompanyId")
	ErrInvalidUUID  = errors.New("InvalidUUID")
	ErrNoMetadata   = errors.New("NoMetadata")
)

// returns user id or the error
func VerifyAuthToken(ctx context.Context, token uuid.UUID) (uuid.UUID, string, error) {
	db, err := db.GetDBPoolConn()
	if err != nil {
		return uuid.Nil, "", err
	}

	var userId uuid.UUID
	var role string
	var expiry_date time.Time
	err = db.QueryRow(ctx, `
    SELECT 
		us.user_id, 
		us.expiry_date, 
		u.role 
    FROM 
		users.user_session us
    JOIN 
		users.users u 
			ON us.user_id = u.id
    WHERE 
		us.id = $1
	`, token).Scan(&userId, &expiry_date, &role)
	if err != nil {
		return uuid.Nil, "", ErrInvalidToken
	}

	if expiry_date.Before(time.Now()) {
		return uuid.Nil, "", ErrExpiredToken
	}

	return userId, role, nil
}

func GetUserId(ctx context.Context) uuid.UUID {
	v := ctx.Value(UserIdKey)
	if v == nil {
		return uuid.Nil
	}
	return v.(uuid.UUID)
}

func GetRole(ctx context.Context) string {
	v := ctx.Value(RoleKey)
	if v == nil {
		return ""
	}
	return v.(string)
}
