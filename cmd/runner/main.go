package main

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goConstants"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goFor"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goIfElse"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goValues"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goVariables"
)

func main() {
	fmt.Println("GoByExample code-along")
	fmt.Println("------------")

	goValues.Values()
	goVariables.Variables()
	goConstants.Constants()
	goFor.ForLoop()
	goIfElse.IfElse()
}
