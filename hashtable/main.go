package main

type Emp struct {
	Id int
	Name string
	Next *Emp
}

type Emplink struct {
	Head *Emp
}

type HashTable struct {
	Links [7]Emplink
}

func (h *HashTable) InsertTable(emp *Emp) {

}

func main () {

}