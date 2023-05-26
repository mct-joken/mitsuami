package main

import (
	"fmt"
	"github.com/mct-joken/mitsuami/server"
)

func main() {
	fmt.Println("Mitsuami - equipment & book management tool\n(C) 2023 Tatsuto Yamamoto")

	server.Start()
}
