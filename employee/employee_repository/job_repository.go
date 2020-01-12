package employee_repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/username/online-service-and-customer-care/employee"
	"gitlab.com/username/online-service-and-customer-care/entity"
)

//JobGormRepo implments Job repository
type JobGormRepo struct {
	dbconn *gorm.DB
}

//NewJobGormRepo is a construtor for JobGromRepo+
func NewJobGormRepo(db *gorm.DB) employee.JobRepository {
	return &JobGormRepo{dbconn: db}
}

// Job retrieves a job of an employee from the database by its id
func (jobRepo *JobGormRepo) Job() (*entity.Employee_job, []error) {
	job := entity.Employee_job{}
	errs := jobRepo.dbconn.First(&job).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &job, errs
}
