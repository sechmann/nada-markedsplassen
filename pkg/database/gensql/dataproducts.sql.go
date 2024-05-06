// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: dataproducts.sql

package gensql

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createDataproduct = `-- name: CreateDataproduct :one
INSERT INTO dataproducts ("name",
                          "description",
                          "group",
                          "teamkatalogen_url",
                          "slug",
                          "team_contact",
                          "team_id")
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7)
RETURNING id, name, description, "group", created, last_modified, tsv_document, slug, teamkatalogen_url, team_contact, team_id
`

type CreateDataproductParams struct {
	Name                  string
	Description           sql.NullString
	OwnerGroup            string
	OwnerTeamkatalogenUrl sql.NullString
	Slug                  string
	TeamContact           sql.NullString
	TeamID                sql.NullString
}

func (q *Queries) CreateDataproduct(ctx context.Context, arg CreateDataproductParams) (Dataproduct, error) {
	row := q.db.QueryRowContext(ctx, createDataproduct,
		arg.Name,
		arg.Description,
		arg.OwnerGroup,
		arg.OwnerTeamkatalogenUrl,
		arg.Slug,
		arg.TeamContact,
		arg.TeamID,
	)
	var i Dataproduct
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Group,
		&i.Created,
		&i.LastModified,
		&i.TsvDocument,
		&i.Slug,
		&i.TeamkatalogenUrl,
		&i.TeamContact,
		&i.TeamID,
	)
	return i, err
}

const dataproductGroupStats = `-- name: DataproductGroupStats :many
SELECT "group",
       count(1) as "count"
FROM "dataproducts"
GROUP BY "group"
ORDER BY "count" DESC
LIMIT $2 OFFSET $1
`

type DataproductGroupStatsParams struct {
	Offs int32
	Lim  int32
}

type DataproductGroupStatsRow struct {
	Group string
	Count int64
}

