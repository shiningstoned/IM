namespace go group

struct CreateGroupRequest {
    1: string ownerUuid
}

struct CommonResponse {
    1: string message
}

struct JoinGroupRequest {
    1: string groupUuid
    2: string myUuid
}

struct GetGroupMemRequest {
    1: string groupUuid
}

struct GetGroupMemResponse {
    1: list<string> groupMemUuid
}

service GroupService {
    CommonResponse CreateGroup(1: CreateGroupRequest req)
    CommonResponse JoinGroup(1: JoinGroupRequest req)
    GetGroupMemResponse GetGroupMem(1: GetGroupMemRequest req)
}