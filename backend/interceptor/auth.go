package interceptor

import (
	"context"
	"errors"
	"hestia/jobfair/api/pb/identity"
	"hestia/jobfair/api/utils/auth"
	"hestia/jobfair/api/utils/herror"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ctxKey string

const UserIdKey ctxKey = "user_id"

// AuthInterceptor is a gRPC interceptor that checks for a valid token in the metadata
func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// Methods to not check
	methods := []string{
		identity.IdentityManagement_Alive_FullMethodName,
		identity.IdentityManagement_Login_FullMethodName,
		identity.IdentityManagement_Logout_FullMethodName,
		identity.IdentityManagement_Register_FullMethodName,
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

	ctx = context.WithValue(ctx, UserIdKey, userId)

	// continue on
	return handler(ctx, req)

}
