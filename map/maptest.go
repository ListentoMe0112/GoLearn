package map_test

import "fmt"

type Stu struct {
	Name    string
	Age     int
	Address string
}

func Quiz() {
	// var mp map[string]bool
	// fmt.Println(mp)
	// mp = make(map[string]bool, 10)
	// mp["lani"] = true
	// fmt.Println(mp)

	// cities := make(map[string]string)
	// cities["no1"] = "beijing"
	// cities["no2"] = "shenzhen"
	// fmt.Println(cities)

	// names := map[string]int{
	// 	"zhangsan": 11,
	// 	"lisi":     12,
	// }
	// names["xieyibo"] = 25
	// fmt.Println(names)

	// stu := make(map[string]string)
	// for i := 0; i < 3; i++ {
	// 	var name string
	// 	var sex string
	// 	fmt.Scanln(&name)
	// 	fmt.Scanln(&sex)
	// 	stu[name] = sex
	// }

	// fmt.Println(stu)
	// val, findRes := stu["xieyibo"]
	// if findRes {
	// 	fmt.Println("research find")
	// 	fmt.Println(val)
	// 	delete(stu, "xieyibo")
	// 	fmt.Println(stu)
	// } else {
	// 	fmt.Println("no research find")
	// }

	// for k, v := range names {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(len(names))

	// monsters := make([]map[string]string, 2)
	// if monsters[0] == nil {
	// 	monsters[0] = make(map[string]string, 2)
	// 	monsters[0]["name"] = "niumowang"
	// 	monsters[0]["age"] = "3000"
	// }

	// if monsters[1] == nil {
	// 	monsters[1] = make(map[string]string, 2)
	// 	monsters[1]["name"] = "yutujing"
	// 	monsters[1]["age"] = "7000"
	// }

	// if monsters[2] == nil {
	// 	monsters[2] = make(map[string]string, 2)
	// 	monsters[2]["name"] = "yutujing"
	// 	monsters[2]["age"] = "7000"
	// }

	// newMonster := map[string]string{
	// 	"name": "newMonster",
	// 	"age":  "321890",
	// 	"test": "test",
	// }

	// monsters = append(monsters, newMonster)

	// fmt.Println(monsters)

	// mpa := make(map[int]int)
	// mpa[100] = 1
	// modify(mpa)
	// fmt.Println(mpa[100])

	// students := make(map[string]Stu, 10)
	// students["192223319"] = Stu{Name: "tom", Age: 18, Address: "beijing"}
	// students["321789322"] = Stu{Name: "mary", Age: 18, Address: "shenzhen"}
	// fmt.Println(students)

	users := make(map[string]map[string]string)
	modifyUser(users, "xieyibo")
	fmt.Println(users)
	modifyUser(users, "xieyibo")
	fmt.Println(users)
}

func modifyUser(users map[string]map[string]string, name string) {
	val, findRes := users[name]
	if findRes {
		fmt.Println("Change pwd")
		val["pwd"] = "88888888"
	} else {
		fmt.Println("Craete new user")
		users[name] = make(map[string]string)
		users[name]["pwd"] = "111111"
		users[name]["nickname"] = "[nickname] " + name
	}
}

// func modify(mp map[int]int) {
// 	mp[100] = 100
// }
