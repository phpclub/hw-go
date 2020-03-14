package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	const ctLabel = "current time:"
	const ntpLabel = "exact time:"
	const ntpHost = "0.beevik-ntp.pool.ntp.org"
	cTime := time.Now()
	//Зададим формат - каким надо быть альтернативно одаренным чтоб придумать такое
	cFormattedTime := cTime.UTC().Format("2006-01-02 15:04:05 -0700 MST")
	fmt.Println(ctLabel, cFormattedTime)
	// Получим сетевое время из ntpHost
	ntpTime, ntpErr := ntp.Time(ntpHost)
	if ntpErr != nil {
		log.Fatalln("NTP Error:", ntpErr)
	}
	cFormattedNtpTime := ntpTime.UTC().Format("2006-01-02 15:04:05 -0700 MST")
	fmt.Println(ntpLabel, cFormattedNtpTime)
}
