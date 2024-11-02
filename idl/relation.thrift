namespace go easy.tok.relation
include "user.thrift"

struct TokRelationActionRequest{
    1:i64 user_id,
    2:string token,
    3:i64 to_user_id,
    4:i32 actionType //1-关注 2-取消关注
}

struct TokRelationActionResponse{
    1:i32 status_code,
    2:string token
}

struct TokRelationFollowListRequest {
    1:i64 user_id,
    2:string token
}

struct TokRelationFollowListResponse {
    1: i32 status_code,       // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: list<user.User> user_list // 用户信息列表
}
struct TokRelationFollowerListRequest {
    1: i64 user_id,           // 用户id
    2: string token           // 用户鉴权token
}

struct TokRelationFollowerListResponse {
    1: i32 status_code,       // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: list<user.User> user_list   // 用户列表
}

service RelationSrv {
    TokRelationActionResponse RelationAction(1: TokRelationActionRequest request), // 关注或取消关注
    TokRelationFollowListResponse RelationFollowList(2: TokRelationFollowListRequest request), // 获取已关注用户的列表
    TokRelationFollowerListResponse RelationFollowerList(3: TokRelationFollowerListRequest request) // 获取粉丝用户列表
}