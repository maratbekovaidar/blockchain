package main

import (
	"blockchain/utils"
	"fmt"
)

func main() {

	//fmt.Println(utils.GetHost())
	u := utils.User{
		Name:       "Aidar Maratbekov",
		Email:      "maratbekovaidar",
		Field:      "Flutter Developer",
		Reputation: "100",
		Optional: map[string]string{
			"experience": "1",
		},
	}
	fileCid := u.UploadUserToIPFS()
	fmt.Printf(fileCid)
}
