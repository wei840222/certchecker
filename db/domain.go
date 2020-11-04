package db

//增刪domain
type Domain struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `json:"name" binding:"required"`
	Host string `gorm:"unique" json:"host" binding:"required"`
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
