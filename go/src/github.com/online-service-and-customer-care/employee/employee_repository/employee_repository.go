package employee_repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"gitlab.com/username/online-service-and-customer-care/employee"
	"gitlab.com/username/online-service-and-customer-care/entity"
)

//EmployeeGromRepo implments emplpyee repository
type EmployeeGormRepo struct {
	dbconn *gorm.DB
}

//NewEmployeeGromRepo isa construtor for EmplpoyeeGromRepo+
func NewEmployeeGormRepo(db *gorm.DB) employee.EmployeeRepository {
	return &EmployeeGormRepo{dbconn: db}
}

// Employees returns all Employees stored in the database
func (emplRepo *EmployeeGormRepo) Employees() ([]entity.Employee, []error) {
	empls := []entity.Employee{}
	fmt.Println(empls)
	errs := emplRepo.dbconn.Find(&empls).GetErrors()
	if len(errs) > 0 {
		fmt.Println("errrrrrrrrrrr")
		return nil, errs
	}
	return empls, errs
}

// Employee retrieves an employee from the database by its id
func (emplRepo *EmployeeGormRepo) Employee(id uint) (*entity.Employee, []error) {
	empl := entity.Employee{}
	errs := emplRepo.dbconn.First(&empl, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &empl, errs
}

// UpdateEmployee updates a given employee in the database
func (emplRepo *EmployeeGormRepo) UpdateEmployee(employee *entity.Employee) (*entity.Employee, []error) {
	empl := employee
	errs := emplRepo.dbconn.Save(empl).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}

// DeleteEmployee deletes a given employee from the database
func (emplRepo *EmployeeGormRepo) DeleteEmployee(id uint) (*entity.Employee, []error) {
	empl, errs := emplRepo.Employee(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = emplRepo.dbconn.Delete(empl, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}

// StoreEmployee stores an employee in the database
func (emplRepo *EmployeeGormRepo) StoreEmployee(emp *entity.Employee) (*entity.Employee, []error) {
	employee := emp
	errs := emplRepo.dbconn.Create(employee).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return employee, errs
}
