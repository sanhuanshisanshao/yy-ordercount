package baseinfo

import (
	"strconv"
	"time"
	"yy-ordercount/util"
)

type field struct {
	Area     string
	FieldNum string
	OpenTime time.Time
}

var UniqueFieldInfo [][]field

func NewFieldInfo() {
	UniqueFieldInfo = make([][]field, 7)

	//江苏 一天87期 10分钟一期 8:35-22:05 ？
	t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 33, 0, 0, time.Local) //提前一分钟
	for i := 0; i < 87; i++ {
		endTime := t.Add(time.Duration(i) * 10 * time.Minute)
		fm := util.GetToday() + "-0" + strconv.Itoa(i+1)
		field := field{Area: "jiangsu", FieldNum: fm, OpenTime: endTime}
		UniqueFieldInfo[0] = append(UniqueFieldInfo[0], field)
	}

	//江西 一天84期 10分钟一期 9:10-23:00
	t = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 9, 8, 0, 0, time.Local) //提前一分钟
	for i := 0; i < 84; i++ {
		endTime := t.Add(time.Duration(i) * 10 * time.Minute)
		fm := util.GetToday() + "-0" + strconv.Itoa(i+1)
		field := field{Area: "jiangxi", FieldNum: fm, OpenTime: endTime}
		UniqueFieldInfo[1] = append(UniqueFieldInfo[1], field)
	}

	//浙江 一天85期 10分钟一期 8:30-10:30 180911-01
	t = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 28, 0, 0, time.Local) //提前一分钟
	for i := 0; i < 85; i++ {
		endTime := t.Add(time.Duration(i) * 10 * time.Minute)
		fm := util.GetToday() + "-0" + strconv.Itoa(i+1)
		field := field{Area: "zhejiang", FieldNum: fm, OpenTime: endTime}
		UniqueFieldInfo[2] = append(UniqueFieldInfo[2], field)
	}

	//广东 一天84期 10分钟一期 9:10-23:00 180911-01
	t = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 9, 8, 0, 0, time.Local) //提前一分钟
	for i := 0; i < 84; i++ {
		endTime := t.Add(time.Duration(i) * 10 * time.Minute)
		fm := util.GetToday() + "-0" + strconv.Itoa(i+1)
		field := field{Area: "guangdong", FieldNum: fm, OpenTime: endTime}
		UniqueFieldInfo[3] = append(UniqueFieldInfo[3], field)
	}

	//上海 一天90期 10分钟一期 9:00-23:50 180911-01
	t = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 58, 0, 0, time.Local) //提前一分钟
	for i := 0; i < 90; i++ {
		endTime := t.Add(time.Duration(i) * 10 * time.Minute)
		fm := util.GetToday() + "-0" + strconv.Itoa(i+1)
		field := field{Area: "shanghai", FieldNum: fm, OpenTime: endTime}
		UniqueFieldInfo[4] = append(UniqueFieldInfo[4], field)
	}

	//福建 一天90期 10分钟一期 8:10-23:00
	t = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 8, 0, 0, time.Local) //提前一分钟
	for i := 0; i < 90; i++ {
		endTime := t.Add(time.Duration(i) * 10 * time.Minute)
		fm := util.GetToday() + "-0" + strconv.Itoa(i+1)
		field := field{Area: "fujian", FieldNum: fm, OpenTime: endTime}
		UniqueFieldInfo[5] = append(UniqueFieldInfo[5], field)
	}

	//山东 一天87期 10分钟一期 8:35-22:55
	t = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 33, 0, 0, time.Local) //提前一分钟
	for i := 0; i < 87; i++ {
		endTime := t.Add(time.Duration(i) * 10 * time.Minute)
		fm := util.GetToday() + "-0" + strconv.Itoa(i+1)
		field := field{Area: "shandong", FieldNum: fm, OpenTime: endTime}
		UniqueFieldInfo[6] = append(UniqueFieldInfo[6], field)
	}

}
