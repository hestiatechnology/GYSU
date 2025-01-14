package scompany

import (
	"context"
	"hestia/jobfair/api/pb/common"
	"hestia/jobfair/api/pb/company"
	"hestia/jobfair/api/utils/db"
	"hestia/jobfair/api/utils/herror"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CompanyServer struct {
	company.UnimplementedCompanyServiceServer
}

func (CompanyServer) CreateCompany(ctx context.Context, c *company.Company) (*common.Id, error) {
	return nil, nil
}
func (CompanyServer) GetCompanies(ctx context.Context, emptypb *emptypb.Empty) (*company.ListCompaniesResponse, error) {
	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Err(err).Msg("Error getting db connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Error getting db connection", herror.DatabaseError, company.CompanyService_ServiceDesc.ServiceName, nil).Err()
	}
	defer db.Close()

	return nil, nil
}
func (CompanyServer) GetCompany(ctx context.Context, id *common.Id) (*company.Company, error) {
	return nil, nil
}
func (CompanyServer) UpdateCompany(ctx context.Context, c *company.Company) (*emptypb.Empty, error) {
	return nil, nil
}
