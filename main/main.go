package main

import (
	_ "GoLearn/array"
	_ "GoLearn/clousure"
	_ "GoLearn/dateTime"
	_ "GoLearn/defer"
	_ "GoLearn/dimensionArray"
	_ "GoLearn/func"
	_ "GoLearn/gostring"
	"GoLearn/inter"
	_ "GoLearn/logic"
	_ "GoLearn/loop"
	_ "GoLearn/map"
	_ "GoLearn/oo"
	_ "GoLearn/operator"
	_ "GoLearn/quiz0405"
	_ "GoLearn/quiz0406"
	_ "GoLearn/recur"
	_ "GoLearn/slice"
	_ "GoLearn/test02"
	_ "GoLearn/wrongcatch"
	"fmt"
	"math/rand"
	"sort"
)

func init() {
	fmt.Println("Befor Main")
}

func main() {
	// acc := clousure.Adapter()
	// fmt.Println(acc(2))
	// fmt.Println(acc(2))
	// fmt.Println(acc(2))

	// jpg := clousure.MakeSuffix(".jpg")
	// fmt.Println(jpg("hello"))
	// fmt.Println(jpg("test.jpg"))
	// defer_test.Test01()
	// fmt.Println(quiz0405.Name)
	// quiz0405.Test01()
	// quiz0405.Test02()
	// quiz0405.Test01()
	// var x int
	// fmt.Scanln(&x)
	// quiz0405.Test03(x)
	// var y int
	// fmt.Scanln(&y)
	// quiz0405.Test04(y)
	// gostring.LenTest()
	// gostring.RuneTest()
	// gostring.AtoiTest()
	// gostring.ItoaTest()
	// gostring.ByteTest()
	// gostring.IntTest()
	// gostring.ContainsTest()
	// gostring.CountTest()
	// gostring.EqualFoldTest()
	// gostring.IndexTest()
	// gostring.LastIndexTest()
	// gostring.ReplaceTest()
	// gostring.SplitTest()
	// gostring.ToLowerTest()
	// gostring.ToUpperTest()
	// gostring.TrimSpaceTest()
	// gostring.TrimTest()
	// gostring.HasPrefixTest()
	// HasSuffix
	// dateTime.Test01()
	// dateTime.Test02()
	// dateTime.Test03()
	// wrongcatch.Test01()
	// wrongcatch.Test02()
	// quiz0406.Test01()
	// quiz0406.Test02()
	// array.Test05()
	// dimensionArray.Test02()
	// map_test.Quiz()
	// p := oo.CreatePerson("xieyibo", 25, 50.0)
	// p.SetName("zhangxiaoyii")
	// fmt.Println(*p)

	// s := oo.CreateStudent("xieyibo", 25, 50.0, "192110100378")
	// s.SetStuId("3271893")
	// s.SetName("zhangxiaoyi")
	// s.Print()
	var stus inter.Students
	for i := 0; i < 5; i++ {
		stu := inter.Student{
			Name:  fmt.Sprintf("Hero~%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: float64(rand.Intn(100)) * 0.337,
		}
		stus = append(stus, stu)
	}

	fmt.Println(stus)

	sort.Sort(stus)

	fmt.Println(stus)
}
