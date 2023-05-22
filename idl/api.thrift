namespace go api

struct Request {

}

struct CreateGroupRequest {
}

struct JoinGroupRequest {
    1: string groupUuid
}

struct LoginRequest {
    1: string username
    2: string password
}

struct LoginResponse {
    1: string token
}

struct RegisterRequest {
    1: string username
    2: string password
}

struct CommonResponse {
    1: string response
}

struct AddFriendRequest {
    1: string frienduuid
}

struct DelFriendRequest {
    1: string frienduuid
}

service ApiService {
    //user service
    LoginResponse Login(1: LoginRequest req)(api.post="/user/login")
    CommonResponse Register(1: RegisterRequest req)(api.post="/user/register")
    CommonResponse AddFriend(1: AddFriendRequest req)(api.post="/user/addfriend")
    CommonResponse DelFriend(1: DelFriendRequest req)(api.post="/user/delfriend")

    CommonResponse WsChat(1: Request req)(api.get="/wschat")

    CommonResponse CreateGroup(1: CreateGroupRequest req)(api.post="/group/create")
    CommonResponse JoinGroup(1: JoinGroupRequest req)(api.post="/group/join")
}


