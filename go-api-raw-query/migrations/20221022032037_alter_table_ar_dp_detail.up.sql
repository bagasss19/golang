ALTER TABLE ar_dp_detail
  RENAME COLUMN amount_in_loc TO amount_loc;

ALTER TABLE ar_dp_detail
  RENAME COLUMN amount_in_doc TO amount_doc;

ALTER TABLE ar_dp_detail
  RENAME COLUMN profit_id TO profit;

ALTER TABLE ar_dp_detail
  RENAME COLUMN orders TO "order";

ALTER TABLE ar_dp_detail
  ADD "text" varchar(50) DEFAULT '',
  ADD payment_ref integer DEFAULT 0,
  DROP COLUMN payment_meet