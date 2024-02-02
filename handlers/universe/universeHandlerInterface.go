package universe

import "github.com/labstack/echo/v4"

type TagHandlersInterface interface {
	//EditUser(c echo.Context) error
	//DeleteUser(c echo.Context) error
	//CreateUser(c echo.Context) error
	//GetUser(c echo.Context) error
	//GetAllUsers(c echo.Context) error
	//Login(c echo.Context) error
	//isUserHavePermissionToActions(roleToFind string, c echo.Context) (bool, int)
	//PostVoteFor(c echo.Context) error
	//PostVoteAgainst(c echo.Context) error
	//DeleteVote(c echo.Context) error
	//ChangeVote(c echo.Context) error

	UniversePOST(c echo.Context) error
	UniverseGET(c echo.Context) error
	UniverseGetByID(c echo.Context) error
	UniversePUT(c echo.Context) error
}