func (q *Queries) DataproductGroupStats(ctx context.Context, arg DataproductGroupStatsParams) ([]DataproductGroupStatsRow, error) {
	rows, err := q.db.QueryContext(ctx, dataproductGroupStats, arg.Offs, arg.Lim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DataproductGroupStatsRow{}
	for rows.Next() {
		var i DataproductGroupStatsRow
		if err := rows.Scan(&i.Group, &i.Count); err != nil {
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

const dataproductKeywords = `-- name: DataproductKeywords :many
SELECT keyword::text, count(1) as "count"
FROM (
	SELECT unnest(ds.keywords) as keyword
	FROM dataproducts dp
    INNER JOIN datasets ds ON ds.dataproduct_id = dp.id
) keywords
WHERE true
AND CASE WHEN coalesce(TRIM($1), '') = '' THEN true ELSE keyword ILIKE $1::text || '%' END
GROUP BY keyword
ORDER BY keywords."count" DESC
LIMIT 15
`

type DataproductKeywordsRow struct {
	Keyword string
	Count   int64
}

func (q *Queries) DataproductKeywords(ctx context.Context, keyword string) ([]DataproductKeywordsRow, error) {
	rows, err := q.db.QueryContext(ctx, dataproductKeywords, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DataproductKeywordsRow{}
	for rows.Next() {
		var i DataproductKeywordsRow
		if err := rows.Scan(&i.Keyword, &i.Count); err != nil {
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

const deleteDataproduct = `-- name: DeleteDataproduct :exec
DELETE
FROM dataproducts
WHERE id = $1
`

func (q *Queries) DeleteDataproduct(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteDataproduct, id)
	return err
}

const getDataproduct = `-- name: GetDataproduct :one
SELECT id, name, description, "group", created, last_modified, tsv_document, slug, teamkatalogen_url, team_contact, team_id
FROM dataproducts
WHERE id = $1
`

func (q *Queries) GetDataproduct(ctx context.Context, id uuid.UUID) (Dataproduct, error) {
	row := q.db.QueryRowContext(ctx, getDataproduct, id)
	var i Dataproduct
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Group,
		&i.Created,
		&i.LastModified,
		&i.TsvDocument,
		&i.Slug,
		&i.TeamkatalogenUrl,
		&i.TeamContact,
		&i.TeamID,
	)
	return i, err
}

const getDataproducts = `-- name: GetDataproducts :many
SELECT id, name, description, "group", created, last_modified, tsv_document, slug, teamkatalogen_url, team_contact, team_id
FROM dataproducts
ORDER BY last_modified DESC
LIMIT $2 OFFSET $1
`

type GetDataproductsParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) GetDataproducts(ctx context.Context, arg GetDataproductsParams) ([]Dataproduct, error) {
	rows, err := q.db.QueryContext(ctx, getDataproducts, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dataproduct{}
	for rows.Next() {
		var i Dataproduct
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Group,
			&i.Created,
			&i.LastModified,
			&i.TsvDocument,
			&i.Slug,
			&i.TeamkatalogenUrl,
			&i.TeamContact,
			&i.TeamID,
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

const getDataproductsByGroups = `-- name: GetDataproductsByGroups :many
SELECT id, name, description, "group", created, last_modified, tsv_document, slug, teamkatalogen_url, team_contact, team_id
FROM dataproducts
WHERE "group" = ANY ($1::text[])
ORDER BY last_modified DESC
`

func (q *Queries) GetDataproductsByGroups(ctx context.Context, groups []string) ([]Dataproduct, error) {
	rows, err := q.db.QueryContext(ctx, getDataproductsByGroups, pq.Array(groups))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dataproduct{}
	for rows.Next() {
		var i Dataproduct
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Group,
			&i.Created,
			&i.LastModified,
			&i.TsvDocument,
			&i.Slug,
			&i.TeamkatalogenUrl,
			&i.TeamContact,
			&i.TeamID,
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

const getDataproductsByIDs = `-- name: GetDataproductsByIDs :many
SELECT id, name, description, "group", created, last_modified, tsv_document, slug, teamkatalogen_url, team_contact, team_id
FROM dataproducts
WHERE id = ANY ($1::uuid[])
ORDER BY last_modified DESC
`

func (q *Queries) GetDataproductsByIDs(ctx context.Context, ids []uuid.UUID) ([]Dataproduct, error) {
	rows, err := q.db.QueryContext(ctx, getDataproductsByIDs, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dataproduct{}
	for rows.Next() {
		var i Dataproduct
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Group,
			&i.Created,
			&i.LastModified,
			&i.TsvDocument,
			&i.Slug,
			&i.TeamkatalogenUrl,
			&i.TeamContact,
			&i.TeamID,
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

const getDataproductsByProductArea = `-- name: GetDataproductsByProductArea :many
SELECT id, name, description, "group", created, last_modified, tsv_document, slug, teamkatalogen_url, team_contact, team_id, team_name, pa_name, pa_id
FROM dataproduct_with_teamkatalogen_view
WHERE team_id = ANY($1::text[])
ORDER BY created DESC
`

func (q *Queries) GetDataproductsByProductArea(ctx context.Context, teamID []string) ([]DataproductWithTeamkatalogenView, error) {
	rows, err := q.db.QueryContext(ctx, getDataproductsByProductArea, pq.Array(teamID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DataproductWithTeamkatalogenView{}
	for rows.Next() {
		var i DataproductWithTeamkatalogenView
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Group,
			&i.Created,
			&i.LastModified,
			&i.TsvDocument,
			&i.Slug,
			&i.TeamkatalogenUrl,
			&i.TeamContact,
			&i.TeamID,
			&i.TeamName,
			&i.PaName,
			&i.PaID,
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

const getDataproductsByTeam = `-- name: GetDataproductsByTeam :many
SELECT id, name, description, "group", created, last_modified, tsv_document, slug, teamkatalogen_url, team_contact, team_id
FROM dataproducts
WHERE team_id = $1
ORDER BY created DESC
`

func (q *Queries) GetDataproductsByTeam(ctx context.Context, teamID sql.NullString) ([]Dataproduct, error) {
	rows, err := q.db.QueryContext(ctx, getDataproductsByTeam, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dataproduct{}
	for rows.Next() {
		var i Dataproduct
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Group,
			&i.Created,
			&i.LastModified,
			&i.TsvDocument,
			&i.Slug,
			&i.TeamkatalogenUrl,
			&i.TeamContact,
			&i.TeamID,
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

const updateDataproduct = `-- name: UpdateDataproduct :one
UPDATE dataproducts
SET "name"              = $1,
    "description"       = $2,
    "slug"              = $3,
    "teamkatalogen_url" = $4,
    "team_contact"      = $5,
    "team_id"           = $6
WHERE id = $7
RETURNING id, name, description, "group", created, last_modified, tsv_document, slug, teamkatalogen_url, team_contact, team_id
`

type UpdateDataproductParams struct {
	Name                  string
	Description           sql.NullString
	Slug                  string
	OwnerTeamkatalogenUrl sql.NullString
	TeamContact           sql.NullString
	TeamID                sql.NullString
	ID                    uuid.UUID
}

func (q *Queries) UpdateDataproduct(ctx context.Context, arg UpdateDataproductParams) (Dataproduct, error) {
	row := q.db.QueryRowContext(ctx, updateDataproduct,
		arg.Name,
		arg.Description,
		arg.Slug,
		arg.OwnerTeamkatalogenUrl,
		arg.TeamContact,
		arg.TeamID,
		arg.ID,
	)
	var i Dataproduct
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Group,
		&i.Created,
		&i.LastModified,
		&i.TsvDocument,
		&i.Slug,
		&i.TeamkatalogenUrl,
		&i.TeamContact,
		&i.TeamID,
	)
	return i, err
}
