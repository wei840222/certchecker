package db

import "gorm.io/gorm"

type Domain struct {
	gorm.Model
	Name string
	Host string `gorm:"unique"`
}

func CreateDomain(d *Domain) error {
	return DB.Create(d).Error
}

func DeleteDomain(id uint) error {
	return DB.Delete(Domain{}, id).Error
}

func ListDomain() ([]*Domain, error) {
	var ds []*Domain
	if err := DB.Find(&ds).Error; err != nil {
		return nil, err
	}
	return ds, nil
}
