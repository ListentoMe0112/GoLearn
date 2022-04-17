package oo

import "fmt"

// type Student struct {
// 	name string
// 	gender string
// 	age int
// 	id int
// 	score float64
// }

// func (stu *Student) say () string {
// 	str := fmt.Sprintf("name: %s, gender: %v, age: %v, id: %v, score: %v\n", stu.name, stu.gender, stu.age, stu.id, stu.score)
// 	return str
// }

// type Box struct {
// 	length float64
// 	width float64
// 	height float64
// }

// func (b *Box) getVolume () float64{
// 	return b.length * b.width * b.height
// }

type person struct {
	name   string
	age    int
	salary float64
}

func CreatePerson(name string, age int, salary float64) *person {
	return &person{
		name:   name,
		age:    age,
		salary: salary,
	}
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) SetName(n string) {
	p.name = n
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetAge(a int) {
	p.age = a
}

func (p *person) GetSalary() float64 {
	return p.salary
}

func (p *person) SetSalary(s float64) {
	p.salary = s
}

type student struct {
	*person
	stuid string
}

func (s *student) SetStuId(id string) {
	s.stuid = id
}

func (s *student) GetStuId() string {
	return s.stuid
}

func (stu *student) SetName(s string) {
	stu.name = "[Student]: " + s
}

func (stu *student) Print() {
	fmt.Println(stu.stuid)
	fmt.Println(*stu.person)
}

func CreateStudent(name string, age int, salary float64, userid string) *student {
	return &student{
		&person{
			name:   name,
			age:    age,
			salary: salary,
		},
		userid,
	}
}
