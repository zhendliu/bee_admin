package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// TableName 设置BackendUser表名
func (a *Company) TableName() string {
	return CompanyName()
}
type Company struct{
	Id   int `orm:"column(id)";json:"id"`
	Name string `orm:"column(name)";json:"name"`
	LicenseNum string `orm:"column(license_num)";json:"license_num"`
	Address  string `orm:"column(address)";json:"address"`
	Person string `orm:"column(person)";json:"person"`
	Capital string `orm:"column(capital)";json:"capital"`
	Scope string `orm:"column(scope)";json:"scope"`
	BeginTime time.Time `orm:"column(begin_time)";json:"begin_time"`
	Duration int `orm:"column(duration)";json:"duration"`
	TName string `orm:"column(tname)";json:"tname"`
	TaxNum string `orm:"column(tax_num)";json:"tax_num"`
	PubName string `orm:"column(pubName)";json:"pub_name"`
	Flag int `orm:"column(flag)";json:"flag"`
	AuditTime time.Time `orm:"column(audit_time)";json:"audit_time"`
	AuditValid int `orm:"column(audit_valid)";json:"audit_valid"`
	AuditResult string `orm:"column(audit_result)";json:"audit_result"`
	WorkMode int `orm:"column(work_mode)";json:"work_mode"`
	Status int `orm:"column(status)";json:"status"`
	Url string `orm:"column(url)";json:"url"`
	Tel string `orm:"column(tel)";json:"tel"`
	Type int8 `orm:"column(type)";json:"type"`
}

func CompanyName()string {
	return TableName("tbl_company")
}


// CoursePageList 获取分页数据
func CompanyPageList(params *CourseQueryParam) ([]*Company) {
	query := orm.NewOrm().QueryTable(CompanyName())
	data := make([]*Company,0)
	//默认排序
	sortOrder := "Id"
	switch params.Sort {
	case "Id":
		sortOrder = "Id"
	case "Seq":
		sortOrder = "Seq"
	}
	if params.Order == "desc" {
		sortOrder = "-" + sortOrder
	}
	query = query.Filter("name__istartswith", params.NameLike)

	query.OrderBy(sortOrder).Limit(params.Limit, params.Offset).All(&data)
	return data
}

func CompanyPageCount(params *CourseQueryParam) ( int64) {
	query := orm.NewOrm().QueryTable(CompanyName())

	//默认排序
	sortOrder := "Id"
	switch params.Sort {
	case "Id":
		sortOrder = "Id"
	case "Seq":
		sortOrder = "Seq"
	}
	if params.Order == "desc" {
		sortOrder = "-" + sortOrder
	}
	query = query.Filter("name__istartswith", params.NameLike)
	total, _ := query.Count()

	return total
}
