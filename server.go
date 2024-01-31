package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

func main() {
	// This is a public sample test API key.
	// Don’t submit any personally identifiable information in requests made with this key.
	// Sign in to see your own test API key embedded in code samples.
	stripe.Key = "sk_test_51OdxYiSAcOZ5DlUs6i4ID7BGQ2X93AcZk2Q8cKMoy9lw8CiGNhtSCXxTodwOlg3HniT1KzFvYXetgDQkdQlyLJ7A00YOv7UhTz"

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	addr := "localhost:4242"
	log.Printf("Listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type item struct {
	id string
}

func calculateOrderAmount(items []item) int64 {
	// Replace this constant with a calculation of the order's amount
	// Calculate the order total on the server to prevent
	// people from directly manipulating the amount on the client
	return 1400
}

type AddressParams struct {
	Line1      string
	Line2      string
	City       string
	State      string
	PostalCode string
	Country    string
}

type CustomerInfo struct {
	Name    string
	Email   string
	Address *AddressParams
	Amount  int
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	var customerid string
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req CustomerInfo

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewDecoder.Decode: %v", err)
		return
	}

	params2 := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "email: " + req.Email + "",
		},
	}
	result := customer.Search(params2)
	if result.Customer().ID == "" {
		params1 := &stripe.CustomerParams{
			Name:  stripe.String("Jenny Rosen"),
			Email: stripe.String("jennyrosen@example.com"),
			Address: &stripe.AddressParams{
				Line1:      stripe.String("123 Main St"),
				Line2:      stripe.String("Apt 4"),
				City:       stripe.String("Cityville"),
				State:      stripe.String("CA"),
				PostalCode: stripe.String("12345"),
				Country:    stripe.String("US"),
			},
		}

		result1, err := customer.New(params1)
		if err != nil {
			log.Println("customer.New : ", err.Error())
		}
		customerid = result1.ID
	} else {
		customerid = result.Customer().ID
	}

	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(req.Amount)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		// In the latest version of the API, specifying the `automatic_payment_methods` parameter is optional because Stripe enables its functionality by default.
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
		Description: stripe.String("tes"),
		Customer:    stripe.String(customerid),
	}

	pi, err := paymentintent.New(params)

	log.Printf("pi.New: %v", pi.ClientSecret)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("pi.New: %v", err)
		return
	}

	writeJSON(w, struct {
		ClientSecret string `json:"clientSecret"`
	}{
		ClientSecret: pi.ClientSecret,
	})
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("io.Copy: %v", err)
		return
	}
}
