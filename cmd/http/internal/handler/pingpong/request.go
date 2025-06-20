package pingpong

type PingPongRequest struct {
	Message string `json:"message"`
}

type PingPongResponse struct {
	Message string `json:"message"`
}
