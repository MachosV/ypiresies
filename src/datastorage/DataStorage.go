package datastorage

import (
	"argies"
	"log"
	"math/rand"
	"models"
	"time"
	"webstorage"
)

var Ypiresies []models.Ypiresia
var Metafrastes []models.Atomo
var Akroates []models.Atomo
var Stratiwtes []models.Atomo
var Year int

const (
	STRATIWTIS  = 0
	METAFRASTIS = 1
	AKROATIS    = 2
	KATHIMERINI = 3
	ARGIA       = 4
	KATH_ARGIA  = 5
)

func InitAll(date time.Time) {
	Year = date.Year()
	argies.InitArgies(Year)
	Akroates = nil
	Metafrastes = nil
	Stratiwtes = nil
	Akroates = LoadAtoma(AKROATIS)
	Metafrastes = LoadAtoma(METAFRASTIS)
	Stratiwtes = LoadAtoma(STRATIWTIS)
	Ypiresies = LoadYpiresies()
	InitAdeies()
	ShuffleAll()
}

func LoadAtoma(typos int) []models.Atomo {
	var adeies []models.Adeia
	var ypiresiesAtomou []models.YpiresiaAtomou
	var atomo models.Atomo
	var atoma []models.Atomo
	db := webstorage.GetDb()
	res, err := db.Query("SELECT id,onoma FROM proswpiko where typos = ? ORDER BY onoma", typos)
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		err = res.Scan(
			&atomo.Id,
			&atomo.Name)
		if err != nil {
			log.Fatal(err)
		}
		atomo.Typos = typos
		atomo.Adeies = adeies
		atomo.YpiresiesAtomou = ypiresiesAtomou
		atoma = append(atoma, atomo)
	}
	res.Close()
	return atoma
}

func LoadYpiresies() []models.Ypiresia {
	var ypiresia models.Ypiresia
	var ypiresies []models.Ypiresia
	db := webstorage.GetDb()
	res, err := db.Query("SELECT id,perigrafi,proswpiko,typos,color FROM ypiresia;")
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		err = res.Scan(
			&ypiresia.Id,
			&ypiresia.Perigrafi,
			&ypiresia.Proswpiko,
			&ypiresia.Typos,
			&ypiresia.Color)
		if err != nil {
			log.Fatal(err)
		}
		ypiresies = append(ypiresies, ypiresia)
	}
	return ypiresies
}

func InitAdeies() {
	for index, _ := range Akroates {
		LoadAdeies(&Akroates[index])
	}
	for index, _ := range Metafrastes {
		LoadAdeies(&Metafrastes[index])
	}
	for index, _ := range Stratiwtes {
		LoadAdeies(&Stratiwtes[index])
	}
}

func LoadAdeies(atomo *models.Atomo) {
	var adeia models.Adeia
	var arxi, telos int64
	db := webstorage.GetDb()
	res, err := db.Query("SELECT aitia,arxi,telos from adeies where atomo=?", atomo.Id)
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		res.Scan(
			&adeia.Aitia,
			&arxi,
			&telos)
		adeia.Arxi = time.Unix(arxi, 0)
		adeia.Telos = time.Unix(telos, 0)
		atomo.AddAdeia(adeia)
	}
}

func Shuffle(slc []models.Atomo) {
	rand.Seed(time.Now().UTC().UnixNano())
	N := len(slc)
	for i := 0; i < N; i++ {
		// choose index uniformly in [i, N-1]
		r := i + rand.Intn(N-i)
		slc[r], slc[i] = slc[i], slc[r]
	}
}

func ShuffleAll() {
	Shuffle(Akroates)
	Shuffle(Metafrastes)
	Shuffle(Stratiwtes)
}
