package errno

var (
    SUCCESS = &Errno{HTTP: 200, Code: "SUCCESS", Msg: "Is OK."}

    INTERNALSERVERERROR = &Errno{HTTP: 500, Code: "INTERNALSERVERERROR", Msg: "Internal server error."}

    NOTFOUNDERROR = &Errno{HTTP: 404, Code: "NOTFOUNDERROR", Msg: "Page not found."}

    PARAMTERBINDERROR = &Errno{HTTP: 400, Code: "PARAMTERBINDERROR", Msg: "Error occurred while binding the request body to the struct."}

    INVALIDPARAMETERERROR = &Errno{HTTP: 400, Code: "INVALIDPARAMETERERROR", Msg: "Parameter verifycation failed."}
)

