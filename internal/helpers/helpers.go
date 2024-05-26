package helpers

import (
	"fmt"
)

func EvaluateErr(err error, panicMsg string) {
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(panicMsg)
	}
}
