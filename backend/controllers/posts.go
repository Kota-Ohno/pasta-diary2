package controllers

import (
	"context"
	"net/http"
	"pasta-diary2-backend/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var postCollection *mongo.Collection = config.DB.Database("blog").Collection("posts")

func GetPosts(c *gin.Context) {
    var posts []bson.M
    cursor, err := postCollection.Find(context.Background(), bson.D{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer cursor.Close(context.Background())
    cursor.All(context.Background(), &posts)
    c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
    var post bson.M
    if err := c.BindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    result, err := postCollection.InsertOne(context.Background(), post)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, result)
}

func UpdatePost(c *gin.Context) {
    id, _ := primitive.ObjectIDFromHex(c.Param("id"))
    var post bson.M
    if err := c.BindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    result, err := postCollection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": post})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, result)
}

func DeletePost(c *gin.Context) {
    id, _ := primitive.ObjectIDFromHex(c.Param("id"))
    result, err := postCollection.DeleteOne(context.Background(), bson.M{"_id": id})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, result)
}
