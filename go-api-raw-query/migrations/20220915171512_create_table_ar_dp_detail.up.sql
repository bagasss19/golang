-- migrate:up
CREATE TABLE IF NOT EXISTS ar_dp_detail (
  dp_detail_id serial PRIMARY KEY,
  dp_id integer NOT NULL,
  amount_in_loc decimal(15,2) DEFAULT 0.00,
  amount_in_doc decimal(15,2) DEFAULT 0.00,
  ppn_code integer DEFAULT 0,
  tax_amount decimal(15,2) DEFAULT 0.00,
  po_number integer DEFAULT 0,
  po_item integer DEFAULT 0,
  assign varchar(50) DEFAULT '',
  payment_block integer DEFAULT 0,
  payment_meet integer DEFAULT 0,
  payment_met integer DEFAULT 0,
  profit_id varchar(50) DEFAULT '',
  due_on timestamp DEFAULT now(),
  orders varchar(50) DEFAULT '',
  status integer NOT NULL,
  reason varchar(255) NOT NULL,
  created_time timestamp NOT NULL,
  created_by varchar(50) NOT NULL,
  last_update timestamp NOT NULL,
  updated_by varchar(50) NOT NULL
)

-- migrate:down