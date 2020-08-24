package	model	
import (	
"fmt"	
"time"	
"github.com/jinzhu/gorm"	
)	

type _CphinfoMgr struct {
	*_BaseMgr
}

// CphinfoMgr open func
func CphinfoMgr(db *gorm.DB) *_CphinfoMgr {
	if db == nil {
		panic(fmt.Errorf("CphinfoMgr need init by db"))
	}
	return &_CphinfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("cphinfo"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CphinfoMgr) GetTableName() string {
	return "cphinfo"
}

// Get 获取
func (obj *_CphinfoMgr) Get() (result Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_CphinfoMgr) Gets() (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_CphinfoMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithHgp hgp获取 
func (obj *_CphinfoMgr) WithHgp(hgp string) Option {
	return optionFunc(func(o *options) { o.query["hgp"] = hgp })
}

// WithHtp htp获取 
func (obj *_CphinfoMgr) WithHtp(htp string) Option {
	return optionFunc(func(o *options) { o.query["htp"] = htp })
}

// WithCtrl ctrl获取 
func (obj *_CphinfoMgr) WithCtrl(ctrl string) Option {
	return optionFunc(func(o *options) { o.query["ctrl"] = ctrl })
}

// WithCycle cycle获取 
func (obj *_CphinfoMgr) WithCycle(cycle int) Option {
	return optionFunc(func(o *options) { o.query["cycle"] = cycle })
}

// WithSw sw获取 
func (obj *_CphinfoMgr) WithSw(sw int) Option {
	return optionFunc(func(o *options) { o.query["sw"] = sw })
}

// WithProtocolid protocolid获取 
func (obj *_CphinfoMgr) WithProtocolid(protocolid int) Option {
	return optionFunc(func(o *options) { o.query["protocolid"] = protocolid })
}

// WithCreatetime createtime获取 
func (obj *_CphinfoMgr) WithCreatetime(createtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtime"] = createtime })
}

// WithInfo info获取 
func (obj *_CphinfoMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}


// GetByOption 功能选项模式获取
func (obj *_CphinfoMgr) GetByOption(opts ...Option) (result Cphinfo, err error) {
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
func (obj *_CphinfoMgr) GetByOptions(opts ...Option) (results []*Cphinfo, err error) {
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
func (obj *_CphinfoMgr)  GetFromID(id uint) (result Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 
func (obj *_CphinfoMgr) GetBatchFromID(ids []uint) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromHgp 通过hgp获取内容  
func (obj *_CphinfoMgr) GetFromHgp(hgp string) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hgp = ?", hgp).Find(&results).Error
	
	return
}

// GetBatchFromHgp 批量唯一主键查找 
func (obj *_CphinfoMgr) GetBatchFromHgp(hgps []string) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hgp IN (?)", hgps).Find(&results).Error
	
	return
}
 
// GetFromHtp 通过htp获取内容  
func (obj *_CphinfoMgr) GetFromHtp(htp string) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("htp = ?", htp).Find(&results).Error
	
	return
}

// GetBatchFromHtp 批量唯一主键查找 
func (obj *_CphinfoMgr) GetBatchFromHtp(htps []string) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("htp IN (?)", htps).Find(&results).Error
	
	return
}
 
// GetFromCtrl 通过ctrl获取内容  
func (obj *_CphinfoMgr) GetFromCtrl(ctrl string) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ctrl = ?", ctrl).Find(&results).Error
	
	return
}

// GetBatchFromCtrl 批量唯一主键查找 
func (obj *_CphinfoMgr) GetBatchFromCtrl(ctrls []string) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ctrl IN (?)", ctrls).Find(&results).Error
	
	return
}
 
// GetFromCycle 通过cycle获取内容  
func (obj *_CphinfoMgr) GetFromCycle(cycle int) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cycle = ?", cycle).Find(&results).Error
	
	return
}

// GetBatchFromCycle 批量唯一主键查找 
func (obj *_CphinfoMgr) GetBatchFromCycle(cycles []int) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cycle IN (?)", cycles).Find(&results).Error
	
	return
}
 
// GetFromSw 通过sw获取内容  
func (obj *_CphinfoMgr) GetFromSw(sw int) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sw = ?", sw).Find(&results).Error
	
	return
}

// GetBatchFromSw 批量唯一主键查找 
func (obj *_CphinfoMgr) GetBatchFromSw(sws []int) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sw IN (?)", sws).Find(&results).Error
	
	return
}
 
// GetFromProtocolid 通过protocolid获取内容  
func (obj *_CphinfoMgr) GetFromProtocolid(protocolid int) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("protocolid = ?", protocolid).Find(&results).Error
	
	return
}

// GetBatchFromProtocolid 批量唯一主键查找 
func (obj *_CphinfoMgr) GetBatchFromProtocolid(protocolids []int) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("protocolid IN (?)", protocolids).Find(&results).Error
	
	return
}
 
// GetFromCreatetime 通过createtime获取内容  
func (obj *_CphinfoMgr) GetFromCreatetime(createtime time.Time) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime = ?", createtime).Find(&results).Error
	
	return
}

// GetBatchFromCreatetime 批量唯一主键查找 
func (obj *_CphinfoMgr) GetBatchFromCreatetime(createtimes []time.Time) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime IN (?)", createtimes).Find(&results).Error
	
	return
}
 
// GetFromInfo 通过info获取内容  
func (obj *_CphinfoMgr) GetFromInfo(info string) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info = ?", info).Find(&results).Error
	
	return
}

// GetBatchFromInfo 批量唯一主键查找 
func (obj *_CphinfoMgr) GetBatchFromInfo(infos []string) (results []*Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info IN (?)", infos).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_CphinfoMgr) FetchByPrimaryKey(id uint ) (result Cphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 

 

	

