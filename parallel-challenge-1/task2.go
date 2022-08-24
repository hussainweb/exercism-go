package training

import (
	"io"
)

type Server struct {
	conn io.WriteCloser
}

func (s *Server) Start() {

}

func (s *Server) sendPing() {

}

func (s *Server) Stop() {
}
