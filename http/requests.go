package http

type RunCommandRequest struct {
	DeviceId      int `json:"DeviceId"`
	ApplicationId int `json:"ApplicationId"`
	SequenceId    int `json:"SequenceId"`
}
