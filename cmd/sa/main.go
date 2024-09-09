package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	serviceAccountEmulator "github.com/navikt/nada-backend/pkg/sa/emulator"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/rs/zerolog"
)

var port = flag.String("port", "8080", "Port to run the HTTP server on")

func main() {
	flag.Parse()

	log := zerolog.New(os.Stdout)

	saEmulator := serviceAccountEmulator.New(log)
	saEmulator.SetPolicy("test", &cloudresourcemanager.Policy{
		Bindings: []*cloudresourcemanager.Binding{
			{
				Role:    "roles/owner",
				Members: []string{fmt.Sprintf("user:%s", "nada@nav.no")},
			},
		},
	})
	router := saEmulator.GetRouter()

	log.Printf("Server starting on port %s...", *port)
	err := http.ListenAndServe(":"+*port, router)
	if err != nil {
		log.Fatal().Err(err).Msg("starting server")
	}
}
