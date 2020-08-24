package	model	
import (	
"fmt"	
"time"	
"github.com/jinzhu/gorm"	
)	

type _CtrlinfoMgr struct {
	*_BaseMgr
}

// CtrlinfoMgr open func
func CtrlinfoMgr(db *gorm.DB) *_CtrlinfoMgr {
	if db == nil {
		panic(fmt.Errorf("CtrlinfoMgr need init by db"))
	}
	return &_CtrlinfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ctrlinfo"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CtrlinfoMgr) GetTableName() string {
	return "ctrlinfo"
}

// Get 获取
func (obj *_CtrlinfoMgr) Get() (result Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_CtrlinfoMgr) Gets() (results []*Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_CtrlinfoMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSensorlist sensorlist获取 
func (obj *_CtrlinfoMgr) WithSensorlist(sensorlist int) Option {
	return optionFunc(func(o *options) { o.query["sensorlist"] = sensorlist })
}

// WithSensorname sensorname获取 
func (obj *_CtrlinfoMgr) WithSensorname(sensorname string) Option {
	return optionFunc(func(o *options) { o.query["sensorname"] = sensorname })
}

// WithInfo info获取 
func (obj *_CtrlinfoMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithCreatetme createtme获取 
func (obj *_CtrlinfoMgr) WithCreatetme(createtme time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtme"] = createtme })
}


// GetByOption 功能选项模式获取
func (obj *_CtrlinfoMgr) GetByOption(opts ...Option) (result Ctrlinfo, err error) {
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
func (obj *_CtrlinfoMgr) GetByOptions(opts ...Option) (results []*Ctrlinfo, err error) {
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
func (obj *_CtrlinfoMgr)  GetFromID(id uint) (result Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 
func (obj *_CtrlinfoMgr) GetBatchFromID(ids []uint) (results []*Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromSensorlist 通过sensorlist获取内容  
func (obj *_CtrlinfoMgr) GetFromSensorlist(sensorlist int) (results []*Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sensorlist = ?", sensorlist).Find(&results).Error
	
	return
}

// GetBatchFromSensorlist 批量唯一主键查找 
func (obj *_CtrlinfoMgr) GetBatchFromSensorlist(sensorlists []int) (results []*Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sensorlist IN (?)", sensorlists).Find(&results).Error
	
	return
}
 
// GetFromSensorname 通过sensorname获取内容  
func (obj *_CtrlinfoMgr) GetFromSensorname(sensorname string) (results []*Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sensorname = ?", sensorname).Find(&results).Error
	
	return
}

// GetBatchFromSensorname 批量唯一主键查找 
func (obj *_CtrlinfoMgr) GetBatchFromSensorname(sensornames []string) (results []*Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sensorname IN (?)", sensornames).Find(&results).Error
	
	return
}
 
// GetFromInfo 通过info获取内容  
func (obj *_CtrlinfoMgr) GetFromInfo(info string) (results []*Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info = ?", info).Find(&results).Error
	
	return
}

// GetBatchFromInfo 批量唯一主键查找 
func (obj *_CtrlinfoMgr) GetBatchFromInfo(infos []string) (results []*Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info IN (?)", infos).Find(&results).Error
	
	return
}
 
// GetFromCreatetme 通过createtme获取内容  
func (obj *_CtrlinfoMgr) GetFromCreatetme(createtme time.Time) (results []*Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtme = ?", createtme).Find(&results).Error
	
	return
}

// GetBatchFromCreatetme 批量唯一主键查找 
func (obj *_CtrlinfoMgr) GetBatchFromCreatetme(createtmes []time.Time) (results []*Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtme IN (?)", createtmes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_CtrlinfoMgr) FetchByPrimaryKey(id uint ) (result Ctrlinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 

 

	

