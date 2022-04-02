package cmd

import (
	"fmt"
	"time"
)

const fmtRFC3339 = time.RFC3339

func InAndParseInLocation() int {
	fmt.Println(time.Local)
	nowUTC := time.Now().UTC()
	fmt.Println(nowUTC.Format(fmtRFC3339))

	fmt.Println(nowUTC.In(time.Local).Format(fmtRFC3339))

	local, err := time.ParseInLocation(fmtRFC3339, nowUTC.Format(fmtRFC3339), time.Local)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	fmt.Println(local.Format(fmtRFC3339))

	fmt.Println(local.UTC().Format(fmtRFC3339))

	return 0
}
