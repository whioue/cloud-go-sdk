package v1

import "time"

type LogDeleteArgs struct {
	LogName string `form:"log_name" json:"log_name" binding:"required"`
}

type LogDownloadArgs struct {
	LogName string `form:"log_name" json:"log_name" binding:"required"`
	MetaUuid
}

type LogDownloadReply struct {
	LogPath string `form:"log_path" json:"log_path"`
}

type LogListArgs struct {
	MetaPage
	MetaUuid
}

type LogListReply struct {
	MetaPage
	MetaTotalCnt
	List []*LogDetailReply `json:"list"`
}

type LogDetailArgs struct {
	LogName string `form:"log_name" json:"log_name" binding:"required"`
	MetaUuid
}

type LogDetailReply struct {
	FileName   string    `json:"file_name"`
	CreateTime time.Time `json:"create_time"`
	Size       int64     `json:"size"`
}

type LogContentReply struct {
	Content string `json:"content"`
}
