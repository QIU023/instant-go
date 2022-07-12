package api

import (
	"log"
	"net/http"
	"strconv"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetInstants(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Get instants error"
	index, err := strconv.ParseInt(c.Query("index"), 0, 64)
	if err != nil {
		log.Println("Parse index error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	instants := []model.Instant{}
	rows, err := database.GetInstants(userID.(string), index, pageSize)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	defer rows.Close(ctx)
	for rows.Next(ctx) {
		var instant model.Instant
		err := rows.Decode(&instant)
		if err != nil {
			log.Println("Database Decode error ", err.Error())
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"code": http.StatusBadRequest, "message": errMsg},
			)
			return
		}
		instants = append(instants, instant)
	}
	if err := rows.Err(); err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": instants})
}

func PostInstant(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Post instant error"
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Println("Bind json failed ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	instant.UserID = userID.(string)
	err := database.PostInstant(instant)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"code": http.StatusCreated,
	})
}

func UpdateInstant(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMSg := "Update instant error"
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Println("Bind json failed ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMSg},
		)
		return
	}
	instant.UserID = userID.(string)
	result, err := database.UpdateInstant(instant)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMSg},
		)
		return
	}
	if result.ModifiedCount == 0 {
		log.Println("No instant updates ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMSg},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusCreated,
	})
}

func LikeInstant(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Like instant error"
	var like model.Like
	if err := c.Bind(&like); err != nil {
		log.Println("Bind json failed ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	like.UserID = userID.(string)
	err := database.LikeInstant(like)
	if err != nil {
		log.Panic("Database Error ", err.Error())
	}
	c.JSON(http.StatusCreated, gin.H{
		"code": http.StatusCreated,
	})
}

func ShareInstant(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Share instant error"
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Println("Bind json failed ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	instant.UserID = userID.(string)
	_, err := database.ShareInstant(instant)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}
