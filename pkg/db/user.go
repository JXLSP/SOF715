package db

import "sof/pkg/meta"

func CreatedUser(request *meta.CreatedUserRequest) error {
    sql, err := DBC.Prepare(`INSERT INTO users (username, password, nickname) VALUE(?, ?, ?)`)
    if err != nil {
        return err
    }

    _, err = sql.Exec(request.UserName, request.Password, request.NickName)
    if err != nil {
        return err
    }

    sql.Close()
    return nil
}

func GetUserByName(username string) (*meta.UserInfoResponse, error) {
    sql, err := DBC.Prepare(`SELECT id, username FROM users WHERE username = ?`)
    if err != nil {
        return nil, err
    }

    var (
        userInfo *meta.UserInfoResponse
        id int32
        userName string
    )
    err = sql.QueryRow(username).Scan(&id, &userName)
    if err != nil {
        return nil, err
    }

    userInfo.ID = id
    userInfo.Username = userName

    return userInfo, nil
}
