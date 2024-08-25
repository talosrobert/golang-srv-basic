-- 1. Create the Database
CREATE DATABASE appl_db;

-- Connect to the newly created database
\c appl_db

-- 2. Create the Schema
CREATE SCHEMA appl;

-- 3. Create the Role
CREATE ROLE appl_role WITH LOGIN PASSWORD 'your_password';

-- 4. Grant Privileges to the Role
GRANT CONNECT ON DATABASE appl_db TO appl_role;
GRANT USAGE ON SCHEMA appl TO appl_role;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA appl TO appl_role;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA appl TO appl_role;
GRANT ALL PRIVILEGES ON SCHEMA appl TO appl_role;

-- 5. Create Tables in the `appl` Schema

CREATE TABLE appl.auction_users (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    display_name varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE
);

CREATE TABLE appl.auction_items (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    starting_price numeric(10, 2) NOT NULL CHECK (starting_price >= 0),
    reserve_price numeric(10, 2) CHECK (reserve_price IS NULL OR reserve_price >= 0),
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    expires_at timestamp with time zone DEFAULT (now() + INTERVAL '6 weeks') NOT NULL,
    seller uuid NOT NULL REFERENCES appl.auction_users(id) ON DELETE SET NULL,
    version smallint DEFAULT 0 CHECK (version >= 0),
);

CREATE TABLE appl.item_comments (
    id serial PRIMARY KEY,
    comment_text text NOT NULL,
    comment_sent_by uuid NOT NULL REFERENCES appl.auction_users(id) ON DELETE CASCADE,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    parent_comment serial,
    CONSTRAINT parent_comment_check CHECK (parent_comment IS NULL OR parent_comment > 0)
);

CREATE TABLE appl.auction_bids (
    id serial PRIMARY KEY,
    item uuid NOT NULL REFERENCES appl.auction_items(id) ON DELETE CASCADE,
    bid_amount numeric(10, 2) NOT NULL CHECK (bid_amount >= 0),
    bid_by uuid NOT NULL REFERENCES appl.auction_users(id) ON DELETE CASCADE,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT bid_amount_check CHECK (bid_amount > 0)
);
