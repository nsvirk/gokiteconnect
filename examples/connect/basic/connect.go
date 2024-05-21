package main

import (
	"fmt"
	"os"

	kiteconnect "github.com/nsvirk/gokiteconnect/v4"
	kitesession "github.com/nsvirk/gokitesession"
)

var (
	userId     string = os.Getenv("KITE_USER_ID")
	password   string = os.Getenv("KITE_PASSWORD")
	totpSecret string = os.Getenv("KITE_TOTP_SECRET")
)

func main() {
	// Create a new Kite session instance
	ks := kitesession.New(userId)

	// Set debug mode
	ks.SetDebug(true)

	// Generate totp value
	totpValue, err := ks.GenerateTotpValue(totpSecret)
	if err != nil {
		fmt.Printf("Error generating totp value: %v", err)
		return
	}

	// Check the inputs values
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Kite User")
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("User ID     	: ", userId)
	fmt.Println("Password     	: ", password)
	fmt.Println("Totp Value  	: ", totpValue)
	fmt.Println("")

	// Get kite session data
	session, err := ks.GenerateSession(password, totpValue)
	if err != nil {
		fmt.Printf("Error generating session: %v", err)
		return
	}

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Kite Session")
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("user_id     	: ", session.UserId)
	fmt.Println("public_token	: ", session.PublicToken)
	fmt.Println("kf_session   	: ", session.KfSession)
	fmt.Println("enctoken    	: ", session.Enctoken)
	fmt.Println("login_time  	: ", session.LoginTime)
	fmt.Println("username   	: ", session.Username)
	fmt.Println("user_shortname	: ", session.UserShortname)
	fmt.Println("avatar_url  	: ", session.AvatarURL)
	fmt.Println("")
	// fmt.Println(session)

	// Get the  enctoken
	enctoken := session.Enctoken

	// Create a new Kite connect instance
	kc := kiteconnect.New(userId)

	// Set enctoken
	kc.SetEnctoken(enctoken)

	// Get user profile
	profile, err := kc.GetUserProfile()
	if err != nil {
		fmt.Printf("Error getting user profile: %v", err)
	}

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("User Profile")
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("user_id 	: ", profile.UserID)
	fmt.Println("username 	: ", profile.UserName)
	fmt.Println("user_short	: ", profile.UserShortName)
	fmt.Println("avatar_url	: ", profile.AvatarURL)
	fmt.Println("user_type	: ", profile.UserType)
	fmt.Println("email		: ", profile.Email)
	fmt.Println("broker		: ", profile.Broker)
	fmt.Println("meta		: ", profile.Meta)
	fmt.Println("products	: ", profile.Products)
	fmt.Println("order_type	: ", profile.OrderTypes)
	fmt.Println("exchanges	: ", profile.Exchanges)
	fmt.Println("")
	fmt.Println("--------------------------------------------------------------")
	// fmt.Println("profile: ", profile)

}
