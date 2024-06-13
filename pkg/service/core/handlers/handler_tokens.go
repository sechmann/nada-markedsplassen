package handlers

import (
	"encoding/json"
	"github.com/navikt/nada-backend/pkg/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type tokenHandler struct {
	tokenService   service.TokenService
	teamTokenCreds string
}

func (h *tokenHandler) GetAllTeamTokens(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if authHeaderParts[1] != h.teamTokenCreds {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenTeamMap, err := h.tokenService.GetNadaTokens(r.Context())
	if err != nil {
		log.WithError(err).Error("getting nada tokens")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payloadBytes, err := json.Marshal(tokenTeamMap)
	if err != nil {
		log.WithError(err).Error("marshalling nada token map reponse")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payloadBytes)
}

func NewTokenHandler(tokenService service.TokenService) *tokenHandler {
	return &tokenHandler{
		tokenService: tokenService,
	}
}
