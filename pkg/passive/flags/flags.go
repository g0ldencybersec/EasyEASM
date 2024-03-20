package flags

import (
	"flag"
	"fmt"
)

func ParsingFlags() string {
	interactive := flag.Bool("i", false, "interactive mode selected")
	//periodic := flag.Bool("p", false, "periodic scan")
	flag.Parse()

	//start the interactive mode with runtime config
	if *interactive {
		fmt.Println("Interactive Mode selected")
		return "interactive"
	}

	// //enter the periodic mode NOT IMPLEMENTED YET
	// if *periodic {
	// 	panic("periodic scan still not implemented")
	// }

	//std run read from config file and doesnt prompt anything to console at runtime
	return "std"
}
