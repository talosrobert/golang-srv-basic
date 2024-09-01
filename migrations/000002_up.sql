-- 1. Add last_minute_bids column to items
ALTER TABLE appl.auction_items
ADD COLUMN IF NOT EXISTS last_minute_bids smallint DEFAULT 100 NOT NULL CHECK (last_minute_bids >= 0);
