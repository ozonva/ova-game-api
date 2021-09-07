-- +goose Up
CREATE TABLE "users"(
    id SERIAL,
    email varchar(255),
    name varchar(255),
    last_name varchar(255) NULL DEFAULT NULL,
    father_name varchar(255) NULL DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "users" ADD CONSTRAINT "user_primary_key_id" PRIMARY KEY (id);

CREATE TABLE "heroes"(
    id UUID PRIMARY KEY,
    user_id int,
    type_hero smallint,
    name varchar(255),
    description text,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "heroes" ADD CONSTRAINT "hero_foreign_key_user_id" FOREIGN KEY (user_id) REFERENCES "users"(id) ON DELETE CASCADE;

-- +goose Down
DROP TABLE "heroes";
DROP TABLE "users";