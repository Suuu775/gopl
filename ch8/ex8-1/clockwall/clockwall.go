package clockwall

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func ClockWall() {
	args := os.Args[1:]
	tz_port_map := make(map[string]int)

	for _, v := range args {
		str_array := strings.Split(v, "=")
		port_str := strings.Split(str_array[1], ":")[1]
		port, err := strconv.Atoi(port_str)
		if err != nil {
			log.Fatal(err)
		}
		tz_port_map[str_array[0]] = port
	}
	for tz, port := range tz_port_map {
		establish_conn(port, tz)
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
