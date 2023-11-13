// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: datasets.sql

package gensql

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/sqlc-dev/pqtype"
)

const createBigqueryDatasource = `-- name: CreateBigqueryDatasource :one
INSERT INTO
  datasource_bigquery (
    "dataset_id",
    "project_id",
    "dataset",
    "table_name",
    "schema",
    "last_modified",
    "created",
    "expires",
    "table_type",
    "pii_tags",
    "pseudo_columns",
    "is_reference"
  )
VALUES
  (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12
  ) RETURNING dataset_id, project_id, dataset, table_name, schema, last_modified, created, expires, table_type, description, pii_tags, missing_since, id, is_reference, pseudo_columns
`

type CreateBigqueryDatasourceParams struct {
	DatasetID     uuid.UUID
	ProjectID     string
	Dataset       string
	TableName     string
	Schema        pqtype.NullRawMessage
	LastModified  time.Time
	Created       time.Time
	Expires       sql.NullTime
	TableType     string
	PiiTags       pqtype.NullRawMessage
	PseudoColumns []string
	IsReference   bool
}

func (q *Queries) CreateBigqueryDatasource(ctx context.Context, arg CreateBigqueryDatasourceParams) (DatasourceBigquery, error) {
	row := q.db.QueryRowContext(ctx, createBigqueryDatasource,
		arg.DatasetID,
		arg.ProjectID,
		arg.Dataset,
		arg.TableName,
		arg.Schema,
		arg.LastModified,
		arg.Created,
		arg.Expires,
		arg.TableType,
		arg.PiiTags,
		pq.Array(arg.PseudoColumns),
		arg.IsReference,
	)
	var i DatasourceBigquery
	err := row.Scan(
		&i.DatasetID,
		&i.ProjectID,
		&i.Dataset,
		&i.TableName,
		&i.Schema,
		&i.LastModified,
		&i.Created,
		&i.Expires,
		&i.TableType,
		&i.Description,
		&i.PiiTags,
		&i.MissingSince,
		&i.ID,
		&i.IsReference,
		pq.Array(&i.PseudoColumns),
	)
	return i, err
}

const createDataset = `-- name: CreateDataset :one
INSERT INTO
  datasets (
    "dataproduct_id",
    "name",
    "description",
    "pii",
    "type",
    "slug",
    "repo",
    "keywords",
    "anonymisation_description",
    "target_user"
  )
VALUES
  (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
  ) RETURNING id, name, description, pii, created, last_modified, type, tsv_document, slug, repo, keywords, dataproduct_id, anonymisation_description, target_user
`

type CreateDatasetParams struct {
	DataproductID            uuid.UUID
	Name                     string
	Description              sql.NullString
	Pii                      PiiLevel
	Type                     DatasourceType
	Slug                     string
	Repo                     sql.NullString
	Keywords                 []string
	AnonymisationDescription sql.NullString
	TargetUser               sql.NullString
}

func (q *Queries) CreateDataset(ctx context.Context, arg CreateDatasetParams) (Dataset, error) {
	row := q.db.QueryRowContext(ctx, createDataset,
		arg.DataproductID,
		arg.Name,
		arg.Description,
		arg.Pii,
		arg.Type,
		arg.Slug,
		arg.Repo,
		pq.Array(arg.Keywords),
		arg.AnonymisationDescription,
		arg.TargetUser,
	)
	var i Dataset
	err := row.Scan(
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
	)
	return i, err
}

const createJoinableViews = `-- name: CreateJoinableViews :one
INSERT INTO
  joinable_views ("name", "owner", "created", "expires")
VALUES
  ($1, $2, $3, $4) RETURNING id, owner, name, created, expires, deleted
`

type CreateJoinableViewsParams struct {
	Name    string
	Owner   string
	Created time.Time
	Expires sql.NullTime
}

func (q *Queries) CreateJoinableViews(ctx context.Context, arg CreateJoinableViewsParams) (JoinableView, error) {
	row := q.db.QueryRowContext(ctx, createJoinableViews,
		arg.Name,
		arg.Owner,
		arg.Created,
		arg.Expires,
	)
	var i JoinableView
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.Created,
		&i.Expires,
		&i.Deleted,
	)
	return i, err
}

