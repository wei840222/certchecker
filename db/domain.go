package db

import "time"

//增刪domain
type Domain struct {
	ID    uint       `gorm:"primarykey" json:"id"`
	Name  string     `json:"name" binding:"required"`
	Host  string     `gorm:"unique" json:"host" binding:"required"`
	Since *time.Time `json:"since,omitempty"`
	End   *time.Time `json:"end,omitempty"`
	Error string     `json:"error,omitempty"`
}

func CreateDomain(d *Domain) error {
	return DB.Create(d).Error
}

func UpdateDomain(id uint, d *Domain) error {
	return DB.Model(Domain{}).Where("id = ?", id).Updates(d).Error
}

func DeleteDomain(id uint) error {
	return DB.Delete(Domain{}, id).Error
}

func DeleteDomainError(id uint) error {
	return DB.Model(Domain{}).Where("id = ?", id).Update("error", "").Error
}

func ListDomain() ([]*Domain, error) {
	var ds []*Domain
	if err := DB.Find(&ds).Error; err != nil {
		return nil, err
	}
	return ds, nil
}
