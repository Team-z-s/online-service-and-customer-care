package employee_service

import (
	"gitlab.com/username/online-service-and-customer-care/employee"
	"gitlab.com/username/online-service-and-customer-care/entity"
)

//EmployeeserviceGorm implments service of employee
type EmployeeServiceGorm struct {
	emplRepo employee.EmployeeRepository
}

//NewEmployeeServiceGorm construstor
func NewEmployeeServiceGorm(emplRepo employee.EmployeeRepository) employee.EmployeeService {
	return &EmployeeServiceGorm{emplRepo: emplRepo}
}

//Employees return all employees on the database
func (es *EmployeeServiceGorm) Employees() ([]entity.Employee, []error) {
	empls, errs := es.emplRepo.Employees()
	if len(errs) > 0 {
		return nil, errs
	}
	return empls, errs
}

//Employee return a single employee
func (es *EmployeeServiceGorm) Employee(id uint) (*entity.Employee, []error) {
	empl, errs := es.emplRepo.Employee(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}

//UpdateEmployee updates
func (es *EmployeeServiceGorm) UpdateEmployee(employee *entity.Employee) (*entity.Employee, []error) {
	empl, errs := es.emplRepo.UpdateEmployee(employee)
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}

//DeleteEmployee delete
func (es *EmployeeServiceGorm) DeleteEmployee(id uint) (*entity.Employee, []error) {
	empl, errs := es.emplRepo.DeleteEmployee(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}

//StoreEmployee store
func (es *EmployeeServiceGorm) StoreEmployee(employee *entity.Employee) (*entity.Employee, []error) {
	empl, errs := es.emplRepo.StoreEmployee(employee)
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}
