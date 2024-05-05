-- Create "users" table
CREATE TABLE "users" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create "emails" table
CREATE TABLE "emails" ("id" uuid NOT NULL, "email" character varying NOT NULL, "is_verified" boolean NOT NULL DEFAULT false, "shasum" character varying NOT NULL GENERATED ALWAYS AS (encode(sha256(email::bytea), 'hex')) STORED, "created_at" timestamptz NOT NULL, "user_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "emails_users_emails" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "emails_email_key" to table: "emails"
CREATE UNIQUE INDEX "emails_email_key" ON "emails" ("email");
-- Create index "emails_shasum_key" to table: "emails"
CREATE UNIQUE INDEX "emails_shasum_key" ON "emails" ("shasum");
-- Create "avatars" table
CREATE TABLE "avatars" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "loc" character varying NOT NULL, "rating" character varying NOT NULL DEFAULT 'G', "email_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "avatars_emails_avatar" FOREIGN KEY ("email_id") REFERENCES "emails" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "avatars_email_id_key" to table: "avatars"
CREATE UNIQUE INDEX "avatars_email_id_key" ON "avatars" ("email_id");
