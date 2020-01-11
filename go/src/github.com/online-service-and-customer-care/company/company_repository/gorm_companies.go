package company_repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/username/online-service-and-customer-care/company"
	"gitlab.com/username/online-service-and-customer-care/entity"
)

// CompanyGormRepo implements company repository interface
type CompanyGormRepo struct {
	conn *gorm.DB
}

// NewCompanyGormRepo returns new object of CompanyGormRepo
func NewCompanyGormRepo(db *gorm.DB) company.CompanyRepository {
	return &CompanyGormRepo{conn: db}
}

// Companies returns all companies stored in the database
func (compRepo *CompanyGormRepo) Companies() ([]entity.Company, []error) {
	comps := []entity.Company{}
	//fmt.Println(cmnts)
	errs := compRepo.conn.Find(&comps).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comps, errs
}

// Company retrieves a company from the database by its id
func (compRepo *CompanyGormRepo) Company(id uint) (*entity.Company, []error) {
	comps := entity.Company{}
	errs := compRepo.conn.First(&comps, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &comps, errs
}

// UpdateCompany updates a given company in the database
func (compRepo *CompanyGormRepo) UpdateCompany(company *entity.Company) (*entity.Company, []error) {
	comp := company
	errs := compRepo.conn.Save(comp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// DeleteCompany deletes a given company from the database
func (compRepo *CompanyGormRepo) DeleteCompany(id uint) (*entity.Company, []error) {
	comp, errs := compRepo.Company(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = compRepo.conn.Delete(comp, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// StoreCompany stores a given company in the database
func (compRepo *CompanyGormRepo) StoreCompany(company *entity.Company) (*entity.Company, []error) {
	comp := company
	errs := compRepo.conn.Create(comp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}
