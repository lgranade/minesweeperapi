-- name: CreateAccount :one
insert into account (
	id, login_name, login_password
) values (
	$1, $2, $3
)
returning *;

-- name: GetAccountByID :one
select * from account
where id = $1;