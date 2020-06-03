package main

import (
	"fmt"

	"github.com/GitterRemote/gcond/condition"
	"github.com/GitterRemote/gcond/configuration"
)

func main() {
	fmt.Println("vim-go")
	cond := condition.NewTrueCondition()
	g := configuration.NewGroup(configuration.Configuration{1, cond, "1"})
	rv := g.Evaluate(condition.NewContext())
	fmt.Printf("rv = %+v\n", rv)
}
