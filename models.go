package netease

//TokenInfo 云通信Token
type TokenInfo struct {
	Token string `json:"token"`
	Accid string `json:"accid"`
	Name  string `json:"name"`
}

//ImUser 云通信用户
type ImUser struct {
	ID        string //网易云通信ID，最大长度32字符，必须保证一个APP内唯一（只允许字母、数字、半角下划线_、@、半角点以及半角-组成，不区分大小写，会统一小写处理，请注意以此接口返回结果中的accid为准）。
	Name      string //网易云通信ID昵称，最大长度64字符，用来PUSH推送时显示的昵称
	Propertys string
	IconURL   string //网易云通信ID头像URL，第三方可选填，最大长度1024
	Token     string
	Sign      string //用户签名，最大长度256字符
	Email     string //用户email，最大长度64字符
	Birthday  string //用户生日，最大长度16字符
	Mobile    string //用户mobile，最大长度32字符
	Gender    int    //用户性别，0表示未知，1表示男，2女表示女
	Extension string //用户名片扩展字段，最大长度1024字符
}

//ImSendMessageOption .
type ImSendMessageOption struct {
	Antispam         bool            //对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。true或false, 默认false。只对消息类型为：100 自定义消息类型 的消息生效。
	AntispamCustom   *AntiSpamCustom //在antispam参数为true时生效。
	Option           *MessageOption  //发消息时特殊指定的行为选项
	Pushcontent      string          //ios推送内容，不超过150字符，option选项中允许推送（push=true），此字段可以指定推送内容
	Payload          string          //ios 推送对应的payload,必须是JSON,不能超过2k字符
	Extension        string          //开发者扩展字段，长度限制1024字符
	ForcePushList    []string        //发送群消息时的强推（@操作）用户列表，格式为JSONArray，如["accid1","accid2"]。若forcepushall为true，则forcepushlist为除发送者外的所有有效群成员
	ForcePushContent string          //发送群消息时，针对强推（@操作）列表forcepushlist中的用户，强制推送的内容
	ForcePushAll     bool            //发送群消息时，强推（@操作）列表是否为群里除发送者外的所有有效成员，true或false，默认为false
	Bid              string          //可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
}

//ImSendAttachMessageOption .
type ImSendAttachMessageOption struct {
	Pushcontent string         //iOS推送内容，第三方自己组装的推送内容,不超过150字符
	Payload     string         //ios 推送对应的payload,必须是JSON,不能超过2k字符
	Sound       string         //如果有指定推送，此属性指定为客户端本地的声音文件名，长度不要超过30个字符，如果不指定，会使用默认声音
	Save        int            //1表示只发在线，2表示会存离线，其他会报414错误。默认会存离线
	Option      *MessageOption //发消息时特殊指定的行为选项
}

//AntiSpamCustom 自定义的反垃圾检测内容, JSON格式，不能超过5000字符
type AntiSpamCustom struct {
	Type int    `json:"type"` //1：文本，2：图片。
	Data string `json:"data"` // 文本内容or图片地址
}

//MessageOption 发消息时特殊指定的行为选项
type MessageOption struct {
	Roam         *bool `json:"roam,omitempty"`         //该消息是否需要漫游，默认true（需要app开通漫游消息功能）
	History      *bool `json:"history,omitempty"`      //该消息是否存云端历史，默认true
	Sendersync   *bool `json:"sendersync,omitempty"`   //该消息是否需要发送方多端同步，默认true
	Push         *bool `json:"push,omitempty"`         //该消息是否需要APNS推送或安卓系统通知栏推送，默认true
	Route        *bool `json:"route,omitempty"`        //该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能)
	Badge        *bool `json:"badge,omitempty"`        //该消息是否需要计入到未读计数中，默认true
	NeedPushNick *bool `json:"needPushNick,omitempty"` //推送文案是否需要带上昵称，不设置该参数时默认true
	Persistent   *bool `json:"persistent,omitempty"`   //是否需要存离线消息，不设置该参数时默认true
}

//TextMessage 文本消息
type TextMessage struct {
	Message string `json:"msg"`
}

//ImageMessage 图片消息
type ImageMessage struct {
	Name      string `json:"name"` //图片name
	Md5       string `json:"md5"`  //图片文件md5
	URL       string `json:"url"`  //生成的url
	Extension string `json:"ext"`  //图片后缀
	Width     uint   `json:"w"`    //宽
	Height    uint   `json:"h"`    //高
	Size      uint   `json:"size"` //图片大小
}

//VoiceMessage 语音消息
type VoiceMessage struct {
	Duration  uint   `json:"dur"`  //语音持续时长ms
	Md5       string `json:"md5"`  //语音文件md5
	URL       string `json:"url"`  //生成的url
	Extension string `json:"ext"`  //语音消息格式，只能是aac格式
	Size      uint   `json:"size"` //语音文件大小
}

//VideoMessage 视频消息
type VideoMessage struct {
	Duration  uint   `json:"dur"`  //视频持续时长ms
	Md5       string `json:"md5"`  //视频文件md5
	URL       string `json:"url"`  //生成的url
	Width     uint   `json:"w"`    //宽
	Height    uint   `json:"h"`    //高
	Extension string `json:"ext"`  //视频格式
	Size      uint   `json:"size"` //视频文件大小
}

//LocationMessage 位置信息
type LocationMessage struct {
	Title     string  `json:"title"` //地理位置title
	Longitude float64 `json:"lng"`   //经度
	Latitude  float64 `json:"lat"`   //纬度
}

//FileMessage 文件消息
type FileMessage struct {
	Name      string `json:"name"` //文件名
	Md5       string `json:"md5"`  //图片文件md5
	URL       string `json:"url"`  //生成的url
	Extension string `json:"ext"`  //语音消息格式，只能是aac格式
	Size      uint   `json:"size"` //语音文件大小
}

//LoginEventCopyInfo 登录事件消息抄送
type LoginEventCopyInfo struct {
	EventType  string `json:"eventType"`  //值为2，表示是登录事件的消息
	AcctID     string `json:"accid"`      //发生登录事件的用户帐号，字符串类型
	IPAdrees   string `json:"clientIp"`   //登录时的ip地址
	ClientType string `json:"clientType"` //客户端类型： AOS、IOS、PC、WINPHONE、WEB、REST，字符串类型
	Code       string `json:"code"`       //登录事件的返回码，可转为Integer类型数据
	SdkVersion string `json:"sdkVersion"` //当前sdk的版本信息，字符串类型
	Time       string `json:"timestamp"`  //登录事件发生时的时间戳，可转为Long型数据
}

