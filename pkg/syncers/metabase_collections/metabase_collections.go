package metabase_collections

import (
	"context"
	"fmt"
	"strings"

	"github.com/navikt/nada-backend/pkg/errs"
	"github.com/navikt/nada-backend/pkg/service"
	"github.com/navikt/nada-backend/pkg/syncers"
	"github.com/rs/zerolog"
)

var _ syncers.Runner = &Runner{}

type Runner struct {
	api     service.MetabaseAPI
	storage service.MetabaseStorage
}

func (r *Runner) Name() string {
	return "MetabaseCollectionsSyncer"
}

func (r *Runner) RunOnce(ctx context.Context, log zerolog.Logger) error {
	err := r.AddRestrictedTagToCollections(ctx, log)
	if err != nil {
		return fmt.Errorf("adding restricted tag to collections: %w", err)
	}

	report, err := r.CollectionsReport(ctx)
	if err != nil {
		return fmt.Errorf("getting collections report: %w", err)
	}

	for _, missing := range report.Missing {
		log.Warn().Fields(map[string]interface{}{
			"dataset_id":    missing.DatasetID,
			"collection_id": missing.CollectionID,
			"database_id":   missing.DatabaseID,
		}).Msg("collection_not_in_metabase")
	}

	for _, dangling := range report.Dangling {
		log.Info().Fields(map[string]interface{}{
			"collection_id":   dangling.ID,
			"collection_name": dangling.Name,
		}).Msg("collection_not_in_database")
	}

	return nil
}

// Dangling means that a collection has been created in metabase but not stored
// in our database
type Dangling struct {
	ID   int
	Name string
}

// Missing means that a collection has been stored in our database but no
// longer exists in metabase
type Missing struct {
	DatasetID    string
	CollectionID int
	DatabaseID   int
}

type CollectionsReport struct {
	Dangling []Dangling
	Missing  []Missing
}

func (r *Runner) CollectionsReport(ctx context.Context) (*CollectionsReport, error) {
	const op errs.Op = "metabase_collections.Runner.CollectionsReport"

	metas, err := r.storage.GetAllMetadata(ctx)
	if err != nil {
		return nil, errs.E(op, err)
	}

	collections, err := r.api.GetCollections(ctx)
	if err != nil {
		return nil, errs.E(op, err)
	}

	collectionByID := make(map[int]*service.MetabaseCollection)
	for _, collection := range collections {
		collectionByID[collection.ID] = collection
	}

	report := &CollectionsReport{}

	for _, meta := range metas {
		if meta.SyncCompleted != nil && meta.CollectionID != nil && *meta.CollectionID != 0 {
			_, ok := collectionByID[*meta.CollectionID]
			if !ok {
				report.Missing = append(report.Missing, Missing{
					DatasetID:    meta.DatasetID.String(),
					CollectionID: *meta.CollectionID,
					DatabaseID:   *meta.DatabaseID,
				})
			}

			delete(collectionByID, *meta.CollectionID)
		}
	}

	for id, collection := range collectionByID {
		report.Dangling = append(report.Dangling, Dangling{
			ID:   id,
			Name: collection.Name,
		})
	}

	return report, nil
}

func (r *Runner) AddRestrictedTagToCollections(ctx context.Context, log zerolog.Logger) error {
	const op errs.Op = "metabase_collections.Runner.AddRestrictedTagToCollections"

	metas, err := r.storage.GetAllMetadata(ctx)
	if err != nil {
		return errs.E(errs.Database, op, err)
	}

	collections, err := r.api.GetCollections(ctx)
	if err != nil {
		return errs.E(errs.IO, op, err)
	}

	names := make([]string, len(collections))

	collectionByID := make(map[int]*service.MetabaseCollection)
	for i, collection := range collections {
		collectionByID[collection.ID] = collection
		names[i] = collection.Name
	}

	log.Info().Msgf("collections: %v", names)

	for _, meta := range metas {
		log.Debug().Msgf("meta: %v", meta)

		if meta.SyncCompleted != nil && meta.CollectionID != nil && *meta.CollectionID != 0 {
			collection, ok := collectionByID[*meta.CollectionID]
			if !ok {
				continue
			}

			err := r.addRestrictedTagIfMissing(ctx, log, collection)
			if err != nil {
				return errs.E(op, fmt.Errorf("adding restricted tag to collection %d: %w", collection.ID, err))
			}
		}
	}

	return nil
}

func (r *Runner) addRestrictedTagIfMissing(ctx context.Context, log zerolog.Logger, collection *service.MetabaseCollection) error {
	const op errs.Op = "metabase_collections.Runner.addRestrictedTagIfMissing"

	if !strings.Contains(collection.Name, service.MetabaseRestrictedCollectionTag) {
		newName := fmt.Sprintf("%s %s", collection.Name, service.MetabaseRestrictedCollectionTag)

		log.Info().Fields(map[string]interface{}{
			"collection_id": collection.ID,
			"existing_name": collection.Name,
			"new_name":      newName,
		}).Msg("adding_restricted_tag")

		err := r.api.UpdateCollection(ctx, &service.MetabaseCollection{
			ID:   collection.ID,
			Name: newName,
		})
		if err != nil {
			return errs.E(op, err)
		}
	}

	return nil
}

func New(api service.MetabaseAPI, storage service.MetabaseStorage) *Runner {
	return &Runner{
		api:     api,
		storage: storage,
	}
}
