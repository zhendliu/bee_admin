package	model	
import (	
"github.com/jinzhu/gorm"	
"fmt"	
)	

type _RedisinfoMgr struct {
	*_BaseMgr
}

// RedisinfoMgr open func
func RedisinfoMgr(db *gorm.DB) *_RedisinfoMgr {
	if db == nil {
		panic(fmt.Errorf("RedisinfoMgr need init by db"))
	}
	return &_RedisinfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("redisinfo"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RedisinfoMgr) GetTableName() string {
	return "redisinfo"
}

// Get 获取
func (obj *_RedisinfoMgr) Get() (result Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_RedisinfoMgr) Gets() (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_RedisinfoMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 
func (obj *_RedisinfoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIP ip获取 
func (obj *_RedisinfoMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithPort port获取 
func (obj *_RedisinfoMgr) WithPort(port string) Option {
	return optionFunc(func(o *options) { o.query["port"] = port })
}

// WithAuth auth获取 
func (obj *_RedisinfoMgr) WithAuth(auth string) Option {
	return optionFunc(func(o *options) { o.query["auth"] = auth })
}

// WithQueuename queuename获取 
func (obj *_RedisinfoMgr) WithQueuename(queuename string) Option {
	return optionFunc(func(o *options) { o.query["queuename"] = queuename })
}

// WithCreatetime createtime获取 
func (obj *_RedisinfoMgr) WithCreatetime(createtime string) Option {
	return optionFunc(func(o *options) { o.query["createtime"] = createtime })
}


// GetByOption 功能选项模式获取
func (obj *_RedisinfoMgr) GetByOption(opts ...Option) (result Redisinfo, err error) {
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
func (obj *_RedisinfoMgr) GetByOptions(opts ...Option) (results []*Redisinfo, err error) {
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
func (obj *_RedisinfoMgr)  GetFromID(id uint) (result Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 
func (obj *_RedisinfoMgr) GetBatchFromID(ids []uint) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_RedisinfoMgr) GetFromName(name string) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_RedisinfoMgr) GetBatchFromName(names []string) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromIP 通过ip获取内容  
func (obj *_RedisinfoMgr) GetFromIP(ip string) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ip = ?", ip).Find(&results).Error
	
	return
}

// GetBatchFromIP 批量唯一主键查找 
func (obj *_RedisinfoMgr) GetBatchFromIP(ips []string) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ip IN (?)", ips).Find(&results).Error
	
	return
}
 
// GetFromPort 通过port获取内容  
func (obj *_RedisinfoMgr) GetFromPort(port string) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("port = ?", port).Find(&results).Error
	
	return
}

// GetBatchFromPort 批量唯一主键查找 
func (obj *_RedisinfoMgr) GetBatchFromPort(ports []string) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("port IN (?)", ports).Find(&results).Error
	
	return
}
 
// GetFromAuth 通过auth获取内容  
func (obj *_RedisinfoMgr) GetFromAuth(auth string) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("auth = ?", auth).Find(&results).Error
	
	return
}

// GetBatchFromAuth 批量唯一主键查找 
func (obj *_RedisinfoMgr) GetBatchFromAuth(auths []string) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("auth IN (?)", auths).Find(&results).Error
	
	return
}
 
// GetFromQueuename 通过queuename获取内容  
func (obj *_RedisinfoMgr) GetFromQueuename(queuename string) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("queuename = ?", queuename).Find(&results).Error
	
	return
}

// GetBatchFromQueuename 批量唯一主键查找 
func (obj *_RedisinfoMgr) GetBatchFromQueuename(queuenames []string) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("queuename IN (?)", queuenames).Find(&results).Error
	
	return
}
 
// GetFromCreatetime 通过createtime获取内容  
func (obj *_RedisinfoMgr) GetFromCreatetime(createtime string) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime = ?", createtime).Find(&results).Error
	
	return
}

// GetBatchFromCreatetime 批量唯一主键查找 
func (obj *_RedisinfoMgr) GetBatchFromCreatetime(createtimes []string) (results []*Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime IN (?)", createtimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_RedisinfoMgr) FetchByPrimaryKey(id uint ) (result Redisinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 

 

	

