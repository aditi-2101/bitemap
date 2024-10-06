CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION IF NOT EXISTS postgis_topology;

CREATE TABLE IF NOT EXISTS "users" (
  "user_id" SERIAL PRIMARY KEY,
  "username" varchar,
  "password" varchar,
  "profile_picture" varchar,
  "biography" text,
  "email" varchar,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "reviews" (
  "review_id" SERIAL PRIMARY KEY,
  "user_id" integer,
  "res_id" integer,
  "review" varchar,
  "rating" real
);

CREATE TABLE IF NOT EXISTS "restaurants" (
  "id" integer PRIMARY KEY,
  "position" integer,
  "name" varchar,
  "score" varchar,
  "ratings" integer,
  "category" varchar,
  "price_range" varchar,
  "full_address" varchar,
  "zip_code" varchar,
  "lat" double precision,
  "long" double precision,
  "geom" geometry
);

ALTER TABLE "reviews" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("res_id") REFERENCES "restaurants" ("id");

INSERT INTO Users (username, email, password, profile_picture, biography)
SELECT
    'User' || generate_series, -- Username
    'user' || generate_series || '@example.com', -- Email
    '$2a$10$jLO/HQ5ECcpmW2HxVlfiK.u0S6Oe58hsNeOIn5XInyXFg61c8kfxG', -- Password
    'https://storage.googleapis.com/cusocial/download.png', -- ProfilePicture
    'Biography for user ' || generate_series -- Biography
FROM generate_series(1, 20);

