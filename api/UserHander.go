package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

// api的接口响应
func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	io.WriteString(w, "create user success")
}

func LoginUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userName := ps.ByName("user_name")
	var repose = userName + "登陆成功"
	io.WriteString(w, repose)
}

