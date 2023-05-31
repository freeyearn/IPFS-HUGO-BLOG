package codes

const (
	//用户错误
	NotLoggedIn        = "1000" // 未登录
	UnauthorizedUserId = "1002" // 非法的用户Id
	Unauthorized       = "1003" // 未授权
	OperationFailure   = "1009" // 操作失败

	// 通用错误
	OK        = "2000" // Success
	NotData   = "2001" // 没有数据
	DataExist = "2002" // 数据已存在
	DataError = "2003" // 数据错误

	// 网络级错误
	ParameterIllegal = "4000" // 参数不合法
	RequestOverDue   = "4001" // 请求已过期
	LoginError       = "4002" // 登录已过期
	AccessDenied     = "4003" // 拒绝访问
	RoutingNotExist  = "4004" // 路由不存在
	PasswordError    = "4005" // 密码错误
	RequestError     = "4006" // 非法访问
	IPError          = "4007" // IP受限

	// 系统级错误
	InternalError = "5000" // 系统错误
	DBError       = "5001" // 数据库错误
	ThirdError    = "5002" // 第三方系统错误
	IOError       = "5003" // IO错误
	UnKnownError  = "5004" // 未知错误
)
