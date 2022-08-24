package training

import (
	"context"
	"io"
	"time"
)

type Server struct {
	conn   io.WriteCloser
	ctx    context.Context
	cancel context.CancelFunc
}

func (s *Server) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	s.ctx, s.cancel = ctx, cancel

	go s.sendPing(ctx)
	go s.sendPong(ctx)
	go s.supervise(ctx)
}

func (s *Server) sendPing(ctx context.Context) {
	tick := time.NewTicker(2 * time.Second)
	for range tick.C {
		select {
		case <-s.ctx.Done():
			return
		default:
		}
		if _, err := s.conn.Write([]byte("Ping")); err != nil {
			s.cancel()
			return
		}
	}
}

func (s *Server) sendPong(ctx context.Context) {
	tick := time.NewTicker(1320 * time.Millisecond)
	for range tick.C {
		select {
		case <-s.ctx.Done():
			return
		default:
		}
		if _, err := s.conn.Write([]byte("Pong")); err != nil {
			s.cancel()
			return
		}
	}
}

func (s *Server) Stop() {
	s.cancel()
}

func (s *Server) supervise(ctx context.Context) {
	<-s.ctx.Done()
	_ = s.conn.Close()
}
