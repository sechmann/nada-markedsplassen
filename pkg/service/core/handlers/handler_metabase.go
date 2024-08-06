package handlers

import (
	"context"
	"fmt"
	"github.com/navikt/nada-backend/pkg/service/core/transport"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/navikt/nada-backend/pkg/auth"
	"github.com/navikt/nada-backend/pkg/errs"
	"github.com/navikt/nada-backend/pkg/service"
)

type MetabaseHandler struct {
	service      service.MetabaseService
	mappingQueue chan uuid.UUID
}

func (h *MetabaseHandler) MapDataset(ctx context.Context, _ *http.Request, in service.DatasetMap) (*transport.Accepted, error) {
	const op errs.Op = "MetabaseHandler.MapDataset"

	id, err := uuid.Parse(chi.URLParamFromCtx(ctx, "id"))
	if err != nil {
		return nil, errs.E(errs.InvalidRequest, op, fmt.Errorf("parsing id: %w", err))
	}

	user := auth.GetUser(ctx)

	err = h.service.CreateMappingRequest(ctx, user, id, in.Services)
	if err != nil {
		return nil, errs.E(op, err)
	}

	h.mappingQueue <- id

	return &transport.Accepted{}, nil
}

func NewMetabaseHandler(service service.MetabaseService, mappingQueue chan uuid.UUID) *MetabaseHandler {
	return &MetabaseHandler{
		service:      service,
		mappingQueue: mappingQueue,
	}
}
