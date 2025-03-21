package netcat3

import (
	"io"
	"log"
	"net"
	"os"
)

func Netcat() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Panicln("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.(*net.TCPConn).CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
