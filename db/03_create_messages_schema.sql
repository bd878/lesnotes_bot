\c lesnotes
CREATE SCHEMA IF NOT EXISTS messages;

CREATE TABLE IF NOT EXISTS messages.messages
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

CREATE TRIGGER created_at_messages_trgr BEFORE UPDATE ON messages.messages FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
CREATE TRIGGER updated_at_messages_trgr BEFORE UPDATE ON messages.messages FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

GRANT USAGE ON SCHEMA messages TO editor;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA messages TO editor;