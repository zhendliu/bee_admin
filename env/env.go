package env

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	iniconf "github.com/clod-moon/goconf"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"itoServer/model"
	"itoServer/server"
	"itoServer/util"
	"os"
	"strconv"
	"sync"
)

/*
	数据库信息描述
*/
type DbInfo struct {
	Host   string
	User   string
	PassWd string
	Port   string
	DbName string
}

/*
	设备信息
*/
type DEVICEMAP struct {
	//DeviceMap map[string]*model.Tbl_wlw_device
	Mx sync.RWMutex
}

/*
	ITO设备配置信息
*/

type ItoInfo struct {
	SendCycle  string
	SensorRGPH string
	SensorHTPH string
}

var (
	DBINFO    DbInfo
	DM        DEVICEMAP
	INF       ItoInfo
	IotLog    *logs.BeeLogger
	DbHandler *gorm.DB
	DbErr     error
)

/*
	初始化整体环境配置
*/
func init() {
	fmt.Println("开始初始化")

	//初始换配置文件
	conf := iniconf.InitConfig("conf/ito.ini")
	//初始化数据库
	DBINFO.User = conf.GetValue("database", "username")
	DBINFO.PassWd = conf.GetValue("database", "password")
	DBINFO.Host = conf.GetValue("database", "host")
	DBINFO.Port = conf.GetValue("database", "port")
	DBINFO.DbName = conf.GetValue("database", "database")
	//初始化数据库
	initDbHandle()

	//数据设备缓存模块
	//DM.DeviceMap = make(map[string]*model.Tbl_wlw_device)

	//获取ITO传感设备配置信息
	INF.SendCycle = conf.GetValue("ito", "sendCycle")
	INF.SensorRGPH = conf.GetValue("ito", "sensorRGPH")
	INF.SensorHTPH = conf.GetValue("ito", "sensorHTPH")

	//获取日志配置

	IotLog = logs.NewLogger(10000) // 创建一个日志记录器，参数为缓冲区的大小
	// 设置配置文件
	jsonConfig := `{
        "filename" : "logs/iot.log", 
		"maxlines":0,
		"maxsize" : 0, 
		"daily": true,
		"daily":true,
		"maxdays":90,
		"color":true
    }`
	IotLog.SetLogger("file", jsonConfig)     // 设置日志记录方式：本地文件记录
	IotLog.SetLevel(logs.LevelInformational) // 设置日志写入缓冲区的等级
	IotLog.EnableFuncCallDepth(true)         // 输出log时能显示输出文件名和行号（非必须）
	//开启协议加载
	initServer()

	//
}

