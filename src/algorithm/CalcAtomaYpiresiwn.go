package algorithm

import (
	"datastorage"
	"models"
	"time"
)

func CalcAtomaYpiresiwn(ypiresiesMina map[time.Time][]models.YpiresiaImeras, date time.Time) {
	GetNextStr := IntFuncStratiwti()
	GetNextMtr := IntFuncMetafrasti()
	GetNextAkr := IntFuncAkroati()
	for d := date; d.Month() == date.Month(); d = d.AddDate(0, 0, 1) {
		for index, ypiresiaImeras := range ypiresiesMina[d] {
			if ypiresiaImeras.Ypiresia.Proswpiko == datastorage.STRATIWTIS {
				ypiresiesMina[d][index].Atomo = GetNextStr(d)
			}
			if ypiresiaImeras.Ypiresia.Proswpiko == datastorage.METAFRASTIS {
				ypiresiesMina[d][index].Atomo = GetNextMtr(d)
			}
			if ypiresiaImeras.Ypiresia.Proswpiko == datastorage.AKROATIS {
				ypiresiesMina[d][index].Atomo = GetNextAkr(d)
			}
		}
	}
}

func IntFuncStratiwti() func(d time.Time) *models.Atomo {
	deiktis := 0
	var previous, cur *models.Atomo
	previous, cur = nil, nil
	length := len(datastorage.Stratiwtes)
	return func(d time.Time) *models.Atomo {
		for {
			deiktis = deiktis % length
			cur = &datastorage.Stratiwtes[deiktis]
			if cur.IsAvailable(d) {
				if cur != previous {
					deiktis = deiktis + 1
					previous = cur
					return cur
				} else {
					previous = nil
					return nil
				}
			} else {
				deiktis = deiktis + 1
			}
		}
	}
}

func IntFuncMetafrasti() func(d time.Time) *models.Atomo {
	deiktis := 0
	var previous, cur *models.Atomo
	previous, cur = nil, nil
	length := len(datastorage.Metafrastes)
	return func(d time.Time) *models.Atomo {
		for {
			deiktis = deiktis % length
			cur = &datastorage.Metafrastes[deiktis]
			if cur.IsAvailable(d) {
				if cur != previous {
					deiktis = deiktis + 1
					previous = cur
					return cur
				} else {
					previous = nil
					return nil
				}
			} else {
				deiktis = deiktis + 1
			}
		}
	}
}

func IntFuncAkroati() func(d time.Time) *models.Atomo {
	deiktis := 0
	var previous, cur *models.Atomo
	previous, cur = nil, nil
	length := len(datastorage.Akroates)
	return func(d time.Time) *models.Atomo {
		for {
			deiktis = deiktis % length
			cur = &datastorage.Akroates[deiktis]
			if cur.IsAvailable(d) {
				if cur != previous {
					deiktis = deiktis + 1
					previous = cur
					return cur
				} else {
					previous = nil
					return nil
				}
			} else {
				deiktis = deiktis + 1
			}
		}
	}
}
