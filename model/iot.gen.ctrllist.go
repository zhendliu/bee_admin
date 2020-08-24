package	model	
import (	
"fmt"	
"time"	
"github.com/jinzhu/gorm"	
)	

type _CtrllistMgr struct {
	*_BaseMgr
}

// CtrllistMgr open func
func CtrllistMgr(db *gorm.DB) *_CtrllistMgr {
	if db == nil {
		panic(fmt.Errorf("CtrllistMgr need init by db"))
	}
	return &_CtrllistMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ctrllist"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CtrllistMgr) GetTableName() string {
	return "ctrllist"
}

// Get 获取
func (obj *_CtrllistMgr) Get() (result Ctrllist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_CtrllistMgr) Gets() (results []*Ctrllist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_CtrllistMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatetime createtime获取 
func (obj *_CtrllistMgr) WithCreatetime(createtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtime"] = createtime })
}

// WithSize size获取 
func (obj *_CtrllistMgr) WithSize(size int) Option {
	return optionFunc(func(o *options) { o.query["size"] = size })
}

// WithName name获取 
func (obj *_CtrllistMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_CtrllistMgr) GetByOption(opts ...Option) (result Ctrllist, err error) {
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
func (obj *_CtrllistMgr) GetByOptions(opts ...Option) (results []*Ctrllist, err error) {
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
func (obj *_CtrllistMgr)  GetFromID(id uint) (result Ctrllist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 
func (obj *_CtrllistMgr) GetBatchFromID(ids []uint) (results []*Ctrllist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}


 
// GetFromCreatetime 通过createtime获取内容  
func (obj *_CtrllistMgr) GetFromCreatetime(createtime time.Time) (results []*Ctrllist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime = ?", createtime).Find(&results).Error
	
	return
}

// GetBatchFromCreatetime 批量唯一主键查找 
func (obj *_CtrllistMgr) GetBatchFromCreatetime(createtimes []time.Time) (results []*Ctrllist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime IN (?)", createtimes).Find(&results).Error
	
	return
}
 
// GetFromSize 通过size获取内容  
func (obj *_CtrllistMgr) GetFromSize(size int) (results []*Ctrllist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("size = ?", size).Find(&results).Error
	
	return
}

// GetBatchFromSize 批量唯一主键查找 
func (obj *_CtrllistMgr) GetBatchFromSize(sizes []int) (results []*Ctrllist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("size IN (?)", sizes).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_CtrllistMgr) GetFromName(name string) (results []*Ctrllist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_CtrllistMgr) GetBatchFromName(names []string) (results []*Ctrllist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_CtrllistMgr) FetchByPrimaryKey(id uint ) (result Ctrllist, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 

 

	

