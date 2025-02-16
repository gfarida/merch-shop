CREATE TABLE IF NOT EXISTS users (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username      VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    balance       INT NOT NULL DEFAULT 1000 CHECK (balance >= 0),
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at    TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_active_users ON users (username) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS transaction_types (
    id    SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS transactions (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sender_id           UUID NOT NULL,
    receiver_id         UUID DEFAULT NULL,
    transaction_type_id INT NOT NULL,
    amount              INT NOT NULL CHECK (amount > 0),
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (receiver_id) REFERENCES users(id),
    FOREIGN KEY (transaction_type_id) REFERENCES transaction_types(id)
);

CREATE INDEX IF NOT EXISTS idx_transactions_sender ON transactions (sender_id);
CREATE INDEX IF NOT EXISTS idx_transactions_receiver ON transactions (receiver_id);

CREATE TABLE IF NOT EXISTS merch (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name       VARCHAR(255) UNIQUE NOT NULL,
    price      INT NOT NULL CHECK (price > 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_active_merch ON merch (name) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS purchases (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    UUID NOT NULL,
    merch_id   UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (merch_id) REFERENCES merch(id)
);

CREATE INDEX IF NOT EXISTS idx_purchases_user_id ON purchases(user_id);

CREATE TABLE IF NOT EXISTS inventory (
    user_id  UUID NOT NULL,
    merch_id UUID NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    PRIMARY KEY (user_id, merch_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (merch_id) REFERENCES merch(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_inventory_user_id ON inventory(user_id);

INSERT INTO transaction_types (id, title) VALUES
    (1, 'coin_transfer'),
    (2, 'purchase');

INSERT INTO merch (id, name, price) VALUES
    (gen_random_uuid(), 't-shirt', 80),
    (gen_random_uuid(), 'cup', 20),
    (gen_random_uuid(), 'book', 50),
    (gen_random_uuid(), 'pen', 10),
    (gen_random_uuid(), 'powerbank', 200),
    (gen_random_uuid(), 'hoody', 300),
    (gen_random_uuid(), 'umbrella', 200),
    (gen_random_uuid(), 'socks', 10),
    (gen_random_uuid(), 'wallet', 50),
    (gen_random_uuid(), 'pink-hoody', 500);
