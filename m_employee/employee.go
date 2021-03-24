package m_employee

import (
	"fmt"
)

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) LeavesRemaining() {
	fmt.Printf("New  struct  %s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) Employee {
	e := Employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

// 继承
// 原因：go用首字母的大小写来确定是共有的还是私有的，
// 也就是一个变量函数等能不能被其他包引用，
// 小写字母开头的（私有）只能包内使用，不能被其他包使用。

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println(" 结构体切片 Contents of Website\n")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func Test_jicheng_main() { // 包外 公有函数 大小写是包内外 共有私有
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post1.details()

	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		posts: []post{post1, post2, post3},
	}
	w.contents()

}
