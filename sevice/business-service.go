package service

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"golang-proj-distributed-systems/domain"
	"golang-proj-distributed-systems/helpers"
	"golang-proj-distributed-systems/repository"
	"strings"

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

func DeleteBusiness(filter map[string]string ){
	logrus.Info("building delete query")
	businesId,err := gocql.ParseUUID(filter["id"])
	if err != nil{
		logrus.Error("Error parsing uuid from string")
		panic("error parsing uuid from string")
	}
	business := new(domain.Business)
	business.Id = businesId
	business.City = filter["city"]
	business.State = filter["state"]
	stmt,names := qb.Delete(helpers.Table).Existing().Where(qb.Eq("id"),qb.Eq("city"),qb.Eq("state")).ToCql()

	q := gocqlx.Query(repository.GetSessionInstance().Query(stmt),names).BindStruct(business)
	logrus.Info("deleting records for table " + helpers.Table + " query " + q.Statement())
	repository.DeleteRecords(q)
	logrus.Info("done deleting records for table " + helpers.Table)
}

func GetBusinessDataById(filter map [string]string) []domain.Business{
	logrus.Info("building select query")
	businesId,err := gocql.ParseUUID(filter["id"])
	if err != nil{
		logrus.Error("Error parsing uuid from string")
		panic("error parsing uuid from string")
	}
	business := new(domain.Business)
	business.Id = businesId
	business.City = filter["city"]
	business.State = filter["state"]
	stmt, names :=qb.Select(helpers.Table).Columns(helpers.SliceToString(strings.Split(helpers.SelectColumns,","))).
		Where(qb.Eq("id"),qb.Eq("city"),qb.Eq("state")).ToCql()
	q := gocqlx.Query(repository.GetSessionInstance().Query(stmt),names).BindStruct(business)
	businessRecords := repository.QueryRecords(q)
	return businessRecords
}

