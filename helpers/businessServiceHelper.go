package helpers

import (
	"golang-proj-distributed-systems/domain"
	"reflect"
	"strings"
)

func GetColumnNames(business *domain.Business) string {
	stringBuilder := strings.Builder{}
	e := reflect.ValueOf(business).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		if i == 0 && e.NumField() == 1 {
			stringBuilder.WriteString(string('"'))
			stringBuilder.WriteString(varName)
			stringBuilder.WriteString(string('"'))
		}
		if i == (e.NumField() - 1) {
			stringBuilder.WriteString(string('"'))
			stringBuilder.WriteString(varName)
			stringBuilder.WriteString(string('"'))
		}
		stringBuilder.WriteString(string('"'))
		stringBuilder.WriteString(varName)
		stringBuilder.WriteString(string('"') + ",")
	}
	return stringBuilder.String()
}

