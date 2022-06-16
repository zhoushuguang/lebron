package xerr

//成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

//全局错误码
const ServerCommonError uint32 = 100001
const ReuqestParamError uint32 = 100002
const TokenExpireError uint32 = 100003
const TokenGenerateError uint32 = 100004
const DbError uint32 = 100005
const DbUpdateAffectedZeroError uint32 = 100006

//用户服务

//订单服务

//商品服务

//支付服务
