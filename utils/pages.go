package utils

func GetPages(limit int64, page int64) (l int64, p int64) {

	if page <= 0 {
		page = 1
	}

	skips := limit * (page - 1)

	return limit, skips
}
