package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"net/textproto"
)

func connect() net.Conn {
	conn, err := net.Dial("tcp", "irc.jpatters.com:6697")
	if err != nil {
		panic(err)
	}

	return conn
}

func disconnect(conn net.Conn) {
	conn.Close()
}

func send(conn net.Conn, message string) {
	fmt.Fprintf(conn, "%s\r\n", message)
}

func main() {
	slog.Info("Hello, World!")

	conn := connect()

	slog.Info("Connected to server", "conn", conn, "addr", conn.RemoteAddr(), "localAddr", conn.LocalAddr())

	send(conn, "USER Test 8 * :Someone")

	tp := textproto.NewReader(bufio.NewReader(conn))

	for {
		line, err := tp.ReadLine()
		if err != nil {
			slog.Error("Error reading line", "err", err)
			if err.Error() == "EOF" {
				break
			}
		}
		fmt.Println(line)
	}

	defer disconnect(conn)

}
