package main

import (
	"github.com/easysoft/zendata/src/test/import/comm"
	"github.com/easysoft/zendata/src/test/import/model"
)

func main() {
	filePath := "data/address/cn.v1.xlsx"
	sheetName := "china"

	db := comm.GetDB()

	records := comm.GetExcelTable(filePath, sheetName)

	for _, record := range records {
		po := model.DataCity{
			Name:    record["city"].(string),
			Code:    record["cityCode"].(string),
			ZipCode: record["zipCode"].(string),
			State:   record["state"].(string),
		}

		db.Save(&po)
	}
}
