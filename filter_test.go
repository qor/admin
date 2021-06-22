package admin_test

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	admin "github.com/TykTechnologies/raava-admin"
	. "github.com/TykTechnologies/raava-admin/tests/dummy"
	qorTestUtils "github.com/qor/qor/test/utils"
)

const FORMTYPE = "application/x-www-form-urlencoded"

func TestDefaultFilter(t *testing.T) {
	initData()
	user := Admin.GetResource("User")
	user.Filter(&admin.Filter{
		Name: "ID",
	})
	//filter bool
	user.Filter(&admin.Filter{
		Name: "Active",
	})
	//filter intiger
	user.Filter(&admin.Filter{
		Name: "Age",
		Config: &admin.SelectOneConfig{
			Collection: []string{"10", "11", "12", "13", "14", "15"},
		},
	})
	//filter string
	user.Filter(&admin.Filter{
		Name: "Role",
		Config: &admin.SelectOneConfig{
			Collection: []string{Role_editor, Role_supervisor, Role_system_administrator, Role_developer},
		},
	})
	//filter belongs_to
	user.Filter(&admin.Filter{
		Name:   "CreditCard",
		Config: &admin.SelectOneConfig{},
	})
	//filter has_one
	user.Filter(&admin.Filter{
		Name:   "Profile",
		Config: &admin.SelectOneConfig{},
	})
	//filter has_many
	user.Filter(&admin.Filter{
		Name:   "Addresses",
		Config: &admin.SelectOneConfig{},
	})
	//filter belongs_to
	user.Filter(&admin.Filter{
		Name:   "Company",
		Config: &admin.SelectOneConfig{},
	})
	//filter many_to_many
	user.Filter(&admin.Filter{
		Name:   "Languages",
		Config: &admin.SelectOneConfig{},
	})

	cases := []struct {
		filter string
		value  interface{}
		expect []string
	}{
		{filter: "ID", value: 2, expect: []string{"user_2"}},
		{filter: "Active", value: true, expect: []string{"user_1", "user_2", "user_3"}},
		{filter: "Age", value: 13, expect: []string{"user_4"}},
		{filter: "Role", value: Role_supervisor, expect: []string{"user_3", "user_4"}},
		{filter: "CreditCard", value: 5, expect: []string{"user_5"}},
		{filter: "Profile", value: 3, expect: []string{"user_3"}},
		{filter: "Addresses", value: 1, expect: []string{"user_1", "user_5"}},
		{filter: "Company", value: 2, expect: []string{"user_5", "user_4", "user_6"}},
		{filter: "Languages", value: 3, expect: []string{"user_2", "user_5"}},
	}

	for _, cv := range cases {
		req, _ := http.NewRequest("GET", fmt.Sprint(server.URL, "/admin/users?"+encodeValues(cv.filter, cv.value)), nil)
		req.Header.Set("Content-Type", FORMTYPE)
		context := Admin.NewContext(nil, req)
		context.Roles = []string{Role_system_administrator}
		context.Resource = user
		context.Searcher = &admin.Searcher{Context: context}
		context.Request.ParseForm()

		if users, err := context.FindMany(); err == nil {
			arr := *users.(*[]*User)
			if len(arr) != len(cv.expect) {
				t.Fatalf("filter: %v, except: %v, got: %v", cv.filter, cv.expect, len(arr))
			}

			for _, v := range arr {
				if !strings.Contains(strings.Join(cv.expect, ","), v.Name) {
					t.Fatalf("filter: %v, except: %v, got: %v", cv.filter, cv.expect, v.Name)
				}
			}
		} else {
			t.Fatal(err)
		}
	}

}

func TestSelectManyConfigFilter(t *testing.T) {
	initData()
	userM := Admin.AddResource(&User{}, &admin.Config{Name: "SelectManyFilter"})
	userM.Filter(&admin.Filter{
		Name: "Age",
		Config: &admin.SelectManyConfig{
			Collection: []string{"10", "11", "12", "13", "14", "15"},
		},
	})

	userM.Filter(&admin.Filter{
		Name: "Role",
		Config: &admin.SelectManyConfig{
			Collection: []string{Role_editor, Role_supervisor, Role_system_administrator, Role_developer},
		},
	})
	//filter belongs_to
	userM.Filter(&admin.Filter{
		Name:   "CreditCard",
		Config: &admin.SelectManyConfig{},
	})
	//filter has_one
	userM.Filter(&admin.Filter{
		Name:   "Profile",
		Config: &admin.SelectManyConfig{},
	})
	//filter has_many
	userM.Filter(&admin.Filter{
		Name:   "Addresses",
		Config: &admin.SelectManyConfig{},
	})
	//filter belongs_to
	userM.Filter(&admin.Filter{
		Name:   "Company",
		Config: &admin.SelectManyConfig{},
	})
	//filter many_to_many
	userM.Filter(&admin.Filter{
		Name:   "Languages",
		Config: &admin.SelectManyConfig{},
	})

	cases := []struct {
		filter string
		value  interface{}
		expect []string
	}{
		{filter: "Age", value: []uint{13, 14, 15, 17}, expect: []string{"user_4", "user_5", "user_6"}},
		{filter: "Role", value: []string{Role_editor, Role_supervisor}, expect: []string{"user_1", "user_2", "user_3", "user_4"}},
		{filter: "CreditCard", value: []uint{5}, expect: []string{"user_5"}},
		{filter: "Profile", value: []uint{3}, expect: []string{"user_3"}},
		{filter: "Addresses", value: []uint{1}, expect: []string{"user_1", "user_5"}},
		{filter: "Company", value: []uint{1, 2}, expect: []string{"user_1", "user_3", "user_2", "user_5", "user_4", "user_6"}},
		{filter: "Languages", value: []uint{1}, expect: []string{"user_1", "user_3", "user_2"}},
	}

	for _, cv := range cases {
		req, _ := http.NewRequest("GET", fmt.Sprint(server.URL, "/admin/select_many_filters?"+encodeValues(cv.filter, cv.value)), nil)
		req.Header.Set("Content-Type", FORMTYPE)
		context := Admin.NewContext(nil, req)
		context.Roles = []string{Role_system_administrator}
		context.Resource = userM
		context.Searcher = &admin.Searcher{Context: context}
		context.Request.ParseForm()

		if users, err := context.FindMany(); err == nil {
			arr := *users.(*[]*User)
			if len(arr) != len(cv.expect) {
				t.Fatalf("filter: %v, except: %v, got: %v", cv.filter, cv.expect, len(arr))
			}

			for _, v := range arr {
				if !strings.Contains(strings.Join(cv.expect, ","), v.Name) {
					t.Fatalf("filter: %v, except: %v, got: %v", cv.filter, cv.expect, v.Name)
				}
			}
		} else {
			t.Fatal(err)
		}
	}
}

