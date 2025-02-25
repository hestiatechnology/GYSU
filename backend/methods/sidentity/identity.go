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
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IdentityServer struct {
	identity.UnimplementedIdentityManagementServiceServer
}

func (IdentityServer) Alive(ctx context.Context, token *common.Token) (*emptypb.Empty, error) {
	id, err := uuid.Parse(token.GetToken())
	if err != nil {
		return nil, herror.StatusWithInfo(codes.InvalidArgument, "Invalid token", herror.AuthInvalidTokenError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Err(err).Msg("Error getting db connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Error getting db connection", herror.DatabaseError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	var expiry_date time.Time
	err = db.QueryRow(ctx, "SELECT expiry_date FROM users.user_session WHERE id = $1", id).Scan(&expiry_date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, herror.StatusWithInfo(codes.Unauthenticated, "User session not found", herror.AuthInvalidTokenError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
		}
		log.Err(err).Msg("Error checking user session")
		return nil, herror.StatusWithInfo(codes.Internal, "Error checking user session", herror.DatabaseError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	if expiry_date.Before(time.Now()) {
		return nil, herror.StatusWithInfo(codes.Unauthenticated, "User session expired", herror.AuthInvalidTokenError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	return &emptypb.Empty{}, nil
}

func (IdentityServer) Login(ctx context.Context, loginRequest *identity.LoginRequest) (*common.Token, error) {
	email := loginRequest.GetEmail()
	password := loginRequest.GetPassword()
	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Err(err).Msg("Error getting db connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Error getting db connection", herror.DatabaseError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		log.Err(err).Msg("Error starting transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Error starting transaction", herror.DatabaseTxError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}
	defer tx.Rollback(ctx)

	var id uuid.UUID
	var hashedPassword string
	err = tx.QueryRow(ctx, "SELECT id, password FROM users.users WHERE email = $1", email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, herror.StatusWithInfo(codes.Unauthenticated, "Wrong email or password", herror.AuthWrongCredentialsError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
		}
		log.Err(err).Msg("Error getting user password")
		return nil, herror.StatusWithInfo(codes.Internal, "Error getting user password", herror.DatabaseError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return nil, herror.StatusWithInfo(codes.Unauthenticated, "Wrong email or password", herror.AuthWrongCredentialsError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	sessionId := uuid.New()
	_, err = tx.Exec(ctx, "INSERT INTO users.user_session (id, user_id) VALUES ($1, $2)", sessionId, id)
	if err != nil {
		log.Err(err).Msg("Error creating user session")
		return nil, herror.StatusWithInfo(codes.Internal, "Error creating user session", herror.DatabaseError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Err(err).Msg("Error committing transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Error committing transaction", herror.DatabaseTxError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	token := &common.Token{}
	token.SetToken(sessionId.String())
	return token, nil
}

func (IdentityServer) Logout(ctx context.Context, token *common.Token) (*emptypb.Empty, error) {
	id, err := uuid.Parse(token.GetToken())
	if err != nil {
		return nil, herror.StatusWithInfo(codes.InvalidArgument, "Invalid token", herror.AuthInvalidTokenError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Err(err).Msg("Error getting db connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Error getting db connection", herror.DatabaseError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		log.Err(err).Msg("Error starting transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Error starting transaction", herror.DatabaseTxError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}
	_, err = tx.Exec(ctx, "DELETE FROM users.user_session WHERE id = $1", id)
	if err != nil {
		log.Err(err).Msg("Error deleting user session")
		return nil, herror.StatusWithInfo(codes.Internal, "Error deleting user session", herror.DatabaseError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	return &emptypb.Empty{}, nil
}

func (IdentityServer) Register(ctx context.Context, registerRequest *identity.RegisterRequest) (*emptypb.Empty, error) {
	email := registerRequest.GetUser().GetEmail()
	name := registerRequest.GetUser().GetName()
	password := registerRequest.GetPassword()
	description := registerRequest.GetUser().GetDescription()
	interests := registerRequest.GetUser().GetInterests()
	contractTypes := registerRequest.GetUser().GetContractTypes()

	fieldViolations := []*errdetails.BadRequest_FieldViolation{}
	if email == "" {
		fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{Field: "user.email", Description: "Email is required", Reason: herror.RequiredFieldError})
	}
	if name == "" {
		fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{Field: "user.name", Description: "Name is required", Reason: herror.RequiredFieldError})
	}
	if password == "" {
		fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{Field: "password", Description: "Password is required", Reason: herror.RequiredFieldError})
	} else if len(password) != 64 {
		fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{Field: "password", Description: "Password must be a SHA-256", Reason: herror.InvalidFieldError})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Err(err).Msg("Error hashing password")
		return nil, herror.StatusWithInfo(codes.Internal, "Error hashing password", herror.InternalError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}
	password = string(hashedPassword)

	if len(fieldViolations) > 0 {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing required fields", fieldViolations).Err()
	}

	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Err(err).Msg("Error getting db connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Error getting db connection", herror.DatabaseError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		log.Err(err).Msg("Error starting transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Error starting transaction", herror.DatabaseTxError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}
	defer tx.Rollback(ctx)

	id := uuid.New()
	_, err = tx.Exec(ctx, "INSERT INTO users.users (id, email, name, password) VALUES ($1, $2, $3, $4)", id, email, name, password)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.ConstraintName == "uq_users_email" {
				return nil, herror.StatusWithInfo(codes.AlreadyExists, "User already exists", herror.UserAlreadyExists, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
			}
		}
		log.Err(err).Msg("Error creating user")
		return nil, herror.StatusWithInfo(codes.Internal, "Error creating user", herror.DatabaseError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}

	if description != "" {
		_, err = tx.Exec(ctx, "UPDATE users.users SET description = $1 WHERE id = $2", description, id)
		if err != nil {
			log.Err(err).Msg("Error creating user description")
			return nil, herror.StatusWithInfo(codes.Internal, "Error creating user description", herror.DatabaseError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
		}
	}

	for _, interest := range interests {
		_, err = tx.Exec(ctx, "INSERT INTO users.user_interest (user_id, interest_id) VALUES ($1, $2)", id, interest.GetId())
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.ConstraintName == "fk_user_interest_interest" {
					return nil, herror.StatusWithInfo(codes.NotFound, "Interest doesn't exist", herror.InvalidInterest, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
				} else if pgErr.ConstraintName == "pk_user_interest" {
					// Already has the interest, ignore
					continue
				}
			}
			log.Err(err).Msg("Error creating user interest")
			return nil, herror.StatusWithInfo(codes.Internal, "Error creating user interest", herror.DatabaseError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
		}
	}

	for _, contractType := range contractTypes {
		_, err = tx.Exec(ctx, "INSERT INTO users.user_contract_type (user_id, contract_type_id) VALUES ($1, $2)", id, contractType.GetId())
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.ConstraintName == "fk_user_contract_type_contract_type" {
					return nil, herror.StatusWithInfo(codes.NotFound, "Contract type doesn't exist", herror.InvalidContractType, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
				} else if pgErr.ConstraintName == "pk_user_contract_type" {
					// Already has the contract type, ignore
					continue
				}
			}
			log.Err(err).Msg("Error creating user contract type")
			return nil, herror.StatusWithInfo(codes.Internal, "Error creating user contract type", herror.DatabaseError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Err(err).Msg("Error committing transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Error committing transaction", herror.DatabaseTxError, identity.IdentityManagementService_ServiceDesc.ServiceName, nil).Err()
	}
	return &emptypb.Empty{}, nil
}
