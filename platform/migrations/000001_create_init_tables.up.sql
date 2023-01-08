-- ./platform/migrations/000001_create_init_tables.up.sql

-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- All timezone databases https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Europe/Berlin";


-- Create Ratings table
CREATE TABLE RATINGS (
    "ID" UUID DEFAULT uuid_generate_v4 (),
    "CREATED_AT" TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    "COMPANY" VARCHAR (255) NOT NULL,
    "AUTHOR" VARCHAR (255) NOT NULL,
    "AUTHOR_ROLE" VARCHAR (255) NOT NULL,
    "COMPANY_RATING" INT NOT NULL DEFAULT 1,
    "DESCRIPTION" TEXT NOT NULL,

    PRIMARY KEY("ID"),
    CONSTRAINT valid_rating
        CHECK ("COMPANY_RATING" <= 5)
);
