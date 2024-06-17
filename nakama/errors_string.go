// Code generated by "stringer -type ErrCode -linecomment -output errors_string.go"; DO NOT EDIT.

package nakama

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StartServerErr-101000]
	_ = x[ActionErr-101001]
	_ = x[SystemErr-101002]
	_ = x[SystemInitFail-101003]
	_ = x[listenConfigErr-101004]
	_ = x[ParamsError-101005]
	_ = x[ParamsErr-101006]
	_ = x[ConnectMysqlErr-101007]
	_ = x[RequestOvertimeErr-101008]
	_ = x[SignErr-101009]
	_ = x[NoSignParamsErr-101010]
	_ = x[GetNoticeConfigErr-101011]
	_ = x[GetGameConfigErr-101012]
	_ = x[GetChannelConfigErr-101013]
	_ = x[GetLogConfigErr-101014]
	_ = x[GetApiConfigErr-101015]
	_ = x[GetDbConfigErr-101016]
	_ = x[GetGRPCConfigErr-101017]
	_ = x[GetSystemConfigErr-101018]
	_ = x[GetNacosConfigErr-101019]
	_ = x[RedisPushErr-101020]
	_ = x[RedisPublishErr-101021]
	_ = x[NeTRequestErr-101022]
	_ = x[RPCRequestErr-101023]
	_ = x[DataSaveErr-101024]
	_ = x[DataAddErr-101025]
	_ = x[DataGetErr-101026]
	_ = x[GetNakamaConfigErr-101027]
	_ = x[PublishDataErr-101028]
	_ = x[DbErr-101029]
	_ = x[DataDeleteErr-101030]
	_ = x[NoTokenErr-101031]
	_ = x[TokenErr-101032]
	_ = x[GetTokenErr-101033]
	_ = x[GetLeaderboardListErr-101034]
	_ = x[GetLeaderboardDetailErr-101035]
	_ = x[ParseJsonDataErr-101036]
	_ = x[GetAccountListErr-101037]
	_ = x[DeleteAccountErr-101038]
	_ = x[EditeAccountErr-101039]
	_ = x[GetAccountDetailErr-101040]
	_ = x[GetAccountBanListErr-101041]
	_ = x[DeleteLeaderboardErr-101042]
	_ = x[AccountUnlinkErr-101043]
	_ = x[GetAccountFriendErr-101044]
	_ = x[DeleteAccountFriendErr-101045]
	_ = x[AccountEnableErr-101046]
	_ = x[AccountDisableErr-101047]
	_ = x[GetMatchDataErr-101048]
	_ = x[GetMatchStateErr-101049]
	_ = x[AccountLoginErr-101050]
	_ = x[AccountTokenExpressErr-101051]
	_ = x[GetGameDataErr-101052]
	_ = x[ExecuteAfterDeleteFunErr-101053]
	_ = x[ExecuteAfterEditFunErr-101054]
	_ = x[ExecuteBeforeEditFunErr-101055]
	_ = x[ExecuteBeforeAddFunErr-101056]
	_ = x[LinkMysqlErr-101057]
	_ = x[GetSingleDataErr-101058]
	_ = x[CreateUploadFileDirErr-101059]
	_ = x[SystemError-101060]
	_ = x[ParamsEmptyError-101061]
	_ = x[ParamsFormatError-101062]
	_ = x[RepeatRequestError-101063]
	_ = x[InitSessionRedisErr-101064]
	_ = x[InitMysqlErr-101065]
	_ = x[InitRedisErr-101066]
	_ = x[GetSystemNoticeConfigErr-101067]
	_ = x[RegisterServerErr-101068]
	_ = x[GetServerErr-101069]
	_ = x[GetConfigErr-101070]
	_ = x[ListenConfigErr-101071]
	_ = x[GetNamingClientErr-101072]
	_ = x[GetConfigClientErr-101073]
	_ = x[GetInstanceErr-101074]
	_ = x[RunModeErr-101075]
	_ = x[SubscribeServerErr-101076]
	_ = x[UnknownServerErr-101077]
	_ = x[RPCLinkErr-101078]
	_ = x[SubscribeDataErr-101079]
	_ = x[NoRecordErr-101080]
	_ = x[PublishErr-101081]
	_ = x[TransDataTypeErr-101082]
}

const _ErrCode_name = "启动服务异常操作异常系统异常系统初始化失败配置文件监控失败参数异常，请检查参数异常，请检查连接数据库异常请求发起时间超时参数签名异常参数签名时间戳或签名为异常获取系统通知配置异常获取游戏配置异常获取发布频道配置异常获取日志配置异常获取Api配置异常获取数据库配置异常获取GRPC配置异常获取系统配置异常获取Nacos配置异常Redis push 数据异常Redis 发布消息异常网络请求异常RPC请求异常DB数据编辑异常DB数据添加异常DB数据获取异常获取Nakama配置异常数据发布异常数据库异常DB数据删除异常无Token认证信息Token认证信息无效获取Token信息异常获取Nakama排行榜数据列表异常获取Nakama排行榜数据详情异常解析Nakama json数据异常获取Nakama账户列表异常删除Nakama账户列表异常编辑Nakama账户列表异常获取Nakama账户详情异常获取Nakama禁用账户列表异常删除Nakama排行榜数据异常删除Nakama账户好友关联异常获取Nakama账户好友异常删除Nakama账户好友异常启用Nakama账户异常禁用Nakama账户异常获取Nakama比赛数据异常获取Nakama比赛状态数据异常Nakama账户登录异常Nakama Token过期异常获取Nakama数据异常执行删除后方法异常执行编辑后方法异常执行编辑前方法异常执行添加前方法异常连接数据库异常获取单条数据异常创建上传目录异常系统异常，请稍后重试参数不可空，请检查参数格式错误，请检查重复请求初始化sessionRedis连接异常初始化系统-连接管理后台数据库异常。初始化系统-连接Redis数据库异常获取系统通知配置文件异常服务注册异常获取服务异常获取配置异常监听配置异常获取服务实例异常获取配置实例异常获取服务实例异常运行模式异常服务监听异常未知服务RPC连接异常定义数据异常数据查询为空！发布消息异常数据类型转换异常"

var _ErrCode_index = [...]uint16{0, 18, 30, 42, 63, 87, 111, 135, 156, 180, 198, 237, 267, 291, 321, 345, 366, 393, 415, 439, 462, 485, 509, 527, 542, 562, 582, 602, 626, 644, 659, 679, 699, 722, 745, 784, 823, 852, 882, 912, 942, 972, 1008, 1041, 1077, 1107, 1137, 1161, 1185, 1215, 1251, 1275, 1299, 1323, 1350, 1377, 1404, 1431, 1452, 1476, 1500, 1530, 1557, 1587, 1599, 1632, 1684, 1726, 1762, 1780, 1798, 1816, 1834, 1858, 1882, 1906, 1924, 1942, 1954, 1969, 1987, 2008, 2026, 2050}

func (i ErrCode) String() string {
	i -= 101000
	if i < 0 || i >= ErrCode(len(_ErrCode_index)-1) {
		return "ErrCode(" + strconv.FormatInt(int64(i+101000), 10) + ")"
	}
	return _ErrCode_name[_ErrCode_index[i]:_ErrCode_index[i+1]]
}
