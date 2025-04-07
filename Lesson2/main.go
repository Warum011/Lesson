package main

import (
	un "Lesson2/unpacking"
	"flag"
	"fmt"
	"os"
)

func main() {

	inputFlag := flag.String("input", "", "Строка для распаковки")
	daemonFlag := flag.Bool("daemon", false, "Запустить в режиме демона")
	flag.Parse()

	if *daemonFlag {
		un.RunDaemon()
	} else if *inputFlag != "" {
		result, err := un.Unpack(*inputFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(result)
	} else {
		fmt.Fprintln(os.Stderr, "Необходимо указать --input или --daemon флаг")
		flag.Usage()
		os.Exit(1)
	}

}
