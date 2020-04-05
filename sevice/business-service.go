package service

import (
	"errors"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"golang-proj-distributed-systems/domain"
	"golang-proj-distributed-systems/helpers"
	"golang-proj-distributed-systems/repository"
	"strings"

	"github.com/scylladb/gocqlx/qb"
	"github.com/sirupsen/logrus"

)

func AddBusiness(business *domain.BusinessRep) (bool,error){
	logrus.Info("starting save record - table : " + helpers.Table)
	businessId ,err := gocql.ParseUUID(business.Id)
	if err != nil{
		logrus.Error("Error parsing uuid from string")
		panic("error parsing uuid from string")
	}
	retrivedBusiness,errBusiness := GetBusinessById(business.Id)
	if errBusiness == nil {
		logrus.Error("error adding business with id " + business.Id + " cause: " + errBusiness.Error())
	}
	if retrivedBusiness != nil {
		return false, errors.New("BusinessAlreadyExists: business with id " + business.Id + " already exists ")
	}
	businessDao := domain.Business{
		Id: businessId,
		State: business.State,
		City: business.City,
		Name: business.Name,
		Pincode: business.Pincode,
	}
	stmt, names := qb.Insert(helpers.Table).Columns(helpers.GetColumnNames(business)).ToCql()
	q := gocqlx.Query(repository.GetSessionInstance().Query(stmt), names).BindStruct(businessDao)
	repository.Save(q)
	logrus.Info("record added successfully to table : " + helpers.Table)
	return true, nil
}

func UpdateBusiness(business *domain.BusinessRep) (bool,error){
	businessId, err := gocql.ParseUUID(business.Id)
	if err != nil{
		logrus.Error("Error parsing uuid from string")
		panic("error parsing uuid from string")
	}
	retrivedBusiness,errBusiness := GetBusinessById(business.Id)
	if errBusiness == nil {
		logrus.Error("error updating business with id " + business.Id + " cause: " + errBusiness.Error())
	}
	if retrivedBusiness ==  nil {
		return false, errors.New("BusinessNotFound: business with id " + business.Id + " does not exists ")
	}
	businessDao := domain.Business{
		Id: businessId,
		State: business.State,
		City: business.City,
		Name: business.Name,
		Pincode: business.Pincode,
	}
	logrus.Info("building update query")
	stmt, names:= qb.Update(helpers.Table).Set(helpers.GetColumnNames(business)).Where(qb.Eq("id"),qb.Eq("city"),qb.Eq("state")).ToCql()
	q := gocqlx.Query(repository.GetSessionInstance().Query(stmt),names).BindStruct(businessDao)
	repository.UpdateRecords(q)
	return true, nil
}

func DeleteBusiness(filter map[string]string ) (bool, error){
	logrus.Info("building delete query")
	businesId,err := gocql.ParseUUID(filter["id"])
	if err != nil{
		logrus.Error("Error parsing uuid from string")
		panic("error parsing uuid from string")
	}
	retrivedBusiness,errBusiness := GetBusinessById(filter["id"])
	if errBusiness == nil {
		logrus.Error("error deleting business with id " + filter["id"] + " cause: " + errBusiness.Error())
	}
	if retrivedBusiness ==  nil {
		return false, errors.New("BusinessNotFound: business with id " + filter["id"] + " does not exists ")
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
	return true, nil
}

func GetBusinessData(filter map [string]string) []domain.Business{
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

func GetBusinessById(id string) ([]domain.Business, error){
	logrus.Info("building select query")
	businesId,err := gocql.ParseUUID(id)
	if err != nil{
		logrus.Error("Error parsing uuid from string")
		panic("error parsing uuid from string")
	}
	business := new(domain.Business)
	business.Id = businesId
	stmt, names :=qb.Select(helpers.Table).Columns(helpers.SliceToString(strings.Split(helpers.SelectColumns,","))).
		Where(qb.Eq("id")).ToCql()
	q := gocqlx.Query(repository.GetSessionInstance().Query(stmt),names).BindStruct(business)
	businessRecords := repository.QueryRecords(q)
	return businessRecords,nil
}


