package main

// Bad structure
type BadEmployee struct {
	name     string
	lastName string
	age      int
	salary   int
}

func (e BadEmployee) GetFullName() string {
	return e.name + " " + e.lastName
}

func (e BadEmployee) PrintSalaryReport() int {
	return e.salary
}

// Correct structure
type Employee struct {
	name     string
	lastName string
	age      int
}

type Salary struct {
	value int
}

func (e Employee) GetFullName() string {
	return e.name + " " + e.lastName
}

func (e Salary) PrintSalaryReport(employee Employee) int {
	// Business logic
	return e.value
}

//Additional
type EmployeeDB struct {
	name     string
	lastName string
	age      int
}

type SalaryDB struct {
	value int
}

func (e EmployeeDB) GetFullName() string {
	// DB query
	dbvalue := e.name + " " + e.lastName
	return dbvalue
}

func (e SalaryDB) PersistPrintSalaryReport(employee Employee) int {
	// Business logic
	// DB query
	dbvalue := e.value
	return dbvalue
}
