package oo

import "fmt"

type Monkey struct {
	Name string
}

func (m *Monkey) Climbing() {
	fmt.Println(m.Name, "is born to climbing")
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

type LittleMonkey struct {
	Monkey
}

func (lm *LittleMonkey) Flying() {
	fmt.Println(lm.Name, "is learning to flying")
}

func (lm *LittleMonkey) Swimming() {
	fmt.Println(lm.Name, "is learning to swimming")
}

type English interface {
	OwnEnglish()
}

type Athlete struct {
	Name string
}

type BascketballAthelete struct {
	Athlete
}

type FootballAthelete struct {
	Athlete
}

type Student struct {
	Name string
}

type Undergraduate struct {
	Student
}

type MiddleStudent struct {
	Student
}

func (p *FootballAthelete) OwnEnglish() {
	fmt.Println(p.Name, "own English Skill")
}

func (p *Undergraduate) OwnEnglish() {
	fmt.Println(p.Name, "own English Skill")
}

func (p *Undergraduate) UgOwnSkill() {
	fmt.Println(p.Name, "own [specail] Skill")
}

func LearnEnglish(e English) {
	e.OwnEnglish()
	temp, ok := e.(*Undergraduate)
	if ok {
		temp.UgOwnSkill()
	}
}
