package sevice

import(
	"github.com/gin-gonic/gin/domain"
	"github.com/gin-gonic/gin/helpers"
	"github.com/gin-gonic/gin/repository"
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