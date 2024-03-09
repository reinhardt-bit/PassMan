package main

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
)

func main() {
	// Generate a new TOTP key
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "YourAppName",
		AccountName: "user@example.com",
	})
	if err != nil {
		panic(err)
	}

	// Print the TOTP key details
	fmt.Printf("TOTP Key:\n")
	fmt.Printf("  Issuer:       %s\n", key.Issuer())
	fmt.Printf("  Account Name: %s\n", key.AccountName())
	fmt.Printf("  Secret:       %s\n", key.Secret())
	fmt.Printf("  URL:          %s\n", key.URL())

	// // Prompt the user to enter the TOTP code
	// var code string
	// fmt.Print("Enter TOTP Code: ")
	// fmt.Scanln(&code)
	//
	// // Validate the entered TOTP code
	// valid := totp.Validate(code, key.Secret())
	// if valid {
	// 	fmt.Println("TOTP code is valid!")
	// } else {
	// 	fmt.Println("TOTP code is invalid.")
	// }
	//
	// Generate OTP codes every 30 seconds
	// for {
	// 	code, remainingSeconds := totp.GenerateCode(key.Secret(), time.Now())
	// 	fmt.Println(time.Now())
	// 	fmt.Println(remainingSeconds)
	// 	// fmt.Printf("\rCurrent OTP: %s (Remaining: %ds)", code, remainingSeconds)
	// 	fmt.Printf("\rCurrent OTP: %s (Remaining: %v)", code, remainingSeconds)
	// 	time.Sleep(1 * time.Second)
	// }
	// Generate OTP codes every 30 seconds
	// for {
	// 	code, expiresAt := totp.GenerateCode(key.Secret(), time.Now())
	// 	// remainingSeconds := int(expiresAt.Sub(time.Now()).Seconds())
	// 	fmt.Printf("this is sub: %v and this is expiresAt: %v and this is code: %v", time.Now().Sub(time.)), expiresAt, code)

	// for remainingSeconds > 0 {
	// 	fmt.Printf("\rCurrent OTP: %s (Remaining: %02ds)", code, remainingSeconds)
	// 	time.Sleep(1 * time.Second)
	// 	remainingSeconds--
	// }
	// Generate OTP codes every 30 seconds
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		code, _ := totp.GenerateCode(key.Secret(), time.Now())
		remainingSeconds := 30

		fmt.Printf("Current OTP: %s (Remaining: %02ds)\n", code, remainingSeconds)

		secondTicker := time.NewTicker(1 * time.Second)
		defer secondTicker.Stop()

		for {
			select {
			case <-secondTicker.C:
				remainingSeconds--
				fmt.Printf("\rCurrent OTP: %s (Remaining: %02ds)", code, remainingSeconds)
				if remainingSeconds == 0 {
					fmt.Println()
					break
				}
			case <-ticker.C:
				fmt.Println()
				goto NextCode
			}
		}

	NextCode:
	}
	// }
}

// 	// Generate OTP codes every 30 seconds
// 	generateOTPCodes(key.Secret(), 30)
// }
//
// func generateOTPCodes(secret string, interval ...int) {
// 	duration := 30
// 	if len(interval) > 0 {
// 		duration = interval[0]
// 	}
//
// 	ticker := time.NewTicker(time.Duration(duration) * time.Second)
// 	defer ticker.Stop()
//
// 	for {
// 		code, _ := totp.GenerateCode(secret, time.Now())
// 		remainingSeconds := duration
//
// 		fmt.Printf("Current OTP: %s (Remaining: %02ds)\n", code, remainingSeconds)
//
// 		secondTicker := time.NewTicker(1 * time.Second)
// 		defer secondTicker.Stop()
//
// 		for {
// 			select {
// 			case <-secondTicker.C:
// 				remainingSeconds--
// 				fmt.Printf("\rCurrent OTP: %s (Remaining: %02ds)", code, remainingSeconds)
// 				if remainingSeconds == 0 {
// 					fmt.Println()
// 					// break
// 					// // totp.GenerateCode(secret, time.Now())
// 				}
// 			case <-ticker.C:
// 				fmt.Println()
// 				return
// 			}
// 		}
// 	}
// }
