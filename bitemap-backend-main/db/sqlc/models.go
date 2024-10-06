// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Restaurant struct {
	ID          int32       `json:"id"`
	Position    *int32      `json:"position"`
	Name        *string     `json:"name"`
	Score       *string     `json:"score"`
	Ratings     *int32      `json:"ratings"`
	Category    *string     `json:"category"`
	PriceRange  *string     `json:"price_range"`
	FullAddress *string     `json:"full_address"`
	ZipCode     *string     `json:"zip_code"`
	Lat         *float64    `json:"lat"`
	Long        *float64    `json:"long"`
	Geom        interface{} `json:"geom"`
}

type Review struct {
	ReviewID int32    `json:"review_id"`
	UserID   *int32   `json:"user_id"`
	ResID    *int32   `json:"res_id"`
	Review   *string  `json:"review"`
	Rating   *float32 `json:"rating"`
}

type User struct {
	UserID         int32            `json:"user_id"`
	Username       *string          `json:"username"`
	Password       *string          `json:"password"`
	ProfilePicture *string          `json:"profile_picture"`
	Biography      *string          `json:"biography"`
	Email          *string          `json:"email"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
}
