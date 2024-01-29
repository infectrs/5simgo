
## Usage

### Client

Firstly, initialize the 5Sim client with:

```golang
func main() {
	apiKey := "ey..."
	sim5Client, err := New5Sim(apiKey)

	if err != nil {
		fmt.Println(err)
		return
	}
}
```

Then you can use the following methods:

```golang
// Retrieve profile information
sim5Client.GetProfileInformation()

// Retrieve orders history
sim5Client.GetOrdersHistory(category string)

// Retrieve payments history
sim5Client.GetPaymentsHistory()

// Buy activation number
sim5Client.BuyActivationNumber(country, operator, product string)

// Buy hosting number
sim5Client.BuyHostingNumber(country, operator, product string)

// Rebuy number
sim5Client.RebuyNumber(product, phoneNumber string)

// Check order (get sms)
sim5Client.CheckOrder(id string)

// Finish order
sim5Client.FinishOrder(id string)

// Cancel order
sim5Client.CancelOrder(id string)

// Ban order
sim5Client.BanOrder(id string)

// Get SMS Inbox List
sim5Client.GetSmsInboxList(id string)
```