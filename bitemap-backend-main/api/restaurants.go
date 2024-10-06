package api

import (
	db "bitemap/db/sqlc"
	"bitemap/token"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Coordinates struct {
	Latitude  *float64 `form:"lat" binding:"required"`
	Longitude *float64 `form:"long" binding:"required"`
	Distance  *int     `form:"distance" binding:"required"`
}

type geom struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type restaurantResponse struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Score       string `json:"score"`
	Ratings     int32  `json:"ratings"`
	Category    string `json:"category"`
	PriceRange  string `json:"price_range"`
	FullAddress string `json:"full_address"`
	ZipCode     string `json:"zip_code"`
	StAsgeojson geom   `json:"geojson"`
}

type userRestaurantResponse struct {
	Status string               `json:"status"`
	Data   []restaurantResponse `json:"data"`
}

type reviewResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type allReviewsResponse struct {
	Status string             `json:"status"`
	Data   []db.GetReviewsRow `json:"data"`
}

type restaurantCuisineResponse struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

// listRestaurants gives a list of restaurants within a certain distance from a given point
func (server *Server) listRestaurants(ctx *gin.Context) {
	var req Coordinates
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.GetRestaurantsParams{
		Lat:      *req.Latitude,
		Long:     *req.Longitude,
		Distance: int32(*req.Distance),
	}

	restaurants, err := server.store.GetRestaurants(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	var response userRestaurantResponse

	for i := range restaurants {
		var jsonStruct geom
		var restData restaurantResponse
		jsonString := restaurants[i].StAsgeojson.(string)
		json.Unmarshal([]byte(jsonString), &jsonStruct)
		restaurants[i].StAsgeojson = jsonStruct
		response.Status = "ok"
		restData.ID = restaurants[i].ID
		restData.Name = *restaurants[i].Name
		restData.Score = *restaurants[i].Score
		restData.Ratings = *restaurants[i].Ratings
		restData.Category = *restaurants[i].Category
		restData.PriceRange = *restaurants[i].PriceRange
		restData.FullAddress = *restaurants[i].FullAddress
		restData.ZipCode = *restaurants[i].ZipCode
		restData.StAsgeojson = jsonStruct
		response.Data = append(response.Data, restData)
	}
	ctx.JSON(http.StatusOK, response)
}

// addReview adds a review to a restaurant
func (server *Server) addReview(ctx *gin.Context) {
	var req db.AddReviewParams
	var ratingReq db.UpdateRatingParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	fmt.Println(req.Rating, req.Review, req.ResID, req.UserID)
	keys := ctx.Keys["authorization_payload"].(*token.Payload)

	req.UserID = &keys.UserID

	review, err := server.store.AddReview(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ratingReq.ID = *req.ResID
	ratingReq.Rating = (float64(*req.Rating))

	_, err = server.store.UpdateRating(ctx, ratingReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	var response reviewResponse
	response.Status = "ok"
	response.Data = *review.Review

	ctx.JSON(http.StatusOK, response)

}

// getRestaurantsByFilter gives a list of restaurants filtered by category and price range
func (server *Server) getRestaurantsByFilter(ctx *gin.Context) {
	var req db.GetRestaurantsByFilterParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var category string = "%" + *req.Category + "%"
	*req.Category = category

	restaurants, err := server.store.GetRestaurantsByFilter(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	var response userRestaurantResponse

	for i := range restaurants {
		var jsonStruct geom
		var restData restaurantResponse
		jsonString := restaurants[i].StAsgeojson.(string)
		json.Unmarshal([]byte(jsonString), &jsonStruct)
		restaurants[i].StAsgeojson = jsonStruct
		response.Status = "ok"
		restData.Name = *restaurants[i].Name
		restData.Score = *restaurants[i].Score
		restData.Ratings = *restaurants[i].Ratings
		restData.Category = *restaurants[i].Category
		restData.PriceRange = *restaurants[i].PriceRange
		restData.FullAddress = *restaurants[i].FullAddress
		restData.ZipCode = *restaurants[i].ZipCode
		restData.StAsgeojson = jsonStruct
		response.Data = append(response.Data, restData)
	}
	ctx.JSON(http.StatusOK, response)
}

// getRestaurantReviews gives a list of reviews for a restaurant
func (server *Server) getRestaurantReviews(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	idInt32 := int32(id)
	reviews, err := server.store.GetReviews(ctx, idInt32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	var response allReviewsResponse
	response.Status = "ok"
	response.Data = reviews

	ctx.JSON(http.StatusOK, response)
}

// getRestaurantCuisines gives a list of cuisines around a point
func (server *Server) getRestaurantCuisines(ctx *gin.Context) {
	var req Coordinates
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.GetRestaurantCuisinesParams{
		Lat:      *req.Latitude,
		Long:     *req.Longitude,
		Distance: int32(*req.Distance),
	}

	cuisines, err := server.store.GetRestaurantCuisines(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	var responseCuisineList []string
	var uniqueCuisines = make(map[string]bool)
	var response restaurantCuisineResponse
	response.Status = "ok"

	for i := range cuisines {
		cuisineList := strings.Split(*cuisines[i], ",")
		for j := range cuisineList {
			if _, ok := uniqueCuisines[cuisineList[j]]; !ok {
				uniqueCuisines[cuisineList[j]] = true
				responseCuisineList = append(responseCuisineList, cuisineList[j])
			}
		}
	}
	response.Data = responseCuisineList
	ctx.JSON(http.StatusOK, response)
}
