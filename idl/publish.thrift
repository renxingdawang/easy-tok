namespace go easy.tok.publish

include"feed.thrift"

struct TokPublishActionRequest {
    1: string token,        // 用户鉴权token
    2: binary data,         // 视频数据
    3: string title         // 视频标题
}

struct TokPublishActionResponse {
    1: i32 status_code,     // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
}

struct TokPublishListRequest {
    1: i64 user_id,         // 用户id
    2: string token         // 用户鉴权token
}

struct TokPublishListResponse {
    1: i32 status_code,     // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: list<feed.TokVideo> video_list // 用户发布的视频列表
}

struct TokDeletePublishRequest {
    1: string token,        // 用户鉴权token
    2: i64 video_id         // 视频id，指要删除的视频
}

struct TokDeletePublishResponse {
    1: i32 status_code,     // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
}

service PublishSrv {
    TokPublishActionResponse PublishAction(1: TokPublishActionRequest request),     // 发布视频操作
    TokPublishListResponse PublishList(1: TokPublishListRequest request),           // 获取用户已发布视频的列表
    TokDeletePublishResponse DeletePublish(1: TokDeletePublishRequest request)      // 删除已发布视频操作
}
