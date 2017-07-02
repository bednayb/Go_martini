// main.go
package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/martini-contrib/binding"
  "net/http"
)

type UserInfo struct {
  Name string `form:"name"`
  Email string `form:"email"`
}

func main() {
  m := martini.Classic()
  // render html templates from templates directory
  m.Use(render.Renderer(render.Options{
    Layout: "layout",
  }))

  m.Get("/", func(r render.Render) {
    r.HTML(http.StatusOK, "index",nil)
  })

  m.Post("/user", binding.Bind(UserInfo{}), func (r render.Render, user UserInfo){

    var retData struct{
      User UserInfo
    }
    retData.User = user

    //r.JSON(http.StatusOK, user)
    r.HTML(http.StatusOK, "user_info",retData)
  })

  m.Get("/user/:userid", func(r render.Render, p martini.Params) {
    var retData struct{
      ID string
    }
    retData.ID = p["userid"]

    r.HTML(http.StatusOK, "user",retData)

    r.JSON(http.StatusOK, retData)
  })

  m.Run()
}