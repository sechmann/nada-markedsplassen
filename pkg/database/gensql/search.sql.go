// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: search.sql

package gensql

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const search = `-- name: Search :many
SELECT
	element_id::uuid,
	element_type::text,
	ts_rank_cd(tsv_document, query) AS rank,
	ts_headline('norwegian', "description", query, 'MinWords=10, MaxWords=20, MaxFragments=2 FragmentDelimiter=" … " StartSel="((START))" StopSel="((STOP))"')::text AS excerpt
FROM
	search,
	websearch_to_tsquery('norwegian', $1) query
WHERE
	(
		CASE
			WHEN array_length($2::text[], 1) > 0 THEN "element_type" = ANY($2)
			ELSE TRUE
		END
	)
	AND (
		CASE
			WHEN array_length($3::text[], 1) > 0 THEN "keywords" && $3
			ELSE TRUE
		END
	)
	AND (
		CASE
			WHEN $1 :: text != '' THEN "tsv_document" @@ query
			ELSE TRUE
		END
	)
	AND (
		CASE
			WHEN array_length($4::text[], 1) > 0 THEN "group" = ANY($4)
			ELSE TRUE
		END
	)
	AND (
		CASE
			WHEN array_length($5::uuid[], 1) > 0 THEN "team_id" = ANY($5)
			ELSE TRUE
		END
	)
	AND (
		CASE
			WHEN array_length($6::text[], 1) > 0 THEN "services" && $6
			ELSE TRUE
		END
	)
ORDER BY rank DESC, created ASC
LIMIT $8 OFFSET $7
`

type SearchParams struct {
	Query   string
	Types   []string
	Keyword []string
	Grp     []string
	TeamID  []uuid.UUID
	Service []string
	Offs    int32
	Lim     int32
}

type SearchRow struct {
	ElementID   uuid.UUID
	ElementType string
	Rank        float32
	Excerpt     string
}

func (q *Queries) Search(ctx context.Context, arg SearchParams) ([]SearchRow, error) {
	rows, err := q.db.QueryContext(ctx, search,
		arg.Query,
		pq.Array(arg.Types),
		pq.Array(arg.Keyword),
		pq.Array(arg.Grp),
		pq.Array(arg.TeamID),
		pq.Array(arg.Service),
		arg.Offs,
		arg.Lim,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SearchRow{}
	for rows.Next() {
		var i SearchRow
		if err := rows.Scan(
			&i.ElementID,
			&i.ElementType,
			&i.Rank,
			&i.Excerpt,
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
