package employee

import "gitlab.com/username/online-service-and-customer-care/entity"

// EmployeeService specifies application employee related services
type EmployeeService interface {
	Employees() ([]entity.Employee, []error)
	Employee(id uint) (*entity.Employee, []error)
	UpdateEmployee(employee *entity.Employee) (*entity.Employee, []error)
	DeleteEmployee(id uint) (*entity.Employee, []error)
	StoreEmployee(employee *entity.Employee) (*entity.Employee, []error)
}

// RoleService speifies application Employee role related services
type RoleService interface {
	Roles() ([]entity.Roles, []error)
	Role(id uint) (*entity.Roles, []error)
	UpdateRole(role *entity.Roles) (*entity.Roles, []error)
	DeleteRole(id uint) (*entity.Roles, []error)
	StoreRole(role *entity.Roles) (*entity.Roles, []error)
}

//JobService specifies application employee job related services
type JobService interface {
	Job() (*entity.Employee_job, []error)
}
