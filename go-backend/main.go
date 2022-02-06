package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	guuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	ID                string
	User_name         string `json:"user_name"`
	Fixed_longitude   int64  `json:"fixed_longitude"`
	Fixed_latitutde   int64  `json:"fixed_latitutd"`
	User_gender       string `json:"user_gender "`
	Body              string `json:"body"`
	Userscreenname    string `json:"userscreenname"`
	Fixed_location    string `json:"fixed_location"`
	Insight_sentiment string `json:"insight_sentiment"`
	Post_retweetcount int64  `json:"post_retweetcount"`
}

type User struct {
	ID       string
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type Login struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
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
		User_name:         data.User_name,
		Fixed_longitude:   data.Fixed_longitude,
		Fixed_latitutde:   data.Fixed_latitutde,
		User_gender:       data.User_gender,
		Body:              data.Body,
		Userscreenname:    data.Userscreenname,
		Fixed_location:    data.Fixed_location,
		Insight_sentiment: data.Insight_sentiment,
		Post_retweetcount: data.Post_retweetcount,
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(userId string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	// atClaims["exp"] = time.Now().Add(time.M * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func LoginUser(c *gin.Context) {
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

	collection := client.Database("backend").Collection("user")
	var u Login
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	user := User{}
	error := collection.FindOne(context.TODO(), bson.M{"email": u.Email}).Decode(&user)
	if error != nil {
		log.Printf("Wrong Email or Password %v\n", error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Wrong Email or Password",
		})
		return
	}

	match := CheckPasswordHash(u.Password, user.Password)
	if !match {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Wrong Email or Password",
		})
		return
	}

	token, err := CreateToken(u.Email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
	return
}

func getUser(c *gin.Context) {
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

	users := []User{}
	collection := client.Database("backend").Collection("user")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	for cur.Next(context.TODO()) {
		var user User
		cur.Decode(&user)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   users,
	})
	return
}

func createUser(c *gin.Context) {
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

	collection := client.Database("backend").Collection("user")
	var user User
	c.BindJSON(&user)
	id := guuid.New().String()
	password, _ := HashPassword(user.Password)
	newUser := User{
		ID:       id,
		Email:    user.Email,
		Password: password,
		Username: user.Username,
	}

	_, error := collection.InsertOne(context.TODO(), newUser)

	if error != nil {
		log.Printf("Error while inserting new user into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created Successfully",
	})
	return
}

func editUser(c *gin.Context) {
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

	collection := client.Database("backend").Collection("user")

	id := c.Param("id")
	var user User
	c.BindJSON(&user)

	newData := bson.M{
		"$set": bson.M{
			"username": user.Username,
		},
	}

	_, error := collection.UpdateOne(context.TODO(), bson.M{"id": id}, newData)
	if error != nil {
		log.Printf("Error, Reason: %v\n", error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "User Edited Successfully",
	})
	return
}

func deleteUser(c *gin.Context) {
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

	collection := client.Database("backend").Collection("user")
	_, error := collection.DeleteOne(context.TODO(), bson.M{"id": c.Param("id")})
	if error != nil {
		log.Printf("Error while deleting a single user, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User deleted successfully",
	})
	return
}

func getSingleUser(c *gin.Context) {
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

	collection := client.Database("backend").Collection("user")

	id := c.Param("id")

	user := User{}
	error := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)
	if error != nil {
		log.Printf("Error while getting a single User, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
	})
	return
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Cors
	r.Use(cors.Default())
	// User
	r.POST("/login", LoginUser)
	r.GET("/user", TokenAuthMiddleware(), getUser)
	r.GET("/user/:id", TokenAuthMiddleware(), getSingleUser)
	r.POST("/user", createUser)
	r.PUT("/user/:id", TokenAuthMiddleware(), editUser)
	r.DELETE("/user/:id", TokenAuthMiddleware(), deleteUser)

	// Data
	r.GET("/test", test)
	r.GET("/data", TokenAuthMiddleware(), getData)
	r.POST("/data", TokenAuthMiddleware(), createData)
	r.DELETE("/data/:id", TokenAuthMiddleware(), deleteData)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
