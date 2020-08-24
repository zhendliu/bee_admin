package server

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/antlabs/timer"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/panjf2000/gnet"
	"itoServer/model"
	"os"
	"sync"
)

// Server [...]
type Server struct {
	Port int    `gorm:"column:port;type:int"`
	Info string `gorm:"column:info;type:text"`
}

type ItoInfo struct {
	ScanTime string      `json:"scan_time"` //采集时间
	IData    []*SaveType `json:"data"`
	Imei     string      `json:"imei"`
}
type DataHandlerType func([]SaveType, *gorm.DB, *os.File, *redis.Pool)

type Sphinfo struct {
	ID          uint   `gorm:"primary_key;column:id;type:int unsigned;not null"`
	Hgp         string `gorm:"column:hgp;type:text"`
	Htp         string `gorm:"column:htp;type:text"`
	Ctrl        string `gorm:"column:ctrl;type:text"`
	Cycle       int    `gorm:"column:cycle;type:int"`
	Sw          int    `gorm:"column:sw;type:int"`
	Protocol    []*model.Sensorinfo
	Database    *gorm.DB
	Tname       string //存储的表名字
	File        *os.File
	Redis       *redis.Pool
	DataHandler *DataHandlerType
	ItoChan     chan *ItoInfo
}

type Cphinfo struct {
	Hgp      string `gorm:"column:hgp;type:text"`
	Htp      string `gorm:"column:htp;type:text"`
	Ctrl     string `gorm:"column:ctrl;type:text"`
	Cycle    int    `gorm:"column:cycle;type:int"`
	Sw       int    `gorm:"column:sw;type:int"`
	Protocol []*model.Ctrlinfo
}

type ConnType struct {
	Conn     gnet.Conn
	ConnTime string
	UpTime   string
	Imei     string
	Register bool            //是否解析过注册包头
	Etype    int             //设备类型 1.传感器，2 控制器
	Rst      string          //注册的包头，可以直接识别协议
	Protocol  *Sphinfo        //连接协议
	CProtocol *Cphinfo
	Task     timer.TimeNoder //时间轮处理句柄
	CStatus  string
}

type SaveType struct {
	Name  string
	Value string
}

//连接描述
type IotConnection struct {
	Conn map[string]*ConnType
	Mx   sync.RWMutex
}

var (
	IC          IotConnection
	IotServer   Server
	SphInfo     map[string]*Sphinfo
	CphInfo     map[string]*Cphinfo
	TM          = timer.NewTimer()
	SetBytes, _ = hex.DecodeString("01030000001445C5")
)

func init() {
	SphInfo = make(map[string]*Sphinfo)
	CphInfo = make(map[string]*Cphinfo)
}

func MessageCenter(itoChan chan *ItoInfo, tname string, db *gorm.DB, file *os.File, redis *redis.Pool) {
	fmt.Println("开启文件处理。。。")

	// INSERT INTO `iot`.`iot` (`imei`, `air_temperature`, `air_humidity`) VALUES ('123', '123', '245');

	for {
		info := <-itoChan
		insert_sql := "INSERT INTO " + tname + " ( "
		val_sql := ") VALUES ("

		info.IData = append(info.IData, &SaveType{Name: "creative", Value: info.ScanTime}, &SaveType{Name: "imei", Value: info.Imei})
		/*
			ScanTime string   `json:"scan_time"` //采集时间
			IData    []*SaveType `json:"data"`
			Imei     string     `json:"imei"`
		*/
		fmt.Println("开始写入数据")
		fmt.Println("info.Imei:", info.Imei)
		fmt.Println("info.ScanTime:", info.ScanTime)
		fmt.Println("info.IData:", info.IData)
		//数据库写入
		for _, v := range info.IData {

			//
			insert_sql = insert_sql + v.Name + ","
			val_sql = val_sql + v.Value + ","
			fmt.Println("传入的数据key:", v.Name)
			fmt.Println("传入的数据val:", v.Value)
		}

		//开始处理末尾sql数据
		insert_sql = insert_sql[:len(insert_sql)-1]
		val_sql = val_sql[:len(val_sql)-1]

		val_sql = val_sql + " );"

		fmt.Println("val_sql:", val_sql)

		fmt.Println("insert_sql:", insert_sql)

		insert_sql = insert_sql + val_sql

		fmt.Println("f_insert_sql:", insert_sql)
		err := db.Exec(insert_sql).Error

		if err != nil {
			fmt.Println("数据库存储数据出错，请确认状态！！！！！！:", err)
		}

		//开始文件存储
		if file != nil {
			saveData, _ := json.Marshal(info)
			// 文件存储
			file.Write(saveData)
			file.Write([]byte("\n"))

		}
	}

}
