create table account (
  id uuid not null,
  login_name text not null,
  login_password text not null,
  created_at timestamp not null default now(),
  primary key (id),
  constraint user_name_uq unique (login_name)
);

create table game (
  id uuid not null,
  account_id uuid not null,
  row_amount int not null,
  column_amount int not null,
  accumulated_seconds int not null,
  board text not null,
  mines int not null,
  mines_left int not null,
  game_status varchar(32) not null,
  created_at timestamp not null default now(),
  primary key (id),
  constraint game_accountid_fk foreign key (account_id) references account (id)
);
