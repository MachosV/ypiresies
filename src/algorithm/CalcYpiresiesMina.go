package algorithm

import (
	"argies"
	"datastorage"
	"models"
	"time"
)

func calcYpiresiesMina(date time.Time) map[time.Time][]models.YpiresiaImeras {
	var ypiresiesMina map[time.Time][]models.YpiresiaImeras
	ypiresiesMina = make(map[time.Time][]models.YpiresiaImeras)
	var ypiresiesImeras []models.YpiresiaImeras
	var ypiresiaImeras models.YpiresiaImeras
	ypiresiaImeras.Atomo = nil
	for d := date; d.Month() == date.Month(); d = d.AddDate(0, 0, 1) {
		if argies.IsArgia(d) {
			for index, ypiresia := range datastorage.Ypiresies {
				if ypiresia.Typos == datastorage.KATH_ARGIA || ypiresia.Typos == datastorage.ARGIA {
					ypiresiaImeras.Ypiresia = &datastorage.Ypiresies[index]
					ypiresiesImeras = append(ypiresiesImeras, ypiresiaImeras)
				}
			}

		} else {
			for index, ypiresia := range datastorage.Ypiresies {
				if ypiresia.Typos == datastorage.KATHIMERINI || ypiresia.Typos == datastorage.KATH_ARGIA {
					ypiresiaImeras.Ypiresia = &datastorage.Ypiresies[index]
					ypiresiesImeras = append(ypiresiesImeras, ypiresiaImeras)
				}
			}
		}
		ypiresiesMina[d] = ypiresiesImeras
		ypiresiesImeras = nil
	}
	return ypiresiesMina
}
