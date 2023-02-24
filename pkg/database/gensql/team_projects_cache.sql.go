// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: team_projects_cache.sql

package gensql

import (
	"context"
)

const addTeamProject = `-- name: AddTeamProject :one
INSERT INTO team_projects ("team",
                           "project")
VALUES (
    $1,
    $2
)
RETURNING team, project
`

type AddTeamProjectParams struct {
	Team    string
	Project string
}

func (q *Queries) AddTeamProject(ctx context.Context, arg AddTeamProjectParams) (TeamProject, error) {
	row := q.db.QueryRowContext(ctx, addTeamProject, arg.Team, arg.Project)
	var i TeamProject
	err := row.Scan(&i.Team, &i.Project)
	return i, err
}

const clearTeamProjectsCache = `-- name: ClearTeamProjectsCache :exec
TRUNCATE team_projects
`

func (q *Queries) ClearTeamProjectsCache(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, clearTeamProjectsCache)
	return err
}

const getTeamProjects = `-- name: GetTeamProjects :many
SELECT team, project
FROM team_projects
`

func (q *Queries) GetTeamProjects(ctx context.Context) ([]TeamProject, error) {
	rows, err := q.db.QueryContext(ctx, getTeamProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TeamProject{}
	for rows.Next() {
		var i TeamProject
		if err := rows.Scan(&i.Team, &i.Project); err != nil {
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
