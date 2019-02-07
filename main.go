package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Get("/ping", func(ctx iris.Context) {
		x:= ctx.URLParam("firstname")
		y:=ctx.URLParam("lastname")
		ctx.JSON(iris.Map{
			"message": "pong",
			"x":ctx.URLParam("firstname"),
			"y":ctx.URLParam("lastname"),
			"fullname":x +" "+ y,
		})
	})

	app.Get("/hai", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "halo",
		})
	})


	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8080"))
}
func main2(){
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

