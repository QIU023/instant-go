package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetUserProfileDetail(c *gin.Context) {
	userID := c.MustGet("UserID")
	targetID := c.Query("userID")
	if targetID == "" {
		targetID = userID.(string)
	}
	errMsg := "Get userinfo error"
	key := strings.Join([]string{"profileDetail", targetID}, ":")
	var user model.User
	if cache, err := database.RedisClient.Get(ctx, key).Result(); err != nil {
		log.Println(errMsg, " ", err.Error())
	} else {
		err := json.Unmarshal([]byte(cache), &user)
		if err != nil {
			log.Println("Unmarshal error ", err.Error())
		} else {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
			return
		}
	}
	err := database.GetUserProfileByID(targetID).Decode(&user)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	cache, err := json.Marshal(user)
	if err != nil {
		log.Println("Marshal error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
	}
	if err := database.RedisClient.Set(ctx, key, cache, redisExpireTime).Err(); err != nil {
		log.Println("Redis error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
}

func GetUserProfile(c *gin.Context) {
	userID := c.MustGet("UserID")
	targetID := c.Query("userID")
	if targetID == "" {
		targetID = userID.(string)
	}
	errMsg := "Get userinfo error"
	key := strings.Join([]string{"profile", targetID}, ":")
	var user model.SimpleUser
	if cache, err := database.RedisClient.Get(ctx, key).Result(); err != nil {
		log.Println(errMsg, " ", err.Error())
	} else {
		err := json.Unmarshal([]byte(cache), &user)
		if err != nil {
			log.Println("Unmarshal error ", err.Error())
		} else {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
			return
		}
	}
	err := database.GetUserProfileByID(targetID).Decode(&user)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	cache, err := json.Marshal(user)
	if err != nil {
		log.Println("Marshal error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
	}
	if err := database.RedisClient.Set(ctx, key, cache, redisExpireTime).Err(); err != nil {
		log.Println("Redis error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
}

func QueryUsers(c *gin.Context) {
	userID := c.MustGet("UserID")
	keyword := c.Query("keyword")
	errMsg := "Query users error"
	if keyword == "" {
		log.Println(errMsg, " Keyword is empty")
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": "Keyword is empty"},
		)
		return
	}
	index, err := strconv.ParseInt(c.Query("index"), 0, 64)
	if err != nil {
		log.Println("Parse index error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	rows, err := database.QueryUsers(userID.(string), keyword, index, pageSize)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	defer rows.Close(ctx)
	users := []model.QueryUser{}
	for rows.Next(ctx) {
		var user model.QueryUser
		err := rows.Decode(&user)
		if err != nil {
			log.Println("Database Decode error ", err.Error())
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"code": http.StatusBadRequest, "message": errMsg},
			)
			return
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": users})
}
