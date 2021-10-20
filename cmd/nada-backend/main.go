package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/navikt/nada-backend/pkg/api"
	"github.com/navikt/nada-backend/pkg/auth"
	"github.com/navikt/nada-backend/pkg/database"
	"github.com/navikt/nada-backend/pkg/database/gensql"
	"github.com/navikt/nada-backend/pkg/graph"
	"github.com/navikt/nada-backend/pkg/graph/generated"
	"github.com/navikt/nada-backend/pkg/metadata"
)

var cfg = DefaultConfig()

const (
	TeamsUpdateFrequency           = 5 * time.Minute
	EnsureAccessUpdateFrequency    = 5 * time.Minute
	TeamProjectsUpdateFrequency    = 5 * time.Minute
	DatasetMetadataUpdateFrequency = 5 * time.Minute
)

func init() {
	flag.StringVar(&cfg.BindAddress, "bind-address", cfg.BindAddress, "Bind address")
	flag.StringVar(&cfg.DBConnectionDSN, "db-connection-dsn", fmt.Sprintf("%v?sslmode=disable", getEnv("NAIS_DATABASE_NADA_BACKEND_NADA_URL", "postgres://postgres:postgres@127.0.0.1:5432/nada")), "database connection DSN")
	flag.StringVar(&cfg.Hostname, "hostname", os.Getenv("HOSTNAME"), "Hostname the application is served from")
	flag.StringVar(&cfg.OAuth2.ClientID, "oauth2-client-id", os.Getenv("CLIENT_ID"), "OAuth2 client ID")
	flag.StringVar(&cfg.OAuth2.ClientSecret, "oauth2-client-secret", os.Getenv("CLIENT_SECRET"), "OAuth2 client secret")
	flag.StringVar(&cfg.ProdTeamProjectsOutputURL, "prod-team-projects-url", cfg.ProdTeamProjectsOutputURL, "URL for json containing prod team projects")
	flag.StringVar(&cfg.DevTeamProjectsOutputURL, "dev-team-projects-url", cfg.DevTeamProjectsOutputURL, "URL for json containing dev team projects")
	flag.StringVar(&cfg.TeamsToken, "teams-token", os.Getenv("GITHUB_READ_TOKEN"), "Token for accessing teams json")
	flag.StringVar(&cfg.LogLevel, "log-level", "info", "which log level to output")
	flag.StringVar(&cfg.CookieSecret, "cookie-secret", "", "Secret used when encrypting cookies")
	flag.BoolVar(&cfg.MockAuth, "mock-auth", false, "Use mock authentication")
	flag.BoolVar(&cfg.SkipMetadataSync, "skip-metadata-sync", false, "Skip metadata sync")
	flag.StringVar(&cfg.GoogleAdminImpersonationSubject, "google-admin-subject", os.Getenv("GOOGLE_ADMIN_IMPERSONATION_SUBJECT"), "Subject to impersonate when accessing google admin apis")
	flag.StringVar(&cfg.ServiceAccountFile, "service-account-file", os.Getenv("GOOGLE_ADMIN_CREDENTIALS_PATH"), "Service account file for accessing google admin apis")
}

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	log := newLogger()

	repo, err := database.New(cfg.DBConnectionDSN)
	if err != nil {
		log.WithError(err).Fatal("setting up database")
	}

	authenticatorMiddleware := auth.MockJWTValidatorMiddleware()
	_ = authenticatorMiddleware
	teamProjectsMapping := &auth.MockTeamProjectsUpdater
	var oauth2Config api.OAuth2
	_ = oauth2Config
	if !cfg.MockAuth {
		teamProjectsMapping = auth.NewTeamProjectsUpdater(cfg.DevTeamProjectsOutputURL, cfg.ProdTeamProjectsOutputURL, cfg.TeamsToken, http.DefaultClient)
		go teamProjectsMapping.Run(ctx, TeamProjectsUpdateFrequency)

		googleGroups, err := metadata.NewGoogleGroups(ctx, cfg.ServiceAccountFile, cfg.GoogleAdminImpersonationSubject)
		if err != nil {
			log.Fatal(err)
		}

		gauth := auth.NewGoogle(cfg.OAuth2.ClientID, cfg.OAuth2.ClientSecret, cfg.Hostname)
		oauth2Config = gauth
		authenticatorMiddleware = gauth.Middleware(googleGroups)
	}

	var gcp api.GCP
	_ = gcp
	var datasetEnricher api.DatasetEnricher = &noopDatasetEnricher{}
	_ = datasetEnricher
	if !cfg.SkipMetadataSync {
		datacatalogClient, err := metadata.NewDatacatalog(ctx)
		if err != nil {
			log.WithError(err).Fatal("creating datacatalog client")
		}
		gcp = datacatalogClient
		de := metadata.New(datacatalogClient, repo, log.WithField("subsystem", "datasetenricher"))
		datasetEnricher = de
		go de.Run(ctx, DatasetMetadataUpdateFrequency)
	}

	log.Info("Listening on :8080")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.New(repo)}))
	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	server := http.Server{
		Addr:    cfg.BindAddress,
		Handler: mux,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.WithError(err).Warn("Shutdown error")
	}
}

func newLogger() *logrus.Logger {
	log := logrus.StandardLogger()
	log.SetFormatter(&logrus.JSONFormatter{})

	l, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(l)
	return log
}

func getEnv(key, fallback string) string {
	if env := os.Getenv(key); env != "" {
		return env
	}
	return fallback
}

type noopDatasetEnricher struct{}

func (n *noopDatasetEnricher) UpdateSchema(ctx context.Context, ds gensql.DatasourceBigquery) error {
	return nil
}
