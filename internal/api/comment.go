package api

import (
	"log"
	"net/http"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {
	var comments []model.Comment
	rows, err := database.GetComments(1)
	for rows.Next() {
		var comment model.Comment
		err := rows.Scan(&comment.CommentID, &comment.CreateTime, &comment.UpdateTime, &comment.UserID, &comment.Content)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": comments})
}

func PostComment(c *gin.Context) {
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error()})
	}
	result, err := database.PostComment(comment);
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Post comment error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK, "data": id,
	})
}

func LikeComment(c *gin.Context) {
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error()})
	}
	result, err := database.PostComment(comment)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Post instant error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK, "data": id,
	})
}

func ShareComment(c *gin.Context) {
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error()})
	}
	result, err :=database.PostComment(comment)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Post instant error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK, "data": id,
	})
}
