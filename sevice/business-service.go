package service

import(
	"golang-proj-distributed-systems/domain"
	"golang-proj-distributed-systems/helpers"
	"golang-proj-distributed-systems/repository"

	"github.com/sirupsen/logrus"

	"reflect"
	"strings"
)


func AddBusiness(business domain.Business) {
	logrus.Info("starting save record - table : " + helpers.Table)
	repository.Save(business)
	logrus.Info("record added successfully to table : " + helpers.Table)
}

func Update(business domain.Business){
	logrus.Info("building update query")
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString(helpers.UPDATE + " ")
	stringBuilder.WriteString(helpers.Table)
	stringBuilder.WriteString(" set ")
	e := reflect.ValueOf(business).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varValue := e.Field(i).String()
		stringBuilder.WriteString(varName + "=" + varValue)
	}
	repository.UpdateRecords(stringBuilder.String())
}

func Delete(businessId string){
	logrus.Info("building delete query")
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString(helpers.DELETE + " from ")
	stringBuilder.WriteString(helpers.Table + " where ")
	stringBuilder.WriteString("id=" + businessId)
	logrus.Info("updating records for table " + helpers.Table)
	repository.DeleteRecords(stringBuilder.String())
	logrus.Info("done updating records for table " + helpers.Table)
}