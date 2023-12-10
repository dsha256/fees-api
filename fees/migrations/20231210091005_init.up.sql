-- create "bills" table
CREATE TABLE "bills" (
  "id" uuid NOT NULL,
  "currency" character varying NOT NULL DEFAULT 'USD',
  "total" bigint NOT NULL DEFAULT 0,
  "is_open" boolean NOT NULL DEFAULT true,
  "closed_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- create "line_items" table
CREATE TABLE "line_items" (
  "id" uuid NOT NULL,
  "name" character varying NOT NULL,
  "price" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "added_at" timestamptz NOT NULL,
  "bill_line_items" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "line_items_bills_line_items" FOREIGN KEY ("bill_line_items") REFERENCES "bills" ("id") ON UPDATE NO ACTION ON DELETE SET NULL
);
