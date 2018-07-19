package pkg

import (
	"fmt"
)

func P(args ...interface{}) {
	for _, v := range args {
		fmt.Printf("%#v\n", v)
	}
}
