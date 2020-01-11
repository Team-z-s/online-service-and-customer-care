package employee_service

import (
	"gitlab.com/username/online-service-and-customer-care/employee"
	"gitlab.com/username/online-service-and-customer-care/entity"
)

//JobServiceGorm implments service of employee
type JobServiceGorm struct {
	jobRepo employee.JobRepository
}

//NewJobServiceGorm construstor
func NewJobServiceGorm(jobRepo employee.JobRepository) employee.JobService {
	return &JobServiceGorm{jobRepo: jobRepo}
}

//Job return a single Job
func (js *JobServiceGorm) Job() (*entity.Employee_job, []error) {
	job, errs := js.jobRepo.Job()
	if len(errs) > 0 {
		return nil, errs
	}
	return job, errs
}
