namespace go easy.tok.comment
include"user.thrift"

struct TokCommentActionRequest{
    1:i64 user_id,
    2:string token,
    3:i64 video_id,
    4:i32 action_type,
    5:optional string comment_text,
    6:optional i64 comment_id//要删除的id
}

struct TokCommentActionResponse {
    1: i32 status_code,            // 状态码，0-成功，其他值-失败
    2: optional string status_msg,  // 返回状态描述
    3: optional Comment comment     // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct TokCommentListRequest {
    1: string token,                // 用户鉴权token
    2: i64 video_id                 // 视频id
}

struct TokCommentListResponse {
    1: i32 status_code,             // 状态码，0-成功，其他值-失败
    2: optional string status_msg,   // 返回状态描述
    3: list<Comment> comment_list    // 评论列表
}

struct Comment {
    1: i64 id,                      // 视频评论id
    2: user.User user,              // 评论用户信息
    3: string content,              // 评论内容
    4: string create_date           // 评论发布日期，格式 mm-dd
}

service CommentSrv {
    TokCommentActionResponse CommentAction(1: TokCommentActionRequest request), // 评论操作
    TokCommentListResponse CommentList(2: TokCommentActionRequest request)       // 返回评论列表
}