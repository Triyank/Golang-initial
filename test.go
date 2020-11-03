package main

import (
	"fmt"
)

var id = 0

type Student struct {
	id                int
	name, class, club string
}

func giveId() int {
	id++
	return id
}

func main() {
	var listOfStudent []Student
index:
	var inputSwitch string
	fmt.Println("Enter Your Choice (1: Add, 2: Edit, 3: Delete, 4: Print, 5: Exit):\n")
	fmt.Scanln(&inputSwitch)

	switch inputSwitch {
	case "1", "ADD", "Add", "add":
		{
			var human Student

			fmt.Println("Enter the following details")
			fmt.Println("Enter the name of the Student:")
			fmt.Scanln(&human.name)

			fmt.Println("enter the class")
			fmt.Scanln(&human.class)

			fmt.Println("Enter the club(Football, Basketball, Badminton)")
			fmt.Scanln(&human.club)

			human.id = giveId()

			listOfStudent = append(listOfStudent, human)
			fmt.Println(listOfStudent)
			goto index
		}

	case "2", "Edit", "EDIT", "edit":
		{
			var editId int
			var editedName, editedClass, editedClub string
			fmt.Println(len(listOfStudent))
			fmt.Println("Enter the Id of the Student:")
			fmt.Scanln(&editId)
			if editId <= len(listOfStudent) {
				editElement := listOfStudent[editId-1]

				var editSwitch string
			editLabel:
				fmt.Println(editElement)
				fmt.Println("Choose the option to edit: \n1. Name \n2. Class\n3. Club \n4. All\n5. Exit")
				fmt.Scanln(&editSwitch)

				switch editSwitch {
				case "1", "Name", "name", "NAME":
					{
						fmt.Println("Enter the new name:")
						fmt.Scanln(&editedName)
						editElement.name = editedName
						goto editLabel
					}
				case "2", "class", "Class", "CLASS":
					{
						fmt.Println("Enter the new class:")
						fmt.Scanln(&editedClass)
						editElement.class = editedClass
						goto editLabel
					}
				case "3", "club", "Club", "CLUB":
					{
						fmt.Println("Enter the new club:")
						fmt.Scanln(&editedClub)
						editElement.club = editedClub
						goto editLabel
					}
				case "4", "all", "All", "ALL":
					{
						fmt.Println("Enter the new name:")
						fmt.Scanln(&editedName)
						editElement.name = editedName

						fmt.Println("Enter the new class:")
						fmt.Scanln(&editedClass)
						editElement.class = editedClass

						fmt.Println("Enter the new club:")
						fmt.Scanln(&editedClub)
						editElement.club = editedClub

					}
				case "5", "exit", "Exit", "EXIT":
					{
						goto index
					}
				default:
					goto index
				}

				listOfStudent[editId-1] = editElement
			} else {
				fmt.Println("invalid input")
			}
			goto index

		}
	case "3", "delete", "Delete", "DELETE":
		{
			var deleteId int

			fmt.Println("Enter the Id of the student")
			fmt.Scanln(&deleteId)

			if deleteId <= len(listOfStudent) {
				listOfStudent = append(listOfStudent[:deleteId-1], listOfStudent[deleteId:]...)
			} else {
				fmt.Println("invalid input")
			}
			goto index
		}
	case "4", "print", "Print", "PRINT":
		{
			var printId int

			fmt.Println("Enter the Id of the student:")
			fmt.Scanln(&printId)

			if printId <= len(listOfStudent) {
				fmt.Println(listOfStudent[printId-1])
			} else {
				fmt.Println("Invalid input")
			}
			goto index
		}
	case "5", "exit", "Exit", "EXIT":
		{
			break
		}
	default:
		{
			fmt.Println(listOfStudent)
			goto index
		}

	}

}
