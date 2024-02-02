package tag

import (
	"github.com/labstack/echo/v4"
	"gormTest/config"
	"gormTest/model"
	"gormTest/service/tag"
	response "gormTest/userResponse"
	"net/http"
)

type TagHandler struct {
	TagInterface tag.TagServiceInterface
}

func NewTagHandler(tagService tag.TagServiceInterface) *TagHandler {
	return &TagHandler{TagInterface: tagService}
}

func (th *TagHandler) TagPOST(c echo.Context) error {
	// Check authentication (assuming authentication is implemented)
	//if !userIsAdmin(c) {
	//	return c.JSON(http.StatusUnauthorized, response.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: &echo.Map{"msg": "Unauthorized"}})
	//}

	// Get the image file and filename from the request
	var tag model.TagInput
	if err := c.Bind(&tag); err != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"msg": "Invalid request body"}})
	}
	//TODO: Create a new image instance in the DB and write file name there

	// Generate a unique path
	ID := config.GeneratePublicID()

	publicID, err, httpStatus := th.TagInterface.CreateTag(tag.Title, ID)
	if err != nil {
		return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "error", Data: &echo.Map{"msg": "Invalid name"}})
	}
	// Save image information to the database (pseudo code)
	//saveImageInfoToDatabase(filename, publicID)

	// Respond with a successful status
	return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "success", Data: &echo.Map{"tagPublicId": publicID}})
}

//func handleImageUpload(c echo.Context) error {
//	// Get the image file from the request
//	image, err := c.FormFile("image")
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "Invalid request"})
//	}
//
//
//	imagePath := filepath.Join("static/", generatePublicID())
//	if err := saveImage(image, imagePath); err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Internal Server Error"})
//	}
//
//	// Respond with the URL to the uploaded image
//	imageURL := fmt.Sprintf("/Users/mishashevnuk/GolandProjects/gormTest/server/static/%s", generatePublicID())
//	return c.JSON(http.StatusCreated, map[string]string{"msg": "Image uploaded successfully", "imageURL": imageURL})
//}

func (th *TagHandler) TagGET(c echo.Context) error {

	// Check authentication (assuming authentication is implemented)
	//if !userIsAdmin(c) {
	//	return c.JSON(http.StatusUnauthorized, response.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: &echo.Map{"msg": "Unauthorized"}})
	//}

	var sortOptions model.SortOptions
	if err := c.Bind(&sortOptions); err != nil {
		return c.JSON(http.StatusBadRequest, response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"msg": "Invalid request body"}})
	}

	tags, err, httpStatus := th.TagInterface.SortTags(sortOptions)
	if err != nil {
		return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "error", Data: &echo.Map{"msg": "Invalid request body"}})
	}

	return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "success", Data: &echo.Map{"tag": tags, "offset": sortOptions.Offset}})
}

func (th *TagHandler) TagGetByID(c echo.Context) error {
	ID := c.Param("id")
	//userID, err := strconv.Atoi(ID)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	//}

	tag, err, httpStatus := th.TagInterface.GetTagByID(ID)
	if err != nil {
		return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(httpStatus, response.UserResponse{Status: httpStatus, Message: "success", Data: &echo.Map{"data": tag}})
}

func (th *TagHandler) TagPUT(c echo.Context) error {
	//if permission, respStatus := uh.isUserHavePermissionToActions(adminRole, c); !permission {
	//	return c.JSON(respStatus, response.UserResponse{Status: respStatus, Message: "error", Data: &echo.Map{"data": "that user has no access to admin actions"}})
	//}

	ID := c.Param("id")

	var newFilename model.ImageUpdate
	if err := c.Bind(&newFilename); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	imageID, err, respStatus := th.TagInterface.EditTag(ID, newFilename.Filename)
	if err != nil {
		return c.JSON(respStatus, response.UserResponse{Status: respStatus, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(respStatus, response.UserResponse{Status: respStatus, Message: "success", Data: &echo.Map{"imagePublicId": imageID}})
}

// todo: implement 4.1 Delete
