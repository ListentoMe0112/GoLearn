package inter

type Student struct {
	Name  string
	Age   int
	Score float64
}

type Students []Student

func (stus Students) Len() int {
	return len(stus)
}

func (stus Students) Less(i, j int) bool {
	return stus[i].Score < stus[j].Score
}

func (stus Students) Swap(i, j int) {
	stus[i], stus[j] = stus[j], stus[i]
}
