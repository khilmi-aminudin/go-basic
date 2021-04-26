package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// var time1 = time.Now()
// 	// fmt.Println(time1)

// 	// var timeDate = time.Date(time.Now().Year(),3,5,15,33,60,7,time.UTC)
// 	// fmt.Println(timeDate)

// 	// Time Parsing
// 	var layoutFormat, value string
// 	var date time.Time

// 	layoutFormat = "2006-01-02 15:04:05"
// 	value = "2015-09-02 08:04:00"
// 	date, _ = time.Parse(layoutFormat, value)
// 	fmt.Println(value, "\t->", date.String())
// 	// 2015-09-02 08:04:00 +0000 UTC

// 	layoutFormat = "02/01/2006 MST"
// 	value = "02/09/2015 WIB"
// 	date, _ = time.Parse(layoutFormat, value)
// 	fmt.Println(value, "\t\t->", date.String())
// 	// 2015-09-02 00:00:00 +0700 WIB

// }

// 	// var now = time.Now()
// 	// fmt.Println(now.Year())

// 	/*
// 		+-----------------------+-------------------+-----------------------------------------------------------------------------------------------------------------------------------+
// 		|		Method			|	Return Type		|											Penjelasan																				|
// 		+-----------------------+-------------------+-----------------------------------------------------------------------------------------------------------------------------------+
// 		|	now.Year()			| int				| Tahun																																|
// 		|	now.YearDay()		| int				| Hari ke-? di mulai awal tahun																										|
// 		|	now.Month()			| int				| Bulan																																|
// 		|	now.Weekday()		| string			| Nama hari. Bisa menggunakan now.Weekday().String() untuk mengambil bentuk string-nya												|
// 		|	now.ISOWeek()		| (int, int)		| Tahun dan minggu ke-? mulai awal tahun																							|
// 		|	now.Day()			| int				| Tanggal																															|
// 		|	now.Hour()			| int				| Jam																																|
// 		|	now.Minute()		| int				| Menit																																|
// 		|	now.Second()		| int				| Detik																																|
// 		|	now.Nanosecond()	| int				| Nano detik																														|
// 		|	now.Local()			| time.Time			| Date-time dalam timezone lokal																									|
// 		|	now.Location()		| *time.Location	| Mengambil informasi lokasi, apakah local atau utc. Bisa menggunakan now.Location().String() untuk mengambil bentuk string-nya		|
// 		|	now.Zone()			| (string, int)		| Mengembalikan informasi timezone offset dalam string dan numerik. Sebagai contoh WIB, 25200										|
// 		|	now.IsZero()		| bool				| Deteksi apakah nilai object now adalah 01 Januari tahun 1, 00:00:00 UTC. Jika iya maka bernilai true								|
// 		|	now.UTC()			| time.Time			| Date-time dalam timezone UTC																										|
// 		|	now.Unix()			| int64				| Date-time dalam format unix time																									|
// 		|	now.UnixNano()		| int64				| Date-time dalam format unix time. Infomasi nano detik juga dimasukkan																|
// 		|	now.String()		| string			| Date-time dalam string																											|
// 		+-----------------------+-------------------+-----------------------------------------------------------------------------------------------------------------------------------+
// 	*/
