package main

import (
	"flag"
	"fmt"
	"github.com/zhekaby/go-generators/common"
	"os"
)

func main() {
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, file := range files {
		p := common.NewParser(file)
		if err := p.Parse(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		w := NewWriter(p)
		w.Write()
	}
}
