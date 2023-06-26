package v1

type TaskCreateArgs struct {
	MetaName
	UserName     string            `form:"user_name" json:"user_name" binding:"required"`
	Label        string            `form:"label" json:"label" binding:"required"`
	ConfigPath   string            `form:"config_path" json:"config_path" binding:"required"`
	ProjectPath  string            `form:"project_path" json:"project_path" binding:"required"`
	ImageVersion string            `form:"image_version" json:"image_version" binding:"required"`
	CommonImage  Image             `form:"common_image" json:"common_image" binding:"required"`
	ActorImage   Image             `form:"actor_image" json:"actor_image" binding:"required"`
	Actor        *TaskBaseResource `form:"actor" json:"actor" binding:"required"`
	Predictor    *TaskBaseResource `form:"predictor" json:"predictor" binding:"required"`
	Learner      *TaskBaseResource `form:"learner" json:"learner" binding:"required"`
}

type TaskCreateReply struct {
	MetaUuid
}

type TaskDeleteArgs struct {
	MetaUuid
}

type TaskDetailArgs struct {
	MetaUuid
}

type TaskDetailReply struct {
	Uuid           string `json:"uuid"`
	Name           string `json:"name"`
	Status         string `json:"status"`
	TrainingTime   int64  `json:"training_time"`
	Cpu            int64  `json:"cpu"`
	Gpu            int64  `json:"gpu"`
	Memory         int64  `json:"memory"`
	TensorboardUrl string `json:"tensorboard_url"`
	CommonImage    string `json:"commonImage,omitempty"`
	ActorImage     string `json:"actorImage,omitempty"`
}

type TaskListArgs struct {
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"page_size" json:"page_size"`
	Status   string `form:"status" json:"status"`
	Name     string `form:"name" json:"name"`
}

type TaskListReply struct {
	MetaPage
	MetaTotalCnt
	List []*TaskDetailReply `json:"list"`
}

type TaskUpdateArgs struct {
	MetaUuid
	MetaName
	CommonImage string `form:"common_image" json:"common_image"`
	ActorImage  string `form:"actor_image" json:"actor_image"`
}

type TaskStartArgs struct {
	MetaUuid
}

type TaskStopArgs struct {
	MetaUuid
	DeleteTensorboard bool
}

type TaskPodDetailReply struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type TaskPodListArgs struct {
	MetaPage
	MetaUuid
}

type TaskPodListReply struct {
	MetaPage
	MetaTotalCnt
	List []*TaskPodDetailReply `json:"list"`
}

type TaskBaseResource struct {
	Number int   `form:"number" json:"number"`
	Cpu    int64 `form:"cpu" json:"cpu"`
	Gpu    int64 `form:"gpu" json:"gpu"`
	Memory int64 `form:"memory" json:"memory"`
}

type Image struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}
