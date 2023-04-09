package utils

import (
	"fmt"
	"time"
)

const (
	mmLayout       = "January"
	yyyyLayout     = "2006"
	ddmmyyyyLayout = "02/01/2006"
	mmyyyyLayout   = "January 2006"
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
