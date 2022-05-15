package must

func NotFail[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}
	return res
}

func NoError(err error) {
	if err != nil {
		panic(err)
	}
}
