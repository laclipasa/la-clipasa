-- reverse: create "comments" table
DROP TABLE "comments";
-- reverse: modify "posts" table
ALTER TABLE "posts" DROP CONSTRAINT "posts_users_posts", DROP COLUMN "user_posts", DROP COLUMN "deleted_at", ALTER COLUMN "categories" TYPE character varying;
