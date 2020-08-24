package util

import (
	"encoding/hex"
	"fmt"
	"strings"
)

/*
	返回CRC校验是否成功
/*/
func CrcOk(data []byte) bool {

	packet_crc := Crc(data[:len(data)-2])
	fmt.Println("CRC:",data)
	fmt.Println("正确答案:",packet_crc)
	fmt.Printf("%X\n",packet_crc)
	fmt.Println("正确答案:",string(packet_crc))

	return string(packet_crc) == string(data[len(data)-2:])
}

/*
	传入 modbus 数据内容，返回 crc校验码
*/
func Crc(data []byte) []byte {
	var crc16 uint16 = 0xffff
	l := len(data)
	for i := 0; i < l; i++ {
		crc16 ^= uint16(data[i])
		for j := 0; j < 8; j++ {
			if crc16&0x0001 > 0 {
				crc16 = (crc16 >> 1) ^ 0xA001
			} else {
				crc16 >>= 1
			}
		}
	}
	packet := make([]byte, 2)
	packet[0] = byte(crc16 & 0xff)
	packet[1] = byte(crc16 >> 8)
	return packet
}

func DecimalToAny(num, n, count int) string {
	num2char := "0123456789abcdef"
	new_num_str := ""
	var remainder int
	var remainder_string string
	for num != 0 {
		remainder = num % n
		remainder_string = string(num2char[remainder])
		new_num_str = remainder_string + new_num_str //注意顺序
		num = num / n
	}
	length := len(new_num_str)
	if length < count { //如果小于8位
		for i := 1; i <= count-length; i++ {
			new_num_str = "0" + new_num_str
		}
	} else {
		return "ERROR"
	}

	return new_num_str
}


func GenSignal(s int)(string){
	//获取数字的最大位移
	offset := GetSignalOffset(s)
	//
	var signalStr = ""
	//var resultMap []int
	for ;offset >=0;offset--{
		if (1 << offset) & s == (1 << offset){
			//可以相除
			signalStr = "1"+signalStr
			fmt.Println(1 << offset)
			//resultMap = append(resultMap,1 << offset)
		}else{
			signalStr = "0" + signalStr
		}
		if offset == 0{
			break
		}
	}
	return signalStr
}


func GetSignalOffset(s int) uint{
	if s==0 {
		return 0
	}
	var offset uint = 0
	for {
		if (1 << offset)>s{
			return offset - 1
		}
		offset++
	}
}


//控制器指令下发
/*
cnum:第几路控制
ctype: 控制状态 1开 0关闭
*/
func GetCtrlStr(cnum int,ctype  int)string{
	ctlSrr :="0105"+DecimalToAny(cnum - 1, 16, 4)
	fmt.Println(strings.ToUpper(ctlSrr))
	if ctype == 0{
		ctlSrr = ctlSrr+"0000"
	}else {
		ctlSrr = ctlSrr+"FF00"
	}
	SetBytes, _ := hex.DecodeString(strings.ToUpper(ctlSrr))
	crc :=fmt.Sprintf("%X",Crc(SetBytes))
	return strings.ToUpper(ctlSrr+crc)
}