package	model	
import (	
"github.com/jinzhu/gorm"	
"fmt"	
)	

type _ServerMgr struct {
	*_BaseMgr
}

// ServerMgr open func
func ServerMgr(db *gorm.DB) *_ServerMgr {
	if db == nil {
		panic(fmt.Errorf("ServerMgr need init by db"))
	}
	return &_ServerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("server"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ServerMgr) GetTableName() string {
	return "server"
}

// Get 获取
func (obj *_ServerMgr) Get() (result Server, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_ServerMgr) Gets() (results []*Server, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithPort port获取 
func (obj *_ServerMgr) WithPort(port int) Option {
	return optionFunc(func(o *options) { o.query["port"] = port })
}

// WithInfo info获取 
func (obj *_ServerMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}


// GetByOption 功能选项模式获取
func (obj *_ServerMgr) GetByOption(opts ...Option) (result Server, err error) {
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
func (obj *_ServerMgr) GetByOptions(opts ...Option) (results []*Server, err error) {
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


// GetFromPort 通过port获取内容  
func (obj *_ServerMgr) GetFromPort(port int) (results []*Server, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("port = ?", port).Find(&results).Error
	
	return
}

// GetBatchFromPort 批量唯一主键查找 
func (obj *_ServerMgr) GetBatchFromPort(ports []int) (results []*Server, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("port IN (?)", ports).Find(&results).Error
	
	return
}
 
// GetFromInfo 通过info获取内容  
func (obj *_ServerMgr) GetFromInfo(info string) (results []*Server, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info = ?", info).Find(&results).Error
	
	return
}

// GetBatchFromInfo 批量唯一主键查找 
func (obj *_ServerMgr) GetBatchFromInfo(infos []string) (results []*Server, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info IN (?)", infos).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

