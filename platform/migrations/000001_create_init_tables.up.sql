-- ./platform/migrations/000001_create_init_tables.up.sql

-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- All timezone databases https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Europe/Berlin";


-- Create Ratings table
CREATE TABLE RATINGS (
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    company VARCHAR (255) NOT NULL,
    author VARCHAR (255) NOT NULL,
    author_role VARCHAR (255) NOT NULL,
    company_rating INT NOT NULL DEFAULT 0,
    decription TEXT NOT NULL,

    PRIMARY KEY(id),
    CONSTRAINT valid_rating
        CHECK (company_rating <= 5)
);
