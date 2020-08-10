package dao

import (
	"encoding/hex"
	"time"

	"github.com/CodyGuo/jumpadmin/internal/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (d *Dao) CreateOrganization(name string) error {
	organization := model.OrgsOrganization{
		Model: &model.Model{
			ID:        genUUID(),
			Name:      name,
			CreatedBy: "jumpadmin",
		},
		DateCreated: time.Now(),
	}
	return organization.Create(d.engine)
}

func (d *Dao) ListOrganization() ([]*model.OrgsOrganization, error) {
	organization := model.OrgsOrganization{}
	return organization.List(d.engine)
}

func (d *Dao) CreateOrganizationUsers(organization, user string) error {
	var (
		org   model.OrgsOrganization
		users model.UsersUser
	)
	if err := d.engine.Where("name = ?", organization).First(&org).Error; err != nil {
		return err
	}
	if err := d.engine.Where("username = ?", user).First(&users).Error; err != nil {
		return err
	}

	ou := model.OrgsOrganizationUsers{
		OrganizationID: org.ID,
		UserID:         users.ID,
	}
	return ou.Create(d.engine)
}

func (d *Dao) DeleteOrganizationUsers(organization, user string) error {
	var (
		org   model.OrgsOrganization
		users model.UsersUser
	)
	if err := d.engine.Where("name = ?", organization).First(&org).Error; err != nil {
		return err
	}
	if err := d.engine.Where("username = ?", user).First(&users).Error; err != nil {
		return err
	}

	ou := model.OrgsOrganizationUsers{
		OrganizationID: org.ID,
		UserID:         users.ID,
	}
	return ou.Delete(d.engine)

}

func (d *Dao) CreateOrganizationAdmins(organization, user string) error {
	var (
		org   model.OrgsOrganization
		users model.UsersUser
	)
	if err := d.engine.Where("name = ?", organization).First(&org).Error; err != nil {
		return err
	}
	if err := d.engine.Where("username = ?", user).First(&users).Error; err != nil {
		return err
	}
	oa := model.OrgsOrganizationAdmins{
		OrganizationID: org.ID,
		UserID:         users.ID,
	}
	return oa.Create(d.engine)

}
func (d *Dao) DeleteOrganizationAdmins(organization, user string) error {
	var (
		org   model.OrgsOrganization
		users model.UsersUser
	)
	if err := d.engine.Where("name = ?", organization).First(&org).Error; err != nil {
		return err
	}
	if err := d.engine.Where("username = ?", user).First(&users).Error; err != nil {
		return err
	}
	oa := model.OrgsOrganizationAdmins{
		OrganizationID: org.ID,
		UserID:         users.ID,
	}
	return oa.Delete(d.engine)
}

func genUUID() string {
	id := uuid.New()
	var text [32]byte
	hex.Encode(text[:], id[:4])
	hex.Encode(text[8:], id[4:])
	return string(text[:])
}
