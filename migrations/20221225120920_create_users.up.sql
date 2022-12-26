CREATE TABLE users (
    id bigserial not null primary key,
    username varchar not null unique,
    pub_key varchar not null
);