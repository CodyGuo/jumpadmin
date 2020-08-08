package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type OrgsOrganization struct {
	*Model
	DateCreated time.Time `josn:"date_created"`
}

func (o OrgsOrganization) TableName() string {
	return "orgs_organization"
}

func (o OrgsOrganization) Create(db *gorm.DB) error {
	return db.Create(&o).Error
}

func (o OrgsOrganization) List(db *gorm.DB) ([]*OrgsOrganization, error) {
	var orgs []*OrgsOrganization
	if err := db.Find(&orgs).Error; err != nil {
		return nil, err
	}
	return orgs, nil
}

type OrgsOrganizationUsers struct {
	ID             uint32 `gorm:"primary_key" json:"id"`
	OrganizationID string `json:"organization_id"`
	UserID         string `json:"user_id"`
}

func (o OrgsOrganizationUsers) TableName() string {
	return "orgs_organization_users"
}

func (o OrgsOrganizationUsers) Create(db *gorm.DB) error {
	return db.Create(&o).Error
}

func (o OrgsOrganizationUsers) Delete(db *gorm.DB) error {
	return db.Where("organization_id = ? and user_id = ?", o.OrganizationID, o.UserID).Delete(&o).Error
}

type OrgsOrganizationAdmins struct {
	ID             uint32 `gorm:"primary_key" json:"id"`
	OrganizationID string `json:"organization_id"`
	UserID         string `json:"user_id"`
}

func (o OrgsOrganizationAdmins) TableName() string {
	return "orgs_organization_admins"
}

func (o OrgsOrganizationAdmins) Create(db *gorm.DB) error {
	return db.Create(&o).Error
}

func (o OrgsOrganizationAdmins) Delete(db *gorm.DB) error {
	return db.Where("organization_id = ? and user_id = ?", o.OrganizationID, o.UserID).Delete(&o).Error
}
