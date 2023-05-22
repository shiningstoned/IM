namespace go message

struct Message {
    1:optional string fromId
    2:string toId
    3:string content
    4:i64 type   //1 表示单聊 2 表示群聊
}

struct SaveMsgRequset {
    1:string fromId
    2:string toId
    3:string content
    4:i64 type
}

struct SaveMsgResponse {
    1: string response
}


service MessageService {
    SaveMsgResponse SaveMessage(1: SaveMsgRequset req)
}