package main

import (
	"bytes"
	"net"
)


type Server struct {
	port int
}
type Client struct {
}
type Queue struct {
}

func initServer(port int) *Server {
	return &Server{
		port: port,
	}
}
func (s *Server) run() error {
	_, err := net.Listen("tcp", string(s.port)); if err != nil {
		return err
	}
	return  nil
}

func main() {
	if err := initServer(8080).run(); err != nil {
		panic(err)
	}
}