package sidentity

import (
	"context"
	"hestia/jobfair/api/pb/identity"
)

type UserServer struct {
	identity.UnimplementedUserManagementServiceServer
}

func (UserServer) AddEducation(context.Context, *identity.AddEducationRequest) (*identity.AddEducationResponse, error) {
	return nil, nil
}
func (UserServer) AddExperience(context.Context, *identity.AddExperienceRequest) (*identity.AddExperienceResponse, error) {
	return nil, nil
}
func (UserServer) DeleteEducation(context.Context, *identity.DeleteEducationRequest) (*identity.DeleteEducationResponse, error) {
	return nil, nil
}
func (UserServer) DeleteExperience(context.Context, *identity.DeleteExperienceRequest) (*identity.DeleteExperienceResponse, error) {
	return nil, nil
}
func (UserServer) DeleteUser(context.Context, *identity.DeleteUserRequest) (*identity.DeleteUserResponse, error) {
	return nil, nil
}
func (UserServer) GetEducations(context.Context, *identity.GetEducationsRequest) (*identity.GetEducationsResponse, error) {
	return nil, nil
}
func (UserServer) GetExperiences(context.Context, *identity.GetExperiencesRequest) (*identity.GetExperiencesResponse, error) {
	return nil, nil
}
func (UserServer) GetUser(context.Context, *identity.GetUserRequest) (*identity.GetUserResponse, error) {
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
