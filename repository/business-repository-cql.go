package repository

import (
	"strings"

	"golang-proj-distributed-systems/domain"
	"golang-proj-distributed-systems/helpers"

	"github.com/sirupsen/logrus"
)

func Save(business domain.Business) {
	logrus.Info("Configuring cql session")
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString(helpers.INSERT)
	stringBuilder.WriteString(" into ")
	stringBuilder.WriteString(helpers.Table)
	stringBuilder.WriteString(" values(?,?,?,?), ")
	logrus.Info("inserting record in the table " + helpers.Table)
	err := getSessionInstance().Query(stringBuilder.String(),business.Id,business.BusinessAddress,business.Telephone).Exec()
	if err !=nil{
		panic("error inserting data record to table " + helpers.Table)
	}
	logrus.Info("done inserting record in the table " + helpers.Table)
}

func FetechRecords(query string) []domain.Business {
	logrus.Info("fetching data for provided query " + query)
	m := map[string]interface{}{}
	var business []domain.Business
	iter := getSessionInstance().Query(query).Iter()
	for iter.MapScan(m){
		business = append(business,domain.Business{
			Id:              m["id"].(string),
			Name:            m["name"].(string),
			BusinessAddress: nil,
			Telephone:       m["telephone"].(int64),
		})
	}
	logrus.Info("Done fetching records from table: " + helpers.Table + " for query: " + query)
	return business
}

func UpdateRecords(query string){
	logrus.Info("update records for table " + helpers.Table + " for query: " + query)
	 err := getSessionInstance().Query(query).Exec()
	 if err != nil{
	 	logrus.Error("error updating table " + helpers.Table + " query: " + query)
	 	panic("error updating table " + helpers.Table + " query: " + query)
	 }
	 logrus.Info(" update complete for table " + helpers.Table + " query: " + query)
}

func DeleteRecords(query string) {
	logrus.Info("deleting records from table " + helpers.Table + " query: " + query)
	err := getSessionInstance().Query(query).Exec()
	if err != nil{
		logrus.Error("error deleting records from table " + helpers.Table + " query: " + query)
		panic("error deleting records from table " + helpers.Table + " query: " + query)
	}
	logrus.Info(" done deleting records from table " + helpers.Table + " query: " + query)
}




