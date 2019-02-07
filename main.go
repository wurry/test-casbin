package main

import (
	"fmt"
	"github.com/casbin/casbin"
)

func main(){
	e := casbin.NewEnforcer("abac_model.conf", "basic_policy.csv")

	sub := "wury" // the user that wants to access a resource.
	obj := "data2" // the resource that is going to be accessed.
	act := "read" // the operation that the user performs on the resource.

	if e.Enforce(sub, obj, act) == true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

