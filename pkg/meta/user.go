package meta

type CreatedUserRequest struct {
    UserName string
    Password string
    NickName string
}

type UserInfoResponse struct {
    ID int32 `json:"id"`
    Username string `json:"username"`
}

