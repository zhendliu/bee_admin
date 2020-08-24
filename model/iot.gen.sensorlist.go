package	model	
import (	
"time"	
"github.com/jinzhu/gorm"	
"fmt"	
)	

type _SensorlistMgr struct {
	*_BaseMgr
}

// SensorlistMgr open func
func SensorlistMgr(db *gorm.DB) *_SensorlistMgr {
	if db == nil {
		panic(fmt.Errorf("SensorlistMgr need init by db"))
	}
	return &_SensorlistMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sensorlist"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SensorlistMgr) GetTableName() string {
	return "sensorlist"
}

// Get 获取
func (obj *_SensorlistMgr) Get() (result Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_SensorlistMgr) Gets() (results []*Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_SensorlistMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatetime createtime获取 
func (obj *_SensorlistMgr) WithCreatetime(createtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtime"] = createtime })
}

// WithSize size获取 
func (obj *_SensorlistMgr) WithSize(size int) Option {
	return optionFunc(func(o *options) { o.query["size"] = size })
}

// WithName name获取 
func (obj *_SensorlistMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTname tname获取 
func (obj *_SensorlistMgr) WithTname(tname string) Option {
	return optionFunc(func(o *options) { o.query["tname"] = tname })
}


// GetByOption 功能选项模式获取
func (obj *_SensorlistMgr) GetByOption(opts ...Option) (result Sensorlist, err error) {
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
func (obj *_SensorlistMgr) GetByOptions(opts ...Option) (results []*Sensorlist, err error) {
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
func (obj *_SensorlistMgr)  GetFromID(id uint) (result Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 
func (obj *_SensorlistMgr) GetBatchFromID(ids []uint) (results []*Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromCreatetime 通过createtime获取内容  
func (obj *_SensorlistMgr) GetFromCreatetime(createtime time.Time) (results []*Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime = ?", createtime).Find(&results).Error
	
	return
}

// GetBatchFromCreatetime 批量唯一主键查找 
func (obj *_SensorlistMgr) GetBatchFromCreatetime(createtimes []time.Time) (results []*Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime IN (?)", createtimes).Find(&results).Error
	
	return
}
 
// GetFromSize 通过size获取内容  
func (obj *_SensorlistMgr) GetFromSize(size int) (results []*Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("size = ?", size).Find(&results).Error
	
	return
}

// GetBatchFromSize 批量唯一主键查找 
func (obj *_SensorlistMgr) GetBatchFromSize(sizes []int) (results []*Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("size IN (?)", sizes).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_SensorlistMgr) GetFromName(name string) (results []*Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_SensorlistMgr) GetBatchFromName(names []string) (results []*Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromTname 通过tname获取内容  
func (obj *_SensorlistMgr) GetFromTname(tname string) (results []*Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("tname = ?", tname).Find(&results).Error
	
	return
}

// GetBatchFromTname 批量唯一主键查找 
func (obj *_SensorlistMgr) GetBatchFromTname(tnames []string) (results []*Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("tname IN (?)", tnames).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_SensorlistMgr) FetchByPrimaryKey(id uint ) (result Sensorlist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 

 

	

