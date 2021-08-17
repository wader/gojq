// gojq - Go implementation of jq
package main

import (
	"os"

	"github.com/wader/gojq/cli"
)

func main() {
	os.Exit(cli.Run())
}
