package server

import "C"
import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/panjf2000/gnet"
	"itoServer/model"
	"itoServer/util"
	"log"
	"strconv"
	"strings"
	"time"
)

type iotServer struct {
	*gnet.EventServer
}

func (iot *iotServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Iot server is listening on %s (multi-cores: %t, loops: %d)\n", srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	//启动写入数据的环境
	//go WriteDB()

	//启动时间轮
	go TimeStart()

	return
}

/*
 消息处理函数
*/
func (iot *iotServer) React(data []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	//获取访问者IP端口
	ipstr := c.RemoteAddr().String()
	fmt.Println("connected from", ipstr)
	reciveStr := string(data)
	two := fmt.Sprintf("%x", data)
	lenght := len(data)

	//判断这个连接过来时信息

	onnType, ok := IC.Conn[c.RemoteAddr().String()]
	if ok{
		//此处处理协议注册
		if onnType.Register == false{

			//循环匹配传感注册包头
			for k,v :=range  SphInfo{
				fmt.Printf("k:%s,v:%s\n",k,v.Hgp)
				fmt.Println("接收的内容为:", reciveStr)
				if len(reciveStr)>=len(k) &&  reciveStr[:len(k)] == k{
					fmt.Printf("命中传感器这个协议\n")
					//命中协议
					//设置imei,标识注册，标识协议
					IC.Mx.Lock()
					if conn, ok := IC.Conn[c.RemoteAddr().String()]; ok {
						conn.ConnTime = time.Now().Format("2006-01-02 15:04:05")
						conn.Imei = reciveStr[len(k):]
						conn.Register = true
						conn.Etype = 1
						conn.Rst = k
						conn.Protocol = v
					}
					IC.Mx.Unlock()
					//处理定时问题,关闭立即停止时间轮
					if v.Sw != 1{
						onnType.Task.Stop()
					}else if v.Cycle !=300 || v.Ctrl !="01030000001445C5"{
						//不同的设计轮，或者指令下发功能
						onnType.Task.Stop()
						setTimerOnConn(c,v.Cycle, v.Ctrl)
					}
				}
			}

			for  k,v  :=range CphInfo{
				fmt.Printf("k:%s,v:%s\n",k,v.Hgp)
				fmt.Println("接收的内容为:", reciveStr)
				if len(reciveStr)>=len(k) &&  reciveStr[:len(k)] == k{
					fmt.Printf("命中控制器协议\n")
					IC.Mx.Lock()
					if conn, ok := IC.Conn[c.RemoteAddr().String()]; ok {
						conn.ConnTime = time.Now().Format("2006-01-02 15:04:05")
						conn.Imei = reciveStr[len(k):]
						conn.Register = true
						conn.Etype = 2
						conn.Rst = k
						conn.CProtocol = v
					}
					IC.Mx.Unlock()
					//处理定时问题,关闭立即停止时间轮
					if v.Sw != 1{
						onnType.Task.Stop()
					}else if v.Cycle !=300 || v.Ctrl !="01030000001445C5"{
						//不同的设计轮，或者指令下发功能
						onnType.Task.Stop()
						setTimerOnConn(c,v.Cycle, v.Ctrl)
					}

				}

			}


		}else{
			//此处理心跳和协议内容
			if onnType.Etype == 1{
				//此处处理传感器控制信息
				//1.心跳
				ProLen :=len(onnType.Protocol.Protocol)
				k :=onnType.Protocol.Htp
				if len(reciveStr)>=len(k) &&  reciveStr[:len(k)] == k{
					//识别到对应的心跳，更新心跳时间
					fmt.Println("识别到心跳包，更新心跳包时间")
					IC.Mx.Lock()
					if conn, ok := IC.Conn[c.RemoteAddr().String()]; ok {
						conn.UpTime = time.Now().Format("2006-01-02 15:04:05")
					}
					IC.Mx.Unlock()

				}else if strings.ToUpper(two[4:6]) == strings.ToUpper(util.DecimalToAny(ProLen * 2, 16, 2)) && util.CrcOk(data[:ProLen * 2 + 5]){
					fmt.Println("识别到正确的协议匹配")

					alsProtocol(onnType.Protocol.Protocol,two,c,onnType.Protocol.ItoChan)
				}
				//开始处理处理协议
			}
			//此处处理控制器信息
			if  onnType.Etype == 2{
				ProLen :=len(onnType.CProtocol.Protocol)
				clen :=util.GetctrlLen(ProLen)
				k :=onnType.CProtocol.Htp
				if len(reciveStr)>=len(k) &&  reciveStr[:len(k)] == k{
					fmt.Println("识别到控制器注册心跳包，更新心跳包时间")
					IC.Mx.Lock()
					if conn, ok := IC.Conn[c.RemoteAddr().String()]; ok {
						conn.UpTime = time.Now().Format("2006-01-02 15:04:05")
					}
					IC.Mx.Unlock()
				}else if strings.ToUpper(two[4:6]) == strings.ToUpper(util.DecimalToAny(clen, 16, 2)) && util.CrcOk(data){
					fmt.Println("识别控制器的协议匹配")

					alsCProtocol(onnType.CProtocol.Protocol,two,c,clen)
				}
				/*fmt.Println( strings.ToUpper(two[4:6]) == strings.ToUpper(util.DecimalToAny(clen, 16, 2)))
				fmt.Println( strings.ToUpper(two[4:6]))
				fmt.Println( strings.ToUpper(util.DecimalToAny(clen, 16, 2)))
				fmt.Println(util.CrcOk(data))*/

			}
		}
	}

	//解析回传设备信息
	//reciveStr := string(data)
	//fmt.Println("接收的内容为:", reciveStr)
	hexdata := fmt.Sprintf("%x", data)
	fmt.Println("hexdata:", hexdata)
	fmt.Println("data:", data)
	fmt.Println("接受的数据长度：", lenght)
	fmt.Println("reciveStr:", reciveStr)


	//测试使用，后期去掉
	IC.Mx.RLock()
	if conn, ok := IC.Conn[c.RemoteAddr().String()]; ok {
		fmt.Println("远端IP:", conn.Conn.RemoteAddr())
		fmt.Println("连接建立时间:", conn.ConnTime)
		fmt.Println("连接心跳时间:", conn.UpTime)
		fmt.Println("IMEI:", conn.Imei)
		fmt.Println("当前在线设备数量:", len(IC.Conn))
	}
	IC.Mx.RUnlock()

	return

}

