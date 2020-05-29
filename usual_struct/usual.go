package usual_struct

type BasePage struct {
	PageSize int    `form:"pageSize"`
	Page     int    `form:"page"`
	Search   string `form:"search"`
	Sort     string `form:"sort"`
	SortDes  string `form:"sortDes"`
}

func CheckParams(bp *BasePage) *BasePage {
	if bp.Page <= 0 {
		bp.Page = 1
	}
	if bp.PageSize <= 0 {
		bp.PageSize = 10
	}
	if bp.Sort == "" {
		bp.Sort = "id"
	}
	if bp.SortDes != "ASC" && bp.SortDes != "DESC" {
		bp.SortDes = "DESC"
	}

	return bp
}
