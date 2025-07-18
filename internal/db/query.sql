-- name: NewUser :one
INSERT INTO users (name, passwd, salt)
VALUES (
        sqlc.arg(name),
        sqlc.arg(passwd),
        sqlc.arg(salt)
    )
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: SetUserAdmin :one
UPDATE users
SET role = "admin"
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: GetUserFromName :one
SELECT *
FROM users
WHERE name = sqlc.arg(name)
LIMIT 1;

-- name: NewCharacter :one
INSERT INTO characters (owner, name)
VALUES (
        sqlc.arg(owner),
        sqlc.arg(name)
    )
RETURNING *;

-- name: GetCharacter :one
SELECT *
FROM characters
WHERE id = sqlc.arg(id)
LIMIT 1;