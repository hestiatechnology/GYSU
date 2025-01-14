package scompany

import (
	"context"
	"hestia/jobfair/api/pb/common"
	"hestia/jobfair/api/pb/company"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CompanyServer struct {
	company.UnimplementedCompanyServiceServer
}

func (CompanyServer) CreateCompany(ctx context.Context, c *company.Company) (*common.Id, error) {
	return nil, nil
}
func (CompanyServer) GetCompanies(ctx context.Context, emptypb *emptypb.Empty) (*company.ListCompaniesResponse, error) {
	return nil, nil
}
func (CompanyServer) GetCompany(ctx context.Context, id *common.Id) (*company.Company, error) {
	return nil, nil
}
func (CompanyServer) UpdateCompany(ctx context.Context, c *company.Company) (*emptypb.Empty, error) {
	return nil, nil
}
