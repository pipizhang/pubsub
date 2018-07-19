package pkg

import (
	"fmt"
)

// P prints debug info
func P(args ...interface{}) {
	for _, v := range args {
		fmt.Printf("%#v\n", v)
	}
}
