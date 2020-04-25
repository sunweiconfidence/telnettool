package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	host := flag.String("host", "localhost", "host")
	port := flag.String("port", ":80", "http listen port")
	flag.Parse()
	dialaddr := strings.Join([]string{*host, *port}, "")
	log.Println(dialaddr)
	conn, err := net.Dial("tcp", dialaddr)
	if err != nil {
		log.Println("connected failed")
		log.Fatal(err)
	} else {
		log.Println("connected")
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
