namespace go easy.tok.user

struct TokUserRegisterRequest {
    1: string username; // 注册用户名，最长32个字符
    2: string password; // 密码，最长32个字符
}

struct TokUserRegisterResponse {
    1: i32 status_code; // 状态码，0-成功，其他值-失败
    2: optional string status_msg; // 返回状态描述
    3: i64 user_id; // 用户id
    4: string token; // 用户鉴权token
}

struct TokUserRequest {
    1: i64 user_id; // 用户id
    2: string token; // 用户鉴权token
}

struct TokUserResponse {
    1: i32 status_code; // 状态码，0-成功，其他值-失败
    2: optional string status_msg; // 返回状态描述
    3: User user; // 用户信息
}

struct User {
    1: i64 id; // 用户id
    2: string name; // 用户名称
    3: optional i64 follow_count; // 关注总数
    4: optional i64 follower_count; // 粉丝总数
    5: bool is_follow; // true-已关注，false-未关注
}

service UserSrv {
    TokUserRegisterResponse  Register(1: TokUserRegisterRequest req);
    TokUserRegisterResponse Login(2: TokUserRegisterRequest req);
    TokUserResponse GetUserById(3: TokUserRequest req);
}
