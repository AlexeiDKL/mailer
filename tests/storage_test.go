package tests

import (
	"fmt"
	"os"
	"testing"

	"dkl.dklsa.mailer/iternal/config"
	"dkl.dklsa.mailer/iternal/storage"
	sqlites "dkl.dklsa.mailer/iternal/storage/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func initConfig() *config.Config {
	os.Setenv("CONFIG_PATH", "../config/local.yaml")
	return config.MustLoad()
}

type TestStorageCompanyStuct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestCompanyTable(t *testing.T) {

	var company TestStorageCompanyStuct
	company.ID = 0
	company.Name = "Test Company"

	cfg := initConfig()
	path := "." + cfg.StoragePath
	fmt.Println(path)
	storag, err := sqlites.CreateCompanyTable(path)
	if err != nil {
		t.Fatalf("Error creating company table: %v", err)
	}
	companyStorage := sqlites.CreateCompanyStorages(storag)

	id, err := companyStorage.Insert(company.Name)
	if err != nil {
		t.Fatalf("Error inserting company: %v", err)
	}
	company.ID = int(id)
	pair := storage.Pair{
		Type:  sqlites.Name,
		Value: company.Name,
	}
	com, err := companyStorage.Select(pair)
	if err != nil {
		t.Fatalf("Error selecting company: %v", err)
	}
	if com.ID != company.ID || com.Name != company.Name {
		t.Fatalf("Company not saved correctly: %+v", com)
	}
	companyStorage.Update(storage.Pair{Type: })

	companyStorage.Delete(pair)
}
