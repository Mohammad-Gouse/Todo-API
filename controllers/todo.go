package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"todo-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var todoCollection *mongo.Collection

func InitController(client *mongo.Client, dbName string) {
	todoCollection = client.Database(dbName).Collection("todos")
}

func CreateTodoHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.ID = primitive.NewObjectID().Hex()
	todo.Created = time.Now()
	todo.Updated = todo.Created

	_, err := todoCollection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, todo)
}

func GetTodosHandler(c *gin.Context) {
	var todos []models.Todo
	userID := c.Query("user_id")
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	createdFrom := c.Query("created_from")
	createdTo := c.Query("created_to")

	filter := bson.M{"user_id": userID}
	if status != "" {
		filter["status"] = status
	}

	// Parse the created_from and created_to dates
	if createdFrom != "" {
		createdFromTime, err := time.Parse(time.RFC3339, createdFrom)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid created_from date"})
			return
		}
		filter["created"] = bson.M{"$gte": createdFromTime}
	}

	if createdTo != "" {
		createdToTime, err := time.Parse(time.RFC3339, createdTo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid created_to date"})
			return
		}
		if _, ok := filter["created"]; ok {
			filter["created"].(bson.M)["$lte"] = createdToTime
		} else {
			filter["created"] = bson.M{"$lte": createdToTime}
		}
	}

	options := options.Find()
	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))
	options.SetSort(bson.D{{Key: "created", Value: 1}})

	cursor, err := todoCollection.Find(context.Background(), filter, options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching todos"})
		return
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &todos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func UpdateTodoHandler(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"title":       todo.Title,
			"description": todo.Description,
			"status":      todo.Status,
			"updated":     time.Now(),
		},
	}

	_, err := todoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

func DeleteTodoHandler(c *gin.Context) {
	id := c.Param("id")

	filter := bson.M{"_id": id}

	_, err := todoCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
