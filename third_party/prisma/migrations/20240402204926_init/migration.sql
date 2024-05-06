-- CreateTable
CREATE TABLE "user" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "email" TEXT NOT NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME
);

-- CreateTable
CREATE TABLE "subscription" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "external_plan_id" TEXT NOT NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME NOT NULL
);

-- CreateTable
CREATE TABLE "site" (
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

-- CreateTable
CREATE TABLE "page" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "type" TEXT NOT NULL DEFAULT 'STATIC',
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

-- CreateIndex
CREATE UNIQUE INDEX "user_email_key" ON "user"("email");
