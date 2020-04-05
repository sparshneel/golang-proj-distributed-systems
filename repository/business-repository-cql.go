package repository

import (
	"github.com/scylladb/gocqlx"
	"golang-proj-distributed-systems/domain"

	"golang-proj-distributed-systems/helpers"

	"github.com/sirupsen/logrus"
)

func Save(query *gocqlx.Queryx) {
	logrus.Info("inserting record in the table " + helpers.Table)
	err := query.ExecRelease()
	if err !=nil{
		logrus.Error("error inserting record to table " + helpers.Table + " cause " + err.Error())
		panic("error inserting record to table " + helpers.Table + " cause: " + err.Error())
	}
	logrus.Info("done inserting record in the table " + helpers.Table)
}

func UpdateRecords(query *gocqlx.Queryx){
	logrus.Info("update records for table " + helpers.Table + " for query: " + query.Statement())
	 err := query.ExecRelease()
	 if err != nil{
	 	logrus.Error("error updating table " + helpers.Table + " query: " + query.Statement() + " cause: " + err.Error())
	 	panic("error updating table " + helpers.Table + " query: " + query.Statement() + " cause: " + err.Error())
	 }
	 logrus.Info(" update complete for table " + helpers.Table + " query: " + query.Statement())
}

func DeleteRecords(query *gocqlx.Queryx) {
	logrus.Info("deleting records from table " + helpers.Table + " query: " + query.Statement())
	err := query.ExecRelease()
	if err != nil{
		logrus.Error("error deleting records from table " + helpers.Table + " query: " + query.Statement() + " cause: " + err.Error())
		panic("error deleting records from table " + helpers.Table + " query: " + query.Statement() + " cause: " + err.Error())
	}
	logrus.Info(" done deleting records from table " + helpers.Table + " query: " + query.Statement())
}

func QueryRecords(query *gocqlx.Queryx) []domain.Business{
	var records []domain.Business
	logrus.Info("fetching records from the table " + helpers.Table + " query: " + query.Statement())
	query.SelectRelease(&records)
	if err != nil{
		logrus.Error("error fetching records from table " + helpers.Table + " query: " + query.Statement() + " cause: " + err.Error())
		panic("error fetching records from table " + helpers.Table + " query: " + query.Statement() + " cause: " + err.Error())
	}
	logrus.Info(" done fetching records form table " +  helpers.Table + " query: " + query.Statement())
    return records
}





