package universe

import (
	"github.com/labstack/echo/v4"
	"gormTest/config"
	"gormTest/model"
	"gormTest/service/universe"
	response "gormTest/userResponse"
	"net/http"
)

type UniverseHandler struct {
	UniverseInterface universe.UniverseServiceInterface
}

func NewUniverseHandler(universeInterface universe.UniverseServiceInterface) *UniverseHandler {
	return &UniverseHandler{UniverseInterface: universeInterface}
}

func (uh *UniverseHandler) UniversePOST(c echo.Context) error {
	// Check authentication (assuming authentication is implemented)
	//if !userIsAdmin(c) {
	//	return c.JSON(http.StatusUnauthorized, response.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: &echo.Map{"msg": "Unauthorized"}})
	//}

	// Get the image file and filename from the request
	var universe model.Universe
	if err := c.Bind(&universe); err != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"msg": "Invalid request body"}})
	}
	//TODO: Create a new image instance in the DB and write file name there

	// Generate a unique path
	ID := config.GeneratePublicID()

	publicID, err, httpStatus := uh.UniverseInterface.CreateUniverse(universe.Title, ID)
	if err != nil {
		return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "error", Data: &echo.Map{"msg": "Invalid name"}})
	}
	// Save image information to the database (pseudo code)
	//saveImageInfoToDatabase(filename, publicID)

	// Respond with a successful status
	return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "success", Data: &echo.Map{"tagPublicId": publicID}})
}

func (uh *UniverseHandler) UniverseGET(c echo.Context) error {

	// Check authentication (assuming authentication is implemented)
	//if !userIsAdmin(c) {
	//	return c.JSON(http.StatusUnauthorized, response.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: &echo.Map{"msg": "Unauthorized"}})
	//}

	var sortOptions model.SortOptions
	if err := c.Bind(&sortOptions); err != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"msg": "Invalid request body"}})
	}

	tags, err, httpStatus := uh.UniverseInterface.SortUniverse(sortOptions)
	if err != nil {
		return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "error", Data: &echo.Map{"msg": "Invalid request body"}})
	}

	return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "success", Data: &echo.Map{"tag": tags, "offset": sortOptions.Offset}})
}

func (uh *UniverseHandler) UniverseGetByID(c echo.Context) error {
	ID := c.Param("id")

	tag, err, httpStatus := uh.UniverseInterface.GetUniverseByID(ID)
	if err != nil {
		return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "success", Data: &echo.Map{"data": tag}})
}

func (uh *UniverseHandler) UniversePUT(c echo.Context) error {
	//if permission, respStatus := uh.isUserHavePermissionToActions(adminRole, c); !permission {
	//	return c.JSON(respStatus, response.UserResponse{Status: respStatus, Message: "error", Data: &echo.Map{"data": "that user has no access to admin actions"}})
	//}

	ID := c.Param("id")

	var newFilename model.ImageUpdate
	if err := c.Bind(&newFilename); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	imageID, err, respStatus := uh.UniverseInterface.EditUniverse(ID, newFilename.Filename)
	if err != nil {
		return c.JSON(respStatus, response.UserResponse{Status: respStatus, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(respStatus, response.UserResponse{Status: respStatus, Message: "success", Data: &echo.Map{"imagePublicId": imageID}})
}

// todo: implement 4.1 Delete
