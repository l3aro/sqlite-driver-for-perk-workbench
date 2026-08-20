package main

import (
	"fmt"
	"os"

	"github.com/l3aro/perk-workbench-plugin-sdk-go/server"
	"github.com/l3aro/sqlite-driver-for-perk-workbench/internal/drivers/sqlite"
)

func main() {
	if err := server.Run(os.Stdin, os.Stdout, sqlite.Factory{}); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
