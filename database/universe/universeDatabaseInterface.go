package universe

import "gormTest/model"

type TagDbInterface interface {
	//FindByID(ID int64) (*User, error)
	InsertUniverse(universe model.Universe) (string, error)
	SelectSortUniverse(options model.SortOptions) (*[]model.Universe, error)
	GetUniverseByPublicID(publicID string) (*model.Universe, error)
	UpdateUniverse(universe model.Universe) error
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
