package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
)

var (
	MongoSession    *mgo.Session
	UsersCollection *mgo.Collection
)

type User struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Username  string `json:"userName"`
	Password  string `json:"password,omitempty"`
}

func (u *User)SaveToDB() error {
	err := UsersCollection.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User)ReadFromDB() ([]User, error) {
	result := []User{}
	err := UsersCollection.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello world")
}

func init() {
	MongoSession, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	MongoSession.SetMode(mgo.Monotonic, true)
	UsersCollection = MongoSession.DB("maejo").C("users")
}

func getUsers(c echo.Context) error {
	user := new(User)
	result, _ := user.ReadFromDB()
	return c.JSON(http.StatusOK, result)
}

func getUserByID(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user.SaveToDB()
	return c.NoContent(http.StatusCreated)
}

func main() {
	defer MongoSession.Close()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.GET("/", index)
	e.GET("/users", getUsers)
	e.GET("/user:id", getUserByID)
	e.POST("/users", saveUser)
	e.Logger.Fatal(e.Start(":8090"))
}
