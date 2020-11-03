-- name: CreateGame :one
insert into game (
	id, account_id, row_amount, column_amount, accumulated_seconds, board, mines, mines_left, cells_stepped, game_status
) values (
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
returning *;

-- name: GetGameByID :one
select * from game
where id = $1;

-- name: UpdateGame :one
update game
set accumulated_seconds = $1, game_status = $2, board = $3, mines_left = $4, cells_stepped = $5
where id = $6
returning *;
