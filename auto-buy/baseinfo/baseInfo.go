package baseinfo

type ids struct {
	Area string
	Gcid int
	Gpid int
}

var UniqueAreaIds []ids //地区ID

func NewAreaIds() {
	UniqueAreaIds = make([]ids, 7)
	UniqueAreaIds[0] = ids{Area: "江苏5连号", Gcid: 5, Gpid: 16}
	UniqueAreaIds[1] = ids{Area: "江西5连号", Gcid: 11, Gpid: 26}
	UniqueAreaIds[2] = ids{Area: "浙江5连号", Gcid: 8, Gpid: 21}
	UniqueAreaIds[3] = ids{Area: "上海5连号", Gcid: 3, Gpid: 11}
	UniqueAreaIds[4] = ids{Area: "福建5连号", Gcid: 12, Gpid: 31}
	UniqueAreaIds[5] = ids{Area: "广东5连号", Gcid: 1, Gpid: 1}
	UniqueAreaIds[6] = ids{Area: "山东5连号", Gcid: 2, Gpid: 6}
}
