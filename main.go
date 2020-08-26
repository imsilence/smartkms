package main

import (
	"github.com/imsilence/smartkms/cmds"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	cmds.Execute()
}
