package admin

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/validations"
)

// Slug is a scanner, could be used as field type, then this field could be used as slug field
type Slug struct {
	Slug string
}

// Scan scan value into Slug
func (slug *Slug) Scan(value interface{}) error {
	if bytes, ok := value.([]byte); ok {
		slug.Slug = string(bytes)
	} else if str, ok := value.(string); ok {
		slug.Slug = str
	} else if strs, ok := value.([]string); ok {
		slug.Slug = strs[0]
	}
	return nil
}

// Value get slug's Value
func (slug Slug) Value() (driver.Value, error) {
	return slug.Slug, nil
}

// ConfigureQorMeta configure slug for qor admin
func (Slug) ConfigureQorMeta(meta resource.Metaor) {
	if meta, ok := meta.(*Meta); ok {
		res := meta.GetBaseResource().(*Resource)
		res.UseTheme("slug")

		slugMetaName := meta.Name
		fieldName := strings.TrimSuffix(meta.Name, "WithSlug")
		if meta := res.GetMeta(fieldName); meta != nil {
			meta.Type = "slug"
		}

		res.AddValidator(&resource.Validator{
			Name: fmt.Sprintf("%v-slug-validator", fieldName),
			Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
				if meta := metaValues.Get(slugMetaName); meta != nil {
					slug := utils.ToString(metaValues.Get(slugMetaName).Value)
					if slug == "" {
						return validations.NewError(record, fieldName, fieldName+"'s slug can't be blank")
					} else if strings.Contains(slug, " ") {
						return validations.NewError(record, fieldName, fieldName+"'s slug can't contains blank string")
					}
				} else {
					if field, ok := context.GetDB().NewScope(record).FieldByName(slugMetaName); ok && field.IsBlank {
						return validations.NewError(record, fieldName, fieldName+"'s slug can't be blank")
					}
				}
				return nil
			},
		})

		res.OverrideIndexAttrs(func() {
			var attrs = res.ConvertSectionToStrings(res.IndexAttrs())
			var hasSlug bool
			for _, attr := range attrs {
				if attr == slugMetaName {
					hasSlug = true
					break
				}
			}

			if !hasSlug {
				res.IndexAttrs(res.IndexAttrs(), "-"+slugMetaName)
			}
		})

		res.OverrideShowAttrs(func() {
			res.ShowAttrs(res.ShowAttrs(), "-"+slugMetaName)
		})

		res.OverrideEditAttrs(func() {
			res.EditAttrs(res.EditAttrs(), "-"+slugMetaName)
		})

		res.OverrideNewAttrs(func() {
			res.NewAttrs(res.NewAttrs(), "-"+slugMetaName)
		})
	}
}
