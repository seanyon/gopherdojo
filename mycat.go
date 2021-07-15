package main

import (
	"flag"
	"fmt"
)

func main() {
	var n =flag.Bool("n",false,"通し番号を付与する")
	flag.Parse()
	var files = flag.Args()

	for i:=0; i<len(files); i++ {
		if *n{//*オプションがあれば
			fmt.Printf("%v: ",i+1)
		}
		fmt.Println(files[i])
	}

}