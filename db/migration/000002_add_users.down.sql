-- ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
-- ALTER TABLE "accpunts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency")

ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "owner_currency_key"; 

ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey"; 

DROP TABLE IF EXISTS "users";