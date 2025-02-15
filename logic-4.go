package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan jumlah cuti bersama: ")
	scanner.Scan()
	sharedLeave, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Masukkan tanggal join karyawan (YYYY-MM-DD): ")
	scanner.Scan()
	joinDate := scanner.Text()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Masukkan tanggal cuti (YYYY-MM-DD): ")
	scanner.Scan()
	leaveDate := scanner.Text()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Masukan durasi cuti : ")
	scanner.Scan()
	durationLeave, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	result, reason := ValidasiCuti(joinDate, leaveDate, sharedLeave, durationLeave)
	fmt.Println(result)
	fmt.Println(reason)
}

func ValidasiCuti(strJoinDate string, strLeaveDate string, sharedLeave int, durationLeave int) (bool, string) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println(err)
		return false, err.Error()
	}

	layout := "2006-01-02"
	joinDate, err := time.ParseInLocation(layout, strJoinDate, loc)
	if err != nil {
		fmt.Println(err)
		return false, err.Error()
	}
	leaveDate, err := time.ParseInLocation(layout, strLeaveDate, loc)
	if err != nil {
		fmt.Println(err)
		return false, err.Error()
	}

	allowedStartDate := joinDate.AddDate(0, 0, 180)
	if leaveDate.Before(allowedStartDate) {
		return false, "Alasan: Karena belum 180 hari sejak tanggal join karyawan"
	}

	endOfYear := time.Date(joinDate.Year(), 12, 31, 0, 0, 0, 0, loc)
	if leaveDate.Year() > joinDate.Year() {
		endOfYear = time.Date(leaveDate.Year(), 12, 31, 0, 0, 0, 0, loc)
	}

	remainingDays := int(endOfYear.Sub(allowedStartDate).Hours() / 24)

	employeeLeave := (remainingDays * (14 - sharedLeave)) / 365

	if durationLeave > 3 {
		return false, "Alasan : Cuti pribadi max 3 hari berturut-turut"
	}

	if durationLeave > employeeLeave {
		return false, fmt.Sprintf("Alasan : Karena hanya boleh mengambil %d hari cuti", employeeLeave)
	}

	return true, ""
}
