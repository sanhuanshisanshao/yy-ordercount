package baseinfo

type ids struct {
	Gcid int
	Gpid int
}

var UniqueAreaIds map[string]ids //地区ID

func NewAreaIds() {
	UniqueAreaIds = make(map[string]ids)
	UniqueAreaIds["guangdong"] = ids{Gcid: 1, Gpid: 1}
	UniqueAreaIds["shanghai"] = ids{Gcid: 3, Gpid: 11}
	UniqueAreaIds["zhejiang"] = ids{Gcid: 8, Gpid: 21}
	UniqueAreaIds["fujian"] = ids{Gcid: 12, Gpid: 31}
	UniqueAreaIds["shandong"] = ids{Gcid: 2, Gpid: 6}
	UniqueAreaIds["jiangsu"] = ids{Gcid: 5, Gpid: 16}
	UniqueAreaIds["jiangxi"] = ids{Gcid: 11, Gpid: 26}
}
