-- migrate:up
CREATE TABLE IF NOT EXISTS ar_transaction_type (
  transaction_type_id serial PRIMARY KEY,
  description varchar(255) NOT NULL,
  status integer NOT NULL,
  created_time timestamp NOT NULL,
  created_by varchar(50) NOT NULL,
  last_update timestamp NOT NULL,
  updated_by varchar(50) NOT NULL
)

-- migrate:down