/*
	服务信息加载
*/
func initServer() {

	sm := model.ServerMgr(DbHandler)
	result, err := sm.Gets()
	CheckErr(err)
	if len(result) > 0 {
		server.IotServer.Port = result[0].Port
		server.IotServer.Info = result[0].Info
	}

	//开启传感器头部信息
	sp := model.SphinfoMgr(DbHandler)
	cl := model.CtrlinfoMgr(DbHandler)
	cp := model.CphinfoMgr(DbHandler)
	se := model.SensorinfoMgr(DbHandler)
	sl := model.SensorlistMgr(DbHandler)
	bi := model.BaseinfoMgr(DbHandler)
	ff := model.FileinfoMgr(DbHandler)
	rs := model.RedisinfoMgr(DbHandler)
	sphinfos, err := sp.Gets()
	CheckErr(err)
	for k, v := range sphinfos {
		//开始循环开机协议
		fmt.Printf("协议：%s：--> %s", k, v)
		ptList ,err :=se.GetFromSensorlist(v.Protocolid)

		//在数据库中检查表结构
		//1 表名：
		//2 数据结构写入
		sensorlist ,err  :=sl.FetchByPrimaryKey(uint(v.Protocolid))
		var  tableName  =sensorlist.Tname
		var fileSlice = make([]string,0)
		for  _,v  :=range   ptList{
			fileSlice =append(fileSlice, v.Savename)
		}
		util.InitTables(DbHandler,tableName,fileSlice)
		//获取协议信息
		CheckErr(err)
		//初始化数据库
		var dbhandler *gorm.DB

		//初始化数据库，通过数据id
		id, _ := strconv.Atoi(v.Databaseid)
		if  v.Databaseid !="" && v.Databaseid !="0"{
			result ,err :=bi.FetchByPrimaryKey(uint(id))
			CheckErr(err)
			dbhandler = getDbHandler(result)
		}
		//查询存储的表名.通过协议
		slist,err :=sl.FetchByPrimaryKey(uint(v.Protocolid))
		//设置存储文件名
		var FileHandler *os.File
		if  v.Fileid !=0 {
			flieinfo,err :=ff.FetchByPrimaryKey(uint(v.Fileid))
			CheckErr(err)
			var ferr error
			FileHandler, ferr = os.OpenFile(flieinfo.Path+"/"+flieinfo.Fname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if ferr != nil {
				IotLog.Error("初始化文件失败file error:%s", ferr)
			}
		}
		//开始获取缓存
		var RedisPool *redis.Pool
		if v.Redisid !=0{
			redisinfo ,err  := rs.FetchByPrimaryKey(uint(v.Redisid))
			CheckErr(err)
			RedisPool =getRedisPool(redisinfo)
		}


		//协议写入
		server.SphInfo[v.Hgp] = &server.Sphinfo{ID: v.ID, Hgp: v.Hgp, Htp: v.Htp, Ctrl: v.Ctrl, Cycle: v.Cycle, Sw: v.Sw, Protocol:ptList,Database:dbhandler,Tname:slist.Tname,File:FileHandler,Redis:RedisPool,ItoChan:make(chan  *server.ItoInfo, 10000)}


		//toChan  chan *ItoInfo ,tanme string,db *gorm.DB,file *os.File,redis *redis.Pool

		//打开数据处理中心
		go server.MessageCenter(server.SphInfo[v.Hgp].ItoChan,slist.Tname,dbhandler,FileHandler,RedisPool)

	}
	//开始获取控制器信息
	cpinfos, err := cp.Gets()
	CheckErr(err)
	for  k,v :=range cpinfos{
		//开始循环开机协议
		fmt.Printf("控制器协议：%s：--> %s", k, v)
		 ctList ,err :=cl.GetFromSensorlist(v.Protocolid)
		CheckErr(err)

		server.CphInfo[v.Hgp] = &server.Cphinfo{Hgp:v.Hgp,Htp:v.Htp,Ctrl:v.Ctrl,Sw:v.Sw,Protocol:ctList}

	}


}

/*
	启动数据连接句柄
*/
func initDbHandle() {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBINFO.User, DBINFO.PassWd, DBINFO.Host, DBINFO.Port, DBINFO.DbName)
	fmt.Println("args:", args)
	DbHandler, DbErr = gorm.Open("mysql", args)
	if DbErr != nil {
		fmt.Printf("connection mysql db error:%s", DbErr)
		panic(DbErr)
	}
	DbHandler.SingularTable(true)
	DbHandler.AutoMigrate(&model.Baseinfo{}, &model.Cphinfo{}, &model.Ctrlinfo{}, &model.Ctrllist{}, &model.Fileinfo{}, &model.Redisinfo{}, &model.Sensorinfo{}, &model.Sensorlist{}, &model.Server{}, &model.Sphinfo{})
	DbHandler.LogMode(true)

}

/*
	获取数据库连接
*/

func getDbHandler(baseInfo model.Baseinfo) *gorm.DB {
	var dbhandler *gorm.DB
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", baseInfo.User, baseInfo.Pwd, baseInfo.IP, baseInfo.Port, baseInfo.Dname)
	fmt.Println("args:", args)
	dbhandler, DbErr = gorm.Open("mysql", args)
	if DbErr != nil {
		fmt.Printf("connection mysql db error:%s", DbErr)
		IotLog.Error("connection mysql db error:%s", DbErr)
	}
	dbhandler.LogMode(true)
	dbhandler.SingularTable(true)
	return  dbhandler
}


/*
	获取redis缓存配置
*/
func getRedisPool(redisinfo model.Redisinfo) *redis.Pool{

	pool := &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {

			c, err := redis.Dial("tcp", redisinfo.IP+":"+ redisinfo.Port)
			if  redisinfo.Auth !=""{
				c.Do("AUTH", redisinfo.Auth)
			}
			return c, err
		},
	}
	return pool
}

func CheckErr(err error) {
	if err != nil {

		IotLog.Error("发现问题:%s", err.Error())
	}
}
