package util

type Paging struct {
	Page   int // 第几页
	Limit  int // 数量
	Offset int // 计算出的偏移值
}

// 方便从URLParam中 解析分页参数
func ParsePaging(ctx *ApiContext) Paging {
	page := ctx.URLParamIntDefault("page", 1)
	limit := ctx.URLParamIntDefault("per_page", 20)
	return Paging{
		page,
		limit,
		(page - 1) * limit,
	}
}
