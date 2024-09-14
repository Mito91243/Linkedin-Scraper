CREATE TABLE IF NOT EXISTS profiles (
id bigserial PRIMARY KEY,
created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
fullname text NOT NULL,
position text NOT NULL,
email text NOT NULL,
link text NOT NULL,
version integer NOT NULL DEFAULT 1
);
