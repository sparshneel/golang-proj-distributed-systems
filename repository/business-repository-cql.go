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
		panic("error inserting data record to table " + helpers.Table)
	}
	logrus.Info("done inserting record in the table " + helpers.Table)
}

func UpdateRecords(query *gocqlx.Queryx){
	logrus.Info("update records for table " + helpers.Table + " for query: " + query.String())
	 err := query.ExecRelease()
	 if err != nil{
	 	logrus.Error("error updating table " + helpers.Table + " query: " + query.String())
	 	panic("error updating table " + helpers.Table + " query: " + query.String())
	 }
	 logrus.Info(" update complete for table " + helpers.Table + " query: " + query.String())
}

func DeleteRecords(query *gocqlx.Queryx) {
	logrus.Info("deleting records from table " + helpers.Table + " query: " + query.String())
	err := query.ExecRelease()
	if err != nil{
		logrus.Error("error deleting records from table " + helpers.Table + " query: " + query.String())
		panic("error deleting records from table " + helpers.Table + " query: " + query.String())
	}
	logrus.Info(" done deleting records from table " + helpers.Table + " query: " + query.String())
}

func QueryRecords(query *gocqlx.Queryx) []domain.Business{
	var records [] domain.Business
	logrus.Info("fetching records from the table " + helpers.Table + " query: " + query.String())
	err := query.GetRelease(&records)
	if err != nil{
		logrus.Info("error fetching records from table " + helpers.Table + " query: " + query.String())
		panic("error fetching records from table " + helpers.Table + " query: " + query.String())
	}
	logrus.Info(" done fetching records form table " +  helpers.Table + " query: " + query.String())
    return records
}





