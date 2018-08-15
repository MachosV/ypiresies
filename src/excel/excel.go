package excel

import (
	"datastorage"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

func BuildExcel(date time.Time) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	_ = row
	var cell *xlsx.Cell
	var err error
	excelFileName := "Υπηρεσίες.xlsx"
	sheetName := (date.Month().String() + " " + strconv.Itoa(datastorage.Year))
	file, err = xlsx.OpenFile(excelFileName)
	if err != nil {
		file = xlsx.NewFile()
		sheet, err = file.AddSheet(sheetName)
		file.Save(excelFileName)
	}
	sheet, err = file.AddSheet((date.Month().String() + " " + strconv.Itoa(datastorage.Year)))
	if err != nil {
		delete(file.Sheet, sheetName)
		for index, filesheet := range file.Sheets {
			if filesheet.Name == sheetName {
				sheet = &xlsx.Sheet{
					Name:     sheetName,
					File:     file,
					Selected: len(file.Sheets) == 0,
				}
				file.Sheets[index] = sheet
			}
		}
		file.Sheet[sheetName] = sheet
	}
	sort.Slice(datastorage.Metafrastes, func(i, j int) bool {
		return datastorage.Metafrastes[i].Name < datastorage.Metafrastes[j].Name
	})
	sort.Slice(datastorage.Akroates, func(i, j int) bool {
		return datastorage.Akroates[i].Name < datastorage.Akroates[j].Name
	})
	sort.Slice(datastorage.Stratiwtes, func(i, j int) bool {
		return datastorage.Stratiwtes[i].Name < datastorage.Stratiwtes[j].Name
	})
	row = sheet.AddRow()
	cell = row.AddCell()
	for d := date; d.Month() == date.Month(); d = d.AddDate(0, 0, 1) {
		cell = row.AddCell()
		if d.Weekday().String() == "Sunday" {
			style := GetStyle("grey")
			cell.SetStyle(style)
		}
		cell.Value = strconv.Itoa(d.Day())
	}
	for _, atomo := range datastorage.Metafrastes {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = atomo.Name
		for d := date; d.Month() == date.Month(); d = d.AddDate(0, 0, 1) {
			cell = row.AddCell()
		}
		for _, ypiresia := range atomo.YpiresiesAtomou {
			cell = row.Cells[ypiresia.Date.Day()]
			cell.Value = ypiresia.OnomaYpiresias
		}
	}
	for _, atomo := range datastorage.Akroates {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = atomo.Name
		for d := date; d.Month() == date.Month(); d = d.AddDate(0, 0, 1) {
			cell = row.AddCell()
		}
		for _, ypiresia := range atomo.YpiresiesAtomou {
			cell = row.Cells[ypiresia.Date.Day()]
			cell.Value = ypiresia.OnomaYpiresias
		}
	}
	for _, atomo := range datastorage.Stratiwtes {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = atomo.Name
		for d := date; d.Month() == date.Month(); d = d.AddDate(0, 0, 1) {
			cell = row.AddCell()
		}
		for _, ypiresia := range atomo.YpiresiesAtomou {
			cell = row.Cells[ypiresia.Date.Day()]
			cell.Value = ypiresia.OnomaYpiresias
		}
	}

	err = file.Save(excelFileName)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
