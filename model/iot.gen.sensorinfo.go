package	model	
import (	
"github.com/jinzhu/gorm"	
"fmt"	
"time"	
)	

type _SensorinfoMgr struct {
	*_BaseMgr
}

// SensorinfoMgr open func
func SensorinfoMgr(db *gorm.DB) *_SensorinfoMgr {
	if db == nil {
		panic(fmt.Errorf("SensorinfoMgr need init by db"))
	}
	return &_SensorinfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sensorinfo"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SensorinfoMgr) GetTableName() string {
	return "sensorinfo"
}

// Get 获取
func (obj *_SensorinfoMgr) Get() (result Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_SensorinfoMgr) Gets() (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_SensorinfoMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSensorlist sensorlist获取 
func (obj *_SensorinfoMgr) WithSensorlist(sensorlist int) Option {
	return optionFunc(func(o *options) { o.query["sensorlist"] = sensorlist })
}

// WithSensorname sensorname获取 
func (obj *_SensorinfoMgr) WithSensorname(sensorname string) Option {
	return optionFunc(func(o *options) { o.query["sensorname"] = sensorname })
}

// WithSavename savename获取 
func (obj *_SensorinfoMgr) WithSavename(savename string) Option {
	return optionFunc(func(o *options) { o.query["savename"] = savename })
}

// WithTransform transform获取 
func (obj *_SensorinfoMgr) WithTransform(transform string) Option {
	return optionFunc(func(o *options) { o.query["transform"] = transform })
}

// WithUnit unit获取 
func (obj *_SensorinfoMgr) WithUnit(unit string) Option {
	return optionFunc(func(o *options) { o.query["unit"] = unit })
}

// WithMaxv maxv获取 
func (obj *_SensorinfoMgr) WithMaxv(maxv string) Option {
	return optionFunc(func(o *options) { o.query["maxv"] = maxv })
}

// WithInfo info获取 
func (obj *_SensorinfoMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithCreatetme createtme获取 
func (obj *_SensorinfoMgr) WithCreatetme(createtme time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtme"] = createtme })
}


// GetByOption 功能选项模式获取
func (obj *_SensorinfoMgr) GetByOption(opts ...Option) (result Sensorinfo, err error) {
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
func (obj *_SensorinfoMgr) GetByOptions(opts ...Option) (results []*Sensorinfo, err error) {
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
func (obj *_SensorinfoMgr)  GetFromID(id uint) (result Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 
func (obj *_SensorinfoMgr) GetBatchFromID(ids []uint) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromSensorlist 通过sensorlist获取内容  
func (obj *_SensorinfoMgr) GetFromSensorlist(sensorlist int) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sensorlist = ?", sensorlist).Find(&results).Error
	
	return
}

// GetBatchFromSensorlist 批量唯一主键查找 
func (obj *_SensorinfoMgr) GetBatchFromSensorlist(sensorlists []int) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sensorlist IN (?)", sensorlists).Find(&results).Error
	
	return
}
 
// GetFromSensorname 通过sensorname获取内容  
func (obj *_SensorinfoMgr) GetFromSensorname(sensorname string) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sensorname = ?", sensorname).Find(&results).Error
	
	return
}

// GetBatchFromSensorname 批量唯一主键查找 
func (obj *_SensorinfoMgr) GetBatchFromSensorname(sensornames []string) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sensorname IN (?)", sensornames).Find(&results).Error
	
	return
}
 
// GetFromSavename 通过savename获取内容  
func (obj *_SensorinfoMgr) GetFromSavename(savename string) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("savename = ?", savename).Find(&results).Error
	
	return
}

// GetBatchFromSavename 批量唯一主键查找 
func (obj *_SensorinfoMgr) GetBatchFromSavename(savenames []string) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("savename IN (?)", savenames).Find(&results).Error
	
	return
}
 
// GetFromTransform 通过transform获取内容  
func (obj *_SensorinfoMgr) GetFromTransform(transform string) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transform = ?", transform).Find(&results).Error
	
	return
}

// GetBatchFromTransform 批量唯一主键查找 
func (obj *_SensorinfoMgr) GetBatchFromTransform(transforms []string) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transform IN (?)", transforms).Find(&results).Error
	
	return
}
 
// GetFromUnit 通过unit获取内容  
func (obj *_SensorinfoMgr) GetFromUnit(unit string) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unit = ?", unit).Find(&results).Error
	
	return
}

// GetBatchFromUnit 批量唯一主键查找 
func (obj *_SensorinfoMgr) GetBatchFromUnit(units []string) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unit IN (?)", units).Find(&results).Error
	
	return
}
 
// GetFromMaxv 通过maxv获取内容  
func (obj *_SensorinfoMgr) GetFromMaxv(maxv string) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("maxv = ?", maxv).Find(&results).Error
	
	return
}

// GetBatchFromMaxv 批量唯一主键查找 
func (obj *_SensorinfoMgr) GetBatchFromMaxv(maxvs []string) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("maxv IN (?)", maxvs).Find(&results).Error
	
	return
}
 
// GetFromInfo 通过info获取内容  
func (obj *_SensorinfoMgr) GetFromInfo(info string) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info = ?", info).Find(&results).Error
	
	return
}

// GetBatchFromInfo 批量唯一主键查找 
func (obj *_SensorinfoMgr) GetBatchFromInfo(infos []string) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info IN (?)", infos).Find(&results).Error
	
	return
}
 
// GetFromCreatetme 通过createtme获取内容  
func (obj *_SensorinfoMgr) GetFromCreatetme(createtme time.Time) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtme = ?", createtme).Find(&results).Error
	
	return
}

// GetBatchFromCreatetme 批量唯一主键查找 
func (obj *_SensorinfoMgr) GetBatchFromCreatetme(createtmes []time.Time) (results []*Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtme IN (?)", createtmes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_SensorinfoMgr) FetchByPrimaryKey(id uint ) (result Sensorinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 

 

	

