package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	//maccounting "hestia/api/methods/accounting"
	//mcompany "hestia/api/methods/company"
	//"hestia/api/methods/idm"
	//minvoicing "hestia/api/methods/invoicing"
	//mtextile "hestia/api/methods/textile"
	//"hestia/api/pb/accounting"
	//"hestia/api/pb/company"
	//"hestia/api/pb/idmanagement"
	//"hestia/api/pb/invoicing"
	//"hestia/api/pb/textile"
	"hestia/jobfair/api/methods/scompany"

	"hestia/jobfair/api/pb/company"
	"hestia/jobfair/api/utils/db"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func checkPrerequisites() error {
	dbUser := os.Getenv("PGUSER")
	dbPass := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")
	dbHost := os.Getenv("PGHOST")

	if dbUser == "" || dbPass == "" || dbName == "" || dbHost == "" {
		log.Fatal().Str("PGUSER", dbUser).Str("PGPASSWORD", dbPass).Str("PGDATABASE", dbName).Str("PGHOST", dbHost).Msg("Missing one or more database environment variables for database connection")
		return fmt.Errorf("missing one or more database environment variables for database connection")
	}

	//b2ReadAppId := os.Getenv("B2_APP_ID")
	//b2ReadAppKey := os.Getenv("B2_APP_KEY")
	//b2WriteAppId := os.Getenv("B2_WRITE_APP_ID")
	//b2WriteAppKey := os.Getenv("B2_WRITE_APP_KEY")
	//
	//if b2ReadAppId == "" || b2ReadAppKey == "" || b2WriteAppId == "" || b2WriteAppKey == "" {
	//	log.Fatal().Msg("Missing one or more Backblaze B2 environment variables for database connection")
	//	return fmt.Errorf("missing one or more Backblaze B2 environment variables for database connection")
	//}

	return nil
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
	// Check if env is set to dev or args --dev is passed
	if strings.ToLower(os.Getenv("ENV")) == "dev" || len(os.Args) > 1 && os.Args[1] == "--dev" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if err := checkPrerequisites(); err != nil {
		log.Fatal().Err(err).Msg("Prerequisites not met")
	}

	// Establish database connection at startup
	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to ping database")
	}

}

func main() {
	PORT := 9000
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	log.Info().Int("port", PORT).Msg("Server listening")
	log.Warn().Msg("ADD INTERCEPTORS, DO NOT FORGET TO ADD THEM")
	s := grpc.NewServer( /*grpc.ChainUnaryInterceptor(interceptor.AuthInterceptor, interceptor.CompanyInterceptor)*/ )

	if strings.ToLower(os.Getenv("ENV")) == "dev" || len(os.Args) > 1 && os.Args[1] == "--dev" {
		log.Info().Msg("Running in development mode")
		log.Info().Msg("Registering reflection service")
		reflection.Register(s)
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Info().Msg("Shutting down gracefully...")
		s.GracefulStop()
		log.Info().Msg("Server stopped")
	}()

	// Service registration
	log.Info().Msg("Registering services")
	//log.Info().Msg("Registering Identity Management service")
	//idmanagement.RegisterIdentityManagementServer(s, &idm.IdentityManagementServer{})
	//log.Info().Msg("Registering Textile service")
	//textile.RegisterTextileServer(s, &mtextile.TextileServer{})
	//log.Info().Msg("Registering Company Management service")
	//company.RegisterCompanyManagementServer(s, &mcompany.CompanyManagementServer{})
	//log.Info().Msg("Registering Accounting Tax service")
	//accounting.RegisterTaxServer(s, &maccounting.TaxServer{})
	//log.Info().Msg("Registering Invoicing service")
	//invoicing.RegisterInvoicingServer(s, &minvoicing.InvoiceServer{})
	// Register Company service
	log.Info().Msg("Registering Company service")
	company.RegisterCompanyServiceServer(s, &scompany.CompanyServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
