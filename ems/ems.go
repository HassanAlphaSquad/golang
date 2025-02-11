package main

import (
	"fmt"
)

type Employee struct {
	ID          int
	Name        string
	Designation string
	Department  string
	Age         int
}

var employees []Employee

func create_employee() {
	var e Employee
	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&e.ID)
	fmt.Print("Enter Name: ")
	fmt.Scan(&e.Name)
	fmt.Print("Enter Designation: ")
	fmt.Scan(&e.Designation)
	fmt.Print("Enter Department: ")
	fmt.Scan(&e.Department)
	fmt.Print("Enter Age: ")
	fmt.Scan(&e.Age)

	employees = append(employees, e)
	fmt.Print("\nEmployee added successfully!\n")
}

func update_employee() {
	var id, index int
	found := false

	fmt.Print("Enter Employee ID to update: ")
	fmt.Scan(&id)

	for i, emp := range employees {
		if emp.ID == id {
			index = i
			found = true
			break
		}
	}

	if !found {
		fmt.Print("\nEmployee not found!\n")
		return
	}

	fmt.Print("Enter New Name: ")
	fmt.Scan(&employees[index].Name)
	fmt.Print("Enter New Designation: ")
	fmt.Scan(&employees[index].Designation)
	fmt.Print("Enter New Department: ")
	fmt.Scan(&employees[index].Department)
	fmt.Print("Enter New Age: ")
	fmt.Scan(&employees[index].Age)

	fmt.Print("\nEmployee updated successfully!\n\n")
}

func view_employees() {
	if len(employees) == 0 {
		fmt.Print("\nNo employees found!\n")
		return
	}

	fmt.Println("\n--- Employee List ---")
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s, Designation: %s, Department: %s, Age: %d\n",
			emp.ID, emp.Name, emp.Designation, emp.Department, emp.Age)
	}
	fmt.Println()
}

func remove_employee() {
	var id int
	fmt.Print("Enter Employee ID to remove: ")
	fmt.Scan(&id)

	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Print("\nEmployee removed successfully!\n\n")
			return
		}
	}

	fmt.Print("\nEmployee not found!\n")
}

func main() {
	for {
		fmt.Println("1. Create Employee")
		fmt.Println("2. Update Employee")
		fmt.Println("3. View All Employees")
		fmt.Println("4. Remove Employee")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			create_employee()
		case 2:
			update_employee()
		case 3:
			view_employees()
		case 4:
			remove_employee()
		case 5:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Print("\nInvalid choice! Try again.\n")
		}
	}
}
