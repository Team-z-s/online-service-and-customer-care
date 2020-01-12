package employee

import "gitlab.com/username/online-service-and-customer-care/entity"

// EmployeeRepository specifies application employee related database operations
type EmployeeRepository interface {
	Employees() ([]entity.Employee, []error)
	Employee(id uint) (*entity.Employee, []error)
	UpdateEmployee(user *entity.Employee) (*entity.Employee, []error)
	DeleteEmployee(id uint) (*entity.Employee, []error)
	StoreEmployee(user *entity.Employee) (*entity.Employee, []error)
}

// RoleRepository speifies application Employee role related database operations
type RoleRepository interface {
	Roles() ([]entity.Roles, []error)
	Role(id uint) (*entity.Roles, []error)
	UpdateRole(role *entity.Roles) (*entity.Roles, []error)
	DeleteRole(id uint) (*entity.Roles, []error)
	StoreRole(role *entity.Roles) (*entity.Roles, []error)
}

//JobRepository specifies applicaton Employee Job related database oprations
type JobRepository interface {
	Job() (*entity.Employee_job, []error)
}
