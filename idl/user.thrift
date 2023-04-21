namespace go user

struct RegisterRequest {
    1: string username
    2: string password
}

struct CommonResponse {
    1: string response
}

struct LoginRequest {
    1: string username
    2: string password
}

struct LoginResponse {
    1: string userid
}

struct AddFriendRequest {
    1: string my_id
    2: string friend_id
}

struct DeleteFriendRequest {
    1: string my_id
    2: string friend_id
}

service UserService {
    LoginResponse Login(1: LoginRequest req)
    CommonResponse Register(1: RegisterRequest req)
    CommonResponse AddFriend(1: AddFriendRequest req)
    CommonResponse DeleteFriend(1: DeleteFriendRequest req)
}