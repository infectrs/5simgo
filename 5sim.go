package main

import (
	"encoding/json"
	"fmt"
)

// Constants
const (
	domain = "5sim.net"
)

// Sim5 is a struct that represents the 5Sim API Client
type Sim5 struct {
	apiKey string
}

// New5Sim creates a new instance of Sim5 with the provided API key
func New5Sim(apiKey string) (*Sim5, error) {

	if apiKey == "" {
		return nil, missingApiKeyError
	}

	return &Sim5{
		apiKey: apiKey,
	}, nil
}

// // GetProfileInformation retrieves user profile information from the 5Sim API
func (s Sim5) GetProfileInformation() (*ProfileInformation, error) {

	client := newHttpClient()

	url := fmt.Sprintf("https://%s/v1/user/profile", domain)
	resp, err := executeRequest(client, url, s.apiKey)

	if err != nil {
		return nil, httpError
	}

	var profileInfo ProfileInformation

	if err := json.Unmarshal(resp, &profileInfo); err != nil {
		return nil, unmarshalError
	}

	return &profileInfo, nil
}

// GetOrdersHistory retrieves user orders history from the 5Sim API based on the specified category
func (s Sim5) GetOrdersHistory(category string) (*OrdersInformation, error) {

	if category == "" {
		return nil, missingArgsError
	}

	url := fmt.Sprintf("https://%s/v1/user/orders?category=%s", domain, category)

	client := newHttpClient()

	resp, err := executeRequest(client, url, s.apiKey)

	if err != nil {
		return nil, httpError
	}

	var ordersInfo OrdersInformation

	if err := json.Unmarshal(resp, &ordersInfo); err != nil {
		return nil, unmarshalError
	}

	return &ordersInfo, nil
}

// GetPaymentsHistory retrieves user payments history from the 5Sim API
func (s Sim5) GetPaymentsHistory() (*PaymentsInformation, error) {
	url := fmt.Sprintf("https://%s/v1/user/payments", domain)

	client := newHttpClient()

	resp, err := executeRequest(client, url, s.apiKey)

	if err != nil {
		return nil, httpError
	}

	var paymentsInfo PaymentsInformation

	if err := json.Unmarshal(resp, &paymentsInfo); err != nil {
		return nil, unmarshalError
	}

	return &paymentsInfo, nil

}

// BuyActivationNumber purchases an activation number from the 5Sim API
func (s Sim5) BuyActivationNumber(country, operator, product string) (*PurchasedActivationNumberInformation, error) {

	if country == "" || operator == "" || product == "" {
		return nil, missingArgsError
	}

	url := fmt.Sprintf("https://%s/v1/user/buy/activation/%s/%s/%s", domain, country, operator, product)

	client := newHttpClient()

	resp, err := executeRequest(client, url, s.apiKey)

	if err != nil {
		return nil, httpError
	}

	var purchasedActivationNumberInfo PurchasedActivationNumberInformation

	if err := json.Unmarshal(resp, &purchasedActivationNumberInfo); err != nil {
		return nil, unmarshalError
	}

	return &purchasedActivationNumberInfo, nil
}

// BuyHostingNumber purchases a hosting number from the 5Sim API
func (s Sim5) BuyHostingNumber(country, operator, product string) (*PurchasedHostingNumberInformation, error) {

	if country == "" || operator == "" || product == "" {
		return nil, missingArgsError
	}

	url := fmt.Sprintf("https://%s/v1/user/buy/hosting/%s/%s/%s", domain, country, operator, product)

	client := newHttpClient()

	resp, err := executeRequest(client, url, s.apiKey)

	if err != nil {
		return nil, httpError
	}

	var purchasedHostingNumberInfo PurchasedHostingNumberInformation

	if err := json.Unmarshal(resp, &purchasedHostingNumberInfo); err != nil {
		return nil, unmarshalError
	}

	return &purchasedHostingNumberInfo, nil

}

// RebuyNumber reuses a number from the 5Sim API
func (s Sim5) RebuyNumber(product, phoneNumber string) error {
	if product == "" || phoneNumber == "" {
		return missingArgsError
	}

	url := fmt.Sprintf("https://%s/v1/user/reuse/%s/%s", domain, product, phoneNumber)

	client := newHttpClient()

	_, err := executeRequest(client, url, s.apiKey)

	if err != nil {
		return httpError
	}

	return nil
}

// CheckOrder checks the status of a specific order using its ID
func (s Sim5) CheckOrder(id string) (*SmsInformation, error) {

	if id == "" {
		return nil, missingArgsError
	}

	url := fmt.Sprintf("https://%s/v1/user/check/%s", domain, id)

	client := newHttpClient()

	resp, err := executeRequest(client, url, s.apiKey)

	if err != nil {
		return nil, httpError
	}

	var smsInfo SmsInformation

	if err := json.Unmarshal(resp, &smsInfo); err != nil {
		return nil, unmarshalError
	}

	return &smsInfo, nil

}

// FinishOrder completes the processing of a specific order using its ID
func (s Sim5) FinishOrder(id string) (*SmsInformation, error) {

	if id == "" {
		return nil, missingArgsError
	}

	url := fmt.Sprintf("https://%s/v1/user/finish/%s", domain, id)

	client := newHttpClient()

	resp, err := executeRequest(client, url, s.apiKey)

	if err != nil {
		return nil, httpError
	}

	var smsInfo SmsInformation

	if err := json.Unmarshal(resp, &smsInfo); err != nil {
		return nil, unmarshalError
	}

	return &smsInfo, nil

}

// CancelOrder cancels a specific order using its ID
func (s Sim5) CancelOrder(id string) (*SmsInformation, error) {

	if id == "" {
		return nil, missingArgsError
	}

	url := fmt.Sprintf("https://%s/v1/user/cancel/%s", domain, id)

	client := newHttpClient()

	resp, err := executeRequest(client, url, s.apiKey)

	if err != nil {
		return nil, httpError
	}

	var smsInfo SmsInformation

	if err := json.Unmarshal(resp, &smsInfo); err != nil {
		return nil, unmarshalError
	}

	return &smsInfo, nil

}

// BanOrder bans a specific order using its ID
func (s Sim5) BanOrder(id string) (*SmsInformation, error) {

	if id == "" {
		return nil, missingArgsError
	}

	url := fmt.Sprintf("https://%s/v1/user/ban/%s", domain, id)

	client := newHttpClient()

	resp, err := executeRequest(client, url, s.apiKey)

	if err != nil {
		return nil, httpError
	}

	var smsInfo SmsInformation

	if err := json.Unmarshal(resp, &smsInfo); err != nil {
		return nil, unmarshalError
	}

	return &smsInfo, nil

}

// GetSmsInboxList retrieves the list of SMS messages in the inbox based on a specific ID
func (s Sim5) GetSmsInboxList(id string) (*SmsInbox, error) {

	if id == "" {
		return nil, missingArgsError
	}

	url := fmt.Sprintf("https://%s/v1/user/sms/inbox/%s", domain, id)

	client := newHttpClient()

	resp, err := executeRequest(client, url, s.apiKey)

	if err != nil {
		return nil, httpError
	}

	var smsInboxList SmsInbox

	if err := json.Unmarshal(resp, &smsInboxList); err != nil {
		return nil, unmarshalError
	}

	return &smsInboxList, nil

}
