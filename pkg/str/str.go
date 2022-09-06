package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

func Snake(s string) string {
	return strcase.ToSnake(s)
}

func Camel(s string) string {
	return strcase.ToCamel(s)
}

func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
