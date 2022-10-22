ALTER TABLE ar_dp_detail
  RENAME COLUMN amount_loc TO amount_in_loc;

ALTER TABLE ar_dp_detail
  RENAME COLUMN amount_doc TO amount_in_doc;

ALTER TABLE ar_dp_detail
  RENAME COLUMN profit TO profit_id;

ALTER TABLE ar_dp_detail
  RENAME COLUMN "order" TO orders;

ALTER TABLE ar_dp_detail
  DROP "text",
  DROP payment_ref,
  ADD COLUMN payment_meet integer DEFAULT 0;