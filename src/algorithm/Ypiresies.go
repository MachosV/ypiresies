package algorithm

import (
	"datastorage"
	"excel"
	"time"
)

func Algorithm(date time.Time) {
	datastorage.InitAll(date)
	ypiresiesMina := calcYpiresiesMina(date)
	CalcAtomaYpiresiwn(ypiresiesMina, date) //complete
	CalcYpiresiesAtomou(ypiresiesMina)
	excel.BuildExcel(date)
}
