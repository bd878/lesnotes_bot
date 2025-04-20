#!/bin/bash
set -e

username=${1?"Usage: user"}

printf "Create with username %s\n" $username

psql -v ON_ERROR_STOP=1 --username "$username" --dbname "lesnotes" <<-EOSQL
CREATE SCHEMA IF NOT EXISTS chats;

CREATE TABLE IF NOT EXISTS chats.chats
(
	id integer NOT NULL,
	lang text NOT NULL,
	type text NOT NULL,
	title text NOT NULL,
	username text NOT NULL,
	first_name text NOT NULL,
	last_name text NOT NULL,
	is_forum bool NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);

CREATE TRIGGER created_at_chats_trgr BEFORE UPDATE ON chats.chats FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
CREATE TRIGGER updated_at_chats_trgr BEFORE UPDATE ON chats.chats FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

GRANT USAGE ON SCHEMA chats TO editor;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA chats TO editor;
EOSQL