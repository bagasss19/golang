-- migrate:up
CREATE TABLE IF NOT EXISTS ar_account_receivable (
  ar_id serial PRIMARY KEY,
  company_id varchar(50) NOT NULL,
  doc_number integer NOT NULL,
  doc_date timestamp DEFAULT now(),
  posting_date timestamp NOT NULL,
  sales_id integer DEFAULT 0,
  outlet_id integer DEFAULT 0,
  collector_id integer DEFAULT 0,
  bank_id integer NOT NULL,
  description varchar(255) DEFAULT '',
  status integer NOT NULL,
  created_time timestamp NOT NULL,
  created_by varchar(50) NOT NULL,
  last_update timestamp NOT NULL,
  updated_by varchar(50) NOT NULL
)

-- migrate:down