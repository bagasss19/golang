package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	ID                string
	user_name         string
	fixed_longitude   int64
	fixed_latitutde   int64
	user_gender       string
	body              string
	userscreenname    string
	fixed_location    string
	insight_sentiment string
	post_retweetcount int64
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
	return
}

func getData(c *gin.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	datas := []Data{}
	collection := client.Database("backend").Collection("data")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Error while getting all data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	for cur.Next(context.TODO()) {
		var data Data
		cur.Decode(&data)
		datas = append(datas, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   datas,
	})
	return
}

func createData(c *gin.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("backend").Collection("data")
	var data Data
	c.BindJSON(&data)
	id := guuid.New().String()

	newData := Data{
		ID:                id,
		user_name:         data.user_name,
		fixed_longitude:   data.fixed_longitude,
		fixed_latitutde:   data.fixed_latitutde,
		user_gender:       data.user_gender,
		body:              data.body,
		userscreenname:    data.userscreenname,
		fixed_location:    data.fixed_location,
		insight_sentiment: data.insight_sentiment,
		post_retweetcount: data.post_retweetcount,
	}

	_, error := collection.InsertOne(context.TODO(), newData)

	if error != nil {
		log.Printf("Error while inserting new data into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Data created Successfully",
	})
	return
}

func deleteData(c *gin.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("backend").Collection("data")
	_, error := collection.DeleteOne(context.TODO(), bson.M{"id": c.Param("id")})
	if error != nil {
		log.Printf("Error while deleting a single data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Data deleted successfully",
	})
	return
}
