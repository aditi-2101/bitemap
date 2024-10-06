package api

import (
	"database/sql"
	"net/http"
	"time"

	db "bitemap/db/sqlc"
	"bitemap/util"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type User struct {
	UserID         int     `db:"user_id" json:"user_id"`
	Username       *string `db:"username" json:"username"`
	Password       *string `db:"password" json:"password"`
	ProfilePicture *string `db:"profile_picture" json:"profile_picture"`
	Biography      *string `db:"biography" json:"biography"`
	Email          *string `db:"email" json:"email"`
	CreatedAt      *string `db:"created_at" json:"created_at"`
}

type createUserRequest struct {
	Username       string `json:"username" binding:"required,alphanum"`
	Password       string `json:"password" binding:"required,min=6"`
	Email          string `json:"email" binding:"required,email"`
	ProfilePicture string `json:"profile_picture"`
	Biography      string `json:"biography"`
}

type userResponse struct {
	UserId    int32     `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		UserId:    user.UserID,
		Username:  *user.Username,
		Email:     *user.Email,
		CreatedAt: user.CreatedAt.Time,
	}
}

func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// createUser creates a new user
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	arg := db.CreateUserParams{
		Username:       &req.Username,
		Password:       &hashedPassword,
		Email:          &req.Email,
		ProfilePicture: &req.ProfilePicture,
		Biography:      &req.Biography,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqError, ok := err.(*pq.Error); ok {
			switch pqError.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

// loginUser logs in a user
func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, &req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, *user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	accessToken, _, err := server.tokenMaker.CreateToken(
		user.UserID,
		*user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)

}
