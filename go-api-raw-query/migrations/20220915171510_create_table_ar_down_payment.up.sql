-- migrate:up
CREATE TABLE IF NOT EXISTS ar_down_payment (
  dp_id serial PRIMARY KEY,
  doc_number integer NOT NULL,
  doc_date timestamp NOT NULL,
  doc_type varchar(50) NOT NULL,
  doc integer DEFAULT 0,
  period integer DEFAULT 0,
  posting_date timestamp NOT NULL,
  company_id varchar(50) NOT NULL,
  currency_id integer NOT NULL,
  amount decimal(15,2) DEFAULT 0.00,
  reference varchar(50) DEFAULT '',
  header_text varchar(50) DEFAULT '',
  translation_date timestamp DEFAULT now(),
  taxreporting_date timestamp DEFAULT now(),
  trading_part varchar(50) DEFAULT '',
  outlet_id integer DEFAULT 0,
  gl_id integer DEFAULT 0,
  trans_type_id integer DEFAULT 0,
  status integer NOT NULL,
  reason varchar(255) NOT NULL,
  created_time timestamp NOT NULL,
  created_by varchar(50) NOT NULL,
  last_update timestamp NOT NULL,
  updated_by varchar(50) NOT NULL
)

-- migrate:down