//LogoutEventCopyInfo 登出事件消息抄送
type LogoutEventCopyInfo struct {
	EventType  string `json:"eventType"`  //值为3，表示是登出事件的消息
	AcctID     string `json:"accid"`      //发生登出事件的用户帐号，字符串类型
	IPAdrees   string `json:"clientIp"`   //登出时的ip地址
	ClientType string `json:"clientType"` //客户端类型： AOS、IOS、PC、WINPHONE、WEB、REST，字符串类型
	Code       string `json:"code"`       //登出事件的返回码，可转为Integer类型数据
	SdkVersion string `json:"sdkVersion"` //当前sdk的版本信息，字符串类型
	Time       string `json:"timestamp"`  //登出事件发生时的时间戳，可转为Long型数据
}

//SenssionCopyInfo 会话类型信息抄送
type SenssionCopyInfo struct {
	EventType      string `json:"eventType"`      //值为1，表示是会话类型的消息
	ConvType       string `json:"convType"`       //会话具体类型：PERSON（二人会话数据）、TEAM（群聊数据）、	CUSTOM_PERSON（个人自定义系统通知）、CUSTOM_TEAM（群组自定义系统通知），字符串类型
	To             string `json:"to"`             //若convType为PERSON或CUSTOM_PERSON，则to为消息接收者的用户账号，字符串类型；若convType为TEAM或CUSTOM_TEAM，则to为tid，即群id，可转为Long型数据
	FromAccount    string `json:"fromAccount"`    //消息发送者的用户账号，字符串类型
	FromClientType string `json:"fromClientType"` //发送客户端类型： AOS、IOS、PC、WINPHONE、WEB、REST，字符串类型
	FromDeviceID   string `json:"fromDeviceId"`   //发送设备id，字符串类型
	FromNick       string `json:"fromNick"`       //发送方昵称，字符串类型
	MsgTimestamp   string `json:"msgTimestamp"`   //消息发送时间，字符串类型
	MsgType        string `json:"msgType"`        //会话具体类型PERSON、TEAM对应的通知消息类型:EXT、PICTURE、AUDIO、VIDEO、LOCATION 、NOTIFICATION、FILE、 //文件消息NETCALL_AUDIO、 //网络电话音频聊天 	NETCALL_VEDIO、 //网络电话视频聊天 	DATATUNNEL_NEW、 //新的数据通道请求通知 	TIPS、 //提示类型消息 	CUSTOM //自定义消息		会话具体类型CUSTOM_PERSON对应的通知消息类型：	FRIEND_ADD、 //加好友	FRIEND_DELETE、 //删除好友	CUSTOM_P2P_MSG、 //个人自定义系统通知		会话具体类型CUSTOM_TEAM对应的通知消息类型：	TEAM_APPLY、 //申请入群	TEAM_APPLY_REJECT、 //拒绝入群申请	TEAM_INVITE、 //邀请进群	TEAM_INVITE_REJECT、 //拒绝邀请	TLIST_UPDATE、 //群信息更新 	CUSTOM_TEAM_MSG、 //群组自定义系统通知
	Body           string `json:"body"`           //消息内容，字符串类型
	Attach         string `json:"attach"`         //附加消息，字符串类型
	MsgidClient    string `json:"msgidClient"`    //客户端生成的消息id，仅在convType为PERSON或TEAM含此字段，字符串类型
	MsgidServer    string `json:"msgidServer"`    //服务端生成的消息id，可转为Long型数据
	ResendFlag     string `json:"resendFlag"`     //重发标记：0不是重发, 1是重发。仅在convType为PERSON或TEAM时含此字段，可转为Integer类型数据
	CustomSafeFlag string `json:"customSafeFlag"` //自定义系统通知消息是否存离线:0：不存，1：存。仅在convType为CUSTOM_PERSON或CUSTOM_TEAM时含此字段，可转为Integer类型数据
	CustomApnsText string `json:"customApnsText"` //自定义系统通知消息推送文本。仅在convType为CUSTOM_PERSON或CUSTOM_TEAM时含此字段，字符串类型
	TMembers       string `json:"tMembers"`       //跟本次群操作有关的用户accid，仅在convType为TEAM或CUSTOM_TEAM时含此字段，字符串类型
	Ext            string `json:"ext"`            //消息扩展字段
	Antispam       string `json:"antispam"`       //标识是否被反垃圾，仅在被反垃圾时才有此字段，可转为Boolean类型数据
	YidunRes       string `json:"yidunRes"`       //易盾反垃圾的原始处理细节，只有接入了相关功能易盾反垃圾的应用才会有这个字段。
}

//AudioCopyInfo 音视频/白板时长消息抄送
type AudioCopyInfo struct {
	ChannelID  string `json:"channelId"`  //通道号
	Createtime string `json:"createtime"` //音视频通话/白板开始的事件, 可转为13位时间戳
	Duration   string `json:"duration"`   //此通通话/白板的通话时长，精确到秒，可转为Integer类型
	EventType  string `json:"eventType"`  //为5，表示是实时音视频/白板时长类型事件
	Live       string `json:"live"`       //是否是互动直播的音视频，0：否，1：是
	Members    string `json:"members"`    //表示通话/白板的参与者：accid为用户帐号；如果是通话的发起者的话，caller字段为true，否则无caller字段；duration表示对应accid用户的单方时长，其中白板消息暂无此单方时长的统计
	Status     string `json:"status"`     //通话/白板状态:SUCCESS：表示正常挂断;TIMEOUT：表示超时;SINGLE_PARTICIPATE：表示只有一个参与者;UNKNOWN：表示未知状态
	Type       string `json:"type"`       //类型：AUDIO：表示音频通话；VEDIO：表示视频通话；DataTunnel：表示白板事件
	Ext        string `json:"ext"`        //音视频发起时的自定义字段，可选，由用户指定
	Running    bool   `json:"running"`    //若为true表示超长时长通话的过程中的抄送，缺省或者false表示普通时长通话的抄送或者超长时长通话的最后一次抄送
}

