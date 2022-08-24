# Challenge

Some of our mentors thought it would be really cool to throw a little challenge around concurrency into the mix!  Consider this some homework that we'll work through in a livestream...

## Task 1

Execute the given tasks in parallel and cancel all others if one returns with an error. If all returned boolean values are true, return true, otherwise false.

```go
func RunParallel(ctx context.Context, tasks ...func() (bool, error)) (bool, error) {
    panic("not implemented")
} 
```

## Task 2

Create a Server object that uses a given conn to regularly send out Pings and Pangs once Started.

Pings are sent every 2sec, Pongs every 1320ms.

It should stop if either the connection is closed (Write returns with error) or Stop is called. The connection must also be closed on Stop.

```go
type Server struct {
    conn io.WriteCloser
}

func (s *Server) Start() {
    panic("not implemented")
}

func (s *Server) Stop() {
    panic("not implemented")
}
```
