package v1

type TaskResourceStartArgs struct {
	ResourceID int64 `form:"resource_id" json:"resource_id"`
}

type TaskResourceStartReply struct {
	ResourceID  int64  `json:"resource_id"`
	ResourceUrl string `json:"resource_url"`
}

type TaskResourceRestartArgs struct {
	ResourceID int64 `form:"resource_id" json:"resource_id"`
}

type TaskResourceStopArgs struct {
	ResourceID int64 `form:"resource_id" json:"resource_id"`
}

type TaskResourceDetailArgs struct {
	ResourceID int64 `form:"resource_id" json:"resource_id"`
}

type TaskResourceDetailReply struct {
	ResourceID     int64  `json:"resource_id"`
	ResourceName   string `json:"resource_name"`
	ResourceStatus string `json:"resource_status"`
	ResourceUrl    string `json:"resource_url"`
}

type TaskResourceTensorboardDetailArgs struct {
	Address string `form:"address" json:"address"`
	Catalog string `form:"catalog" json:"catalog"`
	MetaUuid
	Labels []string `form:"labels" json:"labels"`
}

type TaskResourceTensorboardDetailReply struct {
	LabelResult map[string][][]float64 `json:"label_result"`
	ModelList   []string               `json:"model_list"`
}
