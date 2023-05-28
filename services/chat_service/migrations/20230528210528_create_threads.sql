-- +goose Up
-- +goose StatementBegin
CREATE TABLE "threads" (
    "id" BIGSERIAL PRIMARY KEY,
    "channel_code" text,
    "platform_thread_id" text,
    "platform_customer_id" text,
    "customer_name" text,
    "encoded_customer_name" text,
    "customer_avatar_url" text,
    "unread_count" int4 DEFAULT 0,
    "platform_code" text,
    "seller_id" text,
    "last_message" jsonb,
    "last_message_time" int8,
    "from_type" text,
    "last_message_is_auto_reply" bool DEFAULT false,
    "bot_stop_at" int8,
    "op_source" text,
    "op_source_send_time" int8 DEFAULT 0,
    "inserted_at" timestamp(0) NOT NULL DEFAULT now(),
    "updated_at" timestamp(0) NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX ON "threads" ("channel_code","platform_thread_id" );
CREATE INDEX ON "threads" ("last_message_time" );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "threads";
-- +goose StatementEnd