func alsProtocol(sensorinfos []*model.Sensorinfo, protocolStr string,c gnet.Conn,itoChan chan *ItoInfo) []*SaveType{
	saLice := []*SaveType{}
	fmt.Println("protocolStr:",protocolStr)
	//开始解析协议
	start_index := 6
	for _,v :=range  sensorinfos{
		fmt.Println("start_index:",start_index)
		val, _ := strconv.ParseInt(protocolStr[start_index:start_index+4], 16, 64)
		fmt.Println("protocolStr[start_index:start_index+4]:",protocolStr[start_index:start_index+4])
		fval := float64(val)
		if v.Maxv !="" {
			val, err := strconv.ParseInt(v.Maxv, 16, 64)
			if err !=nil{
				if  fval > float64(val){
					fval = util.GetNegative(protocolStr[start_index:start_index+4])
				}
			}
		}
		if  v.Transform !=""{
			tr ,_ :=strconv.Atoi(strings.Replace( v.Transform[1:], " ", "", -1))
			if v.Transform[:1] =="/"{
				fval = fval / float64(tr)
			}
			if v.Transform[:1] =="*"{
				fval = fval * float64(tr)
			}
		}
		saLice = append(saLice, &SaveType{Name:v.Savename,Value:strconv.FormatFloat(fval, 'f', -1, 64)})
		start_index = start_index + 4
	}
	fmt.Println("开始打印收到协议解析的内容" )
	for k,v := range saLice{
		fmt.Printf("索引:%d,内容:%s\n",k,v)
	}
	var IMEI string
	IC.Mx.RLock()
	if conn, ok := IC.Conn[c.RemoteAddr().String()]; ok {
		IMEI =  conn.Imei
	}
	IC.Mx.RUnlock()
	//开始发送数据到数据中心
	ito := ItoInfo{ScanTime: "'"+time.Now().Format("2006-01-02 15:04:05")+"'", Imei: IMEI,IData:saLice}
	itoChan<-&ito
	return saLice
}




