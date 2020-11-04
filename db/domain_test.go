package db_test

//測試用
import (
	"github.com/wei840222/certchecker/db"

	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func clearDomainTable() {
	if err := db.DB.
		Session(&gorm.Session{AllowGlobalUpdate: true}).
		Unscoped().
		Delete(db.Domain{}).
		Error; err != nil {
		panic(err)
	}
}

func TestCreateDomainAndListDomain(t *testing.T) {
	defer clearDomainTable()

	err := db.CreateDomain(&db.Domain{
		Name: "test",
		Host: "example.com",
	})
	assert.Nil(t, err)

	ds, err := db.ListDomain()
	assert.Nil(t, err)
	assert.Len(t, ds, 1)
	assert.Equal(t, ds[0].Name, "test")
	assert.Equal(t, ds[0].Host, "example.com")
}

func TestDeleteDomain(t *testing.T) {
	defer clearDomainTable()

	d := db.Domain{
		Name: "test",
		Host: "example.com",
	}
	db.CreateDomain(&d)

	err := db.DeleteDomain(d.ID)
	assert.Nil(t, err)

	ds, _ := db.ListDomain()
	assert.Len(t, ds, 0)
}
