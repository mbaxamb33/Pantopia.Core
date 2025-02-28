-- Drop foreign keys
ALTER TABLE "goals" DROP CONSTRAINT "goals_user_id_fkey";
ALTER TABLE "contacts" DROP CONSTRAINT "contacts_user_id_fkey";
ALTER TABLE "users" DROP CONSTRAINT "users_account_id_fkey";

-- Drop tables
DROP TABLE IF EXISTS "goals";
DROP TABLE IF EXISTS "contacts";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "accounts";
