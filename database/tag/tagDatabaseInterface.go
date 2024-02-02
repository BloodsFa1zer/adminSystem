package tag

import "gormTest/model"

type TagDbInterface interface {
	//FindByID(ID int64) (*User, error)
	InsertTag(tag model.Tag) (string, error)
	SelectSortTags(options model.SortOptions) (*[]model.Tag, error)
	GetTagByPublicID(publicID string) (*model.Tag, error)
	UpdateTag(tag model.Tag) error
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
