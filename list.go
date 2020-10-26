package main

import "fmt"

type Node struct {
	key  interface{}
	next *Node
}

type List struct {
	length int
	head   *Node
}

//Inserting
func (item *List) Add(key interface{}) {

	x := &Node{}
	x.key = key
	if item.length == 0 {
		item.head = x
		item.length++
		return
	}

	ptr := item.head

	for i := 0; i < item.length; i++ {
		if ptr.next == nil {
			ptr.next = x
			item.length++
			return
		}
		ptr = ptr.next
	}
}

//PrintAll
func (item *List) ShowAll() {
	x := item.head
	for x != nil {
		fmt.Print(x.key, "-->")
		x = x.next

	}
	fmt.Println("\n")
}

//Print by position
func (item *List) PrintbyPosition(position int) {
	x := item.head

	for i := 0; i < item.length; i++ {
		if position-1 == i {
			fmt.Println(x.key)
		}
		x = x.next
	}

}

//Edit by position
func (item *List) EditByPosition(position int) {
	x := item.head
	var key string
	fmt.Println("Enter the value you want to add in node: ")
	fmt.Scanln(&key)
	for i := 0; i < item.length; i++ {
		if position-1 == i {
			x.key = key
		}
		x = x.next
	}

}

//removing last element
func (item *List) remove() {
	x := item.head
	for i := 0; i < item.length; i++ {
		if x.next.next == nil {
			x.next = nil
			item.length--
			return
		}
		x = x.next
	}
}

func main() {

	item := List{}

	//adding nodes to list
	item.Add("a")
	item.Add(2)
	item.Add(5)
	item.Add(53)
	item.Add(74)
	item.Add(89)
	item.Add(987)
	item.Add(true)
	item.ShowAll()

	//print by position
	item.PrintbyPosition(1)

	//edit by position
	item.EditByPosition(2)

	//showing all nodes
	item.ShowAll()
	item.remove()
	item.remove()
	item.ShowAll()

}
