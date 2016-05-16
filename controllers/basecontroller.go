package controllers

import (
	. "github.com/fishedee/language"
	"github.com/fishedee/web"
	"html/template"
	"net/http"
)

type redirectOut struct {
	url string
}

type BaseController struct {
	web.Controller
}

func (this *BaseController) AutoRender(data interface{}, viewname string) {
	httpRequest := this.Ctx.GetRawRequest().(*http.Request)
	httpResponseWriter := this.Ctx.GetRawResponseWriter().(http.ResponseWriter)

	exception, isException := data.(Exception)
	if isException {
		if exception.GetCode() == 10001 {
			http.Redirect(httpResponseWriter, httpRequest, "/home/login", http.StatusFound)
			return
		} else {
			this.Ctx.Write([]byte(exception.GetMessage()))
		}
		return
	}

	//跳转的输出
	redirectInfo, isRedirect := data.(redirectOut)
	if isRedirect {
		redirectUrl := redirectInfo.url
		http.Redirect(httpResponseWriter, httpRequest, redirectUrl, http.StatusFound)
		return
	}

	//模板的输出
	t, err := template.ParseFiles("views/" + viewname[4:] + ".html")
	if err != nil {
		panic(err)
	}

	t.Execute(httpResponseWriter, data)
}
