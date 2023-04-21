namespace go message

struct Message {
    1:optional string fromId
    2:string toId
    3:string content
}

struct SaveMsgRequset {
    1:string fromId
    2:string toId
    3:string content
}

struct SaveMsgResponse {
    1: string response
}


service MessageService {
    SaveMsgResponse SaveMessage(1: SaveMsgRequset req)
}