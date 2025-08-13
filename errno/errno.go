package errno

// Errno 定义错误码
type Errno struct {
	Code int
	Msg  string
}

// 实现 error 接口
func (e Errno) Error() string {
	return e.Msg
}

// 通用错误
var (
	Success       = Errno{0, "成功"}
	UnknownError  = Errno{1000, "未知错误"}
	InvalidParams = Errno{1001, "参数错误"}
	Timeout       = Errno{1002, "请求超时"}
	Forbidden     = Errno{1003, "权限不足"}
)

// 用户模块 2000+
var (
	UserNotFound         = Errno{2000, "用户不存在"}
	UserAlreadyExist     = Errno{2001, "用户已存在"}
	UserBalanceNotEnough = Errno{2002, "余额不足"}
	UserUnauthorized     = Errno{2003, "用户未登录或未授权"}
)

// 游戏模块 2100+
var (
	SlotSpinFail      = Errno{2100, "Spin 失败"}
	SlotInvalidBet    = Errno{2101, "投注金额无效"}
	SlotBonusNotFound = Errno{2102, "奖励未找到"}
	SlotNotEnoughSpin = Errno{2103, "剩余免费 Spin 不够"}
)

// 系统模块 2200+
var (
	InternalError = Errno{2200, "系统内部错误"}
	Conflict      = Errno{2201, "数据冲突"}
)

// 辅助函数：动态生成错误
func NewErrno(code int, msg string) Errno {
	return Errno{Code: code, Msg: msg}
}
