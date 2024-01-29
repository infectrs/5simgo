
## Usage

### Client

Firstly, initialize the 5Sim client with:

```go
package main

import (
    "fmt"
    "github.com/infectrs/5simgo"
)

func main() {
	apiKey := "ey..."
	sim5Client, err := simgo5.New5Sim(apiKey)

	if err != nil {
		fmt.Println(err)
		return
	}
}
```

Then you can use the following methods:

### Retrieve Profile Information
```go
profile, err := sim5Client.GetProfileInformation()

if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("Profile Information: %v", profile)
```

### Retrieve Order History
```go
var category string = "" // hosting or activation

ordersHistory, err := sim5Client.GetOrdersHistory(category)

if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("Orders History: %v", ordersHistory)
```

### Retrieve Payments History
```go
paymentsHistory, err := sim5Client.GetPaymentsHistory()

if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("Payments History: %v", paymentsHistory)
```

### Purchase an activation phone number
```go
var (
	country string = "any"
	operator string = "any"
	product string = "telegram"
) // for more, visit 5sim website

purchasedActivationNumberInfo, err := sim5Client.BuyActivationNumber(country, operator, product)

if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("Purchased Activation Number Information: %v", purchasedActivationNumberInfo)
```

### Purchase an hosting phone number
```go
var (
	country string = "any"
	operator string = "any"
	product string = "telegram"
) // for more, visit 5sim website

purchasedHostingNumberInfo, err := sim5Client.BuyHostingNumber(country, operator, product)

if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("Purchased Hosting Number Information: %v", purchasedHostingNumberInfo)
```

### Rebuy a phone number
```go
var (
	product string = "telegram" // for more, visit 5sim website
	phoneNumber string = "" // 4-15 digits (without +)
)
	
if err := sim5Client.RebuyNumber(product, phoneNumber); err != nil {
	fmt.Println(err)
	return
}
```

### Check an SMS Order
```go
var orderId string = "123456"

smsInformation, err := sim5Client.CheckOrder(orderId)

if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("SMS Order Information: %v", smsInformation)
```

### Finish an SMS Order
```go
var orderId string = "123456"

smsInformation, err := sim5Client.FinishOrder(orderId)

if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("SMS Order Information: %v", smsInformation)
```

### Cancel an SMS Order
```go
var orderId string = "123456"

smsInformation, err := sim5Client.CancelOrder(orderId)

if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("SMS Order Information: %v", smsInformation)
```

### Ban an SMS Order
```go
var orderId string = "123456"

smsInformation, err := sim5Client.BanOrder(orderId)

if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("SMS Order Information: %v", smsInformation)
```

### Retrieve SMS Inbox List
```go
var orderId string = "123456"

smsInboxList, err := sim5Client.GetSmsInboxList(orderId)

if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("SMS Inbox List: %v", smsInboxList)
```
