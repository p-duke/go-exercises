package main

// Source https://www.digitalocean.com/community/tutorials/how-to-use-interfaces-in-go

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Stringer interface {
	String() string
}

func Print(s Stringer) {
	fmt.Println(s.String())
}

type Book struct {
	Title string
	Author string
	Pages int
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}

func main() {
	a := Article{Title: "hello world", Author: "Johnny Two Tales"}
	Print(a)

	b := Book{
		Title: "Taco Story",
		Author: "Johnny Two Tales",
		Pages: 400,
	}

	Print(b)
}
