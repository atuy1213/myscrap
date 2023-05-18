create table user (
	user_id           uuid PRIMARY KEY,
	access_key_id     uuid NOT NULL,
	secret_access_key uuid NOT NULL,
	created_at        timestamptz NOT NULL,
	updated_at        timestamptz NOT NULL,
);

create table article (
	user_id    uuid PRIMARY KEY,
	url        text NOT NULL,
	title      text NOT NULL,
	is_read    boolean NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL
);