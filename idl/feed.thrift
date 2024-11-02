namespace go easy.tok.feed

include "user.thrift"

struct TokFeedRequest {
    1: optional i64 latest_time,   // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token       // 可选参数，登录用户设置
}

struct TokFeedResponse {
    1: i32 status_code,            // 状态码，0-成功，其他值-失败
    2: optional string status_msg,  // 返回状态描述
    3: list<TokVideo> video_list,      // 视频列表
    4: optional i64 next_time       // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct TokVideoByIdRequest {
    1: i64 video_id,               // 视频唯一标识
    2: i64 search_id               // 搜索唯一标识
}

struct TokVideo {
    1: i64 id,                     // 视频唯一标识
    2: user.User author,           // 视频作者信息
    3: string play_url,            // 视频播放地址
    4: string cover_url,           // 视频封面地址
    5: i64 favorite_count,         // 视频的点赞总数
    6: i64 comment_count,          // 视频的评论总数
    7: bool is_favorite,           // true-已点赞，false-未点赞
    8: string title                // 视频标题
}

service FeedSrv {
    TokFeedResponse GetUserFeed(1: TokFeedRequest request),      // 返回一个视频列表
    TokVideo GetVideoById(2: TokFeedRequest request)                     // 根据视频id返回一个视频
}
