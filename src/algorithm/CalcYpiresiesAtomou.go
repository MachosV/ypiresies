package algorithm

import (
	"datastorage"
	"models"
	"time"
)

func CalcYpiresiesAtomou(ypiresiesMina map[time.Time][]models.YpiresiaImeras) {
	for key, ypiresiesImeras := range ypiresiesMina {
		for _, ypiresiaImeras := range ypiresiesImeras {
			var ypiresiaAtomou models.YpiresiaAtomou
			ypiresiaAtomou.Date = key
			ypiresiaAtomou.OnomaYpiresias = ypiresiaImeras.Ypiresia.Perigrafi
			ypiresiaAtomou.Color = ypiresiaImeras.Ypiresia.Color
			if ypiresiaImeras.Atomo.Typos == datastorage.AKROATIS {
				index := GetIndexOf(datastorage.Akroates, ypiresiaImeras.Atomo.Name)
				(&datastorage.Akroates[index]).AddYpiresia(ypiresiaAtomou)
				continue
			}
			if ypiresiaImeras.Atomo.Typos == datastorage.METAFRASTIS {
				index := GetIndexOf(datastorage.Metafrastes, ypiresiaImeras.Atomo.Name)
				(&datastorage.Metafrastes[index]).AddYpiresia(ypiresiaAtomou)
				continue
			}
			if ypiresiaImeras.Atomo.Typos == datastorage.STRATIWTIS {
				index := GetIndexOf(datastorage.Stratiwtes, ypiresiaImeras.Atomo.Name)
				(&datastorage.Stratiwtes[index]).AddYpiresia(ypiresiaAtomou)
				continue
			}
		}
	}
}

func GetIndexOf(array []models.Atomo, name string) int {
	for index, x := range array {
		if x.Name == name {
			return index
		}
	}
	return -1
}
