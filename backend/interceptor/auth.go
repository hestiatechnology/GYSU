package interceptor

/*
import (
	"context"
	"errors"

	//"hestia/jobfair/api/pb/company"
	//"hestia/jobfair/api/pb/idmanagement"
	//"hestia/jobfair/api/utils/auth"
	"hestia/jobfair/api/utils/db"
	"hestia/jobfair/api/utils/herror"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ctxKey string

// AuthInterceptor is a gRPC interceptor that checks for a valid token in the metadata
func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// Methods to not check
	methods := []string{
		//idmanagement.IdentityManagement_Alive_FullMethodName,
		//idmanagement.IdentityManagement_Login_FullMethodName,
		//idmanagement.IdentityManagement_Logout_FullMethodName,
		//idmanagement.IdentityManagement_Register_FullMethodName,
	}

	// Check if the method is in the list of methods to not check
	for _, method := range methods {
		if method == info.FullMethod {
			return handler(ctx, req)
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		herror.StatusWithInfo(codes.DataLoss, "missing metadata", herror.AuthMissingMetadataError, info.FullMethod, nil).Err()
	}

	authorization := md.Get("authorization")

	if len(authorization) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Missing token")
	}

	token, err := uuid.Parse(authorization[0])
	if err != nil {

		return nil, status.Error(codes.InvalidArgument, "Invalid token")
	}
	userId, err := auth.VerifyAuthToken(ctx, token)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidToken) {
			return nil, herror.StatusWithInfo(codes.Unauthenticated, "Invalid token", herror.AuthInvalidTokenError, info.FullMethod, nil).Err()
		}
		if errors.Is(err, auth.ErrExpiredToken) {
			return nil, herror.StatusWithInfo(codes.Unauthenticated, "Token expired", herror.AuthInvalidTokenError, info.FullMethod, nil).Err()
		}
		return nil, herror.StatusWithInfo(codes.Internal, "failed to verify token", herror.AuthInvalidTokenError, info.FullMethod, nil).Err()
	}

	const userIdKey ctxKey = "user_id"
	ctx = context.WithValue(ctx, userIdKey, userId)

	// continue on
	return handler(ctx, req)

}

func CompanyInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// Methods to not check for company id
	methods := []string{
		idmanagement.IdentityManagement_Alive_FullMethodName,
		idmanagement.IdentityManagement_Login_FullMethodName,
		idmanagement.IdentityManagement_Logout_FullMethodName,
		idmanagement.IdentityManagement_Register_FullMethodName,
		company.CompanyManagement_CreateCompany_FullMethodName,
		company.CompanyManagement_GetCompanies_FullMethodName,
		// TODO: Remove these two later
		company.CompanyManagement_GetCompany_FullMethodName,
		company.CompanyManagement_UpdateCompany_FullMethodName,
	}

	// Check if the method is in the list of methods to not check
	for _, method := range methods {
		if method == info.FullMethod {
			return handler(ctx, req)
		}
	}

	companyId, err := auth.GetCompanyID(ctx)
	if err != nil {
		if errors.Is(err, auth.ErrNoMetadata) {
			return nil, herror.StatusWithInfo(codes.DataLoss, "missing metadata", herror.AuthMissingMetadataError, info.FullMethod, nil).Err()
		}
		if errors.Is(err, auth.ErrNoCompanyId) {
			return nil, herror.StatusWithInfo(codes.InvalidArgument, "missing company id", herror.InvalidCompany, info.FullMethod, nil).Err()
		}
		if errors.Is(err, auth.ErrInvalidUUID) {
			return nil, herror.StatusWithInfo(codes.InvalidArgument, "invalid company id", herror.InvalidCompanyID, info.FullMethod, nil).Err()
		}
		return nil, herror.StatusWithInfo(codes.Internal, "failed to get company id", herror.InternalError, info.FullMethod, nil).Err()
	}

	db, err := db.GetDBPoolConn()
	if err != nil {
		return nil, herror.StatusWithInfo(codes.Internal, "failed to connect to database", herror.DatabaseConnError, info.FullMethod, nil).Err()
	}

	// Get user id from context
	const userIdKey ctxKey = "user_id"
	userId, ok := ctx.Value(userIdKey).(uuid.UUID)
	if !ok {
		log.Error().Msg("failed to get user id from context")
		return nil, herror.StatusWithInfo(codes.Internal, "failed to get user id from context", herror.InternalError, info.FullMethod, nil).Err()
	}

	var count int
	err = db.QueryRow(ctx, "SELECT COUNT(user_id) FROM users.user_company WHERE user_id = $1 AND company_id = $2", userId, companyId).Scan(&count)
	if err != nil {
		return nil, herror.StatusWithInfo(codes.Internal, "failed to query database", herror.DatabaseError, info.FullMethod, nil).Err()
	}

	if count == 0 {
		return nil, herror.StatusWithInfo(codes.PermissionDenied, "company id does not belong to user", herror.InvalidCompany, info.FullMethod, nil).Err()
	}

	// continue on
	return handler(ctx, req)

}
*/
