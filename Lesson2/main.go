package main

import (
	"flag"
	"fmt"
	"os"

	un "Lesson2/unpacking"
)

func main() {

	inputFlag := flag.String("input", "", "Строка для (распаковки/упаковки)")
	daemonFlag := flag.Bool("daemon", false, "Запустить в режиме демона")
	packFlag := flag.Bool("pack", false, "Использовать Pack вместо Unpack")
	flag.Parse()

	if *daemonFlag {
		if *packFlag {
			un.RunPackDaemon()
			return
		}
		un.RunUnpackDaemon()
		return
	}

	if *inputFlag == "" {
		fmt.Fprintln(os.Stderr, "Usage: unpacker [-pack] -input=\"ваша_строка\" [-daemon]\n"+
			"  -daemon — режим демона, игнорирует -input\n"+
			"  -pack   — упаковка вместо распаковки")
		os.Exit(1)
	}

	if *packFlag {
		result := un.Pack(*inputFlag)
		fmt.Println(result)
	} else {
		result, err := un.Unpack(*inputFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(result)
	}
}
