-- name: GetDataproduct :one
SELECT *
FROM dataproducts
WHERE id = @id;

-- name: GetDataproducts :many
SELECT *
FROM dataproducts
ORDER BY last_modified DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: DeleteDataproduct :exec
DELETE
FROM dataproducts
WHERE id = @id;

-- name: CreateDataproduct :one
INSERT INTO dataproducts ("name",
                          "description",
                          "pii",
                          "type",
                          "group",
                          "slug",
                          "repo",
                          "keywords")
VALUES (@name,
        @description,
        @pii,
        @type,
        @owner_group,
        @slug,
        @repo,
        @keywords)
RETURNING *;

-- name: UpdateDataproduct :one
UPDATE dataproducts
SET "name"        = @name,
    "description" = @description,
    "pii"         = @pii,
    "slug"        = @slug,
    "repo"        = @repo,
    "keywords"    = @keywords 
WHERE id = @id
RETURNING *;

-- name: GetBigqueryDatasource :one
SELECT *
FROM datasource_bigquery
WHERE dataproduct_id = @dataproduct_id;

-- name: GetBigqueryDatasources :many
SELECT *
FROM datasource_bigquery;

-- name: CreateBigqueryDatasource :one
INSERT INTO datasource_bigquery ("dataproduct_id",
                                 "project_id",
                                 "dataset",
                                 "table_name")
VALUES (@dataproduct_id,
        @project_id,
        @dataset,
        @table_name)
RETURNING *;

-- name: UpdateBigqueryDatasourceSchema :exec
UPDATE datasource_bigquery
SET "schema" = @schema
WHERE dataproduct_id = @dataproduct_id;
