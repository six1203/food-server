package pagination

func GetPageOffset(pageNum, pageSize int32) int32 {
	return (pageNum - 1) * pageSize
}
