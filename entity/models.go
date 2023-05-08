package entity

type SvcAndVersion struct {
	SvcName   string `json:"svcName"`
	Version   string `json:"version"`
	Namespace string `json:"namespace"`
}

type UpdateTime struct {
	Timestamp int64
}
