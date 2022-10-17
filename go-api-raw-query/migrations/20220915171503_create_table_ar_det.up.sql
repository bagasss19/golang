-- migrate:up
CREATE TABLE IF NOT EXISTS ar_account_receivable_det (
  ar_det_id serial PRIMARY KEY,
  ar_id integer NOT NULL REFERENCES ar_account_receivable(ar_id),
  transaction_id integer NOT NULL,
  invoice timestamp NOT NULL,
  total_payment decimal(15,2) DEFAULT 0.00,
  disc_payment decimal(15,2) DEFAULT 0.00,
  cash_payment decimal(15,2) DEFAULT 0.00,
  giro_num integer DEFAULT 0,
  giro_amount decimal(15,2) DEFAULT 0.00,
  transfer_num integer DEFAULT 0,
  transfer_amount decimal(15,2) DEFAULT 0.00,
  cndn_num integer DEFAULT 0,
  cndn_amount decimal(15,2) DEFAULT 0.00,
  return_num integer DEFAULT 0,
  return_amount decimal(15,2) DEFAULT 0.00,
  status integer NOT NULL
)

-- migrate:down