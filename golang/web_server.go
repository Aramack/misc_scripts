package main

import (
  "net"
)


func main() {
  bind, _ := net.Listen("tcp", ":8080")
  for {
    conn, _ := bind.Accept()
    go handle_request(conn)
  }
}

func handle_request(connection net.Conn) {
  request_buffer := make([]byte, 1024)
  connection.Read(request_buffer)
  connection.Write([]byte("hello internet"))
  connection.Close()
}