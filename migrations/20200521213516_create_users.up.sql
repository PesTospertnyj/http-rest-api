create table users(
    id bigserial not null PRIMARY key,
    email varchar(128) not null UNIQUE,
    encrypted_password VARCHAR(128) not null
);