const createJoinableViewsDatasource = `-- name: CreateJoinableViewsDatasource :one
INSERT INTO
  joinable_views_datasource ("joinable_view_id", "datasource_id")
VALUES
  ($1, $2) RETURNING id, joinable_view_id, datasource_id
`

type CreateJoinableViewsDatasourceParams struct {
	JoinableViewID uuid.UUID
	DatasourceID   uuid.UUID
}

func (q *Queries) CreateJoinableViewsDatasource(ctx context.Context, arg CreateJoinableViewsDatasourceParams) (JoinableViewsDatasource, error) {
	row := q.db.QueryRowContext(ctx, createJoinableViewsDatasource, arg.JoinableViewID, arg.DatasourceID)
	var i JoinableViewsDatasource
	err := row.Scan(&i.ID, &i.JoinableViewID, &i.DatasourceID)
	return i, err
}

const datasetsByMetabase = `-- name: DatasetsByMetabase :many
SELECT
  id, name, description, pii, created, last_modified, type, tsv_document, slug, repo, keywords, dataproduct_id, anonymisation_description, target_user
FROM
  datasets
WHERE
  id IN (
    SELECT
      dataset_id
    FROM
      metabase_metadata
    WHERE
      "deleted_at" IS NULL
  )
ORDER BY
  last_modified DESC
LIMIT
  $2 OFFSET $1
`

type DatasetsByMetabaseParams struct {
	Offs int32
	Lim  int32
}

