// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: story.sql

package gensql

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createStory = `-- name: CreateStory :one
INSERT INTO stories (
	"name",
    "creator",
	"description",
	"keywords",
	"teamkatalogen_url",
    "team_id",
    "group"
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
    $6,
    $7
)
RETURNING id, name, creator, created, last_modified, description, keywords, teamkatalogen_url, team_id, "group"
`

type CreateStoryParams struct {
	Name             string
	Creator          string
	Description      string
	Keywords         []string
	TeamkatalogenUrl sql.NullString
	TeamID           sql.NullString
	OwnerGroup       string
}

func (q *Queries) CreateStory(ctx context.Context, arg CreateStoryParams) (Story, error) {
	row := q.db.QueryRowContext(ctx, createStory,
		arg.Name,
		arg.Creator,
		arg.Description,
		pq.Array(arg.Keywords),
		arg.TeamkatalogenUrl,
		arg.TeamID,
		arg.OwnerGroup,
	)
	var i Story
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Creator,
		&i.Created,
		&i.LastModified,
		&i.Description,
		pq.Array(&i.Keywords),
		&i.TeamkatalogenUrl,
		&i.TeamID,
		&i.Group,
	)
	return i, err
}

const createStoryWithID = `-- name: CreateStoryWithID :one
INSERT INTO stories (
    "id",
	"name",
    "creator",
	"description",
	"keywords",
	"teamkatalogen_url",
    "team_id",
    "group"
) VALUES (
    $1,
	$2,
	$3,
	$4,
	$5,
	$6,
    $7,
    $8
)
RETURNING id, name, creator, created, last_modified, description, keywords, teamkatalogen_url, team_id, "group"
`

type CreateStoryWithIDParams struct {
	ID               uuid.UUID
	Name             string
	Creator          string
	Description      string
	Keywords         []string
	TeamkatalogenUrl sql.NullString
	TeamID           sql.NullString
	OwnerGroup       string
}

func (q *Queries) CreateStoryWithID(ctx context.Context, arg CreateStoryWithIDParams) (Story, error) {
	row := q.db.QueryRowContext(ctx, createStoryWithID,
		arg.ID,
		arg.Name,
		arg.Creator,
		arg.Description,
		pq.Array(arg.Keywords),
		arg.TeamkatalogenUrl,
		arg.TeamID,
		arg.OwnerGroup,
	)
	var i Story
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Creator,
		&i.Created,
		&i.LastModified,
		&i.Description,
		pq.Array(&i.Keywords),
		&i.TeamkatalogenUrl,
		&i.TeamID,
		&i.Group,
	)
	return i, err
}

const deleteStory = `-- name: DeleteStory :exec
DELETE FROM stories
WHERE id = $1
`

func (q *Queries) DeleteStory(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteStory, id)
	return err
}

const getStories = `-- name: GetStories :many
SELECT id, name, creator, created, last_modified, description, keywords, teamkatalogen_url, team_id, "group"
FROM stories
ORDER BY last_modified DESC
`

