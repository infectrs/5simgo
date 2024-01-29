package main

import (
	"time"
)

// ProfileInformation represents user profile information
type ProfileInformation struct {
	ID                      int    `json:"id"`
	Email                   string `json:"email"`
	Vendor                  string `json:"vendor"`
	DefaultForwardingNumber string `json:"default_forwarding_number"`
	Balance                 int    `json:"balance"`
	Rating                  int    `json:"rating"`
	DefaultCountry          struct {
		Name   string `json:"name"`
		Iso    string `json:"iso"`
		Prefix string `json:"prefix"`
	} `json:"default_country"`
	DefaultOperator struct {
		Name string `json:"name"`
	} `json:"default_operator"`
	FrozenBalance int `json:"frozen_balance"`
}

// OrdersInformation represents user orders information
type OrdersInformation struct {
	Data []struct {
		ID        int       `json:"id"`
		Phone     string    `json:"phone"`
		Operator  string    `json:"operator"`
		Product   string    `json:"product"`
		Price     int       `json:"price"`
		Status    string    `json:"status"`
		Expires   time.Time `json:"expires"`
		Sms       []any     `json:"sms"`
		CreatedAt time.Time `json:"created_at"`
		Country   string    `json:"country"`
	} `json:"Data"`
	ProductNames []any `json:"ProductNames"`
	Statuses     []any `json:"Statuses"`
	Total        int   `json:"Total"`
}

// PaymentsInformation represents user payments information
type PaymentsInformation struct {
	Data []struct {
		ID           int       `json:"ID"`
		TypeName     string    `json:"TypeName"`
		ProviderName string    `json:"ProviderName"`
		Amount       int       `json:"Amount"`
		Balance      int       `json:"Balance"`
		CreatedAt    time.Time `json:"CreatedAt"`
	} `json:"Data"`
	PaymentTypes []struct {
		Name string `json:"Name"`
	} `json:"PaymentTypes"`
	PaymentProviders []struct {
		Name string `json:"Name"`
	} `json:"PaymentProviders"`
	Total int `json:"Total"`
}

// PurchasedActivationNumberInformation represents information about a purchased activation number
type PurchasedActivationNumberInformation struct {
	ID               int       `json:"id"`
	Phone            string    `json:"phone"`
	Operator         string    `json:"operator"`
	Product          string    `json:"product"`
	Price            int       `json:"price"`
	Status           string    `json:"status"`
	Expires          time.Time `json:"expires"`
	Sms              any       `json:"sms"`
	CreatedAt        time.Time `json:"created_at"`
	Forwarding       bool      `json:"forwarding"`
	ForwardingNumber string    `json:"forwarding_number"`
	Country          string    `json:"country"`
}

// PurchasedHostingNumberInformation represents information about a purchased hosting number
type PurchasedHostingNumberInformation struct {
	ID      int       `json:"id"`
	Phone   string    `json:"phone"`
	Product string    `json:"product"`
	Price   int       `json:"price"`
	Status  string    `json:"status"`
	Expires time.Time `json:"expires"`
	Sms     []struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Date      time.Time `json:"date"`
		Sender    string    `json:"sender"`
		Text      string    `json:"text"`
		Code      string    `json:"code"`
	} `json:"sms"`
	CreatedAt time.Time `json:"created_at"`
}

// SmsInformation represents information about SMS
type SmsInformation struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Phone     string    `json:"phone"`
	Product   string    `json:"product"`
	Price     int       `json:"price"`
	Status    string    `json:"status"`
	Expires   time.Time `json:"expires"`
	Sms       []struct {
		CreatedAt time.Time `json:"created_at"`
		Date      time.Time `json:"date"`
		Sender    string    `json:"sender"`
		Text      string    `json:"text"`
		Code      string    `json:"code"`
	} `json:"sms"`
	Forwarding       bool   `json:"forwarding"`
	ForwardingNumber string `json:"forwarding_number"`
	Country          string `json:"country"`
}

// SmsInbox represents information about the SMS inbox
type SmsInbox struct {
	Data []struct {
		ID        int       `json:"ID"`
		CreatedAt time.Time `json:"created_at"`
		Date      time.Time `json:"date"`
		Sender    string    `json:"sender"`
		Text      string    `json:"text"`
		Code      string    `json:"code"`
		IsWave    bool      `json:"is_wave"`
		WaveUUID  string    `json:"wave_uuid"`
	} `json:"Data"`
	Total int `json:"Total"`
}
