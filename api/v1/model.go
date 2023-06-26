package v1

import "time"

type ModelDeleteArgs struct{}

type ModelDownloadArgs struct {
	MetaUuid
	Catalog   string `form:"catalog" json:"catalog" binding:"required"`
	ModelName string `form:"model_name" json:"model_name" binding:"required"`
}

type ModelDownloadReply struct {
	ModelPath string `json:"model_path"`
}

type ModelDetailReply struct {
	FileName   string    `json:"file_name"`
	CreateTime time.Time `json:"create_time"`
	Size       int64     `json:"size"`
}

type ModelListArgs struct {
	MetaPage
	MetaUuid
	FileName string `form:"file_name" json:"file_name"`
	IsAsc    bool   `form:"is_asc" json:"is_asc"`
	Catalog  string `form:"catalog" json:"catalog" binding:"required"`
}

type ModelListReply struct {
	MetaPage
	MetaTotalCnt
	List []*ModelDetailReply `json:"list"`
}

type ModelCatalogsDetailReply struct {
	DirName string `json:"dir_name"`
}

type ModelCatalogsArgs struct {
	MetaPage
	MetaUuid
}

type ModelCatalogsReply struct {
	MetaPage
	MetaTotalCnt
	List []*ModelCatalogsDetailReply `json:"list"`
}