//AudioDownloadCopyInfo 音视频/白板文件下载信息抄送
type AudioDownloadCopyInfo struct {
	EventType string `json:"eventType"` //值为6，表示是音视频/白板文件下载信息类型的消息
	// 可转为JSONArray，其中的字段释义如下：
	// caller：是否是此通通话的发起者，若是则为true，若不是则没有此字段，可转为Boolean值
	// channelid：通道号，可转为Long值
	// filename：文件名，直接存储，混合录制文件filename带有"-mix"标记
	// md5：文件的md5值
	// size：文件大小，单位为字符，可转为Long值
	// type：文件的类型（扩展名），包括：实时音频录制文件(aac)、白板录制文件(gz)、实时视频录制文件(mp4)、互动直播视频录制文件(flv)
	// url：文件的下载地址，请不要解析该字段
	// user：用户帐号，若该文件为混合录制文件，则该字段为"0"
	// mix：是否为混合录制文件，true：混合录制文件；false：单人录制文件
	// vid：点播文件id，注意白板录制文件(gz)无此字段
	FileInfo string `json:"fileinfo"`
}

//FileDownloadInfo 单个文件下载信息
type FileDownloadInfo struct {
	Caller    bool   `json:"caller"`    //是否是此通通话的发起者，若是则为true，若不是则没有此字段，可转为Boolean值
	ChannelID string `json:"channelid"` //通道号
	Filename  string `json:"filename"`  //文件名，直接存储，混合录制文件filename带有"-mix"标记
	Md5       string `json:"md5"`       //文件的md5值
	Mix       bool   `json:"mix"`       //是否为混合录制文件，true：混合录制文件；false：单人录制文件
	Size      string `json:"size"`      //size：文件大小，单位为字符，可转为Long值
	Type      string `json:"type"`      //文件的类型（扩展名），包括：实时音频录制文件(aac)、白板录制文件(gz)、实时视频录制文件(mp4)、互动直播视频录制文件(flv)
	Vid       string `json:"vid"`       //点播文件id，注意白板录制文件(gz)无此字段
	URL       string `json:"url"`       //文件的下载地址，请不要解析该字段
	User      string `json:"user"`      //用户帐号，若该文件为混合录制文件，则该字段为"0"
}

//RoomInfo .
type RoomInfo struct {
	RoomID      int64  `json:"cid"`         //房间ID【int64】
	RoomName    string `json:"cname"`       //房间名称
	AcctID      string `json:"accid"`       //房间创建者ID
	Total       int    `json:"total"`       //房间内活跃用户总数
	Mode        int    `json:"mode"`        //房间模式【1：双人、2：多人】
	Status      int    `json:"stats"`       //房间状态【1：初始状态，2：进行中，3：正常结束，4：异常结束】
	CreateTime  int64  `json:"createtime"`  //房间创建时间【int64】
	Destroytime int64  `json:"destroytime"` //房间结束时间【int64】
}

