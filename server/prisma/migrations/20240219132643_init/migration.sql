-- CreateTable
CREATE TABLE "user" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "email" TEXT NOT NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME
);

-- CreateTable
CREATE TABLE "site" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "title" TEXT NOT NULL,
    "user_id" TEXT NOT NULL,
    "repository" TEXT NOT NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME,
    "last_published_at" DATETIME,
    CONSTRAINT "site_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "page" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "title" TEXT NOT NULL,
    "slug" TEXT NOT NULL,
    "dependencies" BLOB NOT NULL,
    "nodes" BLOB NOT NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME,
    "siteId" TEXT,
    "pageId" TEXT,
    CONSTRAINT "page_siteId_fkey" FOREIGN KEY ("siteId") REFERENCES "site" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT "page_pageId_fkey" FOREIGN KEY ("pageId") REFERENCES "page" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);
