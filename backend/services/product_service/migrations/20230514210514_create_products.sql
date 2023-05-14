-- +goose Up
-- +goose StatementBegin
CREATE TABLE "products" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" text,
    "price" integer,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "products";
-- +goose StatementEnd
