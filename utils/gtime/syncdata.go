package gtime

import (
	"context"
	"gf-app/module/base/api"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	graphql "github.com/hasura/go-graphql-client"
)

var Client *graphql.Client

func init()  {

	c := g.Cfg()
	//src := oauth2.StaticTokenSource(
	//	&oauth2.Token{AccessToken: c.GetString("graphql.GRAPHQL_TOKEN")},
	//)
	//
	//httpClient := oauth2.NewClient(context.Background(), src)
	//
	//client := ghttp.NewClient().Header(map[string]string{
	//	"x-hasura-admin-secret":"Xiafan123",
	//})
	//httpClient := client.Client()

	Client = graphql.NewClient(c.GetString("graphql.GRAPHQL_URL"), nil)
}


func synctags()  {
	tags, _ := api.User.GetTag()


	if tags == nil {
		return
	}

	id := garray.NewStrArray()
	name := garray.NewStrArray()

	for _, v := range tags {
		id.Append(gconv.String(v.ID))
		name.Append(v.Name)
	}

	newtagid := "[" + id.Join(",") + "]"

	newname := `["` + name.Join(`","`) + `"]`

	newform := `{"column":1,"schema":{"type":"object","required":["target","content","deadline","task_name","template_id"],"properties":{"target":{"enum":`+ newtagid +`,"type":"array","items":{"type":"string"},"title":"推送用户","enumNames":`+ newname +`,"ui:widget":"multiSelect","description":"向谁推送本条信息"},"content":{"type":"string","title":"任务内容","format":"textarea","ui:options":{}},"deadline":{"type":"string","title":"截止日期","format":"dateTime"},"task_name":{"type":"string","title":"任务名称","ui:options":{}},"template_id":{"enum":["jP7on8ihnhmuzlUuEynfMQq96Vgonwh4CC7IhBHTLl4"],"type":"string","title":"推送模板","enumNames":["默认模板"],"description":"模板能在微信公众号后台添加"}}},"labelWidth":100,"displayType":"row","showDescIcon":true}`

	var m struct {
		Update_wxpolice_wx_form struct {
			Affected_rows      graphql.Int
		} `graphql:"update_wxpolice_wx_form(where: {id: {_eq: 1}}, _set: {form: $form})"`
	}

	variables := map[string]interface{}{
		"form": graphql.String(
			newform,
			),
	}

	err := Client.Mutate(context.Background(),&m,variables)
	if err != nil {
		glog.Info(variables)
		glog.Error(err)
		// Handle error.
	}
	glog.Info("定时同步任务",m.Update_wxpolice_wx_form.Affected_rows)
}