package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/convox/release/version"
)

func main() {
	reqPtr := flag.Bool("require", false, "denote a required version. Lesser versions will first update to this version.")
	pubPtr := flag.Bool("publish", false, "denote a published version. New installs will use the latest published version.")
	flag.Parse()

	switch flag.Arg(0) {
	case "create":
		cmdCreate(*pubPtr, *reqPtr)
	case "update":
		cmdUpdate(*pubPtr, *reqPtr)
	default:
		cmdList()
	}
}

func cmdList() {
	vs, err := version.All()

	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		os.Exit(1)
	}

	v, err := vs.Latest()

	if err != nil {
		fmt.Printf("error: %v\n", err.Error())

		if err.Error() == "no published versions" {
			fmt.Printf("\nLast non-published versions are:\n")

			n := 5
			if len(vs) < 5 {
				n = len(vs)
			}

			for i := len(vs) - 1; i >= len(vs)-n; i-- {
				fmt.Printf("%v\n", vs[i].Display())
			}
		}

		os.Exit(1)
	}

	fmt.Printf("%v\n", v.Display())
}

func cmdCreate(published bool, required bool) {
	vv := flag.Arg(1)

	if vv == "" {
		fmt.Printf("usage: version [-publish] [-require] create 20150906195708\n")
		os.Exit(1)
	}

	v := version.Version{
		Version:   vv,
		Published: published,
		Required:  required,
	}

	v, err := version.AppendVersion(v)

	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(v.Display())
}

func cmdUpdate(published bool, required bool) {
	vv := flag.Arg(1)

	if vv == "" {
		fmt.Printf("usage: version [-publish] [-require] update 20150906195708\n")
		os.Exit(1)
	}

	if required && !published {
		fmt.Printf("error: can not use `-require` without `-publish`\n")
		os.Exit(1)
	}

	v := version.Version{
		Version:   vv,
		Published: published,
		Required:  required,
	}

	v, err := version.UpdateVersion(v)

	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(v.Display())
}
