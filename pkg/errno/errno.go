package errno

import "fmt"

type Errno struct {
    HTTP int
    Code string
    Msg string
}

func (err *Errno) Error() string {
    return err.Msg
}

func (e *Errno) SetMessage(format string, opts any...) *Errno {
    e.Msg = fmt.Sprintf(format, opts...)
    return e
}

func Decoder(err error) (int, string, string) {
    if err == nil {
        return 200, "SUCCESS", "is ok"
    }
    switch typed := err.(type) {
    case *Errno:
        return typed.HTTP, typed.Code, typed.Msg
    default:
    }
    return 500, "InternalError", err.Error()
}
