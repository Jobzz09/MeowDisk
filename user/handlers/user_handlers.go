package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Jobzz09/MeowDisk/models"
	"github.com/Jobzz09/MeowDisk/user/usecase"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

type UserHandlers struct {
	m_userUseCase usecase.UserUseCase
}

func NewUserHandlers(db *sql.DB, redis *redis.Client) UserHandlers {
	t_userUsecase := usecase.NewUserUseCase(db, redis)
	return UserHandlers{m_userUseCase: t_userUsecase}
}

func (userH UserHandlers) Register(ctx echo.Context) error {
	t_userData := models.UserData{}

	readed, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Fatal("Bad request body at register")
		return err
	}

	err = json.Unmarshal(readed, &t_userData)
	if err != nil {
		log.Fatal("Error at unmarshalling request at Register")
		return err
	}

	err = userH.m_userUseCase.Register(t_userData)
	if err != nil {
		log.Fatal("Error at register with user_usecase")
		return err
	}

	return ctx.NoContent(http.StatusOK)

	//TODO add cookie handle
}

func (userH UserHandlers) Login(ctx echo.Context) error {
	var t_userData models.UserData

	readed, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Fatal("Bad request body at Login")
		return err
	}

	err = json.Unmarshal(readed, &t_userData)
	if err != nil {
		log.Fatal("Error at unmarshalling request at Login")
		return err
	}

	err = userH.m_userUseCase.Login(t_userData)
	if err != nil {
		log.Fatal("Error at login with user_usecase")
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	return ctx.NoContent(http.StatusOK)

}

func (userH UserHandlers) Logout(ctx echo.Context) error {
	return ctx.Redirect(http.StatusPermanentRedirect, "google.com")
}

func (userH UserHandlers) Upload(ctx echo.Context) error {
	// Max upload of 10 mb files

	ctx.Request().ParseMultipartForm(10 << 20)

	file, handler, err := ctx.Request().FormFile("file")

	if err != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, err)
	}

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("/tmp/mewdisk", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Println("Successfully Uploaded File")
	return ctx.JSON(http.StatusOK, nil)
}

func (userH UserHandlers) InitHandlers(server *echo.Echo) {
	server.GET("/login", userH.Login)
	server.POST("/register", userH.Register)
	server.POST("/upload", userH.Upload)
	server.DELETE("/logout", userH.Logout)
}

// Marshall Unmarshall
// Coder Decoder
//Json golang
