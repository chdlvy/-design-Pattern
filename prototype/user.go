package prototype

type User struct {
	Name string
	Age  uint8
	M    map[string]string
}

// 克隆
func (u *User) Clone() *User {
	newM := DeepCopy(u.M).(map[string]string)
	return &User{
		Name: u.Name,
		Age:  u.Age,
		M:    newM,
	}
}

// 深拷贝
func DeepCopy(m interface{}) interface{} {
	switch m.(type) {
	case map[string]string:
		newM := make(map[string]string, len(m.(map[string]string)))
		for k, v := range m.(map[string]string) {
			newM[k] = v
		}
		return newM
	default:
		// 报错信息可以根据需要修改
		panic(m)
	}
}
