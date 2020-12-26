package core

import (
	"fmt"
	"github.com/dop251/goja"
)

func RunJs() {
	vm := goja.New()
	_, err := vm.RunString(`function sum(a,b){return a+b;} function mux(a,b){return sum(a, b)*a;}`)
	if err != nil {
		panic(err)
	}
	sum, ok := goja.AssertFunction(vm.Get("mux"))
	if !ok {
		panic("Not a function")
	}

	res, err := sum(goja.Undefined(), vm.ToValue(40), vm.ToValue(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
