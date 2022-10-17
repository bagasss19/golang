-- migrate:up
ALTER TABLE ar_giro
  ADD profit_center varchar(50) DEFAULT '',
  ADD bank_name varchar(50) DEFAULT '',
  ADD type varchar(20) DEFAULT '',
  DROP COLUMN outlet_id

-- migrate:down