package api

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	fb "github.com/huandu/facebook"
	"github.com/labstack/echo"
	"io"
	"makeToon/database"
	"makeToon/handler"
	"net/http"
	"os"
	"strings"
)

var users = map[string]map[string]string {}

type Error struct {
	Message string
	Type string
	Code string
	ErrorSubcode string
	TraceID string
}

func getFBdata(token string) (string, error) {
	fbData, fbErr := fb.Get("/me", fb.Params{
		"access_token": token,
		"fields": "id",
	})

	if fbErr != nil {
		if e, ok := fbErr.(*fb.Error); ok {
			switch e.Type {
			case "OAuthException":
				return "token expired", errors.New(e.Type)
			}
		}
	}

	userId := fbData.GetField("id")

	return userId.(string), fbErr
}

func PutCropPhoto(context echo.Context) error {
	area := context.FormValue("area")
	width := context.FormValue("width")
	height := context.FormValue("height")
	token := context.FormValue("token")
	photo, err := context.FormFile("photo")

	userId, fbErr := getFBdata(token)

	if fbErr != nil {
		fbErrString := fbErr.Error()

		if errorCode := handler.HandleError(fbErrString); errorCode == 401 {
			return context.NoContent(http.StatusUnauthorized)
		}
	}

	if err != nil {
		return err
	}
	source, fileOpenErr := photo.Open()
	if fileOpenErr != nil {
		return fileOpenErr
	}
	defer source.Close()

	uuidRes, uuidErr := uuid.NewV4()
	if uuidErr != nil {
		panic(uuidErr)
	}

	file, _ := os.Open(photo.Filename)

	// dummy 이미지 생성
	fileExtension := strings.Split(photo.Filename, ".")
	fileName := uuidRes.String() + "." + fileExtension[len(fileExtension) - 1]
	newFile, createPhotoErr := os.Create(fileName)
	if createPhotoErr != nil {
		return createPhotoErr
	}
	defer newFile.Close()

	if _, err := io.Copy(newFile, source); err != nil {
		return err
	}

	file, openFileErr := os.Open(fileName)
	if openFileErr != nil {
		return openFileErr
	}

	defer file.Close()

	// aws 업로드
	awsResponse, awsErr := handler.FileUploadTos3(fileName, file)
	if awsErr != nil {
		fmt.Printf("failed to upload file, %v", awsErr)
		return context.NoContent(http.StatusBadRequest)
	}

	// DB에 저장
	database.SetPhoto(userId, area, awsResponse.Location, width, height)

	// dummy 이미지 제거
	_ = os.Remove(fileName)

	return context.NoContent(http.StatusNoContent)
}

func GetMapPhotos(context echo.Context) error {
	token := context.QueryParam("token")
	userId, fbErr := getFBdata(token)

	if fbErr != nil {
		fbErrString := fbErr.Error()
		if errorCode := handler.HandleError(fbErrString); errorCode == 401 {
			return context.NoContent(http.StatusUnauthorized)
		}
	}

	areas := database.GetFunc(userId)

	return context.JSON(http.StatusOK, areas)
}
