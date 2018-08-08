package ws

import "time"

var ResourceDao = &DaoResource{}
var ResourceSummaryDao = &DaoResourceSummary{}
var ResourceDownDao = &DaoResourceDown{}

type Resource struct {
	SiteId, ResScore, UserId, Size                int64
	FileName, ResName, ResDesc, ResPath, Username string
	CreateAt                                      time.Time
	Status                                        int
}
type ResourceDown struct {
	Id                         int64 ``
	UserId, ResourceId, SiteId int64
	Score, AuthorScore         int64
	CreateAt                   time.Time
	Username                   string
}
type DaoResource struct {
	GetBySiteIdAndId func(SiteId, userId, Id int64) (map[string]string, bool, error) `select Rate,Id,fmt(CreateAt)CreateAt,ResDesc,ResName,ResScore,FileName from Resource where SiteId=? and UserId=? and Id=? and Status=1`

	FindBySiteIdAndId func(SiteId, Id int64) (Resource, bool, error)

	ListTitleNew func(SiteId, start int64) ([]map[string]string, error) `select ResName,Id,UserId from Resource where SiteId=? and Status=1 order by Id desc limit ?,20`

	Add func(SiteId, UserId, Size int64, FileName, ResPath, Username string) (int64, error) `INSERT INTO Resource(SiteId,UserId,Size,FileName,CreateAt,Status,ResPath,Username)VALUES(?,?,?,?,now(),0,?,?)`

	Up func(ResName, ResDesc string, ResScore, Id, userId int64) (int64, error) `update Resource set Status=1,ResName=?,ResDesc=?,ResScore=? where Id=? and UserId=?`

	List func(SiteId, UserId int64) ([]map[string]string, error) `select Id,fmt(CreateAt)CreateAt,ResDesc,ResName from Resource where SiteId=? and UserId=? and Status=1`
}

type DaoResourceSummary struct {
	Get func(SiteId, UserId int64) (map[string]string, bool, error) `select * from ResourceSummary where SiteId=? and UserId=?`

	UpUploadItem func(UserId, SiteId, UserId2 int64) (int64, error) `update ResourceSummary set UploadItem=(select count(*) from Resource where UserId=? and Status=1) where SiteId=? and UserId=?`

	Add func(SiteId, UserId, UploadItem, DownItem, UpLimit int64) (int64, error) `INSERT INTO ResourceSummary (SiteId,UserId,UploadItem,DownItem,UpLimit) VALUES (?,?,?,?,?)`
}

type DaoResourceDown struct {
	Add func(SiteId, UserId, ResourceId, Score, AuthorScore int64, Username string) (int64, error) `INSERT INTO ResourceDown(SiteId,UserId,ResourceId,Score,AuthorScore,CreateAt,Username)VALUES(?,?,?,?,?,now(),?) `

	FindByUserIdAndResourceId func(userId, ResourceId int64) (ResourceDown, bool, error)

	Exist func(siteId, userId, ResourceId int64) (int, bool, error) `select count(*) from ResourceDown where SiteId=? and UserId=? and ResourceId=?`
}
