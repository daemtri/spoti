package api

// MulArg 请求参数
type MulArg struct {
	Left  float64 `json:"left" form:"left" valid:"int" help:"左值"`
	Right float64 `json:"right" form:"right" valid:"int" help:"右值"`
}

// MulReply 返回参数
type MulReply struct {
	Result float64 `json:"result"`
}

// AddArg xxx
type AddArg struct {
	Left  int `json:"left" form:"left" valid:"int" help:"左值"`
	Right int `json:"right" form:"right" valid:"int" help:"右值"`
}

// AddReply xxx
type AddReply struct {
	Result int `json:"result"`
}

// DivArg xxx
type DivArg struct {
	Left  int `json:"left" help:"左值"`
	Right int `json:"right" help:"右值"`
}

// DivReply xxx
type DivReply struct {
	Result int `json:"result"`
}

type GetUserInfoArg struct {
	ID int32 `path:"id"`
}

type GetUserInfoReply struct {
	ID      int32  `json:"id"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Account int32  `json:"sex"`
}
