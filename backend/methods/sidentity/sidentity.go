package sidentity

import (
	"context"
	"errors"
	"hestia/jobfair/api/pb/common"
	"hestia/jobfair/api/pb/identity"
	"hestia/jobfair/api/utils/db"
	"hestia/jobfair/api/utils/herror"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IdentityServer struct {
	identity.UnimplementedIdentityManagementServer
}

func (IdentityServer) Alive(ctx context.Context, token *common.Token) (*emptypb.Empty, error) {
	id, err := uuid.Parse(token.GetToken())
	if err != nil {
		return nil, herror.StatusWithInfo(codes.InvalidArgument, "Invalid token", herror.AuthInvalidTokenError, identity.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Err(err).Msg("Error getting db connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Error getting db connection", herror.DatabaseError, identity.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	var expiry_date time.Time
	err = db.QueryRow(ctx, "SELECT expiry_date FROM users.user_session WHERE id = $1", id).Scan(&expiry_date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, herror.StatusWithInfo(codes.Unauthenticated, "User session not found", herror.AuthInvalidTokenError, identity.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
		}
		log.Err(err).Msg("Error checking user session")
		return nil, herror.StatusWithInfo(codes.Internal, "Error checking user session", herror.DatabaseError, identity.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	if expiry_date.Before(time.Now()) {
		return nil, herror.StatusWithInfo(codes.Unauthenticated, "User session expired", herror.AuthInvalidTokenError, identity.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	return &emptypb.Empty{}, nil
}

func (IdentityServer) Login(ctx context.Context, loginRequest *identity.LoginRequest) (*common.Token, error) {
	email := loginRequest.GetEmail()
	password := loginRequest.GetPassword()
	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Err(err).Msg("Error getting db connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Error getting db connection", herror.DatabaseError, identity.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		log.Err(err).Msg("Error starting transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Error starting transaction", herror.DatabaseTxError, identity.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}
	defer tx.Rollback(ctx)

	var id uuid.UUID
	var hashedPassword string
	err = tx.QueryRow(ctx, "SELECT id, password FROM users.users WHERE email = $1", email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, herror.StatusWithInfo(codes.Unauthenticated, "Wrong email or password", herror.AuthWrongCredentialsError, identity.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
		}
		log.Err(err).Msg("Error getting user password")
		return nil, herror.StatusWithInfo(codes.Internal, "Error getting user password", herror.DatabaseError, identity.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return nil, herror.StatusWithInfo(codes.Unauthenticated, "Wrong email or password", herror.AuthWrongCredentialsError, identity.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	sessionId := uuid.New()
	_, err = tx.Exec(ctx, "INSERT INTO users.user_session (id, user_id) VALUES ($1, $2)", sessionId, id)
	if err != nil {
		log.Err(err).Msg("Error creating user session")
		return nil, herror.StatusWithInfo(codes.Internal, "Error creating user session", herror.DatabaseError, identity.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Err(err).Msg("Error committing transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Error committing transaction", herror.DatabaseTxError, identity.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	return &common.Token{Token: sessionId.String()}, nil
}

func (IdentityServer) Logout(ctx context.Context, token *common.Token) (*emptypb.Empty, error) {
	return nil, nil
}

func (IdentityServer) Register(ctx context.Context, registerRequest *identity.RegisterRequest) (*emptypb.Empty, error) {
	return nil, nil
}
