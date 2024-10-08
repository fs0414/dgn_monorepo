package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
  listener, err := net.Listen("tcp", ":9091"); if err != nil {
    panic(err)
  }
  for {
    conn, err := listener.Accept(); if err != nil {
      panic(err)
    }
    go func() {
      defer conn.Close()
      for {
        conn.SetReadDeadline(time.Now().Add(5 * time.Second))
      }
      request, err := http.ReadRequest(bufio.NewReader(conn)); if err != nil {
        panic(err)
      }

      dump, err := httputil.DumpRequest(request, true); if err != nil {
        panic(err)
      }

      fmt.Println(string(dump))

      res := http.Response{
        StatusCode: 200,
        ProtoMajor: 1,
        ProtoMinor: 1,
        Body:       io.NopCloser(
          strings.NewReader("Hello, World!"),
        ),
      }
      res.Write(conn)
      conn.Close()
    }()
  }
}
