-- migrate:up
CREATE TABLE IF NOT EXISTS ar_dp_approval (
  dp_approval_id serial PRIMARY KEY,
  user_id varchar(50) NOT NULL,
  level_id varchar(50) DEFAULT '',
  dp_id integer NOT NULL REFERENCES ar_down_payment(dp_id),
  status integer NOT NULL,
  reason varchar(255) NOT NULL,
  created_time timestamp NOT NULL,
  created_by varchar(50) NOT NULL,
  last_update timestamp NOT NULL,
  updated_by varchar(50) NOT NULL
)

-- migrate:down