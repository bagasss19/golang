-- migrate:up
CREATE TABLE IF NOT EXISTS ar_giro (
  girocek_id serial PRIMARY KEY,
  company_id varchar(50) NOT NULL,
  outlet_id integer DEFAULT 0,
  giro_date timestamp DEFAULT now(),
  giro_num integer DEFAULT 0,
  account_id varchar(50) DEFAULT '',
  account_name varchar(50) DEFAULT '',
  giro_amount decimal(15,2) DEFAULT 0.00,
  due_date timestamp DEFAULT now(),
  status integer NOT NULL,
  created_time timestamp NOT NULL,
  created_by varchar(50) NOT NULL,
  last_update timestamp NOT NULL,
  updated_by varchar(50) NOT NULL
)

-- migrate:down