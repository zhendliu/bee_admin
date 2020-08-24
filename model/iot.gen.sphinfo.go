package	model	
import (	
"github.com/jinzhu/gorm"	
"fmt"	
"time"	
)	

type _SphinfoMgr struct {
	*_BaseMgr
}

// SphinfoMgr open func
func SphinfoMgr(db *gorm.DB) *_SphinfoMgr {
	if db == nil {
		panic(fmt.Errorf("SphinfoMgr need init by db"))
	}
	return &_SphinfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sphinfo"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SphinfoMgr) GetTableName() string {
	return "sphinfo"
}

// Get 获取
func (obj *_SphinfoMgr) Get() (result Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_SphinfoMgr) Gets() (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_SphinfoMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithHgp hgp获取 
func (obj *_SphinfoMgr) WithHgp(hgp string) Option {
	return optionFunc(func(o *options) { o.query["hgp"] = hgp })
}

// WithHtp htp获取 
func (obj *_SphinfoMgr) WithHtp(htp string) Option {
	return optionFunc(func(o *options) { o.query["htp"] = htp })
}

// WithCtrl ctrl获取 
func (obj *_SphinfoMgr) WithCtrl(ctrl string) Option {
	return optionFunc(func(o *options) { o.query["ctrl"] = ctrl })
}

// WithCycle cycle获取 
func (obj *_SphinfoMgr) WithCycle(cycle int) Option {
	return optionFunc(func(o *options) { o.query["cycle"] = cycle })
}

// WithSw sw获取 
func (obj *_SphinfoMgr) WithSw(sw int) Option {
	return optionFunc(func(o *options) { o.query["sw"] = sw })
}

// WithProtocolid protocolid获取 
func (obj *_SphinfoMgr) WithProtocolid(protocolid int) Option {
	return optionFunc(func(o *options) { o.query["protocolid"] = protocolid })
}

// WithDatabaseid databaseid获取 
func (obj *_SphinfoMgr) WithDatabaseid(databaseid string) Option {
	return optionFunc(func(o *options) { o.query["databaseid"] = databaseid })
}

// WithFileid fileid获取 
func (obj *_SphinfoMgr) WithFileid(fileid int) Option {
	return optionFunc(func(o *options) { o.query["fileid"] = fileid })
}

// WithRedisid redisid获取 
func (obj *_SphinfoMgr) WithRedisid(redisid int) Option {
	return optionFunc(func(o *options) { o.query["redisid"] = redisid })
}

// WithCreatetime createtime获取 
func (obj *_SphinfoMgr) WithCreatetime(createtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtime"] = createtime })
}

// WithInfo info获取 
func (obj *_SphinfoMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}


// GetByOption 功能选项模式获取
func (obj *_SphinfoMgr) GetByOption(opts ...Option) (result Sphinfo, err error) {
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
func (obj *_SphinfoMgr) GetByOptions(opts ...Option) (results []*Sphinfo, err error) {
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
func (obj *_SphinfoMgr)  GetFromID(id uint) (result Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 
func (obj *_SphinfoMgr) GetBatchFromID(ids []uint) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromHgp 通过hgp获取内容  
func (obj *_SphinfoMgr) GetFromHgp(hgp string) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hgp = ?", hgp).Find(&results).Error
	
	return
}

// GetBatchFromHgp 批量唯一主键查找 
func (obj *_SphinfoMgr) GetBatchFromHgp(hgps []string) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hgp IN (?)", hgps).Find(&results).Error
	
	return
}
 
// GetFromHtp 通过htp获取内容  
func (obj *_SphinfoMgr) GetFromHtp(htp string) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("htp = ?", htp).Find(&results).Error
	
	return
}

// GetBatchFromHtp 批量唯一主键查找 
func (obj *_SphinfoMgr) GetBatchFromHtp(htps []string) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("htp IN (?)", htps).Find(&results).Error
	
	return
}
 
