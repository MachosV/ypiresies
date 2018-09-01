package argies

import (
	"math"
	"time"
)

var argies []time.Time

func IsArgia(date time.Time) bool {
	for _, d := range argies {
		if date == d {
			return true
		}
	}
	if date.Weekday().String() == "Sunday" || date.Weekday().String() == "Saturday" {
		return true
	}
	return false
}

//initialize resting days
func InitArgies(year int) {
	pasxa := getEasterDate(year)
	argies = append(argies, time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC))
	argies = append(argies, time.Date(year, 1, 6, 0, 0, 0, 0, time.UTC))
	argies = append(argies, time.Date(year, 3, 25, 0, 0, 0, 0, time.UTC))
	argies = append(argies, time.Date(year, 5, 1, 0, 0, 0, 0, time.UTC))
	argies = append(argies, time.Date(year, 8, 15, 0, 0, 0, 0, time.UTC))
	argies = append(argies, time.Date(year, 10, 28, 0, 0, 0, 0, time.UTC))
	argies = append(argies, time.Date(year, 12, 25, 0, 0, 0, 0, time.UTC))
	argies = append(argies, time.Date(year, 12, 26, 0, 0, 0, 0, time.UTC))
	argies = append(argies, pasxa) //pasxa
	argies = append(argies, getKatharaDeytera(pasxa))
	argies = append(argies, getMegaliParaskeyi(pasxa))
	argies = append(argies, getDeyteraPasxa(pasxa))
	argies = append(argies, getAgiouPneymatos(pasxa))
}

func getEasterDate(year int) time.Time {
	var a, b, c, d, e int
	var month time.Month
	var day float64

	a = year % 4
	b = year % 7
	c = year % 19
	d = (19*c + 15) % 30
	e = (2*a + 4*b - d + 34) % 7
	month = time.Month((d + e + 114) / 31)
	day = math.Floor(float64((d+e+114)%31 + 1))
	day = day + 13

	return time.Date(year, month, int(day), 0, 0, 0, 0, time.UTC)
}

func getKatharaDeytera(pasxa time.Time) time.Time {
	return pasxa.Add(-48 * 24 * time.Hour)
}

func getMegaliParaskeyi(pasxa time.Time) time.Time {
	return pasxa.Add(-2 * 24 * time.Hour)
}

func getDeyteraPasxa(pasxa time.Time) time.Time {
	return pasxa.Add(1 * 24 * time.Hour)
}

func getAgiouPneymatos(pasxa time.Time) time.Time {
	return pasxa.Add(50 * 24 * time.Hour)
}
