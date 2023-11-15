package global

func Init() (err error) {
	err = InitRedis()
	if err != nil {
		return
	}

	return
}
