-- name: GetRestaurants :many
SELECT id, name, score, ratings, category, price_range, full_address, zip_code, ST_AsGeoJSON(geom)
FROM restaurants
WHERE ST_DWithin(ST_Transform(geom, 900913), ST_Transform(ST_SetSRID(ST_Point(sqlc.arg(long)::float, sqlc.arg(lat)::float), 4326), 900913), sqlc.arg(distance)::int);

-- name: AddReview :one
INSERT INTO reviews (user_id, res_id, review, rating)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateRating :one
UPDATE restaurants 
SET score=cast(round(cast(((cast(score as float)*cast(ratings as float))+sqlc.arg(rating)::float)/((cast(ratings as float))+1.0) as numeric),1) as varchar), ratings=ratings+1 
WHERE id=$1
RETURNING id;

-- name: GetRestaurantsByFilter :many
SELECT name, score, ratings, category, price_range, full_address, zip_code, ST_AsGeoJSON(geom)
FROM restaurants
WHERE category ilike $1 AND price_range = $2 AND cast(score as float) >= sqlc.arg(ratings)::float AND id = ANY(sqlc.arg(ids)::int[]);

-- name: GetReviews :many
SELECT review_id, review, u.username, r.rating
FROM reviews r
JOIN users u ON r.user_id = u.user_id
WHERE r.res_id = sqlc.arg(res_id)::int;

-- name: GetRestaurantCuisines :many
SELECT category
FROM restaurants
WHERE ST_DWithin(ST_Transform(geom, 900913), ST_Transform(ST_SetSRID(ST_Point(sqlc.arg(long)::float, sqlc.arg(lat)::float), 4326), 900913), sqlc.arg(distance)::int);