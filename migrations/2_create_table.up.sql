CREATE TABLE IF NOT EXISTS company (
    id uuid primary key,
    name varchar(255) not null
);

ALTER TABLE position
    ADD COLUMN company_id uuid references company(id);