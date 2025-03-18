package ex81

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// $env:TZ="Asia/Shanghai" ; ./main.exe -port 8010
func Clock2() {
	nFlag, timezone := handle_arg()
	establish_conn(nFlag, timezone)
}

func handle_arg() (int, string) {
	nFlag := flag.Int("port", 8090, "bind a port")
	flag.Parse()
	timezone, ok := os.LookupEnv("TZ")
	if ok {
		return *nFlag, timezone
	} else {
		return *nFlag, "Asia/Shanghai"
	}
}

func establish_conn(port int, timezone string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, timezone)
	}
}

func handleConn(c net.Conn, timezone string) {
	defer c.Close()
	for {
		location, err := time.LoadLocation(timezone)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.WriteString(c, fmt.Sprintf("timezone:%s,time:%s", timezone, time.Now().In(location).Format(fmt.Sprintf("%s\n", time.RFC3339))))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
