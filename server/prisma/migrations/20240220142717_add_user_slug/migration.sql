/*
  Warnings:

  - You are about to drop the column `pageId` on the `page` table. All the data in the column will be lost.
  - You are about to drop the column `siteId` on the `page` table. All the data in the column will be lost.
  - Added the required column `is_published` to the `site` table without a default value. This is not possible if the table is not empty.

*/
-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_page" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "title" TEXT NOT NULL,
    "slug" TEXT NOT NULL,
    "dependencies" BLOB NOT NULL,
    "nodes" BLOB NOT NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME,
    "site_id" TEXT,
    "page_id" TEXT,
    CONSTRAINT "page_site_id_fkey" FOREIGN KEY ("site_id") REFERENCES "site" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT "page_page_id_fkey" FOREIGN KEY ("page_id") REFERENCES "page" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);
INSERT INTO "new_page" ("created_at", "dependencies", "id", "nodes", "slug", "title", "updated_at") SELECT "created_at", "dependencies", "id", "nodes", "slug", "title", "updated_at" FROM "page";
DROP TABLE "page";
ALTER TABLE "new_page" RENAME TO "page";
CREATE TABLE "new_site" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "title" TEXT NOT NULL,
    "user_id" TEXT NOT NULL,
    "is_published" BOOLEAN NOT NULL,
    "repository" TEXT,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME,
    "last_published_at" DATETIME,
    CONSTRAINT "site_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
INSERT INTO "new_site" ("created_at", "id", "last_published_at", "repository", "title", "updated_at", "user_id") SELECT "created_at", "id", "last_published_at", "repository", "title", "updated_at", "user_id" FROM "site";
DROP TABLE "site";
ALTER TABLE "new_site" RENAME TO "site";
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
