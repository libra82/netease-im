package netease

import (
	"encoding/json"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

/**
 * 功能描述：网易云信 聊天室(聊天室是一项付费拓展能力，需要在选购IM基础功能的情况下增购。)
 *
 * Created with GoLand
 * @Author: david.tao
 * @CreateTime 2022-04-30 17:11
 *
 * 修改历史：(修改人，修改时间，修改原因/内容)
 * https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353
 **/

const (
	neteaseChatroomBaseURL               = "https://api.netease.im/nimserver"
	createChatroomPoint                  = neteaseChatroomBaseURL + "/chatroom/create.action"                  //创建聊天室
	getChatroomInfoPoint                 = neteaseChatroomBaseURL + "/chatroom/get.action"                     //查询聊天室信息
	getBatchChatroomInfoPoint            = neteaseChatroomBaseURL + "/chatroom/getBatch.action"                //批量查询聊天室信息
	updateChatroomInfoPoint              = neteaseChatroomBaseURL + "/chatroom/update.action"                  //更新聊天室信息
	toggleCloseStatChatroomPoint         = neteaseChatroomBaseURL + "/chatroom/toggleCloseStat.action"         //修改聊天室开/关闭状态
	setMemberRoleChatroomPoint           = neteaseChatroomBaseURL + "/chatroom/setMemberRole.action"           //设置聊天室内用户角色
	requestAddrChatroomPoint             = neteaseChatroomBaseURL + "/chatroom/requestAddr.action"             //请求聊天室地址
	sendMsgChatroomPoint                 = neteaseChatroomBaseURL + "/chatroom/sendMsg.action"                 //发送聊天室消息
	addRobotChatroomPoint                = neteaseChatroomBaseURL + "/chatroom/addRobot.action"                //往聊天室内添加机器人
	removeRobotChatroomPoint             = neteaseChatroomBaseURL + "/chatroom/removeRobot.action"             //从聊天室内删除机器人
	cleanRobotChatroomPoint              = neteaseChatroomBaseURL + "/chatroom/cleanRobot.action"              //清空聊天室机器人
	temporaryMuteChatroomPoint           = neteaseChatroomBaseURL + "/chatroom/temporaryMute.action"           //设置临时禁言状态
	muteRoomChatroomPoint                = neteaseChatroomBaseURL + "/chatroom/muteRoom.action"                //将聊天室整体禁言
	topnChatroomPoint                    = neteaseChatroomBaseURL + "/stats/chatroom/topn.action"              //查询聊天室统计指标TopN
	membersByPageChatroomPoint           = neteaseChatroomBaseURL + "/chatroom/membersByPage.action"           //分页获取成员列表
	queryMembersByRoleChatroomPoint      = neteaseChatroomBaseURL + "/chatroom/queryMembersByRole.action"      //根据角色获取固定成员列表
	queryMembersChatroomPoint            = neteaseChatroomBaseURL + "/chatroom/queryMembers.action"            //批量获取在线成员信息
	updateMyRoomRoleChatroomPoint        = neteaseChatroomBaseURL + "/chatroom/updateMyRoomRole.action"        //变更聊天室内的角色信息
	queryUserRoomIdsChatroomPoint        = neteaseChatroomBaseURL + "/chatroom/queryUserRoomIds.action"        //查询用户创建的开启状态聊天室列表
	updateInOutNotificationChatroomPoint = neteaseChatroomBaseURL + "/chatroom/updateInOutNotification.action" //关闭指定聊天室进出通知
	tagTemporaryMuteChatroomPoint        = neteaseChatroomBaseURL + "/chatroom/tagTemporaryMute.action"        //标签禁言
	tagMembersCountChatroomPoint         = neteaseChatroomBaseURL + "/chatroom/tagMembersCount.action"         //查询某个标签下的在线用户数
	tagMembersQueryChatroomPoint         = neteaseChatroomBaseURL + "/chatroom/tagMembersQuery.action"         //根据标签查询在线成员列表
	broadcastChatroomPoint               = neteaseChatroomBaseURL + "/chatroom/broadcast.action"               //聊天室全服广播消息
	recallChatroomPoint                  = neteaseChatroomBaseURL + "/chatroom/recall.action"                  //聊天室消息撤回
	sendMsgToSomeoneChatroomPoint        = neteaseChatroomBaseURL + "/chatroom/sendMsgToSomeone.action"        //发送聊天室定向消息

)

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#创建聊天室
/**
 *@param creator	String	是	聊天室属主的账号accid
 *@param name	String	是	聊天室名称，长度限制128个字符
 *@param announcement	String	否	公告，长度限制4096个字符
 *@param broadcasturl	String	否	直播地址，长度限制1024个字符
 *@param ext	String	否	扩展字段，最长4096字符
 *@param queuelevel	int	否	队列管理权限：0:所有人都有权限变更队列，1:只有主播管理员才能操作变更。默认0
 *@param bid	String	否	反垃圾业务ID，JSON字符串，{"textbid":"","picbid":""}，若不填则使用原来的反垃圾配置
 */
func (c *ImClient) CreateChatRoom(cr *ImChatRoomReq) (*ImChatRoomRes, error) {
	param := map[string]string{"creator": cr.Creator, "name": cr.Name}

	if len(cr.Announcement) > 0 {
		param["announcement"] = cr.Announcement
	}
	if len(cr.Broadcasturl) > 0 {
		param["broadcasturl"] = cr.Broadcasturl
	}
	if len(cr.Ext) > 0 {
		param["ext"] = cr.Ext
	}
	if cr.Queuelevel >= 0 {
		param["queuelevel"] = strconv.Itoa(cr.Queuelevel)
	}
	if len(cr.Bid) > 0 {
		param["bid"] = cr.Bid
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(createChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatRoomRes{}
	err = jsoniter.Unmarshal(*jsonRes["chatroom"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#查询聊天室信息
/**
 *@param roomid	long	是	聊天室id
 *@param needOnlineUserCount	String	否	是否需要返回在线人数，true或false，默认false
 */
func (c *ImClient) GetChatRoomInfo(cr *ImChatRoomInfoReq) (*ImChatRoomInfoRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid))}

	if len(cr.NeedOnlineUserCount) > 0 {
		param["needOnlineUserCount"] = cr.NeedOnlineUserCount
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(getChatroomInfoPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatRoomInfoRes{}
	err = jsoniter.Unmarshal(*jsonRes["chatroom"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#批量查询聊天室信息
/**
 *@param roomid	long	是	聊天室id
 *@param needOnlineUserCount	String	否	是否需要返回在线人数，true或false，默认false
 */
func (c *ImClient) GetBatchChatRoomInfo(cr *ImChatRoomBatchInfoReq) (*ImChatRoomBatchBatchInfoRes, error) {
	param := map[string]string{"roomids": cr.Roomids}

	if len(cr.NeedOnlineUserCount) > 0 {
		param["needOnlineUserCount"] = cr.NeedOnlineUserCount
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(getBatchChatroomInfoPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatrooms := []ImChatRoomInfoRes{}
	err = jsoniter.Unmarshal(*jsonRes["succRooms"], &chatrooms)
	if err != nil {
		return nil, err
	}
	var noExistRooms []int64 //不存在的聊天室id列表
	err = jsoniter.Unmarshal(*jsonRes["noExistRooms"], &noExistRooms)
	if err != nil {
		return nil, err
	}
	var failRooms []int64 //失败的聊天室id,有可能是查的时候有500错误
	err = jsoniter.Unmarshal(*jsonRes["failRooms"], &failRooms)
	if err != nil {
		return nil, err
	}
	chatroom := &ImChatRoomBatchBatchInfoRes{}
	chatroom.SuccRooms = chatrooms
	chatroom.NoExistRooms = noExistRooms
	chatroom.FailRooms = failRooms
	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#更新聊天室信息
/**
 *@param roomid	long	是	聊天室id
 *@param name	String	否	聊天室名称，长度限制128个字符
 *@param announcement	String	否	公告，长度限制4096个字符
 *@param broadcasturl	String	否	直播地址，长度限制1024个字符
 *@param ext	String	否	扩展字段，最长4096字符
 *@param needNotify	String	否	true或false,是否需要发送更新通知事件，默认true
 *@param notifyExt	String	否	通知事件扩展字段，长度限制2048
 *@param queuelevel	int	否	队列管理权限：0:所有人都有权限变更队列，1:只有主播管理员才能操作变更。默认0
 *@param bid	String	否	反垃圾业务ID，JSON字符串，{"textbid":"","picbid":""}，若不填则使用原来的反垃圾配置
 */
func (c *ImClient) UpdateChatRoom(cr *ImChatRoomUpdateReq) (*ImChatRoomRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid))}

	if len(cr.Name) > 0 {
		param["name"] = cr.Name
	}
	if len(cr.Announcement) > 0 {
		param["announcement"] = cr.Announcement
	}
	if len(cr.Broadcasturl) > 0 {
		param["broadcasturl"] = cr.Broadcasturl
	}
	if len(cr.Ext) > 0 {
		param["ext"] = cr.Ext
	}
	if len(cr.NeedNotify) > 0 {
		param["needNotify"] = cr.NeedNotify
	}
	if len(cr.NotifyExt) > 0 {
		param["notifyExt"] = cr.NotifyExt
	}
	if cr.Queuelevel >= 0 {
		param["queuelevel"] = strconv.Itoa(cr.Queuelevel)
	}
	if len(cr.Bid) > 0 {
		param["bid"] = cr.Bid
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(updateChatroomInfoPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatRoomRes{}
	err = jsoniter.Unmarshal(*jsonRes["chatroom"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#修改聊天室开/关闭状态
/**
 *@param roomid	long	是	聊天室id
 *@param operator	String	是	操作者账号，必须是创建者才可以操作
 *@param valid	String	是	true或false，false:关闭聊天室；true:打开聊天室
 */
func (c *ImClient) ChangeChatRoomState(cr *ImChatRoomStateReq) (*ImChatRoomRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "operator": cr.Operator, "valid": cr.Valid}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(toggleCloseStatChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatRoomRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#设置聊天室内用户角色
/**
 *@param roomid	long	是	聊天室id
 *@param operator	String	是	操作者账号accid
 *@param target	String	是	被操作者账号accid
 *@param opt	int	是	操作： 1: 设置为管理员，operator必须是创建者; 2:设置普通等级用户，operator必须是创建者或管理员; -1:设为黑名单用户，operator必须是创建者或管理员; -2:设为禁言用户，operator必须是创建者或管理员
 *@param optvalue	String	是	true或false，true:设置；false:取消设置； 执行“取消”设置后，若成员非禁言且非黑名单，则变成游客
 *@param notifyExt	String	否	通知扩展字段，长度限制2048，请使用json格式
 */
func (c *ImClient) SetChatRoomRole(cr *ImChatRoomRoleReq) (*ImChatRoomRoleRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "operator": cr.Operator, "target": cr.Target, "opt": strconv.Itoa(cr.Opt), "optvalue": cr.Optvalue}

	if len(cr.NotifyExt) > 0 {
		param["notifyExt"] = cr.NotifyExt
	}
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(setMemberRoleChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatRoomRoleRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#请求聊天室地址
/**
 *@param roomid	long	是	聊天室id
 *@param accid	String	是	进入聊天室的账号
 *@param clienttype	int	否	1:weblink（客户端为web端时使用）; 2:commonlink（客户端为非web端时使用）;3:wechatlink(微信小程序使用), 默认1
 *@param clientip	String	否	客户端ip，传此参数时，会根据用户ip所在地区，返回合适的地址rator必须是创建者或管理员; -2:设为禁言用户，operator必须是创建者或管理员
 */
func (c *ImClient) ReqChatRoomAddr(cr *ImChatRoomAddrReq) (*[]string, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "accid": cr.Accid}

	if cr.Clienttype > 0 {
		param["clienttype"] = strconv.Itoa(cr.Clienttype)
	}
	if len(cr.Clientip) > 0 {
		param["clientip"] = cr.Clientip
	}
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(requestAddrChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &[]string{}
	err = jsoniter.Unmarshal(*jsonRes["addr"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#请求聊天室地址
/**
 *@param roomid	long	是	聊天室id
 *@param msgId	String	是	客户端消息id，使用uuid等随机串，msgId相同的消息会被客户端去重
 *@param fromAccid	String	是	消息发出者的账号accid
 *@param msgType	int	是	消息类型：
 * 0: 表示文本消息， 1: 表示图片， 2: 表示语音， 3: 表示视频， 4: 表示地理位置信息， 6: 表示文件， 10: 表示Tips消息， 100: 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
 *@param subType	int	否	自定义消息子类型，大于0
 *@param resendFlag	int	否	重发消息标记，0：非重发消息，1：重发消息，如重发消息会按照msgid检查去重逻辑
 *@param attach	String	否	文本消息：填写消息文案; 其它类型消息，请参考 消息格式示例； 长度限制4096字符
 *@param ext	String	否	消息扩展字段，内容可自定义，请使用JSON格式，长度限制4096字符
 *@param skipHistory	int	否	是否跳过存储云端历史，0：不跳过，即存历史消息；1：跳过，即不存云端历史；默认0
 *@param abandonRatio	int	否	可选，消息丢弃的概率。取值范围[0-9999]； 其中0代表不丢弃消息，9999代表99.99%的概率丢弃消息，默认不丢弃； 注意如果填写了此参数，下面的highPriority参数则会无效； 此参数可用于流控特定业务类型的消息。
 *@param highPriority	Boolean	否	可选，true表示是高优先级消息，云信会优先保障投递这部分消息；false表示低优先级消息。默认false。 强烈建议应用恰当选择参数，以便在必要时，优先保障应用内的高优先级消息的投递。若全部设置为高优先级，则等于没有设置，单个聊天室最多支持每秒10条的高优先级消息，超过的会转为普通消息。 高优先级消息可以设置进入后重发，见needHighPriorityMsgResend参数
 *@param needHighPriorityMsgResend	Boolean	否	可选，true表示会重发消息，false表示不会重发消息。默认true。注:若设置为true， 用户离开聊天室之后重新加入聊天室，在有效期内还是会收到发送的这条消息，目前有效期默认30s。在没有配置highPriority时needHighPriorityMsgResend不生效。
 *@param useYidun	int	否	可选，单条消息是否使用易盾反垃圾，可选值为0。 0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。 若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
 *@param yidunAntiCheating	String	否	可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
 *@param bid	String	否	可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
 *@param antispam	String	否	对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。 true或false, 默认false。 只对消息类型为：100 自定义消息类型 的消息生效。
 *@param notifyTargetTags	String	否	可选，标签表达式，最长128个字符
 *@param antispamCustom	String	否	在antispam参数为true时生效。 自定义的反垃圾检测内容, JSON格式，长度限制同body字段，不能超过5000字符，要求antispamCustom格式如下：{"type":1,"data":"custom content"}
 *字段说明： 1. type: 1：文本，2：图片。 2. data: 文本内容or图片地址。
 *@param env	String	否	所属环境，根据env可以配置不同的抄送地址
 *
 * 为保证用户体验（如避免服务器过载），目前针对消息接收，有两套流控机制。第一套针对普通消息，聊天室用户每秒至多可接收20条，超过部分会因为流控随机丢弃。第二套针对高优先级消息，每秒至多接收10条，超过部分无法保证不丢失。
 * 为避免丢失重要消息（通常为服务端消息），可将发送聊天室消息的 HighPriority 参数设置为 true 实现高优先级接收服务端消息，进而保证高优先级消息流控上限内（每秒10条）的重要消息不丢失。详情请参见本节参数说明中的 HighPriority 参数说明。
 */
func (c *ImClient) SendChatRoomMsg(cr *ImChatRoomSendMsgReq) (*ImChatRoomSendMsgRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "msgId": cr.MsgId, "fromAccid": cr.FromAccid, "msgType": strconv.Itoa(cr.MsgType)}

	if cr.SubType > 0 {
		param["subType"] = strconv.Itoa(cr.SubType)
	}
	if cr.ResendFlag >= 0 {
		param["resendFlag"] = strconv.Itoa(cr.ResendFlag)
	}
	if len(cr.Attach) > 0 {
		param["attach"] = cr.Attach
	}
	if len(cr.Ext) > 0 {
		param["ext"] = cr.Ext
	}
	if cr.SkipHistory >= 0 {
		param["skipHistory"] = strconv.Itoa(cr.ResendFlag)
	}
	if cr.AbandonRatio >= 0 {
		param["abandonRatio"] = strconv.Itoa(cr.AbandonRatio)
	}
	if cr.HighPriority { //默认false
		param["highPriority"] = "true"
	}
	if !cr.NeedHighPriorityMsgResend { //默认true
		param["needHighPriorityMsgResend"] = "false"
	}
	if cr.UseYidun >= 0 {
		param["useYidun"] = strconv.Itoa(cr.UseYidun)
	}
	if len(cr.YidunAntiCheating) > 0 {
		param["yidunAntiCheating"] = cr.YidunAntiCheating
	}
	if len(cr.Bid) > 0 {
		param["bid"] = cr.Bid
	}
	if len(cr.Antispam) > 0 {
		param["antispam"] = cr.Antispam
	}
	if len(cr.NotifyTargetTags) > 0 {
		param["notifyTargetTags"] = cr.NotifyTargetTags
	}
	if len(cr.AntispamCustom) > 0 {
		param["antispamCustom"] = cr.AntispamCustom
	}
	if len(cr.Env) > 0 {
		param["env"] = cr.Env
	}
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(sendMsgChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatRoomSendMsgRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#往聊天室内添加机器人
/**
 *@param roomid	long	是	聊天室id
 *@param accids	JSONArray	是	机器人账号accid列表，必须是有效账号，账号数量上限100个
 *@param roleExt	String	否	机器人信息扩展字段，请使用json格式，长度4096字符
 *@param notifyExt	String	否	机器人进入聊天室通知的扩展字段，请使用json格式，长度2048字符
 */
func (c *ImClient) AddChatRoomRobot(cr *ImChatroomAddRobotReq) (*ImChatroomAddRobotRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "accids": cr.Accids}

	if len(cr.RoleExt) > 0 {
		param["roleExt"] = cr.RoleExt
	}
	if len(cr.NotifyExt) > 0 {
		param["notifyExt"] = cr.NotifyExt
	}
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(addRobotChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatroomAddRobotRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#从聊天室内删除机器人
/**
 *@param roomid	long	是	聊天室id
 *@param accids	JSONArray	是	机器人账号accid列表，必须是有效账号，账号数量上限100个
 */
func (c *ImClient) DelChatRoomRobot(cr *ImChatroomDelRobotReq) (*ImChatroomDelRobotRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "accids": cr.Accids}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(removeRobotChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatroomDelRobotRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#清空聊天室机器人
/**
 *@param roomid	long	是	聊天室id
 *@param notify	boolean	否	是否发送退出聊天室通知消息，默认为false
 */
func (c *ImClient) CleanChatRoomRobot(cr *ImChatroomCleanRobotReq) (*ImChatroomCleanRobotRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid))}

	if cr.Notify { //默认为false
		param["notify"] = "true"
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(cleanRobotChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatroomCleanRobotRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#设置临时禁言状态
/**
 *@param roomid	long	是	聊天室id
 *@param operator	String	是	操作者accid,必须是管理员或创建者
 *@param target	String	是	被禁言的目标账号accid
 *@param muteDuration	long	是	0:解除禁言;>0设置禁言的秒数，不能超过2592000秒(30天)
 *@param needNotify	String	否	操作完成后是否需要发广播，true或false，默认true
 *@param notifyExt	String	否	通知广播事件中的扩展字段，长度限制2048字符
 */
func (c *ImClient) MuteChatRoomTemp(cr *ImChatroomMuteReq) (*ImChatroomMuteRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "operator": cr.Operator, "target": cr.Target, "muteDuration": strconv.Itoa(int(cr.MuteDuration))}

	if len(cr.NeedNotify) > 0 {
		param["needNotify"] = cr.NeedNotify
	}
	if len(cr.NotifyExt) > 0 {
		param["notifyExt"] = cr.NotifyExt
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(temporaryMuteChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatroomMuteRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#将聊天室整体禁言
/**
 *@param roomid	long	是	聊天室id
 *@param operator	String	是	操作者accid,必须是管理员或创建者
 *@param mute	String	是	true或false
 *@param needNotify	String	否	操作完成后是否需要发广播，true或false，默认true
 *@param notifyExt	String	否	通知广播事件中的扩展字段，长度限制2048字符
 */
func (c *ImClient) MuteChatRoom(cr *ImChatroomMuteRoomReq) (*ImChatroomMuteRoomRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "operator": cr.Operator, "mute": cr.Mute}

	if len(cr.NeedNotify) > 0 {
		param["needNotify"] = cr.NeedNotify
	}
	if len(cr.NotifyExt) > 0 {
		param["notifyExt"] = cr.NotifyExt
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(muteRoomChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatroomMuteRoomRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#查询聊天室统计指标TopN
/**
 *@param topn	int	否	topn值，可选值 1~500，默认值100
 *@param timestamp	long	否	需要查询的指标所在的时间坐标点，不提供则默认当前时间，单位秒/毫秒皆可
 *@param period	String	否	统计周期，可选值包括 hour/day, 默认hour
 *@param orderby	String	否	取排序值,可选值 active/enter/message,分别表示按日活排序，进入人次排序和消息数排序， 默认active

接口描述
1、根据时间戳，按指定周期列出聊天室相关指标的TopN列表
2、当天的统计指标需要到第二天才能查询
3、仅支持查询最近30天的统计指标
*/
func (c *ImClient) StatChatRoomTopn(cr *ImChatroomTopnReq) (*[]ImChatroomTopnRes, error) {
	param := map[string]string{}

	if cr.Topn > 0 {
		param["topn"] = strconv.Itoa(cr.Topn)
	}
	if cr.Timestamp > 0 {
		param["timestamp"] = strconv.Itoa(int(cr.Timestamp))
	}
	if len(cr.Period) > 0 {
		param["period"] = cr.Period
	}
	if len(cr.Orderby) > 0 {
		param["orderby"] = cr.Orderby
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(topnChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &[]ImChatroomTopnRes{}
	err = jsoniter.Unmarshal(*jsonRes["data"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#分页获取成员列表
/**
 *@param roomid	long	是	聊天室id
 *@param type	int	是	需要查询的成员类型,0:固定成员;1:非固定成员;2:仅返回在线的固定成员
 *@param endtime	long	是	单位毫秒，按时间倒序最后一个成员的时间戳,0表示系统当前时间
 *@param limit	long	是	返回条数，<=100
 */
func (c *ImClient) PageChatRoomMembers(cr *ImChatroomMembersReq) (*ImChatroomMembersRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "type": strconv.Itoa(cr.Type), "endtime": strconv.Itoa(int(cr.Endtime)), "limit": strconv.Itoa(int(cr.Limit))}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(membersByPageChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatroomMembersRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#根据角色获取固定成员列表
/**
 *@param roomid	long	是	聊天室id
 *@param roles	String	是	设置需要获取的角色,格式示例： {"creator": true,"manager": true,"blacklist": false,"mute": false}
 *  字段说明：
 *  1、creator：聊天室创建者
 *  2、manager：聊天室管理员
 *  3、blacklist：黑名单用户
 *  4、mute：被禁言用户
 *  说明：设置为false或不设置表示不获取相应的角色信息
 */
func (c *ImClient) GetChatRoomMembersByRole(cr *ImChatroomMembersByRoleReq) (*ImChatroomMembersByRoleRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "roles": cr.Roles}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(queryMembersByRoleChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatroomMembersByRoleRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#批量获取在线成员信息
/**
 *@param roomid	long	是	聊天室id
 *@param accids	JSONArray	是	\["abc","def"\], 账号列表，最多200条
 */
func (c *ImClient) GetChatRoomMembersBatch(cr *ImChatroomMembersBatchReq) (*ImChatroomMembersBatchRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "accids": cr.Accids}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(queryMembersChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatroomMembersBatchRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#变更聊天室内的角色信息
/**
 *@param roomid	long	是	聊天室id
 *@param accid	String	是	需要变更角色信息的accid
 *@param save	boolean	否	变更的信息是否需要持久化，默认false，仅对聊天室固定成员生效
 *@param needNotify	boolean	否	是否需要做通知
 *@param notifyExt	String	否	通知的内容，长度限制2048
 *@param nick	String	否	聊天室室内的角色信息：昵称，不超过64个字符
 *@param avator	String	否	聊天室室内的角色信息：头像
 *@param ext	String	否	聊天室室内的角色信息：开发者扩展字段
 *@param bid	String	否	反垃圾业务ID，JSON字符串，{"textbid":"","picbid":""}，若不填则使用原来的反垃圾配置
 */
func (c *ImClient) UpdateChatRoomMemberRole(cr *ImChatroomChangeRoleReq) (bool, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "accid": cr.Accid}

	if cr.Save { //默认false
		param["save"] = "true"
	}
	if cr.NeedNotify { //默认false
		param["needNotify"] = "true"
	}
	if len(cr.NotifyExt) > 0 {
		param["notifyExt"] = cr.NotifyExt
	}
	if len(cr.Nick) > 0 {
		param["nick"] = cr.Nick
	}
	if len(cr.Avator) > 0 {
		param["avator"] = cr.Avator
	}
	if len(cr.Ext) > 0 {
		param["ext"] = cr.Ext
	}
	if len(cr.Bid) > 0 {
		param["bid"] = cr.Bid
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(updateMyRoomRoleChatroomPoint)
	if err != nil {
		return false, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return false, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return false, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return false, errors.New(msg)
	}

	//chatroom := &ImChatroomMembersBatchRes{}
	//err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	//if err != nil {
	//	return nil, err
	//}

	return true, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#查询用户创建的开启状态聊天室列表
/**
 *@param creator	String	是	聊天室创建者accid
 */
func (c *ImClient) GetChatRoomUserRoomIds(cr *ImChatroomUserRoomIdsReq) (*ImChatroomUserRoomIdsRes, error) {
	param := map[string]string{"creator": cr.Creator}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(queryUserRoomIdsChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatroomUserRoomIdsRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#关闭指定聊天室进出通知
/**
 *@param roomid	long	是	聊天室id
 *@param close	boolean	是	true/false, 是否关闭进出通知
 */
func (c *ImClient) CloseChatRoomInOutNotify(cr *ImChatroomInOutNotifyReq) (bool, error) {
	close := "false"
	if cr.Close {
		close = "true"
	}
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "close": close}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(updateInOutNotificationChatroomPoint)
	if err != nil {
		return false, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return false, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return false, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return false, errors.New(msg)
	}

	//chatroom := &ImChatroomMembersBatchRes{}
	//err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	//if err != nil {
	//	return nil, err
	//}

	return true, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#标签禁言
/**
 *@param roomid	long	是	聊天室ID
 *@param operator	string	是	操作者accid，必须是创建者或者管理员
 *@param targetTag	string	是	目标标签
 *@param needNotify	boolean	否	true/false，是否发送禁言通知，默认true
 *@param notifyExt	string	否	禁言通知通知扩展字段
 *@param muteDuration	int	是	禁言时长，单位秒，最长30天，若等于0表示取消禁言
 *@param notifyTargetTags	string	否	禁言通知的目标标签表达式，若缺失则发送给设置了targetTag的人
 */
func (c *ImClient) MuteChatRoomByTag(cr *ImChatroomTagMuteReq) (*ImChatroomTagMuteRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "operator": cr.Operator, "targetTag": cr.TargetTag, "muteDuration": strconv.Itoa(cr.MuteDuration)}

	if !cr.NeedNotify { //默认true
		param["needNotify"] = "false"
	}
	if len(cr.NotifyExt) > 0 {
		param["notifyExt"] = cr.NotifyExt
	}
	if len(cr.NotifyTargetTags) > 0 {
		param["notifyTargetTags"] = cr.NotifyTargetTags
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(tagTemporaryMuteChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatroomTagMuteRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#查询某个标签下的在线用户数
/**
 *@param roomid	long	是	聊天室ID
 *@param tag	string	是	标签
 */
func (c *ImClient) CountChatRoomMemberByTag(cr *ImChatroomTagMemberCountReq) (*ImChatroomTagMemberCountRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "tag": cr.Tag}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(tagMembersCountChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatroomTagMemberCountRes{}
	err = jsoniter.Unmarshal(*jsonRes["data"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#根据标签查询在线成员列表
/**
 *@param roomid	long	是	聊天室ID
 *@param tag	string	是	标签
 *@param endTime	long	是	起始时间，逆序查询，若传0则表示从当前时间往前查
 *@param limit	int	是	条数，最多100
 *
 *根据标签查询在线成员列表，注意多端登录情况下会返回多条记录
 */
func (c *ImClient) PageChatRoomMemberByTag(cr *ImChatroomTagMembersReq) (*ImChatroomTagMembersRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "tag": cr.Tag, "endTime": strconv.Itoa(int(cr.EndTime)), "limit": strconv.Itoa(cr.Limit)}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(tagMembersQueryChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatroomTagMembersRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#聊天室全服广播消息
/**
 *@param msgId	String	是	客户端消息id，使用uuid等随机串，msgId相同的消息会被客户端去重
 *@param fromAccid	String	是	消息发出者的账号accid
 *@param msgType	int	是	消息类型：
 * 0: 表示文本消息， 1: 表示图片， 2: 表示语音， 3: 表示视频， 4: 表示地理位置信息， 6: 表示文件， 10: 表示Tips消息， 100: 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
 *@param subType	int	否	自定义消息子类型，大于0
 *@param resendFlag	int	否	重发消息标记，0：非重发消息，1：重发消息，如重发消息会按照msgid检查去重逻辑
 *@param attach	String	否	文本消息：填写消息文案; 其它类型消息，请参考 消息格式示例； 长度限制4096字符
 *@param ext	String	否	消息扩展字段，内容可自定义，请使用JSON格式，长度限制4096字符
 *@param highPriority	Boolean	否	可选，true表示是高优先级消息，云信会优先保障投递这部分消息；false表示低优先级消息。默认false。 强烈建议应用恰当选择参数，以便在必要时，优先保障应用内的高优先级消息的投递。若全部设置为高优先级，则等于没有设置，单个聊天室最多支持每秒10条的高优先级消息，超过的会转为普通消息。 高优先级消息可以设置进入后重发，见needHighPriorityMsgResend参数
 *@param needHighPriorityMsgResend	Boolean	否	可选，true表示会重发消息，false表示不会重发消息。默认true。注:若设置为true， 用户离开聊天室之后重新加入聊天室，在有效期内还是会收到发送的这条消息，目前有效期默认30s。在没有配置highPriority时needHighPriorityMsgResend不生效。
 *@param useYidun	int	否	可选，单条消息是否使用易盾反垃圾，可选值为0。 0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。 若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
 *@param yidunAntiCheating	String	否	可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
 *@param bid	String	否	可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
 *@param antispam	String	否	对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。 true或false, 默认false。 只对消息类型为：100 自定义消息类型 的消息生效。
 *@param notifyTargetTags	String	否	可选，标签表达式，最长128个字符
 *@param antispamCustom	String	否	在antispam参数为true时生效。 自定义的反垃圾检测内容, JSON格式，长度限制同body字段，不能超过5000字符，要求antispamCustom格式如下：{"type":1,"data":"custom content"}
 *字段说明： 1. type: 1：文本，2：图片。 2. data: 文本内容or图片地址。
 *@param env	String	否	所属环境，根据env可以配置不同的抄送地址
 *
 * 接口描述
1、聊天室全服广播消息，会广播给该应用下所有聊天室的所有成员
2、注意广播消息只能在线广播，不会存历史
3、聊天室全服广播消息1分钟最多发送10条
4、本功能需要开通，请联系商务
*/
func (c *ImClient) BroadcastChatRoom(cr *ImChatRoomBroadcastReq) (*ImChatRoomBroadcastRes, error) {
	param := map[string]string{"msgId": cr.MsgId, "fromAccid": cr.FromAccid, "msgType": strconv.Itoa(cr.MsgType)}

	if cr.SubType > 0 {
		param["subType"] = strconv.Itoa(cr.SubType)
	}
	if cr.ResendFlag >= 0 {
		param["resendFlag"] = strconv.Itoa(cr.ResendFlag)
	}
	if len(cr.Attach) > 0 {
		param["attach"] = cr.Attach
	}
	if len(cr.Ext) > 0 {
		param["ext"] = cr.Ext
	}
	if cr.HighPriority { //默认false
		param["highPriority"] = "true"
	}
	if !cr.NeedHighPriorityMsgResend { //默认true
		param["needHighPriorityMsgResend"] = "false"
	}
	if cr.UseYidun >= 0 {
		param["useYidun"] = strconv.Itoa(cr.UseYidun)
	}
	if len(cr.YidunAntiCheating) > 0 {
		param["yidunAntiCheating"] = cr.YidunAntiCheating
	}
	if len(cr.Bid) > 0 {
		param["bid"] = cr.Bid
	}
	if len(cr.Antispam) > 0 {
		param["antispam"] = cr.Antispam
	}
	if len(cr.NotifyTargetTags) > 0 {
		param["notifyTargetTags"] = cr.NotifyTargetTags
	}
	if len(cr.AntispamCustom) > 0 {
		param["antispamCustom"] = cr.AntispamCustom
	}
	if len(cr.Env) > 0 {
		param["env"] = cr.Env
	}
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(broadcastChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatRoomBroadcastRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#聊天室消息撤回
/**
 *@param roomid	long	是	聊天室id
 *@param msgTimetag	long	是	被撤回消息的时间戳
 *@param fromAcc	String	是	被撤回消息的消息发送者accid
 *@param msgId	String	是	被撤回消息的消息id
 *@param operatorAcc	String	是	消息撤回的操作者accid
 *@param notifyExt	String	否	消息撤回的通知扩展字段，最长1024字符
 *
 * 撤回聊天室内发送的消息，撤回后对应消息的云端历史记录也将一并删除，需要云信IM SDK升级到8.7.0及以上版本。
 */
func (c *ImClient) ReCallChatRoomMsg(cr *ImChatRoomReCallReq) (bool, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "msgTimetag": strconv.Itoa(int(cr.MsgTimetag)), "fromAcc": cr.FromAcc, "msgId": cr.MsgId, "operatorAcc": cr.OperatorAcc}

	if len(cr.NotifyExt) > 0 {
		param["notifyExt"] = cr.NotifyExt
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(recallChatroomPoint)
	if err != nil {
		return false, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return false, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return false, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return false, errors.New(msg)
	}

	//chatroom := &ImChatroomMembersBatchRes{}
	//err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	//if err != nil {
	//	return nil, err
	//}

	return true, nil
}

//https://doc.yunxin.163.com/docs/TM5MzM5Njk/TYyMDI1MTg?platformId=60353#发送聊天室定向消息
/**
 *@param roomid	long	是	聊天室id
 *@param msgId	String	是	客户端消息id，使用uuid等随机串，msgId相同的消息会被客户端去重
 *@param fromAccid	String	是	消息发出者的账号accid
 *@param toAccids	JSONArray	是	消息接收者accid列表，最大100个
 *@param msgType	int	是	消息类型：
 * 0: 表示文本消息， 1: 表示图片， 2: 表示语音， 3: 表示视频， 4: 表示地理位置信息， 6: 表示文件， 10: 表示Tips消息， 100: 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
 *@param subType	int	否	自定义消息子类型，大于0
 *@param resendFlag	int	否	重发消息标记，0：非重发消息，1：重发消息，如重发消息会按照msgid检查去重逻辑
 *@param attach	String	否	文本消息：填写消息文案; 其它类型消息，请参考 消息格式示例； 长度限制4096字符
 *@param ext	String	否	消息扩展字段，内容可自定义，请使用JSON格式，长度限制4096字符
 *@param useYidun	int	否	可选，单条消息是否使用易盾反垃圾，可选值为0。 0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。 若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
 *@param yidunAntiCheating	String	否	可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
 *@param bid	String	否	可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
 *@param antispam	String	否	对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。 true或false, 默认false。 只对消息类型为：100 自定义消息类型 的消息生效。
 *@param antispamCustom	String	否	在antispam参数为true时生效。 自定义的反垃圾检测内容, JSON格式，长度限制同body字段，不能超过5000字符，要求antispamCustom格式如下：{"type":1,"data":"custom content"}
 *字段说明： 1. type: 1：文本，2：图片。 2. data: 文本内容or图片地址。
 *@param env	String	否	所属环境，根据env可以配置不同的抄送地址
 *
 * 往聊天室内某些人发消息  备注：聊天室定向消息不会存历史
 */
func (c *ImClient) SendChatRoomMsgToOne(cr *ImChatRoomSendMsgToOneReq) (*ImChatRoomSendMsgToOneRes, error) {
	param := map[string]string{"roomid": strconv.Itoa(int(cr.Roomid)), "msgId": cr.MsgId, "fromAccid": cr.FromAccid, "toAccids": cr.ToAccids, "msgType": strconv.Itoa(cr.MsgType)}

	if cr.SubType > 0 {
		param["subType"] = strconv.Itoa(cr.SubType)
	}
	if cr.ResendFlag >= 0 {
		param["resendFlag"] = strconv.Itoa(cr.ResendFlag)
	}
	if len(cr.Attach) > 0 {
		param["attach"] = cr.Attach
	}
	if len(cr.Ext) > 0 {
		param["ext"] = cr.Ext
	}
	if cr.UseYidun >= 0 {
		param["useYidun"] = strconv.Itoa(cr.UseYidun)
	}
	if len(cr.YidunAntiCheating) > 0 {
		param["yidunAntiCheating"] = cr.YidunAntiCheating
	}
	if len(cr.Bid) > 0 {
		param["bid"] = cr.Bid
	}
	if len(cr.Antispam) > 0 {
		param["antispam"] = cr.Antispam
	}
	if len(cr.AntispamCustom) > 0 {
		param["antispamCustom"] = cr.AntispamCustom
	}
	if len(cr.Env) > 0 {
		param["env"] = cr.Env
	}
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(sendMsgToSomeoneChatroomPoint)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	chatroom := &ImChatRoomSendMsgToOneRes{}
	err = jsoniter.Unmarshal(*jsonRes["desc"], chatroom)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}
