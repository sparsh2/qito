package main

import (
	"fmt"

	"github.com/sparsh2/pmgr/cmd"
	"github.com/sparsh2/pmgr/common"
)

func main() {
	fmt.Println(common.SecretKey)
	cmd.Execute()
}
