-- +goose Up
CREATE TABLE users (
    id            bigserial PRIMARY KEY,
    phone_number  varchar(18) UNIQUE NOT NULL,
    email         varchar(255) UNIQUE NULL,
    full_name     varchar(255),
    status        smallint, -- | 0 = unverified | 1 = verified | 2 = blocked
    updated_at    timestamptz DEFAULT NOW(),
    created_at    timestamptz DEFAULT NOW()
);

-- +goose Down
DROP TABLE users;