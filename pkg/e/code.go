package e

const (
	SUCCESS       int32 = 0
	ERROR         int32 = 500
	InvalidParams       = 400

	//成员错误
	ErrorExistUser          int32 = 10001
	ErrorNotExistUser       int32 = 10002
	ErrorNotComparePassword int32 = 10003
	ErrorFailEncryption           = 10006
	ErrorNotExistProduct          = 10007
	ErrorNotExistAddress          = 10008

	//视屏错误
	ErrorExistFavorite             int32 = 20001
	ErrorBossCheckTokenTimeout           = 20002
	ErrorBossToken                       = 20003
	ErrorBoss                            = 20004
	ErrorBossInsufficientAuthority       = 20005
	ErrorBossProduct                     = 20006

	// 购物车
	ErrorProductExistCart = 20007
	ErrorProductMoreCart  = 20008

	//管理员错误
	ErrorAuthCheckTokenFail        = 30001 //token 错误
	ErrorAuthCheckTokenTimeout     = 30002 //token 过期
	ErrorAuthToken                 = 30003
	ErrorAuth                      = 30004
	ErrorAuthInsufficientAuthority = 30005
	ErrorReadFile                  = 30006
	ErrorSendEmail                 = 30007
	ErrorCallApi                   = 30008
	ErrorUnmarshalJson             = 30009
	ErrorAdminFindUser             = 30010
	//数据库错误
	ErrorDatabase = 40001

	//对象存储错误
	ErrorOss        = 50001
	ErrorUploadFile = 50002
)
