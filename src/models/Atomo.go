package models

import (
	"time"
)

type Atomo struct {
	Id              int
	Name            string
	Typos           int // CONST ARRAY
	Adeies          []Adeia
	YpiresiesAtomou []YpiresiaAtomou
}

func (atomo *Atomo) IsAvailable(d time.Time) bool {
	for _, dates := range atomo.Adeies {
		if d.After(dates.Arxi.AddDate(0, 0, -1)) && d.Before(dates.Telos.AddDate(0, 0, 1)) {
			return false
		}
	}
	return true
}

func (atomo *Atomo) AddAdeia(adeia Adeia) {
	atomo.Adeies = append(atomo.Adeies, adeia)
}

func (atomo *Atomo) AddYpiresia(ypiresia YpiresiaAtomou) {
	atomo.YpiresiesAtomou = append(atomo.YpiresiesAtomou, ypiresia)
}