func encodeValues(filter string, values interface{}) string {
	key := fmt.Sprintf("filters[%s].Value", filter)
	tv := reflect.ValueOf(values)
	val := make(url.Values)
	switch tv.Kind() {
	case reflect.Slice:
		for i := 0; i < tv.Len(); i++ {
			val.Add(key, fmt.Sprint(tv.Index(i)))
		}
	default:
		val.Add(key, fmt.Sprint(tv))
	}
	return val.Encode()
}

func initData() {
	db.DropTableIfExists("user_languages")
	qorTestUtils.ResetDBTables(db, &Language{}, &User{}, &Company{}, &CreditCard{}, &Address{}, &Profile{})
	coms := []Company{{Name: "Company A"}, {Name: "Company B"}}
	db.Save(&coms[0])
	db.Save(&coms[1])

	lans := []Language{{Name: "en-gb"}, {Name: "en-us"}, {Name: "cn-zh"}, {Name: "ja-jp"}}
	db.Save(&lans[0])
	db.Save(&lans[1])
	db.Save(&lans[2])
	db.Save(&lans[3])

	profiles := []Profile{{Name: "Profile of user_1"}, {Name: "Profile of user_2"}, {Name: "Profile of user_3"}, {Name: "Profile of user_4"}}
	addrs := []Address{{Address1: "Address1"}, {Address1: "Address2"}, {Address1: "Address3"}, {Address1: "Address4"}}
	db.Save(&addrs[0])
	db.Save(&addrs[1])
	db.Save(&addrs[2])
	db.Save(&addrs[3])

	cards := []CreditCard{{Number: "1111111111111"}, {Number: "1211111111111"}, {Number: "1311111111111"}, {Number: "1411111111111"}, {Number: "1511111111111"}, {Number: "1611111111111"}}
	db.Save(&cards[0])
	db.Save(&cards[1])
	db.Save(&cards[2])
	db.Save(&cards[3])
	db.Save(&cards[4])
	db.Save(&cards[5])

	roles := []string{Role_editor, Role_supervisor, Role_system_administrator, Role_developer}
	db.Save(&User{Name: "user_1", Age: 10, Profile: profiles[0], Role: roles[0], Active: true, Addresses: []Address{{Address1: "Address1"}, {Address1: "Address3"}}, CreditCard: cards[0], Company: &coms[0], Languages: []Language{lans[0], lans[1]}})
	db.Save(&User{Name: "user_2", Age: 11, Profile: profiles[1], Role: roles[0], Active: true, Addresses: []Address{{Address1: "Address1"}, {Address1: "Address5"}}, CreditCard: cards[1], Company: &coms[0], Languages: []Language{lans[0], lans[2]}})
	db.Save(&User{Name: "user_3", Age: 12, Profile: profiles[2], Role: roles[1], Active: true, Addresses: []Address{{Address1: "Address1"}}, CreditCard: cards[2], Company: &coms[0], Languages: []Language{lans[0], lans[3]}})
	db.Save(&User{Name: "user_4", Age: 13, Profile: profiles[3], Role: roles[1], Active: false, Addresses: []Address{{Address1: "Address2"}, {Address1: "Address4"}}, CreditCard: cards[3], Company: &coms[1], Languages: []Language{lans[1]}})
	db.Save(&User{Name: "user_5", Age: 14, Role: roles[2], Active: false, Addresses: []Address{{Address1: "Address2"}}, CreditCard: cards[4], Company: &coms[1], Languages: []Language{lans[1], lans[2]}})
	db.Save(&User{Name: "user_6", Age: 15, Role: roles[3], Active: false, Addresses: []Address{{Address1: "Address3"}, {Address1: "Address5"}}, CreditCard: cards[5], Company: &coms[1], Languages: []Language{lans[1], lans[3]}})
}
