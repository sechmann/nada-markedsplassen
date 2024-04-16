// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: datasets_v2.sql

package gensql

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getAccessibleDatasets = `-- name: GetAccessibleDatasets :many
SELECT
  ds.id, ds.name, ds.description, ds.pii, ds.created, ds.last_modified, ds.type, ds.tsv_document, ds.slug, ds.repo, ds.keywords, ds.dataproduct_id, ds.anonymisation_description, ds.target_user,
  dp.slug AS dp_slug,
  dp.name AS dp_name,
  dp.group
FROM
  datasets ds
  LEFT JOIN dataproducts dp ON ds.dataproduct_id = dp.id
  LEFT JOIN dataset_access dsa ON dsa.dataset_id = ds.id
WHERE
  array_length($1::TEXT[], 1) IS NOT NULL AND array_length($1::TEXT[], 1)!=0
  AND dp.group = ANY($1 :: TEXT [])
  OR $2::TEXT IS NOT NULL
  AND dsa.subject = LOWER($2)
  AND revoked IS NULL
  AND (
    expires > NOW()
    OR expires IS NULL
  )
ORDER BY
  ds.last_modified DESC
`

type GetAccessibleDatasetsParams struct {
	Groups    []string
	Requester string
}

type GetAccessibleDatasetsRow struct {
	ID                       uuid.UUID
	Name                     string
	Description              sql.NullString
	Pii                      PiiLevel
	Created                  time.Time
	LastModified             time.Time
	Type                     DatasourceType
	TsvDocument              interface{}
	Slug                     string
	Repo                     sql.NullString
	Keywords                 []string
	DataproductID            uuid.UUID
	AnonymisationDescription sql.NullString
	TargetUser               sql.NullString
	DpSlug                   sql.NullString
	DpName                   sql.NullString
	Group                    sql.NullString
}

func (q *Queries) GetAccessibleDatasets(ctx context.Context, arg GetAccessibleDatasetsParams) ([]GetAccessibleDatasetsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccessibleDatasets, pq.Array(arg.Groups), arg.Requester)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccessibleDatasetsRow{}
	for rows.Next() {
		var i GetAccessibleDatasetsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Pii,
			&i.Created,
			&i.LastModified,
			&i.Type,
			&i.TsvDocument,
			&i.Slug,
			&i.Repo,
			pq.Array(&i.Keywords),
			&i.DataproductID,
			&i.AnonymisationDescription,
			&i.TargetUser,
			&i.DpSlug,
			&i.DpName,
			&i.Group,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllDatasets = `-- name: GetAllDatasets :many
SELECT
  ds_id, ds_name, ds_description, ds_created, ds_last_modified, ds_slug, pii, ds_keywords, bq_id, bq_created, bq_last_modified, bq_expires, bq_description, bq_missing_since, pii_tags, bq_project, bq_dataset, bq_table_name, bq_table_type, pseudo_columns, bq_schema, ds_dp_id, mapping_services, access_id, access_subject, access_granter, access_expires, access_created, access_revoked, access_request_id, mb_database_id
FROM 
  dataset_view
`

func (q *Queries) GetAllDatasets(ctx context.Context) ([]DatasetView, error) {
	rows, err := q.db.QueryContext(ctx, getAllDatasets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DatasetView{}
	for rows.Next() {
		var i DatasetView
		if err := rows.Scan(
			&i.DsID,
			&i.DsName,
			&i.DsDescription,
			&i.DsCreated,
			&i.DsLastModified,
			&i.DsSlug,
			&i.Pii,
			pq.Array(&i.DsKeywords),
			&i.BqID,
			&i.BqCreated,
			&i.BqLastModified,
			&i.BqExpires,
			&i.BqDescription,
			&i.BqMissingSince,
			&i.PiiTags,
			&i.BqProject,
			&i.BqDataset,
			&i.BqTableName,
			&i.BqTableType,
			pq.Array(&i.PseudoColumns),
			&i.BqSchema,
			&i.DsDpID,
			pq.Array(&i.MappingServices),
			&i.AccessID,
			&i.AccessSubject,
			&i.AccessGranter,
			&i.AccessExpires,
			&i.AccessCreated,
			&i.AccessRevoked,
			&i.AccessRequestID,
			&i.MbDatabaseID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDatasetComplete = `-- name: GetDatasetComplete :many
SELECT
  ds_id, ds_name, ds_description, ds_created, ds_last_modified, ds_slug, pii, ds_keywords, bq_id, bq_created, bq_last_modified, bq_expires, bq_description, bq_missing_since, pii_tags, bq_project, bq_dataset, bq_table_name, bq_table_type, pseudo_columns, bq_schema, ds_dp_id, mapping_services, access_id, access_subject, access_granter, access_expires, access_created, access_revoked, access_request_id, mb_database_id
FROM
  dataset_view
WHERE
  ds_id = $1
`

func (q *Queries) GetDatasetComplete(ctx context.Context, id uuid.UUID) ([]DatasetView, error) {
	rows, err := q.db.QueryContext(ctx, getDatasetComplete, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DatasetView{}
	for rows.Next() {
		var i DatasetView
		if err := rows.Scan(
			&i.DsID,
			&i.DsName,
			&i.DsDescription,
			&i.DsCreated,
			&i.DsLastModified,
			&i.DsSlug,
			&i.Pii,
			pq.Array(&i.DsKeywords),
			&i.BqID,
			&i.BqCreated,
			&i.BqLastModified,
			&i.BqExpires,
			&i.BqDescription,
			&i.BqMissingSince,
			&i.PiiTags,
			&i.BqProject,
			&i.BqDataset,
			&i.BqTableName,
			&i.BqTableType,
			pq.Array(&i.PseudoColumns),
			&i.BqSchema,
			&i.DsDpID,
			pq.Array(&i.MappingServices),
			&i.AccessID,
			&i.AccessSubject,
			&i.AccessGranter,
			&i.AccessExpires,
			&i.AccessCreated,
			&i.AccessRevoked,
			&i.AccessRequestID,
			&i.MbDatabaseID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
