package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/krageon/php/ast/printer"
	"github.com/krageon/php/parser"
)

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {

		fmt.Println(arg)
		fmt.Println()

		src, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		p := printer.NewPrinter(os.Stdout)
		file, err := parser.NewParser().Parse("test.php", string(src))
		if err != nil {
			log.Fatal(err)
		}
		for _, node := range file.Nodes {
			p.PrintNode(node)
		}
	}
}