// GetFromCtrl 通过ctrl获取内容  
func (obj *_SphinfoMgr) GetFromCtrl(ctrl string) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ctrl = ?", ctrl).Find(&results).Error
	
	return
}

// GetBatchFromCtrl 批量唯一主键查找 
func (obj *_SphinfoMgr) GetBatchFromCtrl(ctrls []string) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ctrl IN (?)", ctrls).Find(&results).Error
	
	return
}
 
// GetFromCycle 通过cycle获取内容  
func (obj *_SphinfoMgr) GetFromCycle(cycle int) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cycle = ?", cycle).Find(&results).Error
	
	return
}

// GetBatchFromCycle 批量唯一主键查找 
func (obj *_SphinfoMgr) GetBatchFromCycle(cycles []int) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cycle IN (?)", cycles).Find(&results).Error
	
	return
}
 
// GetFromSw 通过sw获取内容  
func (obj *_SphinfoMgr) GetFromSw(sw int) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sw = ?", sw).Find(&results).Error
	
	return
}

// GetBatchFromSw 批量唯一主键查找 
func (obj *_SphinfoMgr) GetBatchFromSw(sws []int) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sw IN (?)", sws).Find(&results).Error
	
	return
}
 
// GetFromProtocolid 通过protocolid获取内容  
func (obj *_SphinfoMgr) GetFromProtocolid(protocolid int) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("protocolid = ?", protocolid).Find(&results).Error
	
	return
}

// GetBatchFromProtocolid 批量唯一主键查找 
func (obj *_SphinfoMgr) GetBatchFromProtocolid(protocolids []int) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("protocolid IN (?)", protocolids).Find(&results).Error
	
	return
}
 
// GetFromDatabaseid 通过databaseid获取内容  
func (obj *_SphinfoMgr) GetFromDatabaseid(databaseid string) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("databaseid = ?", databaseid).Find(&results).Error
	
	return
}

// GetBatchFromDatabaseid 批量唯一主键查找 
func (obj *_SphinfoMgr) GetBatchFromDatabaseid(databaseids []string) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("databaseid IN (?)", databaseids).Find(&results).Error
	
	return
}
 
// GetFromFileid 通过fileid获取内容  
func (obj *_SphinfoMgr) GetFromFileid(fileid int) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("fileid = ?", fileid).Find(&results).Error
	
	return
}

// GetBatchFromFileid 批量唯一主键查找 
func (obj *_SphinfoMgr) GetBatchFromFileid(fileids []int) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("fileid IN (?)", fileids).Find(&results).Error
	
	return
}
 
// GetFromRedisid 通过redisid获取内容  
func (obj *_SphinfoMgr) GetFromRedisid(redisid int) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("redisid = ?", redisid).Find(&results).Error
	
	return
}

// GetBatchFromRedisid 批量唯一主键查找 
func (obj *_SphinfoMgr) GetBatchFromRedisid(redisids []int) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("redisid IN (?)", redisids).Find(&results).Error
	
	return
}
 
// GetFromCreatetime 通过createtime获取内容  
func (obj *_SphinfoMgr) GetFromCreatetime(createtime time.Time) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime = ?", createtime).Find(&results).Error
	
	return
}

// GetBatchFromCreatetime 批量唯一主键查找 
func (obj *_SphinfoMgr) GetBatchFromCreatetime(createtimes []time.Time) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("createtime IN (?)", createtimes).Find(&results).Error
	
	return
}
 
// GetFromInfo 通过info获取内容  
func (obj *_SphinfoMgr) GetFromInfo(info string) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info = ?", info).Find(&results).Error
	
	return
}

// GetBatchFromInfo 批量唯一主键查找 
func (obj *_SphinfoMgr) GetBatchFromInfo(infos []string) (results []*Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info IN (?)", infos).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_SphinfoMgr) FetchByPrimaryKey(id uint ) (result Sphinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 

 

	

