package main

import "net"

type Server struct {
    listenAddr string 
    ln         net.Listener
    quitCh     chan struct{} 
}

func NewServer(listenAddr string) *Server {
    return &Server{
        listenAddr: listenAddr,
        quitCh: make(chan struct{}),
    }
}

func (s *Server) Start() error {
    ln, err := net.Listen("tcp", "8000")
    if err != nil {
        return err
    }
    defer ln.Close()
    s.ln = ln

    <-s.quitCh
    return nil
}

func main() {
    print("yppp")
}