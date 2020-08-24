package	model	
import (	
"github.com/jinzhu/gorm"	
"fmt"	
"time"	
)	

type _FileinfoMgr struct {
	*_BaseMgr
}

// FileinfoMgr open func
func FileinfoMgr(db *gorm.DB) *_FileinfoMgr {
	if db == nil {
		panic(fmt.Errorf("FileinfoMgr need init by db"))
	}
	return &_FileinfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fileinfo"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FileinfoMgr) GetTableName() string {
	return "fileinfo"
}

// Get 获取
func (obj *_FileinfoMgr) Get() (result Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_FileinfoMgr) Gets() (results []*Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_FileinfoMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 
func (obj *_FileinfoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPath path获取 
func (obj *_FileinfoMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithFname fname获取 
func (obj *_FileinfoMgr) WithFname(fname string) Option {
	return optionFunc(func(o *options) { o.query["fname"] = fname })
}

// WithCreatetime createtime获取 
func (obj *_FileinfoMgr) WithCreatetime(createtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtime"] = createtime })
}


// GetByOption 功能选项模式获取
func (obj *_FileinfoMgr) GetByOption(opts ...Option) (result Fileinfo, err error) {
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
func (obj *_FileinfoMgr) GetByOptions(opts ...Option) (results []*Fileinfo, err error) {
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
func (obj *_FileinfoMgr)  GetFromID(id uint) (result Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 
func (obj *_FileinfoMgr) GetBatchFromID(ids []uint) (results []*Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_FileinfoMgr) GetFromName(name string) (results []*Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_FileinfoMgr) GetBatchFromName(names []string) (results []*Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromPath 通过path获取内容  
func (obj *_FileinfoMgr) GetFromPath(path string) (results []*Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("path = ?", path).Find(&results).Error
	
	return
}

// GetBatchFromPath 批量唯一主键查找 
func (obj *_FileinfoMgr) GetBatchFromPath(paths []string) (results []*Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("path IN (?)", paths).Find(&results).Error
	
	return
}
 
// GetFromFname 通过fname获取内容  
func (obj *_FileinfoMgr) GetFromFname(fname string) (results []*Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("fname = ?", fname).Find(&results).Error
	
	return
}

// GetBatchFromFname 批量唯一主键查找 
func (obj *_FileinfoMgr) GetBatchFromFname(fnames []string) (results []*Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("fname IN (?)", fnames).Find(&results).Error
	
	return
}
 
// GetFromCreatetime 通过createtime获取内容  
func (obj *_FileinfoMgr) GetFromCreatetime(createtime time.Time) (results []*Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime = ?", createtime).Find(&results).Error
	
	return
}

// GetBatchFromCreatetime 批量唯一主键查找 
func (obj *_FileinfoMgr) GetBatchFromCreatetime(createtimes []time.Time) (results []*Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime IN (?)", createtimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_FileinfoMgr) FetchByPrimaryKey(id uint ) (result Fileinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 

 

	

