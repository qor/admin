package admin_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	admin "github.com/TykTechnologies/raava-admin"
	. "github.com/TykTechnologies/raava-admin/tests/dummy"
	"github.com/jinzhu/gorm"
)

var (
	server       *httptest.Server
	db           *gorm.DB
	Admin        *admin.Admin
	adminHandler http.Handler
)

func init() {
	Admin = NewDummyAdmin()
	adminHandler = Admin.NewServeMux("/admin")
	db = Admin.DB
	server = httptest.NewServer(adminHandler)
}

func TestMain(m *testing.M) {
	// Create universal logged-in user for test.
	createLoggedInUser()
	retCode := m.Run()

	os.Exit(retCode)
}

func createLoggedInUser() *User {
	user := User{Name: LoggedInUserName, Role: Role_system_administrator}
	if err := db.Save(&user).Error; err != nil {
		panic(err)
	}

	return &user
}
