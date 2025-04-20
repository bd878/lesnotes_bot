#!/bin/bash
set -e

username=${1?"Usage: user postgres_db"}
postgres_db=${2?"Usage: user postgres_db"}

printf "Setup with params username=%s postgres_db=%s\n" $username $postgres_db

psql --username "$username" --dbname "$postgres_db" <<-EOSQL
CREATE DATABASE lesnotes;

CREATE USER editor WITH ENCRYPTED PASSWORD 'md5e1cd61afab63c461ab483e94c917a39f';

GRANT CONNECT ON DATABASE lesnotes TO editor;
EOSQL

psql --username "$username" --dbname "lesnotes" <<-EOSQL
CREATE OR REPLACE FUNCTION created_at_trigger()
RETURNS TRIGGER AS \$\$
BEGIN
	NEW.created_at := OLD.created_at;
	RETURN NEW;
END
\$\$ language plpgsql;

CREATE OR REPLACE FUNCTION updated_at_trigger()
RETURNS TRIGGER AS \$\$
BEGIN
	IF row(NEW.*) IS DISTINCT FROM row(OLD.*) THEN
		NEW.updated_at = NOW();
		RETURN NEW;
	ELSE
		RETURN OLD;
	END IF;
END;
\$\$ language 'plpgsql';
EOSQL