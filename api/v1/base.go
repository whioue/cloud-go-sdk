package v1

type MetaID struct {
	ID int64 `form:"id" json:"id" binding:"required"`
}

type MetaUuid struct {
	Uuid string `form:"uuid" json:"uuid"`
}

type MetaName struct {
	Name string `form:"name" json:"name"`
}

type MetaTaskID struct {
	TaskID int64 `form:"task_id" json:"task_id"`
}

type MetaPage struct {
	Page     int `form:"page,default=1" json:"page"`
	PageSize int `form:"page_size,default=10" json:"page_size"`
}

type MetaTotalCnt struct {
	TotalCnt int64 `json:"total_cnt"`
}