func (q *Queries) GetStories(ctx context.Context) ([]Story, error) {
	rows, err := q.db.QueryContext(ctx, getStories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Story{}
	for rows.Next() {
		var i Story
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Creator,
			&i.Created,
			&i.LastModified,
			&i.Description,
			pq.Array(&i.Keywords),
			&i.TeamkatalogenUrl,
			&i.TeamID,
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

const getStoriesByGroups = `-- name: GetStoriesByGroups :many
SELECT id, name, creator, created, last_modified, description, keywords, teamkatalogen_url, team_id, "group"
FROM stories
WHERE "group" = ANY ($1::text[])
ORDER BY last_modified DESC
`

func (q *Queries) GetStoriesByGroups(ctx context.Context, groups []string) ([]Story, error) {
	rows, err := q.db.QueryContext(ctx, getStoriesByGroups, pq.Array(groups))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Story{}
	for rows.Next() {
		var i Story
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Creator,
			&i.Created,
			&i.LastModified,
			&i.Description,
			pq.Array(&i.Keywords),
			&i.TeamkatalogenUrl,
			&i.TeamID,
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

const getStoriesByIDs = `-- name: GetStoriesByIDs :many
SELECT id, name, creator, created, last_modified, description, keywords, teamkatalogen_url, team_id, "group"
FROM stories
WHERE id = ANY ($1::uuid[])
ORDER BY last_modified DESC
`

func (q *Queries) GetStoriesByIDs(ctx context.Context, ids []uuid.UUID) ([]Story, error) {
	rows, err := q.db.QueryContext(ctx, getStoriesByIDs, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Story{}
	for rows.Next() {
		var i Story
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Creator,
			&i.Created,
			&i.LastModified,
			&i.Description,
			pq.Array(&i.Keywords),
			&i.TeamkatalogenUrl,
			&i.TeamID,
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

const getStoriesByProductArea = `-- name: GetStoriesByProductArea :many
SELECT id, name, creator, created, last_modified, description, keywords, teamkatalogen_url, team_id, "group"
FROM stories
WHERE team_id = ANY($1::text[])
ORDER BY last_modified DESC
`

func (q *Queries) GetStoriesByProductArea(ctx context.Context, teamID []string) ([]Story, error) {
	rows, err := q.db.QueryContext(ctx, getStoriesByProductArea, pq.Array(teamID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Story{}
	for rows.Next() {
		var i Story
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Creator,
			&i.Created,
			&i.LastModified,
			&i.Description,
			pq.Array(&i.Keywords),
			&i.TeamkatalogenUrl,
			&i.TeamID,
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

const getStoriesByTeam = `-- name: GetStoriesByTeam :many
SELECT id, name, creator, created, last_modified, description, keywords, teamkatalogen_url, team_id, "group"
FROM stories
WHERE team_id = $1
ORDER BY last_modified DESC
`

func (q *Queries) GetStoriesByTeam(ctx context.Context, teamID sql.NullString) ([]Story, error) {
	rows, err := q.db.QueryContext(ctx, getStoriesByTeam, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Story{}
	for rows.Next() {
		var i Story
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Creator,
			&i.Created,
			&i.LastModified,
			&i.Description,
			pq.Array(&i.Keywords),
			&i.TeamkatalogenUrl,
			&i.TeamID,
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

const getStoriesNumberByTeam = `-- name: GetStoriesNumberByTeam :one
SELECT COUNT(*) as "count"
FROM stories
WHERE team_id = $1
`

func (q *Queries) GetStoriesNumberByTeam(ctx context.Context, teamID sql.NullString) (int64, error) {
	row := q.db.QueryRowContext(ctx, getStoriesNumberByTeam, teamID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getStory = `-- name: GetStory :one
SELECT id, name, creator, created, last_modified, description, keywords, teamkatalogen_url, team_id, "group"
FROM stories
WHERE id = $1
`

func (q *Queries) GetStory(ctx context.Context, id uuid.UUID) (Story, error) {
	row := q.db.QueryRowContext(ctx, getStory, id)
	var i Story
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Creator,
		&i.Created,
		&i.LastModified,
		&i.Description,
		pq.Array(&i.Keywords),
		&i.TeamkatalogenUrl,
		&i.TeamID,
		&i.Group,
	)
	return i, err
}

const replaceStoriesTag = `-- name: ReplaceStoriesTag :exec
UPDATE stories
SET "keywords" = array_replace(keywords, $1, $2)
`

type ReplaceStoriesTagParams struct {
	TagToReplace interface{}
	TagUpdated   interface{}
}

func (q *Queries) ReplaceStoriesTag(ctx context.Context, arg ReplaceStoriesTagParams) error {
	_, err := q.db.ExecContext(ctx, replaceStoriesTag, arg.TagToReplace, arg.TagUpdated)
	return err
}

const updateStory = `-- name: UpdateStory :one
UPDATE stories
SET
	"name" = $1,
	"description" = $2,
	"keywords" = $3,
	"teamkatalogen_url" = $4,
    "team_id" = $5,
    "group" = $6
WHERE id = $7
RETURNING id, name, creator, created, last_modified, description, keywords, teamkatalogen_url, team_id, "group"
`

type UpdateStoryParams struct {
	Name             string
	Description      string
	Keywords         []string
	TeamkatalogenUrl sql.NullString
	TeamID           sql.NullString
	OwnerGroup       string
	ID               uuid.UUID
}

func (q *Queries) UpdateStory(ctx context.Context, arg UpdateStoryParams) (Story, error) {
	row := q.db.QueryRowContext(ctx, updateStory,
		arg.Name,
		arg.Description,
		pq.Array(arg.Keywords),
		arg.TeamkatalogenUrl,
		arg.TeamID,
		arg.OwnerGroup,
		arg.ID,
	)
	var i Story
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Creator,
		&i.Created,
		&i.LastModified,
		&i.Description,
		pq.Array(&i.Keywords),
		&i.TeamkatalogenUrl,
		&i.TeamID,
		&i.Group,
	)
	return i, err
}
