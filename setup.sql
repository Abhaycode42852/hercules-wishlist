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

    name VARCHAR(30) UNIQUE NOT NULL,

    yield FLOAT NOT NULL CHECK (yield >= 0),

    frequency VARCHAR(20) NOT NULL,

    rating VARCHAR(4) NOT NULL,

    min_units INT NOT NULL CHECK (min_units > 0),

    max_units INT NOT NULL CHECK (max_units >= min_units),

    maturity_date DATE NOT NULL,

    issuer VARCHAR(30) NOT NULL,

    isin VARCHAR(20) UNIQUE NOT NULL,

    coupon_rate DECIMAL(5,2) NOT NULL CHECK (coupon_rate >= 0),

    logo_url TEXT,

    min_investment NUMERIC(12,2) CHECK (min_investment >= 0),

    face_value DECIMAL(12,2) NOT NULL CHECK (face_value > 0),

    sector VARCHAR(50),
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

INSERT INTO wishlists(
    w_id,
    name,
    size
)
VALUES (
    gen_random_uuid(),
    'Default Wishlist',
    0
);