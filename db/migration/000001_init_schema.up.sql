CREATE TABLE accounts (
  id BIGSERIAL PRIMARY KEY,
  owner VARCHAR NOT NULL,
  balance NUMERIC(15, 2) NOT NULL,
  currency VARCHAR NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE entries (
  id BIGSERIAL PRIMARY KEY,
  account_id BIGINT NOT NULL,
  amount NUMERIC NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE transfers (
  id BIGSERIAL PRIMARY KEY,
  from_account_id BIGINT NOT NULL,
  to_account_id BIGINT NOT NULL,
  amount NUMERIC NOT NULL CHECK (amount > 0),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ON accounts (owner);
CREATE INDEX ON entries (account_id);
CREATE INDEX ON transfers (from_account_id);
CREATE INDEX ON transfers (to_account_id);
CREATE INDEX ON transfers (from_account_id, to_account_id);

COMMENT ON COLUMN entries.amount IS 'Positive or negative amount representing deposits or withdrawals';
COMMENT ON COLUMN transfers.amount IS 'Positive amount representing transfer value';

ALTER TABLE entries ADD FOREIGN KEY (account_id) REFERENCES accounts (id);
ALTER TABLE transfers ADD FOREIGN KEY (from_account_id) REFERENCES accounts (id);
ALTER TABLE transfers ADD FOREIGN KEY (to_account_id) REFERENCES accounts (id);
