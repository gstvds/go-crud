-- CreateEnum
CREATE TYPE "CHANNEL" AS ENUM ('EMAIL', 'PUSH_NOTIFICATION');

-- CreateTable
CREATE TABLE "contacts" (
    "id" VARCHAR(36) NOT NULL,
    "channel" "CHANNEL" NOT NULL,
    "enabled" BOOLEAN NOT NULL DEFAULT true,
    "receiver" VARCHAR(200) NOT NULL,
    "user_id" VARCHAR(36) NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "contacts_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "contacts_id_key" ON "contacts"("id");

-- AddForeignKey
ALTER TABLE "contacts" ADD CONSTRAINT "contacts_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
