package services

import "log"


type GreeterService interface {
	SayHello(name string) (string, error)
}

type greeterService struct {
}

func NewGreeterService() GreeterService {
	return &greeterService{}
}

func (s *greeterService) SayHello(name string) (string, error) {
	log.Println("GreeterService", name)
	return "GreeterService", nil
}