func alsCProtocol(ctrlinfos []*model.Ctrlinfo, protocolStr string,c gnet.Conn,clen int) {

	fmt.Println("protocolStr:",protocolStr)
	//开始解析协议
	start_index := 6
	val, _ := strconv.ParseInt(protocolStr[start_index:start_index+clen *2 ], 16, 64)

	singalStr :=util.GenSignal(int(val))
	for i:=len(singalStr) ;i<len(ctrlinfos);i++{
		singalStr = singalStr +"0"
	}
	IC.Mx.Lock()
	if conn, ok := IC.Conn[c.RemoteAddr().String()]; ok {
		conn.CStatus = singalStr
	}
	IC.Mx.Unlock()
}


/*
	连接打开时
*/
func (iot *iotServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	//加入时间片，注意在识别到具体协议时，再次修正
	task := TM.ScheduleFunc(300*time.Second, func() {
		fmt.Printf("开始下发指令:[%s]\n", c.RemoteAddr())
		if c == nil {
			return
		}
		err := c.AsyncWrite(SetBytes)
		fmt.Printf("时间[%s]:[%s]:下发指令:%x\n", time.Now().Format("2006-01-02 15:04:05"), c.RemoteAddr(), SetBytes)
		if err != nil {
			c.Close()
		}
	})
	//接入连接管理器
	IC.Mx.Lock()
	//更新了连接和连接建立时间
	IC.Conn[c.RemoteAddr().String()] = &ConnType{Conn: c, ConnTime: time.Now().Format("2006-01-02 15:04:05"), Task: task}
	IC.Mx.Unlock()
	return
}

/*
	连接退出时
*/
func (iot *iotServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	fmt.Println("退出:", c.RemoteAddr().String())
	//时间片的退出
	IC.Mx.RLock()
	if conn, ok := IC.Conn[c.RemoteAddr().String()]; ok {
		conn.Task.Stop()
		fmt.Println("一个调度任务退出了:", conn.Task)
	}
	IC.Mx.RUnlock()
	//连接队列退出
	IC.Mx.Lock()
	delete(IC.Conn, c.RemoteAddr().String())
	IC.Mx.Unlock()
	return
}


/*
	连接打开时
*/
func  setTimerOnConn(c gnet.Conn, cycle int, ctrl string) (out []byte, action gnet.Action) {
	//加入时间片，注意在识别到具体协议时，再次修正
	fmt.Printf("开始修正这个时间片\n")
	SetByte, _ := hex.DecodeString(ctrl)
			task := TM.ScheduleFunc(time.Duration(cycle) * time.Second, func() {
				if c == nil {
					return
		}
		err := c.AsyncWrite(SetByte)
		fmt.Printf("时间[%s]:[%s]:下发指令:%x\n", time.Now().Format("2006-01-02 15:04:05"), c.RemoteAddr(), SetByte)
		if err != nil {
			c.Close()
		}
	})

	//接入连接管理器
	IC.Mx.Lock()
	//更新了连接和连接建立时间
	IC.Conn[c.RemoteAddr().String()].Task=task
	IC.Mx.Unlock()
	fmt.Printf("结束修正这个时间片\n")
	return
}


/*
	时间轮模块启动
*/
func TimeStart() {
	TM.Run()
}

/*
	服务软件开启
*/
func IotStart() {
	var port int
	var multicore bool

	IC.Conn = make(map[string]*ConnType)

	flag.IntVar(&port, "port", IotServer.Port, "--port 9000")
	flag.BoolVar(&multicore, "multicore", true, "--multicore true")
	flag.Parse()
	iot := new(iotServer)
	log.Fatal(gnet.Serve(iot, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(multicore)))
}
