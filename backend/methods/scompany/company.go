package scompany

import (
	"context"
	"hestia/jobfair/api/pb/company"
)

type CompanyServer struct {
	company.UnimplementedCompanyServiceServer
}

func (CompanyServer) CreateCompany(ctx context.Context, req *company.CreateCompanyRequest) (*company.CreateCompanyResponse, error) {
	return nil, nil
}
func (CompanyServer) GetCompanies(ctx context.Context, req *company.GetCompaniesRequest) (*company.GetCompaniesResponse, error) {
	//db, err := db.GetDBPoolConn()
	//if err != nil {
	//	log.Err(err).Msg("Error getting db connection")
	//	return nil, herror.StatusWithInfo(codes.Internal, "Error getting db connection", herror.DatabaseError, company.CompanyService_ServiceDesc.ServiceName, nil).Err()
	//}

	return nil, nil
}
func (CompanyServer) GetCompany(ctx context.Context, req *company.GetCompanyRequest) (*company.GetCompanyResponse, error) {
	return nil, nil
}
func (CompanyServer) UpdateCompany(ctx context.Context, req *company.UpdateCompanyRequest) (*company.UpdateCompanyResponse, error) {
	return nil, nil
}
