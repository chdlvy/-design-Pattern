package singleton

type Hungry struct {
}

var hungryInstance *Hungry

func init() {
	hungryInstance = &Hungry{}
}

func getHungryInstance() *Hungry {
	return hungryInstance
}
