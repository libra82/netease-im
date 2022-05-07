package tests

import (
	"encoding/json"
	"github.com/libra82/netease-im"
	"testing"
)

/**
 * 功能描述：
 *
 * Created with GoLand
 * @Author: david.tao
 * @CreateTime 2022-05-01 10:44
 *
 * 修改历史：(修改人，修改时间，修改原因/内容)
 *
 **/

func Test_CreateChatroom(t *testing.T) { //{1274643577 true  一场篮球比赛   1}
	chatroom := &netease.ImChatRoomReq{Creator: "1", Name: "一场篮球比赛"}
	room, err := client.CreateChatRoom(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_GetChatRoomInfo(t *testing.T) { //{"roomid":1274643577,"valid":true,"muted":false,"announcement":"聊天室倡导文明健康聊天聊天环境，禁止任何联系方式、群、广告，违规者禁言或封号。","name":"一场足球比赛","broadcasturl":"","onlineusercount":0,"ext":"","creator":"1","queuelevel":0,"ionotify":true}
	chatroom := &netease.ImChatRoomInfoReq{Roomid: 1274643577, NeedOnlineUserCount: "true"}
	room, err := client.GetChatRoomInfo(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_GetBatchChatRoomInfo(t *testing.T) { //{"noExistRooms":[123,456],"succRooms":[{"roomid":1274643577,"valid":true,"muted":false,"announcement":"聊天室倡导文明健康聊天聊天环境，禁止任何联系方式、群、广告，违规者禁言或封号。","name":"一场足球比赛","broadcasturl":"","onlineusercount":0,"ext":"","creator":"1","queuelevel":0,"ionotify":true}],"failRooms":[]}
	chatroom := &netease.ImChatRoomBatchInfoReq{Roomids: "[1274643577,123,456]", NeedOnlineUserCount: "true"}
	room, err := client.GetBatchChatRoomInfo(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_UpdateChatRoom(t *testing.T) { //{"roomid":1274643577,"valid":true,"announcement":"聊天室倡导文明健康聊天聊天环境，禁止任何联系方式、群、广告，违规者禁言或封号。","name":"一场足球比赛","broadcasturl":"","ext":"","creator":"1"}
	chatroom := &netease.ImChatRoomUpdateReq{Roomid: 1274643577, Name: "一场足球比赛", Announcement: "聊天室倡导文明健康聊天聊天环境，禁止任何联系方式、群、广告，违规者禁言或封号。"}
	room, err := client.UpdateChatRoom(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_ChangeChatRoomState(t *testing.T) { //{"roomid":1274643577,"valid":false,"announcement":"聊天室倡导文明健康聊天聊天环境，禁止任何联系方式、群、广告，违规者禁言或封号。","name":"一场足球比赛","broadcasturl":"","ext":"","creator":"1"}
	chatroom := &netease.ImChatRoomStateReq{Roomid: 1274643577, Operator: "1", Valid: "true"}
	room, err := client.ChangeChatRoomState(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_SetChatRoomRole(t *testing.T) { //{"roomid":1274643577,"level":0,"accid":"2","type":"MANAGER"}
	chatroom := &netease.ImChatRoomRoleReq{Roomid: 1274643577, Operator: "1", Target: "2", Opt: 1, Optvalue: "true"}
	room, err := client.SetChatRoomRole(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_ReqChatRoomAddr(t *testing.T) { //["chatweblink03.netease.im:443","chatweblink10.netease.im:443"]
	chatroom := &netease.ImChatRoomAddrReq{Roomid: 1274643577, Accid: "2"}
	room, err := client.ReqChatRoomAddr(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_SendChatRoomMsg(t *testing.T) { //{"time":"1651826273092","fromAvator":"","msgid_client":"12341234123","fromClientType":"REST","attach":"文本消息：填写消息文案; 其它类型消息，请参考 消息格式示例； 长度限制4096字符","roomId":"1274643577","fromAccount":"1","fromNick":"david.tao","type":"0","ext":"","highPriorityFlag":0,"msgAbandonFlag":""}
	chatroom := &netease.ImChatRoomSendMsgReq{Roomid: 1274643577, MsgId: "12341234123", FromAccid: "2", MsgType: 0, Attach: "文本消息：填写消息文案; 其它类型消息，请参考 消息格式示例； 长度限制4096字符"}
	room, err := client.SendChatRoomMsg(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_AddChatRoomRobot(t *testing.T) { //{"failAccids":"[]","successAccids":"[\"1\",\"2\"]","oldAccids":"[]"}
	chatroom := &netease.ImChatroomAddRobotReq{Roomid: 1274643577, Accids: "[1,2]"}
	room, err := client.AddChatRoomRobot(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_DelChatRoomRobot(t *testing.T) { //{"failAccids":"[]","successAccids":"[\"2\"]"}
	chatroom := &netease.ImChatroomDelRobotReq{Roomid: 1274643577, Accids: "[2]"}
	room, err := client.DelChatRoomRobot(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_CleanChatRoomRobot(t *testing.T) { //{"size":1}
	chatroom := &netease.ImChatroomCleanRobotReq{Roomid: 1274643577, Notify: true}
	room, err := client.CleanChatRoomRobot(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_MuteChatRoomTemp(t *testing.T) { //{"muteDuration":300}
	chatroom := &netease.ImChatroomMuteReq{Roomid: 1274643577, Operator: "1", Target: "2", MuteDuration: 300}
	room, err := client.MuteChatRoomTemp(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_MuteChatRoom(t *testing.T) { //{"success":true}
	chatroom := &netease.ImChatroomMuteRoomReq{Roomid: 1274643577, Operator: "1", Mute: "true"}
	room, err := client.MuteChatRoom(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_StatChatRoomTopn(t *testing.T) { //[]
	chatroom := &netease.ImChatroomTopnReq{}
	room, err := client.StatChatRoomTopn(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_PageChatRoomMembers(t *testing.T) { //{"data":[{"roomid":1274643577,"accid":"2","nick":"joan","avator":"","ext":"","type":"MANAGER","level":0,"onlineStat":true,"enterTime":1651824690091,"blacklisted":false,"muted":false,"tempMuted":false,"tempMuteTtl":0,"isRobot":true,"robotExpirAt":0},{"roomid":1274643577,"accid":"1","nick":"david.tao","avator":"","ext":"","type":"CREATOR","level":0,"onlineStat":true,"enterTime":1651815459343,"blacklisted":false,"muted":false,"tempMuted":false,"tempMuteTtl":0,"isRobot":true,"robotExpirAt":0}]}
	chatroom := &netease.ImChatroomMembersReq{Roomid: 1274643577, Type: 0, Endtime: 0, Limit: 100}
	room, err := client.PageChatRoomMembers(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_GetChatRoomMembersByRole(t *testing.T) { //{"data":[{"roomid":1274643577,"accid":"2","nick":"joan","avator":"","ext":"","type":"MANAGER","level":0,"onlineStat":true,"enterTime":1651830018736,"blacklisted":false,"muted":false,"tempMuted":false,"tempMuteTtl":0,"isRobot":true,"robotExpirAt":0},{"roomid":1274643577,"accid":"1","nick":"david.tao","avator":"","ext":"","type":"CREATOR","level":0,"onlineStat":true,"enterTime":1651830018724,"blacklisted":false,"muted":false,"tempMuted":false,"tempMuteTtl":0,"isRobot":true,"robotExpirAt":0}]}
	chatroom := &netease.ImChatroomMembersByRoleReq{Roomid: 1274643577, Roles: "{\"creator\": true,\"manager\": true,\"blacklist\": false,\"mute\": false}"}
	room, err := client.GetChatRoomMembersByRole(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_GetChatRoomMembersBatch(t *testing.T) { //{"data":[{"roomid":1274643577,"accid":"1","nick":"david.tao","type":"CREATOR","onlineStat":true},{"roomid":1274643577,"accid":"2","nick":"joan","type":"MANAGER","onlineStat":true}]}
	chatroom := &netease.ImChatroomMembersBatchReq{Roomid: 1274643577, Accids: "[1,2]"}
	room, err := client.GetChatRoomMembersBatch(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_UpdateChatRoomMemberRole(t *testing.T) { //{"data":[{"roomid":1274643577,"accid":"1","nick":"david.tao","type":"CREATOR","onlineStat":true},{"roomid":1274643577,"accid":"2","nick":"joan","type":"MANAGER","onlineStat":true}]}
	chatroom := &netease.ImChatroomChangeRoleReq{Roomid: 1274643577, Accid: "2", Save: true, Nick: "joan.xiong"}
	room, err := client.UpdateChatRoomMemberRole(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_GetChatRoomUserRoomIds(t *testing.T) { //{"roomids":["1274643577"]}
	chatroom := &netease.ImChatroomUserRoomIdsReq{Creator: "1"}
	room, err := client.GetChatRoomUserRoomIds(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_CloseChatRoomInOutNotify(t *testing.T) { //TODO：403 app io notify closed
	chatroom := &netease.ImChatroomInOutNotifyReq{Roomid: 1274643577, Close: false}
	room, err := client.CloseChatRoomInOutNotify(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_MuteChatRoomByTag(t *testing.T) { //{"muteDuration":300}
	chatroom := &netease.ImChatroomTagMuteReq{Roomid: 1274643577, Operator: "1", TargetTag: "xxxTag", MuteDuration: 300}
	room, err := client.MuteChatRoomByTag(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_CountChatRoomMemberByTag(t *testing.T) { //{"tag":"xxxTag","onlineUserCount":0}
	chatroom := &netease.ImChatroomTagMemberCountReq{Roomid: 1274643577, Tag: "xxxTag"}
	room, err := client.CountChatRoomMemberByTag(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_PageChatRoomMemberByTag(t *testing.T) { //{"data":[]}
	chatroom := &netease.ImChatroomTagMembersReq{Roomid: 1274643577, Tag: "xxxTag", EndTime: 0, Limit: 100}
	room, err := client.PageChatRoomMemberByTag(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_BroadcastChatRoom(t *testing.T) { //403 no right to send chatroom broadcast msg!  本功能需要开通，请联系商务
	chatroom := &netease.ImChatRoomBroadcastReq{MsgId: "12341234123", FromAccid: "2", MsgType: 0, Attach: "文本消息：填写消息文案; 其它类型消息，请参考 消息格式示例； 长度限制4096字符"}
	room, err := client.BroadcastChatRoom(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_ReCallChatRoomMsg(t *testing.T) { //true
	chatroom := &netease.ImChatRoomReCallReq{Roomid: 1274643577, MsgTimetag: 1648090381, FromAcc: "2", MsgId: "12341234123", OperatorAcc: "1"}
	room, err := client.ReCallChatRoomMsg(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}

func Test_SendChatRoomMsgToOne(t *testing.T) { //{"time":"1651894983250","fromAvator":"","msgid_client":"12341234123","fromClientType":"REST","attach":"文本消息：填写消息文案; 其它类型消息，请参考 消息格式示例； 长度限制4096字符","roomId":"1274643577","fromAccount":"1","fromNick":"david.tao","type":"0","ext":""}
	chatroom := &netease.ImChatRoomSendMsgToOneReq{Roomid: 1274643577, MsgId: "12341234123", FromAccid: "1", ToAccids: "[2]", MsgType: 0, Attach: "文本消息：填写消息文案; 其它类型消息，请参考 消息格式示例； 长度限制4096字符"}
	room, err := client.SendChatRoomMsgToOne(chatroom)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(room)
	t.Log(string(d))
}
