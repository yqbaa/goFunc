package controllers

import (
	"github.com/astaxie/beego"
	"math"

	"strconv"
)

/*
 * JSON 返回数据结构
 */
type JSON_DATA struct {
	Data interface{}
	Code int
	Msg  string
}
type BaseController struct {
	beego.Controller
}

/*
 * @param data interface 返回的数据
 */
func (this *BaseController) success(data interface{}) {
	json := JSON_DATA{Data: data, Code: 0, Msg: ""}
	this.Data["json"] = json
	this.ServeJSON()

}

/*
 * @param msg string 错误提示
 * @param code int 错误码
 */
func (this *BaseController) error(msg string, code int) {
	json := JSON_DATA{Data: "", Code: code, Msg: msg}
	this.Data["json"] = json
	this.ServeJSON()
}

/*
 * 分页信息
 */
func (this *BaseController) page(total int64, currentPage int64, perPage int64) map[string]interface{} {
	var page map[string]interface{}
	page = make(map[string]interface{})
	page["total"] = total
	page["perPage"] = perPage
	page["currentPage"] = currentPage
	page["totalPage"] = math.Ceil(float64(total) / float64(perPage))
	return page
}

func (this *BaseController) getPage() int {
	pageName := beego.AppConfig.String("pageName")
	page, _ := strconv.Atoi(this.GetString(pageName))
	if page < 1 {
		page = 1
	}
	return page
}
