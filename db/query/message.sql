-- name: StoreMessage :one
INSERT INTO messages (
  message_content
) VALUES (
  $1
) RETURNING *;