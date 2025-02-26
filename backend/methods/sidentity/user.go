package sidentity

import (
	"context"
	"hestia/jobfair/api/pb/identity"
	"hestia/jobfair/api/utils/auth"
	"hestia/jobfair/api/utils/db"
	"hestia/jobfair/api/utils/herror"

	"google.golang.org/grpc/codes"
)

type UserServer struct {
	identity.UnimplementedUserManagementServiceServer
}

// Can only get a user if
// 1. the user is an admin
// 2. the user is a staff
// 3. the user is a company AND the user has visited the company
// 4. the user is himself
func (UserServer) GetUser(ctx context.Context, req *identity.GetUserRequest) (*identity.GetUserResponse, error) {
	role := auth.GetRole(ctx)
	if role == "" {
		return nil, herror.StatusWithInfo(codes.Internal, "user role not found", "USER_NO_ROLE", identity.UserManagementService_GetUser_FullMethodName, nil).Err()
	}
	reqUserId := req.GetId()
	userId := auth.GetUserId(ctx)
	db, err := db.GetDBPoolConn()
	if err != nil {
		return nil, herror.StatusWithInfo(codes.Internal, "database connection error", "DB_CONN_ERROR", identity.UserManagementService_GetUser_FullMethodName, nil).Err()
	}

	if role == auth.CompanyRole {
		// check if the user has visited the company
		var count int
		db.QueryRow(ctx, `
			SELECT
				COUNT(DISTINCT cu.company_id)
			FROM company.company_users cu
			JOIN company.visit v
				ON cu.company_id = v.company_id
			WHERE
				cu.user_id = $1 AND v.user_id = $2`,
			userId,
			reqUserId,
		).Scan(&count)

		if count == 0 {
			return nil, herror.StatusWithInfo(codes.PermissionDenied, "user has not visited the company", "USER_NOT_VISITED_COMPANY", identity.UserManagementService_GetUser_FullMethodName, nil).Err()
		}
	}

	if role != auth.AdminRole && role != auth.StaffRole && role != auth.CompanyRole && reqUserId.String() != userId.String() {
		return nil, herror.StatusWithInfo(codes.PermissionDenied, "unauthorized to get user", "USER_UNAUTHORIZED", identity.UserManagementService_GetUser_FullMethodName, nil).Err()
	}

	var user *identity.User
	db.QueryRow(ctx, `
		SELECT
			id,
			`).Scan(&user)
	return nil, nil
}

func (UserServer) GetUsers(context.Context, *identity.GetUsersRequest) (*identity.GetUsersResponse, error) {
	return nil, nil
}

func (UserServer) SearchUsers(context.Context, *identity.SearchUsersRequest) (*identity.SearchUsersResponse, error) {
	return nil, nil
}

func (UserServer) UpdateUser(context.Context, *identity.UpdateUserRequest) (*identity.UpdateUserResponse, error) {
	return nil, nil
}

func (UserServer) DeleteUser(context.Context, *identity.DeleteUserRequest) (*identity.DeleteUserResponse, error) {
	return nil, nil
}

func (UserServer) AddEducation(context.Context, *identity.AddEducationRequest) (*identity.AddEducationResponse, error) {
	return nil, nil
}

func (UserServer) DeleteEducation(context.Context, *identity.DeleteEducationRequest) (*identity.DeleteEducationResponse, error) {
	return nil, nil
}

func (UserServer) GetEducations(context.Context, *identity.GetEducationsRequest) (*identity.GetEducationsResponse, error) {
	return nil, nil
}

func (UserServer) AddExperience(context.Context, *identity.AddExperienceRequest) (*identity.AddExperienceResponse, error) {
	return nil, nil
}

func (UserServer) DeleteExperience(context.Context, *identity.DeleteExperienceRequest) (*identity.DeleteExperienceResponse, error) {
	return nil, nil
}

func (UserServer) GetExperiences(context.Context, *identity.GetExperiencesRequest) (*identity.GetExperiencesResponse, error) {
	return nil, nil
}