type ImChatRoomReq struct {
	Creator      string `json:"creator"`      //是	聊天室属主的账号accid
	Name         string `json:"name"`         //是	聊天室名称，长度限制128个字符
	Announcement string `json:"announcement"` //否	公告，长度限制4096个字符
	Broadcasturl string `json:"broadcasturl"` //否	直播地址，长度限制1024个字符
	Ext          string `json:"ext"`          //否	扩展字段，最长4096字符
	Queuelevel   int    `json:"queuelevel"`   //否	队列管理权限：0:所有人都有权限变更队列，1:只有主播管理员才能操作变更。默认0
	Bid          string `json:"bid"`          //否	反垃圾业务ID，JSON字符串，{"textbid":"","picbid":""}，若不填则使用原来的反垃圾配置
}
type ImChatRoomRes struct {
	Roomid       int    `json:"roomid"`       //房间ID
	Valid        bool   `json:"valid"`        //有效性
	Announcement string `json:"announcement"` //公告，长度限制4096个字符
	Name         string `json:"name"`         //聊天室名称，长度限制128个字符
	Broadcasturl string `json:"broadcasturl"` //直播地址，长度限制1024个字符
	Ext          string `json:"ext"`          //扩展字段，最长4096字符
	Creator      string `json:"creator"`      //聊天室属主的账号accid
}
type ImChatRoomInfoReq struct {
	Roomid              int64  `json:"roomid"` //是	聊天室id
	NeedOnlineUserCount string //否  是否需要返回在线人数，true或false，默认false
}
type ImChatRoomInfoRes struct {
	Roomid          int    `json:"roomid"`          //房间ID
	Valid           bool   `json:"valid"`           //有效性
	Muted           bool   `json:"muted"`           //聊天室是否处于全体禁言状态，全体禁言时仅管理员和创建者可以发言
	Announcement    string `json:"announcement"`    //公告，长度限制4096个字符
	Name            string `json:"name"`            //聊天室名称，长度限制128个字符
	Broadcasturl    string `json:"broadcasturl"`    //直播地址，长度限制1024个字符
	Onlineusercount int    `json:"onlineusercount"` //在线人数
	Ext             string `json:"ext"`             //扩展字段，最长4096字符
	Creator         string `json:"creator"`         //聊天室属主的账号accid
	Queuelevel      int    `json:"queuelevel"`      //否	队列管理权限：0:所有人都有权限变更队列，1:只有主播管理员才能操作变更。默认0
	Ionotify        bool   `json:"ionotify"`        // 聊天室进出通知是否开启
}
type ImChatRoomBatchInfoReq struct {
	Roomids             string `json:"roomids"` //是	多个roomid，格式为：["6001","6002","6003"]（JSONArray对应的roomid，如果解析出错，会报414错误），限20个roomid
	NeedOnlineUserCount string //否  是否需要返回在线人数，true或false，默认false
}
type ImChatRoomBatchBatchInfoRes struct {
	NoExistRooms []int64             `json:"noExistRooms"` //不存在的聊天室id列表
	SuccRooms    []ImChatRoomInfoRes `json:"succRooms"`
	FailRooms    []int64             `json:"failRooms"` //失败的聊天室id,有可能是查的时候有500错误
}
type ImChatRoomUpdateReq struct {
	Roomid       int64  `json:"roomid"`       //是	聊天室id
	Name         string `json:"name"`         //否	聊天室名称，长度限制128个字符
	Announcement string `json:"announcement"` //否	公告，长度限制4096个字符
	Broadcasturl string `json:"broadcasturl"` //否	直播地址，长度限制1024个字符
	Ext          string `json:"ext"`          //否	扩展字段，最长4096字符
	NeedNotify   string `json:"needNotify"`   //否	true或false,是否需要发送更新通知事件，默认true
	NotifyExt    string `json:"notifyExt"`    //否	通知事件扩展字段，长度限制2048
	Queuelevel   int    `json:"queuelevel"`   //否	队列管理权限：0:所有人都有权限变更队列，1:只有主播管理员才能操作变更。默认0
	Bid          string `json:"bid"`          //否	反垃圾业务ID，JSON字符串，{"textbid":"","picbid":""}，若不填则使用原来的反垃圾配置
}
type ImChatRoomStateReq struct {
	Roomid   int64  `json:"roomid"`   //是	聊天室id
	Operator string `json:"operator"` //是	操作者账号，必须是创建者才可以操作
	Valid    string `json:"valid"`    //是	true或false，false:关闭聊天室；true:打开聊天室
}
type ImChatRoomRoleReq struct {
	Roomid   int64  `json:"roomid"`   //是	聊天室id
	Operator string `json:"operator"` //是	操作者账号accid
	Target   string `json:"target"`   //是	被操作者账号accid
	Opt      int    `json:"opt"`      //是	操作：
	//1: 设置为管理员，operator必须是创建者
	//2: 设置普通等级用户，operator必须是创建者或管理员
	//-1:设为黑名单用户，operator必须是创建者或管理员
	//-2:设为禁言用户，operator必须是创建者或管理员
	Optvalue  string `json:"optvalue"`  //是	true或false，true:设置；false:取消设置； 执行“取消”设置后，若成员非禁言且非黑名单，则变成游客
	NotifyExt string `json:"notifyExt"` //否	通知扩展字段，长度限制2048，请使用json格式
}
type ImChatRoomRoleRes struct {
	Roomid int64  `json:"roomid"` //聊天室id
	Level  int    `json:"level"`
	Accid  string `json:"accid"` //聊天室创建者id
	Type   string `json:"type"`  //返回的type字段可能为：
	//LIMITED,          //受限用户,黑名单+禁言
	//COMMON,           //普通固定成员
	//CREATOR,          //创建者
	//MANAGER,          //管理员
	//TEMPORARY,        //临时用户,非固定成员
}
type ImChatRoomAddrReq struct {
	Roomid     int64  `json:"roomid"`     //是	聊天室id
	Accid      string `json:"accid"`      //是	进入聊天室的账号
	Clienttype int    `json:"clienttype"` //否	1:weblink（客户端为web端时使用）; 2:commonlink（客户端为非web端时使用）;3:wechatlink(微信小程序使用), 默认1
	Clientip   string `json:"clientip"`   //否	客户端ip，传此参数时，会根据用户ip所在地区，返回合适的地址
}
type ImChatRoomSendMsgReq struct {
	Roomid    int64  `json:"roomid"`    //是	聊天室id
	MsgId     string `json:"msgId"`     //是	客户端消息id，使用uuid等随机串，msgId相同的消息会被客户端去重
	FromAccid string `json:"fromAccid"` //是	消息发出者的账号accid
	MsgType   int    `json:"msgType"`   //是	消息类型：
	//0: 表示文本消息，
	//1: 表示图片，
	//2: 表示语音，
	//3: 表示视频，
	//4: 表示地理位置信息，
	//6: 表示文件，
	//10: 表示Tips消息，
	//100: 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
	SubType      int    `json:"subType"`      //否	自定义消息子类型，大于0
	ResendFlag   int    `json:"resendFlag"`   //否	重发消息标记，0：非重发消息，1：重发消息，如重发消息会按照msgid检查去重逻辑
	Attach       string `json:"attach"`       //否	文本消息：填写消息文案; 其它类型消息，请参考 消息格式示例； 长度限制4096字符
	Ext          string `json:"ext"`          //否	消息扩展字段，内容可自定义，请使用JSON格式，长度限制4096字符
	SkipHistory  int    `json:"skipHistory"`  //否	是否跳过存储云端历史，0：不跳过，即存历史消息；1：跳过，即不存云端历史；默认0
	AbandonRatio int    `json:"abandonRatio"` //否	可选，消息丢弃的概率。取值范围[0-9999]；
	//其中0代表不丢弃消息，9999代表99.99%的概率丢弃消息，默认不丢弃；
	//注意如果填写了此参数，下面的highPriority参数则会无效；
	//此参数可用于流控特定业务类型的消息。
	HighPriority bool `json:"highPriority"` //否	可选，true表示是高优先级消息，云信会优先保障投递这部分消息；false表示低优先级消息。默认false。
	//强烈建议应用恰当选择参数，以便在必要时，优先保障应用内的高优先级消息的投递。若全部设置为高优先级，则等于没有设置，单个聊天室最多支持每秒10条的高优先级消息，超过的会转为普通消息。 高优先级消息可以设置进入后重发，见needHighPriorityMsgResend参数
	NeedHighPriorityMsgResend bool `json:"needHighPriorityMsgResend"` //否	可选，true表示会重发消息，false表示不会重发消息。默认true。注:若设置为true， 用户离开聊天室之后重新加入聊天室，在有效期内还是会收到发送的这条消息，目前有效期默认30s。在没有配置highPriority时needHighPriorityMsgResend不生效。
	UseYidun                  int  `json:"useYidun"`                  //否	可选，单条消息是否使用易盾反垃圾，可选值为0。
	//0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。
	//若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
	YidunAntiCheating string `json:"yidunAntiCheating"` //否	可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
	Bid               string `json:"bid"`               //否	可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
	Antispam          string `json:"antispam"`          //否	对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。 true或false, 默认false。 只对消息类型为：100 自定义消息类型 的消息生效。
	NotifyTargetTags  string `json:"notifyTargetTags"`  //否	可选，标签表达式，最长128个字符
	AntispamCustom    string `json:"antispamCustom"`    //否	在antispam参数为true时生效。
	//自定义的反垃圾检测内容, JSON格式，长度限制同body字段，不能超过5000字符，要求antispamCustom格式如下：
	//{"type":1,"data":"custom content"}
	//字段说明：
	//1. type: 1：文本，2：图片。
	//2. data: 文本内容or图片地址。
	Env string `json:"env"` //否	所属环境，根据env可以配置不同的抄送地址
}

