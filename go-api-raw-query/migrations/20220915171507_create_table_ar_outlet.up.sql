-- migrate:up
CREATE TABLE IF NOT EXISTS ar_outlet (
  outlet_id serial PRIMARY KEY,
  name_outlet varchar(50) NOT NULL,
  address varchar(255) DEFAULT '',
  phone varchar(15) DEFAULT '',
  status integer NOT NULL,
  created_time timestamp NOT NULL,
  created_by varchar(50) NOT NULL,
  last_update timestamp NOT NULL,
  updated_by varchar(50) NOT NULL
)

-- migrate:down