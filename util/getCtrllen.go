package util

func GetctrlLen(len int)int{
	i := 0
	for  len >0{
		len  = len -8
		i++
	}
	return i
}


