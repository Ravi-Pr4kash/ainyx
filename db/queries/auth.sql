-- name: RegisterUser :one
INSERT INTO auth_users (name, email, password_hash)
VALUES ($1, $2, $3)
RETURNING id, name, email, password_hash, created_at;

-- name: GetUserByEmail :one
SELECT id, name, email, password_hash, created_at
FROM auth_users
WHERE email = $1
LIMIT 1;
