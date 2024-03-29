package api

import (
	"gf-app/module/base/service"
	"gf-app/utils/resp"
	"github.com/gogf/gf/net/ghttp"
	"github.com/silenceper/wechat/v2/officialaccount/basic"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/silenceper/wechat/v2/officialaccount/user"
)

var (
	Robot *service.WxRobot
	Template *message.Template
	User *user.User
	Basic *basic.Basic
)

func init()  {
	Robot = service.Start("WxRobot-wuxi")
	Template = Robot.Account.GetTemplate()
	User = Robot.Account.GetUser()
	Basic = basic.NewBasic(Robot.Account.GetContext())
}

type TemplateMessage struct {
	ToUser     string                       `json:"touser"`          // 必须, 接受者OpenID
	TemplateID string                       `json:"template_id"`     // 必须, 模版ID
	URL        string                       `json:"url,omitempty"`   // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	Color      string                       `json:"color,omitempty"` // 可选, 整个消息的颜色, 可以不设置
	Data       map[string]*TemplateDataItem `json:"data"`            // 必须, 模板数据

	MiniProgram struct {
		AppID    string `json:"appid"`    //所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
		PagePath string `json:"pagepath"` //所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
	} `json:"miniprogram"` //可选,跳转至小程序地址
}

//TemplateDataItem 模版内某个 .DATA 的值
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// 基准消息处理
//func MessageHandler(r *ghttp.Request) {
//	r.Response.Write(Robot.Name,"\n",Robot.CreateTime)
//	server := Robot.GetServer(r.Response.Writer.RawWriter(),r.Request)
//	server.SkipValidate(true)
//	//设置接收消息的处理方法
//	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
//
//		//回复消息：演示回复用户发送的消息
//		text := message.NewText(msg.Content)
//
//		//TODO
//		//阻止消息回复，自定义消息回复方法务必去掉这行
//		text.Content = ""
//
//		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
//
//		//article1 := message.NewArticle("测试图文1", "图文描述", "", "")
//		//articles := []*message.Article{article1}
//		//news := message.NewNews(articles)
//		//return &message.Reply{MsgType: message.MsgTypeNews, MsgData: news}
//
//		//voice := message.NewVoice(mediaID)
//		//return &message.Reply{MsgType: message.MsgTypeVoice, MsgData: voice}
//
//		//
//		//video := message.NewVideo(mediaID, "标题", "描述")
//		//return &message.Reply{MsgType: message.MsgTypeVideo, MsgData: video}
//
//		//music := message.NewMusic("标题", "描述", "音乐链接", "HQMusicUrl", "缩略图的媒体id")
//		//return &message.Reply{MsgType: message.MsgTypeMusic, MsgData: music}
//
//		//多客服消息转发
//		//transferCustomer := message.NewTransferCustomer("")
//		//return &message.Reply{MsgType: message.MsgTypeTransfer, MsgData: transferCustomer}
//	})
//
//	//处理消息接收以及回复
//	err := server.Serve()
//	if err != nil {
//		glog.Printf("Serve Error, err=%+v", err)
//		return
//	}
//	//发送回复的消息
//	err = server.Send()
//	if err != nil {
//		glog.Printf("Send Error, err=%v", err)
//		return
//	}
//
//}

//模板发送接口
func TemplateHandler(r *ghttp.Request) {
	templateMessage := &message.TemplateMessage{}
	r.Parse(templateMessage)
	msgID, err := Template.Send(templateMessage)
	//templateMessage.ToUser

	if err != nil {
		r.Response.WriteJsonExit(resp.Error(err.Error()))
	}

	r.Response.WriteJson(msgID)

}

//获取用户openid列表接口
func ListAllUserOpenIDs(r *ghttp.Request) {
	openidList, err := User.ListAllUserOpenIDs()

	if err != nil {
		r.Response.WriteJsonExit(resp.Error(err.Error()))
	}

	r.Response.WriteJson(resp.Succ(openidList))

}

//获取tag列表
func GetTagList(r *ghttp.Request) {
	tags, err := User.GetTag()

	if err != nil {
		r.Response.WriteJsonExit(resp.Error(err.Error()))
	}

	r.Response.WriteJson(resp.Succ(tags))
}

//获取tag下的用户
func OpenIDListByTag(r *ghttp.Request) {
	tagid := r.GetInt32("tagid")
	openidlist, err := User.OpenIDListByTag(tagid)

	if err != nil {
		r.Response.WriteJsonExit(resp.Error(err.Error()))
	}

	r.Response.WriteJson(resp.Succ(openidlist))
}

//获取template列表
func ListTemplate(r *ghttp.Request) {

	templatelist, err := Template.List()

	if err != nil {
		r.Response.WriteJsonExit(resp.Error(err.Error()))
	}

	r.Response.WriteJson(resp.Succ(templatelist))
}

//清空调用次数
func ClearQuota(r *ghttp.Request) {

	err := Basic.ClearQuota()

	if err != nil {
		r.Response.WriteJsonExit(resp.Error(err.Error()))
	}

	r.Response.WriteJson(resp.Succ("ok"))
}