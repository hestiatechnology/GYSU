package auth

import (
	"context"
	"errors"
	"hestia/jobfair/api/utils/db"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("TokenExpired")
	ErrInvalidToken = errors.New("TokenInvalid")
	ErrNoCompanyId  = errors.New("NoCompanyId")
	ErrInvalidUUID  = errors.New("InvalidUUID")
	ErrNoMetadata   = errors.New("NoMetadata")
)

// returns user id or the error
func VerifyAuthToken(ctx context.Context, token uuid.UUID) (uuid.UUID, error) {
	db, err := db.GetDBPoolConn()
	if err != nil {
		return uuid.Nil, err
	}

	var userId uuid.UUID
	var expiry_date time.Time
	err = db.QueryRow(ctx, "SELECT user_id, expiry_date FROM users.user_session WHERE id = $1", token).Scan(&userId, &expiry_date)
	if err != nil {
		return uuid.Nil, ErrInvalidToken
	}

	if expiry_date.Before(time.Now()) {
		return uuid.Nil, ErrExpiredToken
	}

	return userId, nil
}
