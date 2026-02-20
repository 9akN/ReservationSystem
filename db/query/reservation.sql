-- name: GetReservation :one
SELECT * FROM reservations
WHERE id = $1;

-- name: ListReservations :many
SELECT * FROM reservations
ORDER BY date, time;

-- name: CreateReservation :one
INSERT INTO reservations (
    name, email, phone, date, time, guests, notes
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateReservation :one
UPDATE reservations
SET
    name       = $2,
    email      = $3,
    phone      = $4,
    date       = $5,
    time       = $6,
    guests     = $7,
    notes      = $8,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteReservation :exec
DELETE FROM reservations
WHERE id = $1;
