package cservice

import (
	"gitlab.com/username/online-service-and-customer-care/company"
	"gitlab.com/username/online-service-and-customer-care/entity"
)

// CompanyService implements menu.CommentService interface
type CompanyService struct {
	companyRepo company.CompanyRepository
}

// NewCompanyService returns a new CompanyService object
func NewCompanyService(compRepo company.CompanyRepository) company.CompanyService {
	return &CompanyService{companyRepo: compRepo}
}

// Companies returns all stored companies
func (cs *CompanyService) Companies() ([]entity.Company, []error) {
	comps, errs := cs.companyRepo.Companies()
	if len(errs) > 0 {
		return nil, errs
	}
	return comps, errs
}

// Company retrieves stored company by its id
func (cs *CompanyService) Company(id uint) (*entity.Company, []error) {
	comp, errs := cs.companyRepo.Company(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// UpdateCompany updates a given company
func (cs *CompanyService) UpdateCompany(company *entity.Company) (*entity.Company, []error) {
	comp, errs := cs.companyRepo.UpdateCompany(company)
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// DeleteCompany deletes a given comment
func (cs *CompanyService) DeleteCompany(id uint) (*entity.Company, []error) {
	comp, errs := cs.companRepo.DeleteCompany(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return comp, errs
}

// StoreComment stores a given comment
func (cs *CompanyService) StoreCompany(company *entity.Company) (*entity.Company, []error) {
	return nil, nil
}
