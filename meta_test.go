package admin_test

import (
	"reflect"
	"testing"

	. "github.com/TykTechnologies/raava-admin/tests/dummy"

	admin "github.com/TykTechnologies/raava-admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
)

func TestTextInput(t *testing.T) {
	user := Admin.AddResource(&User{})
	meta := user.GetMeta("Name")

	if meta.Label != "Name" {
		t.Error("default label not set")
	}

	if meta.GetFieldName() != "Name" {
		t.Error("default Alias is not same as field Name")
	}

	if meta.Type != "string" {
		t.Error("default Type is not string")
	}
}

func TestDefaultMetaType(t *testing.T) {
	var (
		user        = Admin.AddResource(&User{})
		booleanMeta = user.GetMeta("Active")
		timeMeta    = user.GetMeta("RegisteredAt")
		numberMeta  = user.GetMeta("Age")
		fileMeta    = user.GetMeta("Avatar")
	)

	if booleanMeta.Type != "checkbox" {
		t.Error("boolean field doesn't set as checkbox")
	}

	if timeMeta.Type != "datetime" {
		t.Error("time field doesn't set as datetime")
	}

	if numberMeta.Type != "number" {
		t.Error("number field doesn't set as number")
	}

	if fileMeta.Type != "file" {
		t.Error("file field doesn't set as file")
	}
}

func TestRelationFieldMetaType(t *testing.T) {
	userRecord := &User{}
	db.Create(userRecord)

	user := Admin.AddResource(&User{})

	userProfileMeta := user.GetMeta("Profile")

	if userProfileMeta.Type != "single_edit" {
		t.Error("has_one relation doesn't generate single_edit type meta")
	}

	userAddressesMeta := user.GetMeta("Addresses")

	if userAddressesMeta.Type != "collection_edit" {
		t.Error("has_many relation doesn't generate collection_edit type meta")
	}

	userLanguagesMeta := user.GetMeta("Languages")

	if userLanguagesMeta.Type != "select_many" {
		t.Error("many_to_many relation doesn't generate select_many type meta")
	}
}

func TestGetStringMetaValue(t *testing.T) {
	user := Admin.AddResource(&User{})
	stringMeta := user.GetMeta("Name")

	UserName := "user name"
	userRecord := &User{Name: UserName}
	db.Create(&userRecord)
	value := stringMeta.GetValuer()(userRecord, &qor.Context{Config: &qor.Config{DB: db}})

	if value.(string) != UserName {
		t.Error("resource's value doesn't get")
	}
}

func TestGetStructMetaValue(t *testing.T) {
	user := Admin.AddResource(&User{})
	structMeta := user.GetMeta("CreditCard")

	creditCard := CreditCard{
		Number: "123456",
		Issuer: "bank",
	}

	userRecord := &User{CreditCard: creditCard}
	db.Create(&userRecord)

	value := structMeta.GetValuer()(userRecord, &qor.Context{Config: &qor.Config{DB: db}})
	creditCardValue := reflect.Indirect(reflect.ValueOf(value))

	if creditCardValue.FieldByName("Number").String() != "123456" || creditCardValue.FieldByName("Issuer").String() != "bank" {
		t.Error("struct field value doesn't get")
	}
}

func TestGetSliceMetaValue(t *testing.T) {
	user := Admin.AddResource(&User{})
	sliceMeta := user.GetMeta("Addresses")

	address1 := &Address{Address1: "an address"}
	address2 := &Address{Address1: "another address"}

	userRecord := &User{Addresses: []Address{*address1, *address2}}
	db.Create(&userRecord)

	value := sliceMeta.GetValuer()(userRecord, &qor.Context{Config: &qor.Config{DB: db}})
	addresses := reflect.Indirect(reflect.ValueOf(value))

	if addresses.Index(0).FieldByName("Address1").String() != "an address" || addresses.Index(1).FieldByName("Address1").String() != "another address" {
		t.Error("slice field value doesn't get")
	}
}

func TestStringMetaSetter(t *testing.T) {
	user := Admin.AddResource(&User{})
	meta := user.GetMeta("Name")

	UserName := "new name"
	userRecord := &User{Name: UserName}
	db.Create(&userRecord)

	metaValue := &resource.MetaValue{
		Name:  "User.Name",
		Value: UserName,
		Meta:  meta,
	}

	meta.GetSetter()(userRecord, metaValue, &qor.Context{Config: &qor.Config{DB: db}})
	if userRecord.Name != UserName {
		t.Error("resource's value doesn't set")
	}
}

