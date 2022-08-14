package controllers

import "github.com/gin-gonic/gin"

//type of Response object
type Response struct {
  Code  int     `json:"code"`
  Msg   string  `json:"msg"`
  Data  interface{}  `json:"data"`
}

// send an http response
func HttpResponse(ctx *gin.Context, code int, msg string, data interface{}) {
  ctx.JSON(
    code, Response{
      Code: code,
      Msg: msg,
      Data: data,
    },
  )
  return

}
