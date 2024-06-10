package single_responsibility

import "fmt"

// employee struct has two responsibilities
// 1. get employee da and get employee details
// 2. save the employee data into a database
type Employee struct {
    Id int
    Name string
}

func (e *Employee) GetDetails() string {
    return fmt.Sprintf("Id: %d - Name: %s", e.Id, e.Name)
}

func (e *Employee) SaveEmployee() {
    //Code to saze employee details into a database
}

// Can Refactor like this....
// EmployeeSingleRes struct only has one responsibility managing employee data
// if save will use employeedb , easy to read test and maintain :)
type EmployeeSingleRes struct {
    Id int
    Name string
}

func (e * EmployeeSingleRes) GetDetails() string {
    return fmt.Sprintf("Id: %d - Name: %s", e.Id, e.Name)
}

type EmployeeDb struct {
    // Database related fields
}
func (db *EmployeeDb) SaveEmployee(e *EmployeeSingleRes) {
    // Code to save employee details into a database
}

// Applying SRP to func example
func ParseAndSave(filename string, db *EmployeeDb) {
    // 1. Read the file
    // 2. Parse the file content to get employee details
    // 3. Save the parsed data into the database
}

// it will look like this SRP
func ReadFile(filename string) []byte {
    // Code to read the file.
}

func ParseEmployee(data byte[]) *EmployeeSingleRes {
    // Code to parse the data to get employee details.
}

func (db *EmployeeDb) SaveEmployee2(e *EmployeeSingleRes) {
    // Cade to save employee details into database
}


