
-- +migrate Up
CREATE TYPE service.transaction_status AS ENUM('SUCCESSFUL', 'PENDING', 'FAILED');

CREATE TABLE IF NOT EXISTS service.transactions (
   id uuid PRIMARY KEY,
   account_id uuid NOT NULL REFERENCES service.accounts(id),
   ephemeral_account_id uuid NOT NULL REFERENCES service.ephemeral_accounts(id),
   external_id VARCHAR,
   amount INTEGER NOT NULL,
   currency VARCHAR NOT NULL,
   account_name VARCHAR,
   account_number VARCHAR,
   bank_name VARCHAR,
   status service.transaction_status NOT NULL,
   provider VARCHAR,
   provider_response JSONB,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS service.transactions;
DROP TYPE IF EXISTS service.transaction_status;
