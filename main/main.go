package main

import (
	_ "GoLearn/array"
	_ "GoLearn/clousure"
	_ "GoLearn/dateTime"
	_ "GoLearn/defer"
	_ "GoLearn/dimensionArray"
	_ "GoLearn/func"
	_ "GoLearn/gostring"
	_ "GoLearn/inter"
	_ "GoLearn/logic"
	_ "GoLearn/loop"
	_ "GoLearn/map"
	"GoLearn/oo"
	_ "GoLearn/operator"
	_ "GoLearn/quiz0405"
	_ "GoLearn/quiz0406"
	_ "GoLearn/recur"
	_ "GoLearn/slice"
	_ "GoLearn/test02"
	_ "GoLearn/wrongcatch"
	"fmt"
	_ "math/rand"
	_ "sort"
)

func init() {
	fmt.Println("Befor Main")
}

func TypeJudge(items ...interface{}) {
	for idx, val := range items {
		switch val.(type) {
		case bool:
			fmt.Printf("%d th arg's type is bool, the value is %v\n", idx, val)
		case oo.Student:
			fmt.Printf("%d th arg's type is Student, the value is %v\n", idx, val)
		case *oo.Student:
			fmt.Printf("%d th arg's type is *Student, the value is %v\n", idx, val)
		default:
			fmt.Printf("%d th arg's type is unkonwn, the value is %v\n", idx, val)
		}

	}
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
	// l := oo.LittleMonkey{}
	// l.Name = "Little Monkey"
	// l.Climbing()
	// var ba oo.BirdAble = &l
	// ba.Flying()
	// var fa oo.FishAble = &l
	// fa.Swimming()

	ug := oo.Undergraduate{}
	ug.Name = "dajkdsa"
	ug1 := oo.Undergraduate{}
	ug1.Name = "ewuqi"
	fa := oo.FootballAthelete{}
	fa.Name = "jihi"

	var les [3]oo.English

	les[0] = &ug
	les[1] = &ug1
	les[2] = &fa
	for _, v := range les {
		v.OwnEnglish()

	}

	var n1 bool = true
	var n2 int32 = 30
	var ns string = "tom"

	var stu oo.Student = oo.Student{"xieyibo"}
	var stu1 *oo.Student = &oo.Student{"zhangxiaogou"}
	TypeJudge(n1, n2, ns, ug, stu, stu1)
}
