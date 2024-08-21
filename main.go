package main

import (
	"app/lasagna"
	"fmt"
)

func main() {

	fmt.Print("Input OvenTime:")

	// รับค่าจำนวนนาทีที่ต้องการให้เตาอบ
	var ovenTime int
	fmt.Scanln(&ovenTime)

	// กำหนดฟังก์ชัน RemainingOvenTime ที่รับพารามิเตอร์สองตัว
	// 1. พารามิเตอร์แรกคือเวลาที่ต้องการให้เตาอบ
	// 2. พารามิเตอร์ตัวที่สองคือเวลาที่เอาลาซานญ่าออกจากเตาอบ
	// 3. return เวลาที่เหลือในการเตาอบ
	remaingTime := lasagna.RemainingOvenTime(ovenTime, 30)
	fmt.Println("Remaining time in oven:", remaingTime, " minutes")

	// กำหนดฟังก์ชัน PreparationTime
	// 1. กรอกจำนวนชั้นของลาซานญ่าเป็นพารามิเตอร์
	// 2. แต่ละชั้นใช้เวลา 2 นาทีในการเตรียม
	// 3. return จำนวนเวลาทั้งหมดที่ใช้
	prepareTime := lasagna.PreparationTime(2)
	fmt.Println("Preparation time:", prepareTime, " minutes")

	// กำหนดฟังก์ชัน ElapsedTime ที่รับพารามิเตอร์สองตัว
	// 1. พารามิเตอร์แรกคือจำนวนชั้นที่คุณเพิ่มลงในลาซานญ่า
	// 2. พารามิเตอร์ตัวที่สองคือจำนวนนาทีที่ลาซานญ่าอยู่ในเตาอบ
	// 3. return จำนวนเวลาที่ใช้ทั้งหมด
	elapsedTime := lasagna.ElapsedTime(3, 20)
	fmt.Println("Elapsed time:", elapsedTime, " minutes")

}
