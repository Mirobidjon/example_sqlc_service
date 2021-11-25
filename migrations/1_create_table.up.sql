CREATE TYPE "attribute_types" AS ENUM (
    'datetime',
    'text',
    'number'
);

CREATE TABLE "profession"
(
    "id"   uuid primary key,
    "name" varchar not null
);

CREATE TABLE "position"
(
    "id"            uuid primary key,
    "name"          varchar not null,
    "profession_id" uuid    not null
);

CREATE TABLE "attribute"
(
    "id"   uuid primary key,
    "name" varchar         not null,
    "type" attribute_types not null
);

CREATE TABLE "position_attributes"
(
    "id"           uuid primary key,
    "attribute_id" uuid    not null,
    "position_id"  uuid    not null,
    "value"        varchar not null
);

ALTER TABLE "position"
    ADD FOREIGN KEY ("profession_id") REFERENCES "profession" ("id");

ALTER TABLE "position_attributes"
    ADD FOREIGN KEY ("attribute_id") REFERENCES "attribute" ("id");

ALTER TABLE "position_attributes"
    ADD FOREIGN KEY ("position_id") REFERENCES "position" ("id");
