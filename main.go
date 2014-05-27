package main

import (
    "net"
    "io"
    "fmt"
)

func main() {
    socket, _ := net.Listen("tcp", ":8080")
    for {
        client, _ := socket.Accept()
        go func() {
            fmt.Println(client.RemoteAddr())
            io.Copy(client, client)
            client.Close()
        }()
    }
}