// TODO: waiting for Juice to explain logic here. spent too much time on this..
func TestManyToManyMetaSetter(t *testing.T) {
	// userRecord := &User{Name: "A user"}
	// db.Create(&userRecord)

	// en := &Language{Name: "EN"}
	// cn := &Language{Name: "CN"}
	// db.Create(&en)
	// db.Create(&cn)

	// user := Admin.AddResource(&User{})
	// meta := &admin.Meta{Name: "Languages", Type: "select_many", Collection: [][]string{{fmt.Sprintf("%v", en.Id), en.Name}, {fmt.Sprintf("%v", cn.Id), cn.Name}}}
	// user.Meta(meta)

	// metaValue := &resource.MetaValue{
	// 	Name:  "User.Languages",
	// 	Meta:  meta,
	// 	Value: []int{en.Id, cn.Id},
	// }
	// meta.Setter(userRecord, metaValue, &qor.Context{Config: &qor.Config{DB: db}})

	// if len(userRecord.Languages) != 2 {
	// 	t.Error("many to many resource's value doesn't set")
	// }
}

func TestNestedField(t *testing.T) {
	profileModel := Profile{
		Name:  "Qor",
		Sex:   "Female",
		Phone: Phone{Num: "1024"},
	}
	userModel := &User{Profile: profileModel}
	db.Create(userModel)

	user := Admin.AddResource(&User{})
	profileNameMeta := &admin.Meta{Name: "Profile.Name"}
	user.Meta(profileNameMeta)
	profileSexMeta := &admin.Meta{Name: "Profile.Sex"}
	user.Meta(profileSexMeta)
	phoneNumMeta := &admin.Meta{Name: "Profile.Phone.Num"}
	user.Meta(phoneNumMeta)

	userModel.Profile = Profile{}
	valx := phoneNumMeta.GetValuer()(userModel, &qor.Context{Config: &qor.Config{DB: db}})
	if val, ok := valx.(string); !ok || val != profileModel.Phone.Num {
		t.Errorf("Profile.Phone.Num: got %q; expect %q", val, profileModel.Phone.Num)
	}
	if userModel.Profile.Name != profileModel.Name {
		t.Errorf("Profile.Name: got %q; expect %q", userModel.Profile.Name, profileModel.Name)
	}
	if userModel.Profile.Sex != profileModel.Sex {
		t.Errorf("Profile.Sex: got %q; expect %q", userModel.Profile.Sex, profileModel.Sex)
	}
	if userModel.Profile.Phone.Num != profileModel.Phone.Num {
		t.Errorf("Profile.Phone.Num: got %q; expect %q", userModel.Profile.Phone.Num, profileModel.Phone.Num)
	}

	mvs := &resource.MetaValues{
		Values: []*resource.MetaValue{
			{
				Name:  "Profile.Name",
				Value: "Qor III",
				Meta:  profileNameMeta,
			},
			{
				Name:  "Profile.Sex",
				Value: "Male",
				Meta:  profileSexMeta,
			},
			{
				Name:  "Profile.Phone.Num",
				Value: "2048",
				Meta:  phoneNumMeta,
			},
		},
	}
	profileNameMeta.GetSetter()(userModel, mvs.Values[0], &qor.Context{Config: &qor.Config{DB: db}})
	if userModel.Profile.Name != mvs.Values[0].Value {
		t.Errorf("Profile.Name: got %q; expect %q", userModel.Profile.Name, mvs.Values[0].Value)
	}
	profileSexMeta.GetSetter()(userModel, mvs.Values[1], &qor.Context{Config: &qor.Config{DB: db}})
	if userModel.Profile.Sex != mvs.Values[1].Value {
		t.Errorf("Profile.Sex: got %q; expect %q", userModel.Profile.Sex, mvs.Values[1].Value)
	}
	phoneNumMeta.GetSetter()(userModel, mvs.Values[2], &qor.Context{Config: &qor.Config{DB: db}})
	if userModel.Profile.Phone.Num != mvs.Values[2].Value {
		t.Errorf("Profile.Phone.Num: got %q; expect %q", userModel.Profile.Phone.Num, mvs.Values[2].Value)
	}
}
