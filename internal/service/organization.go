package service

import "jumpadmin/internal/model"

type CreateOrganizationRequest struct {
	Name string
}

type OrganizationUsersRequest struct {
	Organization string
	Username     string
}

type OrganizationAdminsRequest struct {
	Organization string
	Username     string
}

func (svc *Service) CreateOrganization(param *CreateOrganizationRequest) error {
	return svc.dao.CreateOrganization(param.Name)
}

func (svc *Service) ListOrganization() ([]*model.OrgsOrganization, error) {
	return svc.dao.ListOrganization()
}

func (svc *Service) CreateOrganizationUsers(param *OrganizationUsersRequest) error {
	return svc.dao.CreateOrganizationUsers(param.Organization, param.Username)
}

func (svc *Service) DeleteOrganizationUsers(param *OrganizationUsersRequest) error {
	return svc.dao.DeleteOrganizationUsers(param.Organization, param.Username)
}

func (svc *Service) CreateOrganizationAdmins(param *OrganizationAdminsRequest) error {
	return svc.dao.CreateOrganizationAdmins(param.Organization, param.Username)
}
func (svc *Service) DeleteOrganizationAdmins(param *OrganizationAdminsRequest) error {
	return svc.dao.DeleteOrganizationAdmins(param.Organization, param.Username)
}
