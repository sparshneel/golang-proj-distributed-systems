package service

import (
	"github.com/scylladb/gocqlx"
	"golang-proj-distributed-systems/domain"
	"golang-proj-distributed-systems/helpers"
	"golang-proj-distributed-systems/repository"

	"github.com/scylladb/gocqlx/qb"
	"github.com/sirupsen/logrus"
)


func AddBusiness(business *domain.Business) {
	logrus.Info("starting save record - table : " + helpers.Table)
	stmt, names := qb.Insert(helpers.Table).Columns(helpers.GetColumnNames(business)).ToCql()
	q := gocqlx.Query(repository.GetSessionInstance().Query(stmt), names).BindStruct(business)
	repository.Save(q)
	logrus.Info("record added successfully to table : " + helpers.Table)
}

func UpdateBusiness(business *domain.Business){
	logrus.Info("building update query")
	stmt, names:= qb.Update(helpers.Table).Set(helpers.GetColumnNames(business)).Where(qb.Eq("id")).ToCql()
	q := gocqlx.Query(repository.GetSessionInstance().Query(stmt),names).BindStruct(business)
	repository.UpdateRecords(q)
}

func DeleteBusiness(business *domain.Business){
	logrus.Info("building delete query")
	stmt,names := qb.Delete(helpers.Table).Existing().Where(qb.Eq("id")).ToCql()
	q := gocqlx.Query(repository.GetSessionInstance().Query(stmt),names).BindStruct(business)
	logrus.Info("deleting records for table " + helpers.Table)
	repository.DeleteRecords(q)
	logrus.Info("done deleting records for table " + helpers.Table)
}

func GetBusinessData(){

}

