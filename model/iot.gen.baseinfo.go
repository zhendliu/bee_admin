package	model	
import (	
"fmt"	
"time"	
"github.com/jinzhu/gorm"	
)	

type _BaseinfoMgr struct {
	*_BaseMgr
}

// BaseinfoMgr open func
func BaseinfoMgr(db *gorm.DB) *_BaseinfoMgr {
	if db == nil {
		panic(fmt.Errorf("BaseinfoMgr need init by db"))
	}
	return &_BaseinfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("baseinfo"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BaseinfoMgr) GetTableName() string {
	return "baseinfo"
}

// Get 获取
func (obj *_BaseinfoMgr) Get() (result Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_BaseinfoMgr) Gets() (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_BaseinfoMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 
func (obj *_BaseinfoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIP ip获取 
func (obj *_BaseinfoMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithPort port获取 
func (obj *_BaseinfoMgr) WithPort(port string) Option {
	return optionFunc(func(o *options) { o.query["port"] = port })
}

// WithUser user获取 
func (obj *_BaseinfoMgr) WithUser(user string) Option {
	return optionFunc(func(o *options) { o.query["user"] = user })
}

// WithPwd pwd获取 
func (obj *_BaseinfoMgr) WithPwd(pwd string) Option {
	return optionFunc(func(o *options) { o.query["pwd"] = pwd })
}

// WithDname dname获取 
func (obj *_BaseinfoMgr) WithDname(dname string) Option {
	return optionFunc(func(o *options) { o.query["dname"] = dname })
}

// WithCreatetime createtime获取 
func (obj *_BaseinfoMgr) WithCreatetime(createtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtime"] = createtime })
}


// GetByOption 功能选项模式获取
func (obj *_BaseinfoMgr) GetByOption(opts ...Option) (result Baseinfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_BaseinfoMgr) GetByOptions(opts ...Option) (results []*Baseinfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	
	return
}
//////////////////////////enume case ////////////////////////////////////////////


// GetFromID 通过id获取内容  
func (obj *_BaseinfoMgr)  GetFromID(id uint) (result Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 
func (obj *_BaseinfoMgr) GetBatchFromID(ids []uint) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_BaseinfoMgr) GetFromName(name string) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_BaseinfoMgr) GetBatchFromName(names []string) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromIP 通过ip获取内容  
func (obj *_BaseinfoMgr) GetFromIP(ip string) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ip = ?", ip).Find(&results).Error
	
	return
}

// GetBatchFromIP 批量唯一主键查找 
func (obj *_BaseinfoMgr) GetBatchFromIP(ips []string) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ip IN (?)", ips).Find(&results).Error
	
	return
}
 
// GetFromPort 通过port获取内容  
func (obj *_BaseinfoMgr) GetFromPort(port string) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("port = ?", port).Find(&results).Error
	
	return
}

// GetBatchFromPort 批量唯一主键查找 
func (obj *_BaseinfoMgr) GetBatchFromPort(ports []string) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("port IN (?)", ports).Find(&results).Error
	
	return
}
 
// GetFromUser 通过user获取内容  
func (obj *_BaseinfoMgr) GetFromUser(user string) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user = ?", user).Find(&results).Error
	
	return
}

// GetBatchFromUser 批量唯一主键查找 
func (obj *_BaseinfoMgr) GetBatchFromUser(users []string) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user IN (?)", users).Find(&results).Error
	
	return
}
 
// GetFromPwd 通过pwd获取内容  
func (obj *_BaseinfoMgr) GetFromPwd(pwd string) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pwd = ?", pwd).Find(&results).Error
	
	return
}

// GetBatchFromPwd 批量唯一主键查找 
func (obj *_BaseinfoMgr) GetBatchFromPwd(pwds []string) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pwd IN (?)", pwds).Find(&results).Error
	
	return
}
 
// GetFromDname 通过dname获取内容  
func (obj *_BaseinfoMgr) GetFromDname(dname string) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("dname = ?", dname).Find(&results).Error
	
	return
}

// GetBatchFromDname 批量唯一主键查找 
func (obj *_BaseinfoMgr) GetBatchFromDname(dnames []string) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("dname IN (?)", dnames).Find(&results).Error
	
	return
}
 
// GetFromCreatetime 通过createtime获取内容  
func (obj *_BaseinfoMgr) GetFromCreatetime(createtime time.Time) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime = ?", createtime).Find(&results).Error
	
	return
}

// GetBatchFromCreatetime 批量唯一主键查找 
func (obj *_BaseinfoMgr) GetBatchFromCreatetime(createtimes []time.Time) (results []*Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime IN (?)", createtimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_BaseinfoMgr) FetchByPrimaryKey(id uint ) (result Baseinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 

 

	

