package controllers

import (
	"bee_admin/models"
	"bee_admin/utils"
)

type CompanyController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *CompanyController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "DataList", "UpdateSeq")

}

func (c *CompanyController) Index() {


	Name:=c.Ctx.Input.Query("name")

	var params models.CourseQueryParam
	if Name !=""{
		params.NameLike = Name
	}
	models.CompanyPageCount(&params)
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)

	//获取数据列表和总数
	p := utils.NewPaginator(c.Ctx.Request, models.CompanyPageCount(&params))

	c.Data["Page"] = p
	params.Offset = p.Offset()
	params.Limit = p.PerPageNums
	data := models.CompanyPageList(&params)

	c.Data["Row"] = data
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("BackendUserController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("BackendUserController", "Delete")
	c.Data["Name"] = Name
	//页面模板设置
	c.setTpl()


}
