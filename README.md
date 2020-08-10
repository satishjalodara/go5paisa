# 5paisa Golang SDK

Golang SDK for 5paisa APIs natively written in VB .NET

#### Documentation

GoDoc at https://godoc.org/github.com/5paisa/go5paisa


#### Features

-   Order placement, modification and cancellation
-   Fetching user info including holdings, positions, margin and order book.
-   Fetching order status and trade information.

### Usage

```go
package main

import (
	"fmt"
	"go5paisa/src"
)

func main() {

	conf := &go5paisa.Config{
		AppName:       "YOUR_APP_NAME_HERE",
		AppSource:     "YOUR_APP_SOURCE_HERE",
		UserID:        "YOUR_USER_ID_HERE",
		Password:      "YOUR_PASSWORD_HERE",
		UserKey:       "YOUR_USER_KEY_HERE",
		EncryptionKey: "YOUR_ENCRYPTION_KEY_HERE",
	}
	appConfig := go5paisa.Init(conf)
	client, err := go5paisa.Login(appConfig, "xyz@gmail.com", "password", "YYYYMMDD")
	if err != nil {
		panic(err)
	}
	holdings, err := client.GetHoldings()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", holdings)
}


```


#### TODO
 - Write tests.
 - Update docs

Feel free to contribute and open issues :)
