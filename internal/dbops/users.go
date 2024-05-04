package dbops

import (
	"context"
	"database/sql"
	"sof/pkg/meta"
)

type UserOps interface {
    CreatedUser(ctx context.Context, request *meta.CreatedUserRequest) error
    GetUserByName(ctx context.Context, username string) (*meta.UserInfoResponse, error)
    ExistUser(ctx context.Context, username string) bool
}

type users struct {
    db *sql.DB
}

func (us *users) CreatedUser(ctx context.Context, request *meta.CreatedUserRequest) error {
    sql, err := us.db.Prepare(`INSERT INTO users (username, password, nickname) VALUES(?, ?, ?)`)
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

func (us *users) GetUserByName(ctx context.Context, username string) (*meta.UserInfoResponse, error) {
    var (
        userInfo *meta.UserInfoResponse
        id int32
        userName string
    )

    sql, err := us.db.Prepare(`SELECT id, username FROM users WHERE username = ?`)
    if err != nil {
        return nil, err
    }

    err = sql.QueryRow(username).Scan(&id, &userName)
    if err != nil {
        return nil, err
    }

    userInfo.ID = id
    userInfo.Username = userName

    return userInfo, nil
}

func (us *users) ExistUser(ctx context.Context, username string) bool {
    var id int32
    sql, err := us.db.Prepare(`SELECT id FROM users WHERE username = ?`)
    if err != nil {
        return false
    }

    err = sql.QueryRow(username).Scan(&id)
    if err != nil {
        return false
    }

    if id == 0 {
        return false
    }

    return true
}
