CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- ==========================
-- Wishlist Table
-- ==========================

CREATE TABLE IF NOT EXISTS wishlists (
    w_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name VARCHAR(50) NOT NULL,

    size INT NOT NULL DEFAULT 0,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT wishlist_size_check
    CHECK(size <= 10)
);

-- ==========================
-- Bonds Table
-- ==========================

CREATE TABLE IF NOT EXISTS bonds (
    b_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name VARCHAR(30) NOT NULL,

    yield FLOAT NOT NULL,

    frequency VARCHAR(20) NOT NULL,

    rating VARCHAR(4) NOT NULL,

    min_units INT NOT NULL,

    max_units INT NOT NULL,

    maturity_date DATE NOT NULL,

    issuer VARCHAR(30) NOT NULL,

    isin VARCHAR(20) UNIQUE NOT NULL
);

-- ==========================
-- Wishlist Bond Mapping
-- ==========================

CREATE TABLE IF NOT EXISTS wishlist_bonds (
    w_id UUID NOT NULL,

    b_id UUID NOT NULL,

    PRIMARY KEY (w_id, b_id),

    CONSTRAINT fk_wishlist
        FOREIGN KEY (w_id)
        REFERENCES wishlists(w_id)
        ON DELETE CASCADE,

    CONSTRAINT fk_bond
        FOREIGN KEY (b_id)
        REFERENCES bonds(b_id)
        ON DELETE CASCADE
);