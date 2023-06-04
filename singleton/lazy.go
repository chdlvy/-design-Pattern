package singleton

type LazyObj struct {
}

var LazyInstance *LazyObj

func getLazyInstance() *LazyObj {
	if LazyInstance == nil {
		return nil
	}
	return &LazyObj{}
}
