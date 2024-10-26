package listener

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Server struct {
	connections map[net.Conn]bool
	wg          sync.WaitGroup
	mu          sync.Mutex
}

func NewServer() *Server {
	return &Server{
		connections: make(map[net.Conn]bool),
	}
}

func (s *Server) Listen(port string) error {
	address := ":" + port
	lit, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("error listening: %w", err)
	}
	defer lit.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("Received interrupt signal. Closing connections...")
		s.closeConnections(lit)
	}()

	for {
		con, err := lit.Accept()
		if err != nil {
			// Check if the error is due to a closed listener, then break the loop
			if opErr, ok := err.(*net.OpError); ok && opErr.Err.Error() == "use of closed network connection" {
				fmt.Println("Listener closed. Stopping accept loop.")
				break
			}
			fmt.Println("Error Accepting: ", err)
			continue
		}

		s.mu.Lock()
		s.connections[con] = true
		s.mu.Unlock()

		s.wg.Add(1)
		go s.handleConnection(con)
	}

	// Wait for all connections to close before exiting
	s.wg.Wait()
	return nil
}

func (s *Server) handleConnection(con net.Conn) {
	defer func() {
		con.Close()
		s.mu.Lock()
		delete(s.connections, con)
		s.mu.Unlock()
		s.wg.Done()
	}()

	buf := make([]byte, 4096)
	for {
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println("Error Reading Buffer: ", err)
			break
		}

		fmt.Printf("%s", buf[:n])
	}
}

func (s *Server) closeConnections(lit net.Listener) {
	s.mu.Lock()
	for conn := range s.connections {
		conn.Close()
	}
	s.mu.Unlock()
	lit.Close()
	s.wg.Wait()
}
