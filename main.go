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

	app.Get("/casbin", func(ctx iris.Context) {
		e := casbin.NewEnforcer("abac_model.conf", "basic_policy.csv")
		sub:= ctx.URLParam("sub")
		obj:= ctx.URLParam("obj")
		act:= ctx.URLParam("act")
		env:= ctx.URLParam("env")
		if e.Enforce(sub, obj, act, env) == true {
			ctx.JSON(iris.Map{
				"message": "halo",
				"value":"true",
			})
		}else{
			ctx.JSON(iris.Map{
				"message": "halo",
				"value":"false",
			})
		}
	})

	app.Post("/casbin", Casbin)
	app.Post("/casbinJSON", CasbinJSON)

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

func Casbin(ctx iris.Context) {
	e := casbin.NewEnforcer("abac_model.conf", "basic_policy.csv")
	sub:= ctx.PostValue("sub")
	obj:= ctx.PostValue("obj")
	act:= ctx.PostValue("act")
	env:= ctx.PostValue("env")
	if e.Enforce(sub, obj, act, env) == true {
		ctx.JSON(iris.Map{
			"message": "halo",
			"value":"true",
		})
	}else{
		ctx.JSON(iris.Map{
			"message": "halo",
			"value":"false",
		})
	}
}

type Req struct{
	Sub string `json:"sub"`
	Obj string `json:"obj"`
	Act string `json:"act"`
	Env string `json:"env"`
}

func CasbinJSON(ctx iris.Context) {

	var r Req
	err:= ctx.ReadJSON(&r)
	if err!=nil{
		ctx.JSON(iris.Map{
			"message": "data not valid",
		})
		return
	}

	e := casbin.NewEnforcer("abac_model.conf", "basic_policy.csv")
	sub:= r.Sub
	obj:= r.Obj
	act:= r.Act
	env:= r.Env
	if e.Enforce(sub, obj, act, env) == true {
		ctx.JSON(iris.Map{
			"message": "halo",
			"value":"true",
		})
	}else{
		ctx.JSON(iris.Map{
			"message": "halo",
			"value":"false",
		})
	}
}