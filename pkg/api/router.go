package api

import (
	"context"
	"embed"
	"net/http"

	"github.com/navikt/nada-backend/pkg/auth"
	"github.com/navikt/nada-backend/pkg/database"
	"github.com/navikt/nada-backend/pkg/database/gensql"
	"github.com/navikt/nada-backend/pkg/openapi"
	"github.com/sirupsen/logrus"
)

//go:embed swagger/*
var swagger embed.FS

type DatasetEnricher interface {
	UpdateSchema(ctx context.Context, ds gensql.DatasourceBigquery) error
}

func NewRouter(repo *database.Repo,
	oauth2Config OAuth2,
	log *logrus.Entry,
	projectsMapping *auth.TeamProjectsUpdater,
	gcp GCP,
	datasetEnricher DatasetEnricher,
	middlewares ...openapi.MiddlewareFunc,
) http.Handler {
	/*	corsMW := cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowCredentials: true,
		})

		srv := New(repo, oauth2Config, log.WithField("subsystem", "api"), projectsMapping, gcp, datasetEnricher)

		latencyHistBuckets := []float64{.001, .005, .01, .025, .05, .1, .5, 1, 3, 5}
		prometheusMiddleware := PrometheusMiddleware("nada-backend", latencyHistBuckets...)
		prometheusMiddleware.Initialize("/api/v1/", http.MethodGet, http.StatusOK)

		baseRouter := chi.NewRouter()
		baseRouter.Use(prometheusMiddleware.Handler())
		baseRouter.Use(corsMW)
		baseRouter.Get("/api/login", srv.Login)
		baseRouter.Get("/api/oauth2/callback", srv.Callback)
		baseRouter.Get("/internal/isalive", func(rw http.ResponseWriter, r *http.Request) {})
		baseRouter.Get("/internal/isready", func(rw http.ResponseWriter, r *http.Request) {})
		baseRouter.Get("/internal/metrics", promhttp.Handler().(http.HandlerFunc))
		baseRouter.Get("/api/spec", func(rw http.ResponseWriter, r *http.Request) {
			spec, err := openapi.GetSwagger()
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			_ = json.NewEncoder(rw).Encode(spec)
		})

		baseRouter.Handle("/api/*", http.StripPrefix("/api/", http.FileServer(http.FS(swagger))))

		router := openapi.HandlerWithOptions(srv, openapi.ChiServerOptions{BaseRouter: baseRouter, BaseURL: "/api", Middlewares: middlewares})
		return router*/
	return nil
}
