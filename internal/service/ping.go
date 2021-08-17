package service

type PingService interface {
	Ping() string
}
type pingService struct {
}

func NewPingService() PingService {
	return &pingService{}
}

func (s *pingService) Ping() string {
	return "pong"
}
