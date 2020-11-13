package api

import (
	"gf-app/module/base/service"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

var Robot service.WxRobot

func init()  {
	Robot.Start("WxRobot-wuxi")
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
func MessageHandler(r *ghttp.Request) {
	r.Response.Write(Robot.Name,"\n",Robot.CreateTime)
	server := Robot.GetServer(r.Response.Writer.RawWriter(),r.Request)
	server.SkipValidate(true)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//TODO
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeImage, MsgData: text}

		//article1 := message.NewArticle("测试图文1", "图文描述", "", "")
		//articles := []*message.Article{article1}
		//news := message.NewNews(articles)
		//return &message.Reply{MsgType: message.MsgTypeNews, MsgData: news}

		//voice := message.NewVoice(mediaID)
		//return &message.Reply{MsgType: message.MsgTypeVoice, MsgData: voice}

		//
		//video := message.NewVideo(mediaID, "标题", "描述")
		//return &message.Reply{MsgType: message.MsgTypeVideo, MsgData: video}

		//music := message.NewMusic("标题", "描述", "音乐链接", "HQMusicUrl", "缩略图的媒体id")
		//return &message.Reply{MsgType: message.MsgTypeMusic, MsgData: music}

		//多客服消息转发
		//transferCustomer := message.NewTransferCustomer("")
		//return &message.Reply{MsgType: message.MsgTypeTransfer, MsgData: transferCustomer}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		glog.Printf("Serve Error, err=%+v", err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		glog.Printf("Send Error, err=%v", err)
		return
	}

}

func TemplateHandler(r *ghttp.Request) {
	Robot.Account.GetTemplate()

	//message.TemplateMessage{}
}