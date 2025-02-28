CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "contacts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar,
  "phone" varchar,
  "company_name" varchar,
  "address" text,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "goals" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "description" text,
  "type" varchar DEFAULT 'standard',
  "target_value" numeric,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "contacts" ("email");

COMMENT ON COLUMN "goals"."name" IS 'e.g., "Increase meeting rates"';

COMMENT ON COLUMN "goals"."type" IS 'standard or custom';

COMMENT ON COLUMN "goals"."target_value" IS 'e.g., number of meetings to set';

ALTER TABLE "users" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "contacts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "goals" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