type ImChatRoomSendMsgRes struct {
	Time             string `json:"time"`             //"1456396333115",
	FromAvator       string `json:"fromAvator"`       //"http://b12026.nos.netease.com/MTAxMTAxMA==/bmltYV84NDU4OF8xNDU1ODczMjA2NzUwX2QzNjkxMjI2LWY2NmQtNDQ3Ni0E2LTg4NGE4MDNmOGIwMQ==",
	MsgidClient      string `json:"msgid_client"`     //"c9e6c306-804f-4ec3-b8f0-573778829419",
	FromClientType   string `json:"fromClientType"`   //"REST",
	Attach           string `json:"attach"`           //"This+is+test+msg",
	RoomId           string `json:"roomId"`           //"36",
	FromAccount      string `json:"fromAccount"`      //"zhangsan",
	FromNick         string `json:"fromNick"`         //"张三",
	Type             string `json:"type"`             //"0",
	Ext              string `json:"ext"`              //"",
	HighPriorityFlag int    `json:"highPriorityFlag"` //1, //高优先级消息标记，不带此标记表示非高优先级
	MsgAbandonFlag   string `json:"msgAbandonFlag"`   //"1" //消息被丢弃标记，传abandonRatio参数时才会返回此标记，不返回此标记代表未被丢弃
}
type ImChatroomAddRobotReq struct {
	Roomid    int64  `json:"roomid"`    //是	聊天室id
	Accids    string `json:"accids"`    //是	机器人账号accid列表，必须是有效账号，账号数量上限100个
	RoleExt   string `json:"roleExt"`   //否	机器人信息扩展字段，请使用json格式，长度4096字符
	NotifyExt string `json:"notifyExt"` //否	机器人进入聊天室通知的扩展字段，请使用json格式，长度2048字符
}
type ImChatroomAddRobotRes struct {
	FailAccids    string `json:"failAccids"`    //"[\"hzzhangsan\"]",
	SuccessAccids string `json:"successAccids"` //"[\"hzlisi\"]",
	OldAccids     string `json:"oldAccids"`     //"[\"hzwangwu\"]"
}
type ImChatroomDelRobotReq struct {
	Roomid int64  `json:"roomid"` //是	聊天室id
	Accids string `json:"accids"` //是	机器人账号accid列表，必须是有效账号，账号数量上限100个
}
type ImChatroomDelRobotRes struct {
	FailAccids    string `json:"failAccids"`    //"[\"hzzhangsan\"]",
	SuccessAccids string `json:"successAccids"` //"[\"hzlisi\"]",
}
type ImChatroomCleanRobotReq struct {
	Roomid int64 `json:"roomid"` //是	聊天室id
	Notify bool  `json:"notify"` //否	是否发送退出聊天室通知消息，默认为false
}
type ImChatroomCleanRobotRes struct {
	Size int `json:"size"` //清空机器人的个数
}
type ImChatroomMuteReq struct {
	Roomid       int64  `json:"roomid"`       //是	聊天室id
	Operator     string `json:"operator"`     //是	操作者accid,必须是管理员或创建者
	Target       string `json:"target"`       //是	被禁言的目标账号accid
	MuteDuration int64  `json:"muteDuration"` //是	0:解除禁言;>0设置禁言的秒数，不能超过2592000秒(30天)
	NeedNotify   string `json:"needNotify"`   //否	操作完成后是否需要发广播，true或false，默认true
	NotifyExt    string `json:"notifyExt"`    //否	通知广播事件中的扩展字段，长度限制2048字符
}
type ImChatroomMuteRes struct {
	MuteDuration int64 `json:"muteDuration"` //禁言的秒数
}
type ImChatroomMuteRoomReq struct {
	Roomid     int64  `json:"roomid"`     //是	聊天室id
	Operator   string `json:"operator"`   //是	操作者accid,必须是管理员或创建者
	Mute       string `json:"mute"`       //是	true或false
	NeedNotify string `json:"needNotify"` //否	操作完成后是否需要发广播，true或false，默认true
	NotifyExt  string `json:"notifyExt"`  //否	通知广播事件中的扩展字段，长度限制2048字符
}
type ImChatroomMuteRoomRes struct {
	Success bool `json:"success"` //操作是否成功
}
type ImChatroomTopnReq struct {
	Topn      int    `json:"topn"`      //否	topn值，可选值 1~500，默认值100
	Timestamp int64  `json:"timestamp"` //否	需要查询的指标所在的时间坐标点，不提供则默认当前时间，单位秒/毫秒皆可
	Period    string `json:"period"`    //否	统计周期，可选值包括 hour/day, 默认hour
	Orderby   string `json:"orderby"`   //否	取排序值,可选值 active/enter/message,分别表示按日活排序，进入人次排序和消息数排序， 默认active
}
type ImChatroomTopnRes struct {
	ActiveNums int    `json:"activeNums"` // 该聊天室内的活跃数
	Datetime   int    `json:"datetime"`   // 统计时间点，单位秒，按天统计的是当天的0点整点；按小时统计的是指定小时的整点
	EnterNums  int    `json:"enterNums"`  // 进入人次数量
	Msgs       int    `json:"msgs"`       // 聊天室内发生的消息数
	Period     string `json:"period"`     // 统计周期，HOUR表示按小时统计；DAY表示按天统计
	RoomId     int64  `json:"roomId"`     // 聊天室ID号
}
type ImChatroomMembersReq struct {
	Roomid  int64 `json:"roomid"`  //是	聊天室id
	Type    int   `json:"type"`    //是	需要查询的成员类型,0:固定成员;1:非固定成员;2:仅返回在线的固定成员
	Endtime int64 `json:"endtime"` //是	单位毫秒，按时间倒序最后一个成员的时间戳,0表示系统当前时间
	Limit   int64 `json:"limit"`   //是	返回条数，<=100
}
type ImChatroomMembersRes struct {
	Data []ImChatroomMembersRes_ `json:"data"`
}
type ImChatroomMembersRes_ struct {
	Roomid       int64  `json:"roomid"`       //聊天室id
	Accid        string `json:"accid"`        //用户accid
	Nick         string `json:"nick"`         //聊天室内的昵称
	Avator       string `json:"avator"`       //聊天室内的头像
	Ext          string `json:"ext"`          //开发者扩展字段
	Type         string `json:"type"`         //角色类型： UNSET（未设置）， LIMITED（受限用户，黑名单或禁言）， COMMON（普通固定成员）， CREATOR（创建者）， MANAGER（管理员）， TEMPORARY（临时用户,非固定成员）
	Level        int    `json:"level"`        //成员级别（若未设置成员级别，则无此字段）
	OnlineStat   bool   `json:"onlineStat"`   //是否在线
	EnterTime    int64  `json:"enterTime"`    //进入聊天室的时间点
	Blacklisted  bool   `json:"blacklisted"`  //是否在黑名单中（若未被拉黑，则无此字段）
	Muted        bool   `json:"muted"`        //是否被禁言（若未被禁言，则无此字段）
	TempMuted    bool   `json:"tempMuted"`    //是否被临时禁言（若未被临时禁言，则无此字段）
	TempMuteTtl  int64  `json:"tempMuteTtl"`  //临时禁言的解除时长,单位秒（若未被临时禁言，则无此字段）
	IsRobot      bool   `json:"isRobot"`      //是否是聊天室机器人（若不是机器人，则无此字段）
	RobotExpirAt int    `json:"robotExpirAt"` //机器人失效的时长，单位秒（若不是机器人，则无此字段）
}
type ImChatroomMembersByRoleReq struct {
	Roomid int64  `json:"roomid"` //是	聊天室id
	Roles  string `json:"roles"`  //是	设置需要获取的角色,格式示例： {"creator": true,"manager": true,"blacklist": false,"mute": false}
	//字段说明：
	//1、creator：聊天室创建者
	//2、manager：聊天室管理员
	//3、blacklist：黑名单用户
	//4、mute：被禁言用户
	//说明：设置为false或不设置表示不获取相应的角色信息
}
type ImChatroomMembersByRoleRes struct {
	Data []ImChatroomMembersByRoleRes_ `json:"data"`
}
type ImChatroomMembersByRoleRes_ struct {
	Roomid       int64  `json:"roomid"`       //聊天室id
	Accid        string `json:"accid"`        //用户accid
	Nick         string `json:"nick"`         //聊天室内的昵称
	Avator       string `json:"avator"`       //聊天室内的头像
	Ext          string `json:"ext"`          //开发者扩展字段
	Type         string `json:"type"`         //角色类型：UNSET（未设置），LIMITED（受限用户，黑名单或禁言），COMMON（普通固定成员），CREATOR（创建者），MANAGER（管理员），TEMPORARY（临时用户,非固定成员）
	Level        int    `json:"level"`        //成员级别（若未设置成员级别，则无此字段）
	OnlineStat   bool   `json:"onlineStat"`   //是否在线
	EnterTime    int64  `json:"enterTime"`    //进入聊天室的时间点
	Blacklisted  bool   `json:"blacklisted"`  //是否在黑名单中（若未被拉黑，则无此字段）
	Muted        bool   `json:"muted"`        //是否被禁言（若未被禁言，则无此字段）
	TempMuted    bool   `json:"tempMuted"`    //是否被临时禁言（若未被临时禁言，则无此字段）
	TempMuteTtl  int64  `json:"tempMuteTtl"`  //临时禁言的解除时长,单位秒（若未被临时禁言，则无此字段）
	IsRobot      bool   `json:"isRobot"`      //是否是聊天室机器人（若不是机器人，则无此字段）
	RobotExpirAt int    `json:"robotExpirAt"` //机器人失效的时长，单位秒（若不是机器人，则无此字段）
}
type ImChatroomMembersBatchReq struct {
	Roomid int64  `json:"roomid"` //是	聊天室id
	Accids string `json:"accids"` //是	\["abc","def"\], 账号列表，最多200条
}
type ImChatroomMembersBatchRes struct {
	Data []ImChatroomMembersBatchRes_ `json:"data"`
}
type ImChatroomMembersBatchRes_ struct {
	Roomid     int64  `json:"roomid"`
	Accid      string `json:"accid"`
	Nick       string `json:"nick"`
	Type       string `json:"type"` //COMMON:普通成员(固定成员)；CREATOR:聊天室创建者；MANAGER:聊天室管理员；TEMPORARY:临时用户(非聊天室固定成员)；ANONYMOUS:匿名用户(未注册账号)；LIMITED:受限用户(黑名单+禁言)
	OnlineStat bool   `json:"onlineStat"`
}
type ImChatroomChangeRoleReq struct {
	Roomid     int64  `json:"roomid"`     //是	聊天室id
	Accid      string `json:"accid"`      //是	需要变更角色信息的accid
	Save       bool   `json:"save"`       //否	变更的信息是否需要持久化，默认false，仅对聊天室固定成员生效
	NeedNotify bool   `json:"needNotify"` //否	是否需要做通知
	NotifyExt  string `json:"notifyExt"`  //否	通知的内容，长度限制2048
	Nick       string `json:"nick"`       //否	聊天室室内的角色信息：昵称，不超过64个字符
	Avator     string `json:"avator"`     //否	聊天室室内的角色信息：头像
	Ext        string `json:"ext"`        //否	聊天室室内的角色信息：开发者扩展字段
	Bid        string `json:"bid"`        //否	反垃圾业务ID，JSON字符串，{"textbid":"","picbid":""}，若不填则使用原来的反垃圾配置
}
type ImChatroomUserRoomIdsReq struct {
	Creator string `json:"creator"` //是	聊天室创建者accid
}
type ImChatroomUserRoomIdsRes struct {
	Roomids []string `json:"roomids"`
}
type ImChatroomInOutNotifyReq struct {
	Roomid int64 `json:"roomid"` //是	聊天室ID
	Close  bool  `json:"close"`  //是	true/false, 是否关闭进出通知
}
type ImChatroomTagMuteReq struct {
	Roomid           int64  `json:"roomid"`           //是	聊天室ID
	Operator         string `json:"operator"`         //是	操作者accid，必须是创建者或者管理员
	TargetTag        string `json:"targetTag"`        //是	目标标签
	NeedNotify       bool   `json:"needNotify"`       //否	true/false，是否发送禁言通知，默认true
	NotifyExt        string `json:"notifyExt"`        //否	禁言通知通知扩展字段
	MuteDuration     int    `json:"muteDuration"`     //是	禁言时长，单位秒，最长30天，若等于0表示取消禁言
	NotifyTargetTags string `json:"notifyTargetTags"` //否	禁言通知的目标标签表达式，若缺失则发送给设置了targetTag的人
}
type ImChatroomTagMuteRes struct {
	MuteDuration int64 `json:"muteDuration"` //禁言时长，若取消禁言，则返回上次禁言的剩余禁言时长
}
type ImChatroomTagMemberCountReq struct {
	Roomid int64  `json:"roomid"` //是	聊天室ID
	Tag    string `json:"tag"`    //是	标签
}
type ImChatroomTagMemberCountRes struct {
	Tag             string `json:"tag"`             //标签
	OnlineUserCount int    `json:"onlineUserCount"` //在线用户数
}
type ImChatroomTagMembersReq struct {
	Roomid  int64  `json:"roomid"`  //是	聊天室ID
	Tag     string `json:"tag"`     //是	标签
	EndTime int64  `json:"endTime"` //是	起始时间，逆序查询，若传0则表示从当前时间往前查
	Limit   int    `json:"limit"`   //是	条数，最多100
}
type ImChatroomTagMembersRes struct {
	Data []ImChatroomTagMembersRes_ `json:"data"`
}
type ImChatroomTagMembersRes_ struct {
	Roomid           int64  `json:"roomid"`
	Accid            string `json:"accid"`
	Nick             string `json:"nick"`
	Avator           string `json:"avator"`
	Ext              string `json:"ext"`
	Type             string `json:"type"`
	Level            int    `json:"level"`
	OnlineStat       bool   `json:"onlineStat"`
	EnterTime        string `json:"enterTime"`
	Blacklisted      bool   `json:"blacklisted"`
	Muted            bool   `json:"muted"`
	TempMuted        bool   `json:"tempMuted"`
	TempMuteTtl      int    `json:"tempMuteTtl"`
	IsRobot          bool   `json:"isRobot"`
	RobotExpirAt     int    `json:"robotExpirAt"`
	Tags             string `json:"tags"`
	NotifyTargetTags string `json:"notifyTargetTags"` //"{\"tag\": \"abc\"} and {\"tag\": \"def\"}"
}
type ImChatRoomBroadcastReq struct {
	//Roomid    int64  `json:"roomid"`    //是	聊天室id
	MsgId     string `json:"msgId"`     //是	客户端消息id，使用uuid等随机串，msgId相同的消息会被客户端去重
	FromAccid string `json:"fromAccid"` //是	消息发出者的账号accid
	MsgType   int    `json:"msgType"`   //是	消息类型：
	//0: 表示文本消息，
	//1: 表示图片，
	//2: 表示语音，
	//3: 表示视频，
	//4: 表示地理位置信息，
	//6: 表示文件，
	//10: 表示Tips消息，
	//100: 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
	SubType    int    `json:"subType"`    //否	自定义消息子类型，大于0
	ResendFlag int    `json:"resendFlag"` //否	重发消息标记，0：非重发消息，1：重发消息，如重发消息会按照msgid检查去重逻辑
	Attach     string `json:"attach"`     //否	文本消息：填写消息文案; 其它类型消息，请参考 消息格式示例； 长度限制4096字符
	Ext        string `json:"ext"`        //否	消息扩展字段，内容可自定义，请使用JSON格式，长度限制4096字符
	//SkipHistory  int    `json:"skipHistory"`  //否	是否跳过存储云端历史，0：不跳过，即存历史消息；1：跳过，即不存云端历史；默认0
	//AbandonRatio int    `json:"abandonRatio"` //否	可选，消息丢弃的概率。取值范围[0-9999]；
	//其中0代表不丢弃消息，9999代表99.99%的概率丢弃消息，默认不丢弃；
	//注意如果填写了此参数，下面的highPriority参数则会无效；
	//此参数可用于流控特定业务类型的消息。
	HighPriority bool `json:"highPriority"` //否	可选，true表示是高优先级消息，云信会优先保障投递这部分消息；false表示低优先级消息。默认false。
	//强烈建议应用恰当选择参数，以便在必要时，优先保障应用内的高优先级消息的投递。若全部设置为高优先级，则等于没有设置，单个聊天室最多支持每秒10条的高优先级消息，超过的会转为普通消息。 高优先级消息可以设置进入后重发，见needHighPriorityMsgResend参数
	NeedHighPriorityMsgResend bool `json:"needHighPriorityMsgResend"` //否	可选，true表示会重发消息，false表示不会重发消息。默认true。注:若设置为true， 用户离开聊天室之后重新加入聊天室，在有效期内还是会收到发送的这条消息，目前有效期默认30s。在没有配置highPriority时needHighPriorityMsgResend不生效。
	UseYidun                  int  `json:"useYidun"`                  //否	可选，单条消息是否使用易盾反垃圾，可选值为0。
	//0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。
	//若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
	YidunAntiCheating string `json:"yidunAntiCheating"` //否	可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
	Bid               string `json:"bid"`               //否	可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
	Antispam          string `json:"antispam"`          //否	对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。 true或false, 默认false。 只对消息类型为：100 自定义消息类型 的消息生效。
	NotifyTargetTags  string `json:"notifyTargetTags"`  //否	可选，标签表达式，最长128个字符
	AntispamCustom    string `json:"antispamCustom"`    //否	在antispam参数为true时生效。
	//自定义的反垃圾检测内容, JSON格式，长度限制同body字段，不能超过5000字符，要求antispamCustom格式如下：
	//{"type":1,"data":"custom content"}
	//字段说明：
	//1. type: 1：文本，2：图片。
	//2. data: 文本内容or图片地址。
	Env string `json:"env"` //否	所属环境，根据env可以配置不同的抄送地址
}
type ImChatRoomBroadcastRes struct {
	Time             string `json:"time"`             //"1456396333115",
	FromAvator       string `json:"fromAvator"`       //"http://b12026.nos.netease.com/MTAxMTAxMA==/bmltYV84NDU4OF8xNDU1ODczMjA2NzUwX2QzNjkxMjI2LWY2NmQtNDQ3Ni0E2LTg4NGE4MDNmOGIwMQ==",
	MsgidClient      string `json:"msgid_client"`     //"c9e6c306-804f-4ec3-b8f0-573778829419",
	FromClientType   string `json:"fromClientType"`   //"REST",
	Attach           string `json:"attach"`           //"This+is+test+msg",
	RoomId           string `json:"roomId"`           //"36",
	FromAccount      string `json:"fromAccount"`      //"zhangsan",
	FromNick         string `json:"fromNick"`         //"张三",
	Type             string `json:"type"`             //"0",
	Ext              string `json:"ext"`              //"",
	HighPriorityFlag int    `json:"highPriorityFlag"` //1, //高优先级消息标记，不带此标记表示非高优先级
	//MsgAbandonFlag   string `json:"msgAbandonFlag"`   //"1" //消息被丢弃标记，传abandonRatio参数时才会返回此标记，不返回此标记代表未被丢弃
}
type ImChatRoomReCallReq struct {
	Roomid      int64  `json:"roomid"`      //是	聊天室id
	MsgTimetag  int64  `json:"msgTimetag"`  //是	被撤回消息的时间戳
	FromAcc     string `json:"fromAcc"`     //是	被撤回消息的消息发送者accid
	MsgId       string `json:"msgId"`       //是	被撤回消息的消息id
	OperatorAcc string `json:"operatorAcc"` //是	消息撤回的操作者accid
	NotifyExt   string `json:"notifyExt"`   //否	消息撤回的通知扩展字段，最长1024字符
}
type ImChatRoomSendMsgToOneReq struct {
	Roomid    int64  `json:"roomid"`    //是	聊天室id
	MsgId     string `json:"msgId"`     //是	客户端消息id，使用uuid等随机串，msgId相同的消息会被客户端去重
	FromAccid string `json:"fromAccid"` //是	消息发出者的账号accid
	ToAccids  string `json:"toAccids"`  //是	消息接收者accid列表，最大100个  ["acc1","acc2"]
	MsgType   int    `json:"msgType"`   //是	消息类型：
	//0: 表示文本消息，
	//1: 表示图片，
	//2: 表示语音，
	//3: 表示视频，
	//4: 表示地理位置信息，
	//6: 表示文件，
	//10: 表示Tips消息，
	//100: 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
	SubType    int    `json:"subType"`    //否	自定义消息子类型，大于0
	ResendFlag int    `json:"resendFlag"` //否	重发消息标记，0：非重发消息，1：重发消息，如重发消息会按照msgid检查去重逻辑
	Attach     string `json:"attach"`     //否	文本消息：填写消息文案; 其它类型消息，请参考 消息格式示例； 长度限制4096字符
	Ext        string `json:"ext"`        //否	消息扩展字段，内容可自定义，请使用JSON格式，长度限制4096字符
	//SkipHistory  int    `json:"skipHistory"`  //否	是否跳过存储云端历史，0：不跳过，即存历史消息；1：跳过，即不存云端历史；默认0
	//AbandonRatio int    `json:"abandonRatio"` //否	可选，消息丢弃的概率。取值范围[0-9999]；
	//其中0代表不丢弃消息，9999代表99.99%的概率丢弃消息，默认不丢弃；
	//注意如果填写了此参数，下面的highPriority参数则会无效；
	//此参数可用于流控特定业务类型的消息。
	//HighPriority bool `json:"highPriority"` //否	可选，true表示是高优先级消息，云信会优先保障投递这部分消息；false表示低优先级消息。默认false。
	////强烈建议应用恰当选择参数，以便在必要时，优先保障应用内的高优先级消息的投递。若全部设置为高优先级，则等于没有设置，单个聊天室最多支持每秒10条的高优先级消息，超过的会转为普通消息。 高优先级消息可以设置进入后重发，见needHighPriorityMsgResend参数
	//NeedHighPriorityMsgResend bool `json:"needHighPriorityMsgResend"` //否	可选，true表示会重发消息，false表示不会重发消息。默认true。注:若设置为true， 用户离开聊天室之后重新加入聊天室，在有效期内还是会收到发送的这条消息，目前有效期默认30s。在没有配置highPriority时needHighPriorityMsgResend不生效。
	UseYidun int `json:"useYidun"` //否	可选，单条消息是否使用易盾反垃圾，可选值为0。
	//0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。
	//若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
	YidunAntiCheating string `json:"yidunAntiCheating"` //否	可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
	Bid               string `json:"bid"`               //否	可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
	Antispam          string `json:"antispam"`          //否	对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。 true或false, 默认false。 只对消息类型为：100 自定义消息类型 的消息生效。
	//NotifyTargetTags  string `json:"notifyTargetTags"`  //否	可选，标签表达式，最长128个字符
	AntispamCustom string `json:"antispamCustom"` //否	在antispam参数为true时生效。
	//自定义的反垃圾检测内容, JSON格式，长度限制同body字段，不能超过5000字符，要求antispamCustom格式如下：
	//{"type":1,"data":"custom content"}
	//字段说明：
	//1. type: 1：文本，2：图片。
	//2. data: 文本内容or图片地址。
	Env string `json:"env"` //否	所属环境，根据env可以配置不同的抄送地址
}
type ImChatRoomSendMsgToOneRes struct {
	Time           string `json:"time"`           //"1456396333115",
	FromAvator     string `json:"fromAvator"`     //"http://b12026.nos.netease.com/MTAxMTAxMA==/bmltYV84NDU4OF8xNDU1ODczMjA2NzUwX2QzNjkxMjI2LWY2NmQtNDQ3Ni0E2LTg4NGE4MDNmOGIwMQ==",
	MsgidClient    string `json:"msgid_client"`   //"c9e6c306-804f-4ec3-b8f0-573778829419",
	FromClientType string `json:"fromClientType"` //"REST",
	Attach         string `json:"attach"`         //"This+is+test+msg",
	RoomId         string `json:"roomId"`         //"36",
	FromAccount    string `json:"fromAccount"`    //"zhangsan",
	FromNick       string `json:"fromNick"`       //"张三",
	Type           string `json:"type"`           //"0",
	Ext            string `json:"ext"`            //"",
	//HighPriorityFlag int    `json:"highPriorityFlag"` //1, //高优先级消息标记，不带此标记表示非高优先级
	//MsgAbandonFlag   string `json:"msgAbandonFlag"`   //"1" //消息被丢弃标记，传abandonRatio参数时才会返回此标记，不返回此标记代表未被丢弃
}
