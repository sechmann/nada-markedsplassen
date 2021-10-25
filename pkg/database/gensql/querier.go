// Code generated by sqlc. DO NOT EDIT.

package gensql

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateBigqueryDatasource(ctx context.Context, arg CreateBigqueryDatasourceParams) (DatasourceBigquery, error)
	CreateCollection(ctx context.Context, arg CreateCollectionParams) (Collection, error)
	CreateCollectionElement(ctx context.Context, arg CreateCollectionElementParams) error
	CreateDataproduct(ctx context.Context, arg CreateDataproductParams) (Dataproduct, error)
	CreateDataproductRequester(ctx context.Context, arg CreateDataproductRequesterParams) error
	DeleteCollection(ctx context.Context, id uuid.UUID) error
	DeleteCollectionElement(ctx context.Context, arg DeleteCollectionElementParams) error
	DeleteDataproduct(ctx context.Context, id uuid.UUID) error
	DeleteDataproductRequester(ctx context.Context, arg DeleteDataproductRequesterParams) error
	GetBigqueryDatasource(ctx context.Context, dataproductID uuid.UUID) (DatasourceBigquery, error)
	GetBigqueryDatasources(ctx context.Context) ([]DatasourceBigquery, error)
	GetCollection(ctx context.Context, id uuid.UUID) (Collection, error)
	GetCollectionElements(ctx context.Context, collectionID uuid.UUID) ([]Dataproduct, error)
	GetCollections(ctx context.Context, arg GetCollectionsParams) ([]Collection, error)
	GetCollectionsByIDs(ctx context.Context, ids []uuid.UUID) ([]Collection, error)
	GetDataproduct(ctx context.Context, id uuid.UUID) (Dataproduct, error)
	GetDataproductRequesters(ctx context.Context, dataproductID uuid.UUID) ([]string, error)
	GetDataproducts(ctx context.Context, arg GetDataproductsParams) ([]Dataproduct, error)
	GetDataproductsByIDs(ctx context.Context, ids []uuid.UUID) ([]Dataproduct, error)
	Search(ctx context.Context, arg SearchParams) ([]SearchRow, error)
	UpdateBigqueryDatasourceSchema(ctx context.Context, arg UpdateBigqueryDatasourceSchemaParams) error
	UpdateCollection(ctx context.Context, arg UpdateCollectionParams) (Collection, error)
	UpdateDataproduct(ctx context.Context, arg UpdateDataproductParams) (Dataproduct, error)
}

var _ Querier = (*Queries)(nil)
