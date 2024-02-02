package image

import (
	"github.com/labstack/echo/v4"
	"gormTest/model"
	"gormTest/service/image"
	response "gormTest/userResponse"
	"net/http"
)

type ImageHandler struct {
	imageService image.ImageServiceInterface
}

func NewImageHandler(service image.ImageServiceInterface) *ImageHandler {
	return &ImageHandler{imageService: service}
}

// TODO: create a func to generate json response

func (ih *ImageHandler) ImagePOST(c echo.Context) error {
	// Check authentication (assuming authentication is implemented)
	//if !userIsAdmin(c) {
	//	return c.JSON(http.StatusUnauthorized, response.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: &echo.Map{"msg": "Unauthorized"}})
	//}

	// Get the image file and filename from the request
	image, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"msg": "Invalid request"}})
	}
	filename := c.FormValue("filename")

	publicID, err, httpStatus := ih.imageService.CreateImage(filename, image)
	if err != nil {
		return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "error", Data: &echo.Map{"msg": "Invalid name"}})
	}
	// Save image information to the database (pseudo code)
	//saveImageInfoToDatabase(filename, publicID)

	// Respond with a successful status
	return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "success", Data: &echo.Map{"imagePublicId": publicID}})
}

func (ih *ImageHandler) ImagesGET(c echo.Context) error {

	// Check authentication (assuming authentication is implemented)
	//if !userIsAdmin(c) {
	//	return c.JSON(http.StatusUnauthorized, response.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: &echo.Map{"msg": "Unauthorized"}})
	//}

	var sortOptions model.SortOptions
	if err := c.Bind(&sortOptions); err != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"msg": "Invalid request body"}})
	}

	images, err, httpStatus := ih.imageService.SortImages(sortOptions)
	if err != nil {
		return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "error", Data: &echo.Map{"msg": "Invalid request body"}})
	}

	return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "success", Data: &echo.Map{"tags": images, "offset": sortOptions.Offset}})
}

func (ih *ImageHandler) ImageGetByID(c echo.Context) error {
	ID := c.Param("id")
	//userID, err := strconv.Atoi(ID)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	//}

	image, err, httpStatus := ih.imageService.GetImageByID(ID)
	if err != nil {
		return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "success", Data: &echo.Map{"data": image}})
}

func (ih *ImageHandler) ImagePUT(c echo.Context) error {
	//if permission, respStatus := uh.isUserHavePermissionToActions(adminRole, c); !permission {
	//	return c.JSON(respStatus, response.UserResponse{Status: respStatus, Message: "error", Data: &echo.Map{"data": "that user has no access to admin actions"}})
	//}

	ID := c.Param("id")

	var newFilename model.ImageUpdate
	if err := c.Bind(&newFilename); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	imageID, err, respStatus := ih.imageService.EditImage(ID, newFilename.Filename)
	if err != nil {
		return c.JSON(respStatus, response.UserResponse{Status: respStatus, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(respStatus, response.UserResponse{Status: respStatus, Message: "success", Data: &echo.Map{"imagePublicId": imageID}})
}

// todo: implement 4.1 Delete
