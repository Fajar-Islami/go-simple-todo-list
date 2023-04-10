package utils

import (
	"fmt"
	"log"
	"time"
)

const (
	mmLayout       = "January"
	yyyyLayout     = "2006"
	ddmmyyyyLayout = "02/01/2006"
	mmyyyyLayout   = "January 2006"
	formatteddated = "2006-01-02T15:04:05.000Z"
)

func ShortDateFromString(ds string) (time.Time, error) {
	t, err := time.Parse(ddmmyyyyLayout, ds)
	if err != nil {
		return t, fmt.Errorf("Data format must be DD/MM/YYYY")
	}
	return t, nil
}

func ShortDateFromDate(date time.Time) string {
	return date.Format(ddmmyyyyLayout)
}

func DateFormatter(date time.Time) time.Time {
	res, err := time.Parse(formatteddated, date.Format(formatteddated))
	if err != nil {
		log.Println(err)
	}
	return res
}
