
-- +migrate Up
ALTER TABLE service.account_holders ADD COLUMN IF NOT EXISTS provider_response JSONB;

ALTER TABLE service.accounts ADD COLUMN IF NOT EXISTS provider_response JSONB;
ALTER TABLE service.accounts ADD COLUMN IF NOT EXISTS is_dormant BOOLEAN NOT NULL DEFAULT TRUE;

CREATE TYPE service.ephemeral_account_status AS ENUM('ACTIVE', 'EXPIRED');

CREATE TABLE IF NOT EXISTS service.ephemeral_accounts (
   id uuid PRIMARY KEY,
   account_id uuid NOT NULL REFERENCES service.accounts(id),
   amount INTEGER,
   status service.ephemeral_account_status NOT NULL DEFAULT 'ACTIVE',
   expires_at TIMESTAMP NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL
);

-- +migrate Down
ALTER TABLE service.account_holders DROP COLUMN IF EXISTS provider_response;

ALTER TABLE service.accounts DROP COLUMN IF EXISTS provider_response;
ALTER TABLE service.accounts DROP COLUMN IF EXISTS is_dormant;

DROP TABLE IF EXISTS service.ephemeral_accounts;
