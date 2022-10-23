-- migrate:up
CREATE TABLE IF NOT EXISTS ar_account_receivable (
  ar_id serial PRIMARY KEY,
  company_code varchar(50) NOT NULL,
  doc_date timestamp DEFAULT now(),
  posting_date timestamp NOT NULL,
  sales_id integer DEFAULT 0,
  outlet_id integer DEFAULT 0,
  collector_id integer DEFAULT 0,
  bank_id integer NOT NULL,
  description varchar(255) DEFAULT '',
  invoice_number integer NOT NULL,
  invoice_type varchar(15) NOT NULL,
  invoice_value decimal(15,2) NOT NULL,
  invoice_date timestamp NOT NULL,
  total_payment decimal(15,2) DEFAULT 0.00,
  discount_payment decimal(15,2) DEFAULT 0.00,
  cash_payment decimal(15,2) DEFAULT 0.00,
  giro_number integer DEFAULT 0,
  giro_payment decimal(15,2) DEFAULT 0.00,
  transfer_number integer DEFAULT 0,
  transfer_payment decimal(15,2) DEFAULT 0.00,
  cndn_number integer DEFAULT 0,
  cndn_payment decimal(15,2) DEFAULT 0.00,
  return_number integer DEFAULT 0,
  return_payment decimal(15,2) DEFAULT 0.00,
  down_payment_number integer DEFAULT 0,
  down_payment decimal(15,2) DEFAULT 0.00,
  status integer DEFAULT 0,
  created_time timestamp NOT NULL,
  created_by varchar(50) NOT NULL,
  last_update timestamp NOT NULL,
  updated_by varchar(50) NOT NULL
)

-- migrate:down