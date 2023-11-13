// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: joinable_views.sql

package gensql

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

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

const getJoinableViewWithDataset = `-- name: GetJoinableViewWithDataset :many
SELECT
    dsrc.project_id as bq_project,
    dsrc.dataset as bq_dataset,
    dsrc.table_name as bq_table,
    datasets.id as dataset_id,
    jv.id as joinable_view_id,
    dp.group,
    jv.name as joinable_view_name,
    jv.created as joinable_view_created,
    jv.expires as joinable_view_expires
FROM
    (
        (
            joinable_views jv
            INNER JOIN joinable_views_datasource jvds ON jv.id = jvds.joinable_view_id
        )
        INNER JOIN (
            (
                datasource_bigquery dsrc
                INNER JOIN datasets ON dsrc.dataset_id = datasets.id
            )
        ) ON jvds.datasource_id = dsrc.id
    )
    INNER JOIN dataproducts dp ON datasets.dataproduct_id = dp.id
WHERE
    jv.id = $1
`

type GetJoinableViewWithDatasetRow struct {
	BqProject           string
	BqDataset           string
	BqTable             string
	DatasetID           uuid.UUID
	JoinableViewID      uuid.UUID
	Group               string
	JoinableViewName    string
	JoinableViewCreated time.Time
	JoinableViewExpires sql.NullTime
}

func (q *Queries) GetJoinableViewWithDataset(ctx context.Context, id uuid.UUID) ([]GetJoinableViewWithDatasetRow, error) {
	rows, err := q.db.QueryContext(ctx, getJoinableViewWithDataset, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetJoinableViewWithDatasetRow{}
	for rows.Next() {
		var i GetJoinableViewWithDatasetRow
		if err := rows.Scan(
			&i.BqProject,
			&i.BqDataset,
			&i.BqTable,
			&i.DatasetID,
			&i.JoinableViewID,
			&i.Group,
			&i.JoinableViewName,
			&i.JoinableViewCreated,
			&i.JoinableViewExpires,
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
