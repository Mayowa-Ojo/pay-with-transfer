
-- +migrate Up
CREATE TABLE IF NOT EXISTS service.account_holders (
   id uuid PRIMARY KEY,
   first_name VARCHAR NOT NULL,
   last_name VARCHAR NOT NULL,
   email VARCHAR UNIQUE,
   phone VARCHAR UNIQUE,
   provider_id VARCHAR NOT NULL,
   provider_code VARCHAR,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS service.accounts (
   id uuid PRIMARY KEY,
   account_holder_id uuid NOT NULL REFERENCES service.account_holders(id),
   account_name VARCHAR NOT NULL,
   account_number VARCHAR NOT NULL UNIQUE,
   bank_name VARCHAR,
   currency VARCHAR NOT NULL,
   provider_id VARCHAR NOT NULL,
   provider VARCHAR NOT NULL,
   is_active BOOLEAN NOT NULL DEFAULT FALSE,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS service.accounts;
DROP TABLE IF EXISTS service.account_holders;
