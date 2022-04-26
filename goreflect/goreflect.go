package goreflect

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	// 通过反射获取传入变量的type, kind, val

	// 获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	// 获取reflect.Val reflect.Value 和 int并不是同类型
	rVal := reflect.ValueOf(b)

	// rVal += 2  invalid operation: rVal += 2 (mismatched types reflect.Value and int)
	fmt.Println(rVal)
	fmt.Printf("%v\n", rVal.Kind())

	// 修改b的值
	rVal.Elem().SetInt(101)
}

type Student struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	rtype := reflect.TypeOf(b)
	fmt.Println(rtype) // goreflect.Student
	rval := reflect.ValueOf(b)
	iV := rval.Interface()
	fmt.Printf("iV = %v, iV = %T\n", iV, iV) // iV = {tom 20}, iV = goreflect.Student
	fmt.Println(rtype.Kind(), rval.Kind())   // struct struct
	stu, ok := iV.(Student)
	if ok {
		fmt.Println(stu.Name)
	}
}

func quiz01(v interface{}) {
	rVal := reflect.ValueOf(v)
	rType := rVal.Type()
	rKind := rVal.Kind()
	fmt.Println(rType, rKind)

	iv := rVal.Interface()
	var num float64 = iv.(float64)

	fmt.Println(num)
}

func quiz02(s interface{}) {
	rVal := reflect.ValueOf(s)
	rType := rVal.Type()
	rKind := rVal.Kind()
	fmt.Println(rType, rKind)
	rVal.Elem().SetString("mary")
	fmt.Println("In quiz02: ", rVal.Elem())
}

type Monster struct {
	Name   string `json:"name"`
	Age    int    `json:"monster_age"`
	Score  float32
	Gender string
}

func (m Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(m)
	fmt.Println("---end---")
}

func (m Monster) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

func (m Monster) Set(name string, age int, score float32, gender string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Gender = gender
}

func quiz03(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	// kd := val.Kind()
	// if kd != reflect.Struct {
	// 	fmt.Println("expect struct")
	// 	return
	// }
	val.Method(1).Call(nil)
	num := val.Elem().NumField()
	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: val = %v\n", i, val.Elem().Field(i))
		tagVal := typ.Elem().Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag = %v\n", i, tagVal)
		} else {
			fmt.Printf("Field %d: tag = %v\n", i, typ.Elem().Field(i).Name)
		}
	}

	val.Elem().Field(0).SetString("baixiangjing")

	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	// method is sorted by ASCII

	// val.Method(0).Call(2, 3) error: want ([]reflect.Value)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println(res[0].Int())
}

func funcAdaptor() {
	call1 := func(v1 int, v2 int) {
		// t.Log(v1, v2)
		fmt.Println(v1, v2)
	}

	call2 := func(v1 int, v2 int, s string) {
		// t.Log(v1, v2, s)
		fmt.Println(v1, v2, s)
	}

	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)

	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(inValue)
	}

	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")
}

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Println(name, " has finished sub, ", c.Num1, " - ", c.Num2, " = ", c.Num1-c.Num2)

}

func quiz04(c interface{}) {
	val := reflect.ValueOf(c)
	typ := reflect.TypeOf(c)
	num := val.Elem().NumField()
	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++ {
		var a int
		fmt.Printf("Please enter Num%d", i)
		fmt.Scanln(&a)
		fmt.Printf("Field %d: val = %v\n", i, val.Elem().Field(i))
		val.Elem().Field(i).SetInt(int64(a))
		fmt.Printf("Field %d: tag = %v\n", i, typ.Elem().Field(i).Name)
	}
	var params []reflect.Value
	params = append(params, reflect.ValueOf("Tom"))
	val.Method(0).Call(params)
}

func Test01() {
	var num int = 100
	// reflectTest01(num)
	reflectTest01(&num)
	fmt.Println(num)
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)

	fmt.Println("************************test01************************")
	var v float64 = 1.2
	quiz01(v)
	fmt.Println("************************test02************************")
	s := "tom"
	quiz02(&s)
	fmt.Println("In main: s = ", s)
	fmt.Println("************************test03************************")
	var a Monster = Monster{
		Name:  "Huangshulangjing",
		Age:   400,
		Score: 30.8,
	}
	quiz03(&a)
	a.Print()
	fmt.Println("************************test04************************")
	funcAdaptor()
	fmt.Println("************************test05************************")
	c := Cal{}
	quiz04(&c)
}
