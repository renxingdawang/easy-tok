namespace go easy.tok.favorite

include "feed.thrift"

struct TokFavoriteActionRequest {
    1: i64 user_id,            // 用户id
    2: string token,           // 用户鉴权token
    3: i64 video_id,           // 视频id
    4: i32 action_type         // 1-点赞，2-取消点赞
}

struct TokFavoriteActionResponse {
    1: i32 status_code,        // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
}

struct TokFavoriteListRequest {
    1: i64 user_id,            // 用户id
    2: string token            // 用户鉴权token
}

struct TokFavoriteListResponse {
    1: i32 status_code,        // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: list<feed.TokVideo> video_list // 用户点赞视频列表
}

service FavoriteSrv {
    TokFavoriteActionResponse FavoriteAction(1: TokFavoriteActionRequest request), // 点赞或取消点赞
    TokFavoriteListResponse FavoriteList(2: TokFavoriteListRequest request)        // 返回点赞视频列表
}
