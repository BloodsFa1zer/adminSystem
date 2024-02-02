package image

import (
	"gormTest/model"
)

type ImageDbInterface interface {
	//FindByID(ID int64) (*User, error)
	InsertImage(img model.Image) error
	SelectSortImages(options model.SortOptions) (*[]model.Image, error)
	GetImageByPublicID(publicID string) (*model.Image, error)
	UpdateImage(img model.Image) (uint, error)
	//UpdateUser(ID int64, user User) (int64, error)
	//FindUsers() (*[]User, error)
	//DeleteUserByID(ID int64) error
	//FindByNicknameToGetUserPassword(nickname string) (*User, error)
	//WriteUserVotes(userID, voterID, voteValue int) error
	//WithdrawVote(userID, voterID int) error
	//ChangeVote(userID, voterID int) error
	//IsSuchVoteExists(userID, voterID int) (bool, error)
	//GetUserLastVoteTime(voterID int) (string, error)
}
