package sidentity

import (
	"context"
	"hestia/jobfair/api/pb/common"
	"hestia/jobfair/api/pb/identity"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	identity.UnimplementedUserManagementServer
}

func (UserServer) GetUsers(context.Context, *emptypb.Empty) (*identity.UserList, error) {
	return nil, nil
}

func (UserServer) GetUser(context.Context, *common.Id) (*identity.User, error) {
	return nil, nil
}

func (UserServer) UpdateUser(context.Context, *identity.User) (*emptypb.Empty, error) {
	return nil, nil
}

func (UserServer) DeleteUser(context.Context, *common.Id) (*emptypb.Empty, error) {
	return nil, nil
}

func (UserServer) GetEducations(context.Context, *common.Id) (*identity.EducationList, error) {
	return nil, nil
}

func (UserServer) GetExperiences(context.Context, *common.Id) (*identity.ExperienceList, error) {
	return nil, nil
}

func (UserServer) AddEducation(context.Context, *identity.Education) (*emptypb.Empty, error) {
	return nil, nil
}

func (UserServer) AddExperience(context.Context, *identity.Experience) (*emptypb.Empty, error) {
	return nil, nil
}

func (UserServer) DeleteEducation(context.Context, *common.Id) (*emptypb.Empty, error) {
	return nil, nil
}

func (UserServer) DeleteExperience(context.Context, *common.Id) (*emptypb.Empty, error) {
	return nil, nil
}
