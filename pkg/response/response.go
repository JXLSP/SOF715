package response

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
    Code int64
    Message string
    Data any
}

func WriteResponse(w http.ResponseWriter, data BaseResponse) {
    d, err := json.Marshal(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")

    w.Write(d)
}

