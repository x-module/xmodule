// Code generated by "stringer -type ErrCode -linecomment -output errors_string.go"; DO NOT EDIT.

package internal

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Success-200]
	_ = x[GetConfigErr-10000]
	_ = x[InitRedisErr-10001]
	_ = x[ConnectMysqlErr-10002]
	_ = x[StartServerErr-101003]
	_ = x[ActionErr-101004]
	_ = x[SystemErr-101005]
	_ = x[SystemInitFail-101006]
	_ = x[listenConfigErr-101007]
	_ = x[ParamsError-101008]
	_ = x[ParamsErr-101009]
	_ = x[RequestOvertimeErr-101010]
	_ = x[SignErr-101011]
	_ = x[NoSignParamsErr-101012]
	_ = x[GetNoticeConfigErr-101013]
	_ = x[GetGameConfigErr-101014]
	_ = x[GetChannelConfigErr-101015]
	_ = x[GetLogConfigErr-101016]
	_ = x[GetApiConfigErr-101017]
	_ = x[GetDbConfigErr-101018]
	_ = x[GetGRPCConfigErr-101019]
	_ = x[GetSystemConfigErr-101020]
	_ = x[GetNacosConfigErr-101021]
	_ = x[RedisPushErr-101022]
	_ = x[RedisPublishErr-101023]
	_ = x[NeTRequestErr-101024]
	_ = x[RPCRequestErr-101025]
	_ = x[DataSaveErr-101026]
	_ = x[DataAddErr-101027]
	_ = x[DataGetErr-101028]
	_ = x[GetNakamaConfigErr-101029]
	_ = x[PublishDataErr-101030]
	_ = x[DbErr-101031]
	_ = x[DataDeleteErr-101032]
	_ = x[NoTokenErr-101033]
	_ = x[TokenErr-101034]
	_ = x[GetTokenErr-101035]
	_ = x[GetLeaderboardListErr-101036]
	_ = x[GetLeaderboardDetailErr-101037]
	_ = x[ParseJsonDataErr-101038]
	_ = x[GetAccountListErr-101039]
	_ = x[DeleteAccountErr-101040]
	_ = x[EditeAccountErr-101041]
	_ = x[GetAccountDetailErr-101042]
	_ = x[GetAccountBanListErr-101043]
	_ = x[DeleteLeaderboardErr-101044]
	_ = x[AccountUnlinkErr-101045]
	_ = x[GetAccountFriendErr-101046]
	_ = x[DeleteAccountFriendErr-101047]
	_ = x[AccountEnableErr-101048]
	_ = x[AccountDisableErr-101049]
	_ = x[GetMatchDataErr-101050]
	_ = x[GetMatchStateErr-101051]
	_ = x[AccountLoginErr-101052]
	_ = x[AccountTokenExpressErr-101053]
	_ = x[GetGameDataErr-101054]
	_ = x[ExecuteAfterDeleteFunErr-101055]
	_ = x[ExecuteAfterEditFunErr-101056]
	_ = x[ExecuteBeforeEditFunErr-101057]
	_ = x[ExecuteBeforeAddFunErr-101058]
	_ = x[LinkMysqlErr-101059]
	_ = x[GetSingleDataErr-101060]
	_ = x[CreateUploadFileDirErr-101061]
	_ = x[SystemError-101062]
	_ = x[ParamsEmptyError-101063]
	_ = x[ParamsFormatError-101064]
	_ = x[RepeatRequestError-101065]
	_ = x[InitSessionRedisErr-101066]
	_ = x[InitMysqlErr-101067]
	_ = x[GetSystemNoticeConfigErr-101068]
	_ = x[RegisterServerErr-101069]
	_ = x[GetServerErr-101070]
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

const (
	_ErrCode_name_0 = "Success"
	_ErrCode_name_1 = "获取配置异常初始化系统-连接Redis数据库异常连接数据库异常"
	_ErrCode_name_2 = "启动服务异常操作异常系统异常系统初始化失败配置文件监控失败参数异常，请检查参数异常，请检查请求发起时间超时参数签名异常参数签名时间戳或签名为异常获取系统通知配置异常获取游戏配置异常获取发布频道配置异常获取日志配置异常获取Api配置异常获取数据库配置异常获取GRPC配置异常获取系统配置异常获取Nacos配置异常Redis push 数据异常Redis 发布消息异常网络请求异常RPC请求异常DB数据编辑异常DB数据添加异常DB数据获取异常获取Nakama配置异常数据发布异常数据库异常DB数据删除异常无Token认证信息Token认证信息无效获取Token信息异常获取Nakama排行榜数据列表异常获取Nakama排行榜数据详情异常解析Nakama json数据异常获取Nakama账户列表异常删除Nakama账户列表异常编辑Nakama账户列表异常获取Nakama账户详情异常获取Nakama禁用账户列表异常删除Nakama排行榜数据异常删除Nakama账户好友关联异常获取Nakama账户好友异常删除Nakama账户好友异常启用Nakama账户异常禁用Nakama账户异常获取Nakama比赛数据异常获取Nakama比赛状态数据异常Nakama账户登录异常Nakama Token过期异常获取Nakama数据异常执行删除后方法异常执行编辑后方法异常执行编辑前方法异常执行添加前方法异常连接数据库异常获取单条数据异常创建上传目录异常系统异常，请稍后重试参数不可空，请检查参数格式错误，请检查重复请求初始化sessionRedis连接异常初始化系统-连接管理后台数据库异常。获取系统通知配置文件异常服务注册异常获取服务异常监听配置异常获取服务实例异常获取配置实例异常获取服务实例异常运行模式异常服务监听异常未知服务RPC连接异常定义数据异常数据查询为空！发布消息异常数据类型转换异常"
)

var (
	_ErrCode_index_1 = [...]uint8{0, 18, 60, 81}
	_ErrCode_index_2 = [...]uint16{0, 18, 30, 42, 63, 87, 111, 135, 159, 177, 216, 246, 270, 300, 324, 345, 372, 394, 418, 441, 464, 488, 506, 521, 541, 561, 581, 605, 623, 638, 658, 678, 701, 724, 763, 802, 831, 861, 891, 921, 951, 987, 1020, 1056, 1086, 1116, 1140, 1164, 1194, 1230, 1254, 1278, 1302, 1329, 1356, 1383, 1410, 1431, 1455, 1479, 1509, 1536, 1566, 1578, 1611, 1663, 1699, 1717, 1735, 1753, 1777, 1801, 1825, 1843, 1861, 1873, 1888, 1906, 1927, 1945, 1969}
)

func (i ErrCode) String() string {
	switch {
	case i == 200:
		return _ErrCode_name_0
	case 10000 <= i && i <= 10002:
		i -= 10000
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	case 101003 <= i && i <= 101082:
		i -= 101003
		return _ErrCode_name_2[_ErrCode_index_2[i]:_ErrCode_index_2[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
