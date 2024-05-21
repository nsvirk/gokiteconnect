# The Kite Connect API Go client [Custom version]

The customised Go client for communicating with the Kite Connect API.

Kite Connect is a set of REST-like APIs that expose many capabilities required
to build a complete investment and trading platform. Execute orders in real
time, manage user portfolio, stream live market data (WebSockets), and more,
with the simple HTTP API collection.

## Documentation

- [Client API documentation - GoDoc](https://godoc.org/github.com/nsvirk/gokiteconnect)

## Installation

```
go get github.com/nsvirk/gokiteconnect/v4
```

## Kite connect usage

```go
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
```

## Kite ticker usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	kiteconnect "github.com/nsvirk/gokiteconnect/v4"
	kitemodels "github.com/nsvirk/gokiteconnect/v4/models"
	kiteticker "github.com/nsvirk/gokiteconnect/v4/ticker"
	kitesession "github.com/nsvirk/gokitesession"
)

var (
	userId     string   = getEnv("KITE_USER_ID", "my_user_id")
	password   string   = getEnv("KITE_PASSWORD", "my_password")
	totpSecret string   = getEnv("KITE_TOTP_SECRET", "my_totp_secret")
	instTokens []uint32 = append([]uint32{}, 256265, 264969, 5633, 779521, 408065, 738561, 895745)
)

var (
	ticker *kiteticker.Ticker
)

// Triggered when any error is raised
func onError(err error) {
	fmt.Println("Error: ", err)
}

// Triggered when websocket connection is closed
func onClose(code int, reason string) {
	fmt.Println("Close: ", code, reason)
}

// Triggered when connection is established and ready to send and accept data
func onConnect() {
	fmt.Println("Connected")
	fmt.Println("Subscribing to", instTokens)
	err := ticker.Subscribe(instTokens)
	if err != nil {
		fmt.Println("err: ", err)
	}
	// Set subscription mode for given list of tokens
	// Default mode is Quote
	err = ticker.SetMode(kiteticker.ModeFull, instTokens)
	if err != nil {
		fmt.Println("err: ", err)
	}
}

// Triggered when tick is recevived
func onTick(tick kitemodels.Tick) {
	// fmt.Println("Tick: ", tick)
	tickJson, err := json.Marshal(tick)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(tickJson))
	fmt.Println("-------------------------------------------")
}

// Triggered when reconnection is attempted which is enabled by default
func onReconnect(attempt int, delay time.Duration) {
	fmt.Printf("Reconnect attempt %d in %fs\n", attempt, delay.Seconds())
}

// Triggered when maximum number of reconnect attempt is made and the program is terminated
func onNoReconnect(attempt int) {
	fmt.Printf("Maximum no of reconnect attempt reached: %d", attempt)
}

// Triggered when order update is received
func onOrderUpdate(order kiteconnect.Order) {
	fmt.Printf("Order: %s", order.OrderID)
}

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

	// Get kite session data
	session, err := ks.GenerateSession(password, totpValue)
	if err != nil {
		fmt.Printf("Error generating session: %v", err)
		return
	}

	// Create new Kite ticker instance
	ticker = kiteticker.New(userId, session.Enctoken)

	// Assign callbacks
	ticker.OnError(onError)
	ticker.OnClose(onClose)
	ticker.OnConnect(onConnect)
	ticker.OnReconnect(onReconnect)
	ticker.OnNoReconnect(onNoReconnect)
	ticker.OnTick(onTick)
	ticker.OnOrderUpdate(onOrderUpdate)

	// Start the connection
	ticker.Serve()
}

// getEnv returns the value of the environment variable provided.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

```

## Examples

Check [examples folder](https://github.com/nsvirk/gokiteconnect/tree/master/examples) for more examples.

You can run the following after updating the API Keys in the examples:

## Run unit tests

```
go test -v
```
