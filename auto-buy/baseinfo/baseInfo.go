package baseinfo

type ids struct {
	Area string
	Gcid int
	Gpid int
}

var UniqueAreaIds []ids //地区ID

func NewAreaIds() {
	UniqueAreaIds = make([]ids, 7)
	UniqueAreaIds[0] = ids{Area: "jiangsu", Gcid: 5, Gpid: 16}
	UniqueAreaIds[1] = ids{Area: "jiangxi", Gcid: 11, Gpid: 26}
	UniqueAreaIds[2] = ids{Area: "zhejiang", Gcid: 8, Gpid: 21}
	UniqueAreaIds[3] = ids{Area: "shanghai", Gcid: 3, Gpid: 11}
	UniqueAreaIds[4] = ids{Area: "fujian", Gcid: 12, Gpid: 31}
	UniqueAreaIds[5] = ids{Area: "guangdong", Gcid: 1, Gpid: 1}
	UniqueAreaIds[6] = ids{Area: "shandong", Gcid: 2, Gpid: 6}
}
