package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"sof/internal/services"
	"sof/pkg/db"
	"sof/pkg/meta"
	"sof/pkg/response"

	"github.com/julienschmidt/httprouter"
)

type UserControllers struct {
    b services.Serviceser
}

func (ctrl *UserControllers) CreatedUserController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    res, _ := io.ReadAll(r.Body)

    request := &meta.CreatedUserRequest{}

    if err := json.Unmarshal(res, request); err != nil {
        rs := response.BaseResponse{Code: 400, Message: "参数错误", Data: []string{}}
        response.WriteResponse(w, rs)
        return
    }

    if err := ctrl.b.Users().CreatedUser(r.Context(), request); err != nil {}
        rs := response.BaseResponse{Code: 500, Message: err.Error(), Data: []string{}}
        response.WriteResponse(w, rs)
        return
    }

    data := response.BaseResponse{Code: 200, Message: "用户创建成功", Data: []string{}}
    response.WriteResponse(w, data)
}
