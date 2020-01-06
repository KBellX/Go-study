package libs

// 转换成接口

type ToItf struct {
}

func (self *ToItf) ArrInt2Itf(slice []int) (itf []interface{}) {
	for _, v := range slice {
		itf = append(itf, v)
	}

	return
}

// 看情况，是否有需要加map2Itf
