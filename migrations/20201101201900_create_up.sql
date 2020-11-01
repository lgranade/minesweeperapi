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
  account_id not null,
  row_amount integer not null,
  column_amount integer not null,
  accumulated_seconds integer not null,
  board text not null,
  game_status smallint not null,
  created_at timestamp not null default now(),
  primary key (id),
  constraint game_accountid_fk foreign key (account_id) references account (id)
);