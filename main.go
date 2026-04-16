package main

import (
	"context"
	"fmt"
	"os"
	"socialProject_backend/httpserver"

	"github.com/jackc/pgx/v5"
)

func main() {
	parents, parents_close := context.WithCancel(context.Background())
	postgres, _ := context.WithCancel(parents)
	defer parents_close()

	fmt.Println("Starting server")

	conn_url := os.Getenv("CONN_STRING")
	conn, err := pgx.Connect(postgres, conn_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error create conn: %v\n", err)
		return
	}

	if er := conn.Ping(postgres); er != nil {
		fmt.Fprintf(os.Stderr, "Error ping: %v\n", err)
		return
	}

	handler := httpserver.CreateHTTPHandler(conn, parents)
	server := httpserver.CreateServer(handler)

	server.StartHTTPServer()
}