func (q *Queries) DatasetsByMetabase(ctx context.Context, arg DatasetsByMetabaseParams) ([]Dataset, error) {
	rows, err := q.db.QueryContext(ctx, datasetsByMetabase, arg.Offs, arg.Lim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dataset{}
	for rows.Next() {
		var i Dataset
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

const deleteDataset = `-- name: DeleteDataset :exec
DELETE FROM
  datasets
WHERE
  id = $1
`

func (q *Queries) DeleteDataset(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteDataset, id)
	return err
}

const getAccessiblePseudoDatasetsByUser = `-- name: GetAccessiblePseudoDatasetsByUser :many
WITH owned_dp AS(
  SELECT
    dp.id
  FROM
    dataproducts dp
  WHERE
    dp.group = ANY($2 :: text [])
)
SELECT
  included_ds.id AS dataset_id,
  included_ds.name AS name,
  sbq.project_id AS bq_project_id,
  sbq.dataset AS bq_dataset_id,
  sbq.table_name AS bq_table_id,
  sbq.id AS bq_datasource_id
FROM
  (
    (
      SELECT
        ds.id AS id,
        ds.name AS name,
        ds.dataproduct_id AS dataproduct_id
      FROM
        datasets ds
        INNER JOIN dataset_access da ON ds.id = da.dataset_id
      WHERE
        da.subject = ANY($1 :: text [])
        AND (
          da.revoked IS NULL
          AND(
            da.expires IS NULL
            OR da.expires > CURRENT_TIMESTAMP
          )
        )
      GROUP BY
        ds.id
    )
    UNION
    (
      SELECT
        ds.id AS id,
        ds.name AS name,
        ds.dataproduct_id AS dataproduct_id
      FROM
        datasets ds
        INNER JOIN owned_dp ON ds.dataproduct_id = owned_dp.id
    )
  ) AS included_ds
  INNER JOIN datasource_bigquery AS sbq ON included_ds.id = sbq.dataset_id
  AND sbq.is_reference = TRUE
`

type GetAccessiblePseudoDatasetsByUserParams struct {
	AccessSubjects []string
	OwnerSubjects  []string
}

type GetAccessiblePseudoDatasetsByUserRow struct {
	DatasetID      uuid.UUID
	Name           string
	BqProjectID    string
	BqDatasetID    string
	BqTableID      string
	BqDatasourceID uuid.UUID
}

func (q *Queries) GetAccessiblePseudoDatasetsByUser(ctx context.Context, arg GetAccessiblePseudoDatasetsByUserParams) ([]GetAccessiblePseudoDatasetsByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccessiblePseudoDatasetsByUser, pq.Array(arg.AccessSubjects), pq.Array(arg.OwnerSubjects))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccessiblePseudoDatasetsByUserRow{}
	for rows.Next() {
		var i GetAccessiblePseudoDatasetsByUserRow
		if err := rows.Scan(
			&i.DatasetID,
			&i.Name,
			&i.BqProjectID,
			&i.BqDatasetID,
			&i.BqTableID,
			&i.BqDatasourceID,
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

const getBigqueryDatasource = `-- name: GetBigqueryDatasource :one
SELECT
  dataset_id, project_id, dataset, table_name, schema, last_modified, created, expires, table_type, description, pii_tags, missing_since, id, is_reference, pseudo_columns
FROM
  datasource_bigquery
WHERE
  dataset_id = $1
  AND is_reference = $2
`

type GetBigqueryDatasourceParams struct {
	DatasetID   uuid.UUID
	IsReference bool
}

func (q *Queries) GetBigqueryDatasource(ctx context.Context, arg GetBigqueryDatasourceParams) (DatasourceBigquery, error) {
	row := q.db.QueryRowContext(ctx, getBigqueryDatasource, arg.DatasetID, arg.IsReference)
	var i DatasourceBigquery
	err := row.Scan(
		&i.DatasetID,
		&i.ProjectID,
		&i.Dataset,
		&i.TableName,
		&i.Schema,
		&i.LastModified,
		&i.Created,
		&i.Expires,
		&i.TableType,
		&i.Description,
		&i.PiiTags,
		&i.MissingSince,
		&i.ID,
		&i.IsReference,
		pq.Array(&i.PseudoColumns),
	)
	return i, err
}

const getBigqueryDatasources = `-- name: GetBigqueryDatasources :many
SELECT
  dataset_id, project_id, dataset, table_name, schema, last_modified, created, expires, table_type, description, pii_tags, missing_since, id, is_reference, pseudo_columns
FROM
  datasource_bigquery
`

func (q *Queries) GetBigqueryDatasources(ctx context.Context) ([]DatasourceBigquery, error) {
	rows, err := q.db.QueryContext(ctx, getBigqueryDatasources)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DatasourceBigquery{}
	for rows.Next() {
		var i DatasourceBigquery
		if err := rows.Scan(
			&i.DatasetID,
			&i.ProjectID,
			&i.Dataset,
			&i.TableName,
			&i.Schema,
			&i.LastModified,
			&i.Created,
			&i.Expires,
			&i.TableType,
			&i.Description,
			&i.PiiTags,
			&i.MissingSince,
			&i.ID,
			&i.IsReference,
			pq.Array(&i.PseudoColumns),
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

const getDataset = `-- name: GetDataset :one
SELECT
  id, name, description, pii, created, last_modified, type, tsv_document, slug, repo, keywords, dataproduct_id, anonymisation_description, target_user
FROM
  datasets
WHERE
  id = $1
`

func (q *Queries) GetDataset(ctx context.Context, id uuid.UUID) (Dataset, error) {
	row := q.db.QueryRowContext(ctx, getDataset, id)
	var i Dataset
	err := row.Scan(
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
	)
	return i, err
}

const getDatasets = `-- name: GetDatasets :many
SELECT
  id, name, description, pii, created, last_modified, type, tsv_document, slug, repo, keywords, dataproduct_id, anonymisation_description, target_user
FROM
  datasets
ORDER BY
  last_modified DESC
LIMIT
  $2 OFFSET $1
`

type GetDatasetsParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) GetDatasets(ctx context.Context, arg GetDatasetsParams) ([]Dataset, error) {
	rows, err := q.db.QueryContext(ctx, getDatasets, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dataset{}
	for rows.Next() {
		var i Dataset
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

const getDatasetsByGroups = `-- name: GetDatasetsByGroups :many
SELECT
  id, name, description, pii, created, last_modified, type, tsv_document, slug, repo, keywords, dataproduct_id, anonymisation_description, target_user
FROM
  datasets
WHERE
  "group" = ANY ($1 :: text [])
ORDER BY
  last_modified DESC
`

func (q *Queries) GetDatasetsByGroups(ctx context.Context, groups []string) ([]Dataset, error) {
	rows, err := q.db.QueryContext(ctx, getDatasetsByGroups, pq.Array(groups))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dataset{}
	for rows.Next() {
		var i Dataset
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

const getDatasetsByIDs = `-- name: GetDatasetsByIDs :many
SELECT
  id, name, description, pii, created, last_modified, type, tsv_document, slug, repo, keywords, dataproduct_id, anonymisation_description, target_user
FROM
  datasets
WHERE
  id = ANY ($1 :: uuid [])
ORDER BY
  last_modified DESC
`

func (q *Queries) GetDatasetsByIDs(ctx context.Context, ids []uuid.UUID) ([]Dataset, error) {
	rows, err := q.db.QueryContext(ctx, getDatasetsByIDs, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dataset{}
	for rows.Next() {
		var i Dataset
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

const getDatasetsByUserAccess = `-- name: GetDatasetsByUserAccess :many
SELECT
  id, name, description, pii, created, last_modified, type, tsv_document, slug, repo, keywords, dataproduct_id, anonymisation_description, target_user
FROM
  datasets
WHERE
  id = ANY (
    SELECT
      dataset_id
    FROM
      dataset_access
    WHERE
      "subject" = LOWER($1)
      AND revoked IS NULL
      AND (
        expires > NOW()
        OR expires IS NULL
      )
  )
ORDER BY
  last_modified DESC
`

func (q *Queries) GetDatasetsByUserAccess(ctx context.Context, id string) ([]Dataset, error) {
	rows, err := q.db.QueryContext(ctx, getDatasetsByUserAccess, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dataset{}
	for rows.Next() {
		var i Dataset
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

const getDatasetsForOwner = `-- name: GetDatasetsForOwner :many
SELECT
  ds.id, ds.name, ds.description, ds.pii, ds.created, ds.last_modified, ds.type, ds.tsv_document, ds.slug, ds.repo, ds.keywords, ds.dataproduct_id, ds.anonymisation_description, ds.target_user
FROM
  datasets ds
WHERE
  dataproduct_id IN (
    SELECT
      id
    FROM
      dataproducts dp
    WHERE
      dp.group = ANY($1 :: text [])
  )
`

func (q *Queries) GetDatasetsForOwner(ctx context.Context, groups []string) ([]Dataset, error) {
	rows, err := q.db.QueryContext(ctx, getDatasetsForOwner, pq.Array(groups))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dataset{}
	for rows.Next() {
		var i Dataset
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

const getDatasetsInDataproduct = `-- name: GetDatasetsInDataproduct :many
SELECT
  id, name, description, pii, created, last_modified, type, tsv_document, slug, repo, keywords, dataproduct_id, anonymisation_description, target_user
FROM
  datasets
WHERE
  dataproduct_id = $1
`

func (q *Queries) GetDatasetsInDataproduct(ctx context.Context, dataproductID uuid.UUID) ([]Dataset, error) {
	rows, err := q.db.QueryContext(ctx, getDatasetsInDataproduct, dataproductID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dataset{}
	for rows.Next() {
		var i Dataset
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

const getJoinableViewsForOwner = `-- name: GetJoinableViewsForOwner :many
SELECT
  jv.id AS id,
  jv.name AS name,
  jv.owner AS owner,
  jv.created AS created,
  jv.expires AS expires,
  bq.project_id AS project_id,
  bq.dataset AS dataset_id,
  bq.table_name AS table_id
FROM
  (
    joinable_views jv
    INNER JOIN (
      joinable_views_datasource jds
      INNER JOIN datasource_bigquery bq ON jds.datasource_id = bq.id
    ) ON jv.id = jds.joinable_view_id
  )
WHERE
  jv.owner = $1
  AND (
    jv.expires IS NULL
    OR jv.expires > NOW()
  )
`

type GetJoinableViewsForOwnerRow struct {
	ID        uuid.UUID
	Name      string
	Owner     string
	Created   time.Time
	Expires   sql.NullTime
	ProjectID string
	DatasetID string
	TableID   string
}

func (q *Queries) GetJoinableViewsForOwner(ctx context.Context, owner string) ([]GetJoinableViewsForOwnerRow, error) {
	rows, err := q.db.QueryContext(ctx, getJoinableViewsForOwner, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetJoinableViewsForOwnerRow{}
	for rows.Next() {
		var i GetJoinableViewsForOwnerRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Owner,
			&i.Created,
			&i.Expires,
			&i.ProjectID,
			&i.DatasetID,
			&i.TableID,
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

const getJoinableViewsForReferenceAndUser = `-- name: GetJoinableViewsForReferenceAndUser :many
SELECT
  a.id as id,
  a.name as dataset
FROM
  joinable_views a
  JOIN joinable_views_datasource b ON a.id = b.joinable_view_id
  JOIN datasource_bigquery c ON b.datasource_id = c.id
WHERE
  owner = $1
  AND c.dataset_id = $2
`

type GetJoinableViewsForReferenceAndUserParams struct {
	Owner           string
	PseudoDatasetID uuid.UUID
}

type GetJoinableViewsForReferenceAndUserRow struct {
	ID      uuid.UUID
	Dataset string
}

func (q *Queries) GetJoinableViewsForReferenceAndUser(ctx context.Context, arg GetJoinableViewsForReferenceAndUserParams) ([]GetJoinableViewsForReferenceAndUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getJoinableViewsForReferenceAndUser, arg.Owner, arg.PseudoDatasetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetJoinableViewsForReferenceAndUserRow{}
	for rows.Next() {
		var i GetJoinableViewsForReferenceAndUserRow
		if err := rows.Scan(&i.ID, &i.Dataset); err != nil {
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

const getJoinableViewsWithReference = `-- name: GetJoinableViewsWithReference :many
SELECT
  a.owner as owner,
  a.id as joinable_view_id,
  a.name as joinable_view_dataset,
  c.dataset_id as pseudo_view_id,
  c.project_id as pseudo_project_id,
  c.dataset as pseudo_dataset,
  c.table_name as pseudo_table,
  a.expires as expires
FROM
  joinable_views a
  JOIN joinable_views_datasource b ON a.id = b.joinable_view_id
  JOIN datasource_bigquery c ON b.datasource_id = c.id
WHERE
  a.deleted IS NULL
`

type GetJoinableViewsWithReferenceRow struct {
	Owner               string
	JoinableViewID      uuid.UUID
	JoinableViewDataset string
	PseudoViewID        uuid.UUID
	PseudoProjectID     string
	PseudoDataset       string
	PseudoTable         string
	Expires             sql.NullTime
}

func (q *Queries) GetJoinableViewsWithReference(ctx context.Context) ([]GetJoinableViewsWithReferenceRow, error) {
	rows, err := q.db.QueryContext(ctx, getJoinableViewsWithReference)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetJoinableViewsWithReferenceRow{}
	for rows.Next() {
		var i GetJoinableViewsWithReferenceRow
		if err := rows.Scan(
			&i.Owner,
			&i.JoinableViewID,
			&i.JoinableViewDataset,
			&i.PseudoViewID,
			&i.PseudoProjectID,
			&i.PseudoDataset,
			&i.PseudoTable,
			&i.Expires,
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

const getOwnerGroupOfDataset = `-- name: GetOwnerGroupOfDataset :one
SELECT
  d.group as group
FROM
  dataproducts d
WHERE
  d.id = (
    SELECT
      dataproduct_id
    FROM
      datasets ds
    WHERE
      ds.id = $1
  )
`

func (q *Queries) GetOwnerGroupOfDataset(ctx context.Context, datasetID uuid.UUID) (string, error) {
	row := q.db.QueryRowContext(ctx, getOwnerGroupOfDataset, datasetID)
	var group string
	err := row.Scan(&group)
	return group, err
}

const replaceDatasetsTag = `-- name: ReplaceDatasetsTag :exec
UPDATE
  datasets
SET
  "keywords" = array_replace(keywords, $1, $2)
`

type ReplaceDatasetsTagParams struct {
	TagToReplace interface{}
	TagUpdated   interface{}
}

func (q *Queries) ReplaceDatasetsTag(ctx context.Context, arg ReplaceDatasetsTagParams) error {
	_, err := q.db.ExecContext(ctx, replaceDatasetsTag, arg.TagToReplace, arg.TagUpdated)
	return err
}

const setJoinableViewDeleted = `-- name: SetJoinableViewDeleted :exec
UPDATE
  joinable_views
SET
  deleted = NOW()
WHERE
  id = $1
`

func (q *Queries) SetJoinableViewDeleted(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, setJoinableViewDeleted, id)
	return err
}

const updateBigqueryDatasource = `-- name: UpdateBigqueryDatasource :exec
UPDATE
  datasource_bigquery
SET
  "pii_tags" = $1,
  "pseudo_columns" = $2
WHERE
  dataset_id = $3
`

type UpdateBigqueryDatasourceParams struct {
	PiiTags       pqtype.NullRawMessage
	PseudoColumns []string
	DatasetID     uuid.UUID
}

func (q *Queries) UpdateBigqueryDatasource(ctx context.Context, arg UpdateBigqueryDatasourceParams) error {
	_, err := q.db.ExecContext(ctx, updateBigqueryDatasource, arg.PiiTags, pq.Array(arg.PseudoColumns), arg.DatasetID)
	return err
}

const updateBigqueryDatasourceMissing = `-- name: UpdateBigqueryDatasourceMissing :exec
UPDATE
  datasource_bigquery
SET
  "missing_since" = NOW()
WHERE
  dataset_id = $1
`

func (q *Queries) UpdateBigqueryDatasourceMissing(ctx context.Context, datasetID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, updateBigqueryDatasourceMissing, datasetID)
	return err
}

const updateBigqueryDatasourceSchema = `-- name: UpdateBigqueryDatasourceSchema :exec
UPDATE
  datasource_bigquery
SET
  "schema" = $1,
  "last_modified" = $2,
  "expires" = $3,
  "description" = $4,
  "missing_since" = null,
  "pseudo_columns" = CASE
    WHEN $5 IS NOT NULL THEN $5
    ELSE "pseudo_columns"
  END
WHERE
  dataset_id = $6
`

type UpdateBigqueryDatasourceSchemaParams struct {
	Schema        pqtype.NullRawMessage
	LastModified  time.Time
	Expires       sql.NullTime
	Description   sql.NullString
	PseudoColumns []string
	DatasetID     uuid.UUID
}

func (q *Queries) UpdateBigqueryDatasourceSchema(ctx context.Context, arg UpdateBigqueryDatasourceSchemaParams) error {
	_, err := q.db.ExecContext(ctx, updateBigqueryDatasourceSchema,
		arg.Schema,
		arg.LastModified,
		arg.Expires,
		arg.Description,
		pq.Array(arg.PseudoColumns),
		arg.DatasetID,
	)
	return err
}

const updateDataset = `-- name: UpdateDataset :one
UPDATE
  datasets
SET
  "name" = $1,
  "description" = $2,
  "pii" = $3,
  "slug" = $4,
  "repo" = $5,
  "keywords" = $6,
  "dataproduct_id" = $7,
  "anonymisation_description" = $8,
  "target_user" = $9
WHERE
  id = $10 RETURNING id, name, description, pii, created, last_modified, type, tsv_document, slug, repo, keywords, dataproduct_id, anonymisation_description, target_user
`

type UpdateDatasetParams struct {
	Name                     string
	Description              sql.NullString
	Pii                      PiiLevel
	Slug                     string
	Repo                     sql.NullString
	Keywords                 []string
	DataproductID            uuid.UUID
	AnonymisationDescription sql.NullString
	TargetUser               sql.NullString
	ID                       uuid.UUID
}

func (q *Queries) UpdateDataset(ctx context.Context, arg UpdateDatasetParams) (Dataset, error) {
	row := q.db.QueryRowContext(ctx, updateDataset,
		arg.Name,
		arg.Description,
		arg.Pii,
		arg.Slug,
		arg.Repo,
		pq.Array(arg.Keywords),
		arg.DataproductID,
		arg.AnonymisationDescription,
		arg.TargetUser,
		arg.ID,
	)
	var i Dataset
	err := row.Scan(
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
	)
	return i, err
}
