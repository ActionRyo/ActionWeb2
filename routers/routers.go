package controllers

import (
	"ActionWeb2/controllers"
	"github.com/fishedee/web"
)

func init() {
	web.InitRoute("home", &controllers.MainController{})
}
