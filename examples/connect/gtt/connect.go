package main

import (
	"log"

	kiteconnect "github.com/nsvirk/gokiteconnect/v4"
)

const (
	userId   string = "my_userId"
	enctoken string = "my_enctoken"
)

func main() {
	// Create a new Kite connect instance
	kc := kiteconnect.New(userId)

	// Set access token
	kc.SetEnctoken(enctoken)

	log.Println("Fetching GTTs...")
	orders, err := kc.GetGTTs()
	if err != nil {
		log.Fatalf("Error getting GTTs: %v", err)
	}
	log.Printf("gtt: %v", orders)

	log.Println("Placing GTT...")
	// Place GTT
	gttResp, err := kc.PlaceGTT(kiteconnect.GTTParams{
		Tradingsymbol:   "INFY",
		Exchange:        "NSE",
		LastPrice:       800,
		TransactionType: kiteconnect.TransactionTypeBuy,
		Trigger: &kiteconnect.GTTSingleLegTrigger{
			TriggerParams: kiteconnect.TriggerParams{
				TriggerValue: 1,
				Quantity:     1,
				LimitPrice:   1,
			},
		},
	})
	if err != nil {
		log.Fatalf("error placing gtt: %v", err)
	}

	log.Println("placed GTT trigger_id = ", gttResp.TriggerID)

	log.Println("Fetching details of placed GTT...")

	order, err := kc.GetGTT(gttResp.TriggerID)
	if err != nil {
		log.Fatalf("Error getting GTTs: %v", err)
	}
	log.Printf("gtt: %v", order)

	log.Println("Modify existing GTT...")

	gttModifyResp, err := kc.ModifyGTT(gttResp.TriggerID, kiteconnect.GTTParams{
		Tradingsymbol:   "INFY",
		Exchange:        "NSE",
		LastPrice:       800,
		TransactionType: kiteconnect.TransactionTypeBuy,
		Trigger: &kiteconnect.GTTSingleLegTrigger{
			TriggerParams: kiteconnect.TriggerParams{
				TriggerValue: 2,
				Quantity:     2,
				LimitPrice:   2,
			},
		},
	})
	if err != nil {
		log.Fatalf("error placing gtt: %v", err)
	}

	log.Println("modified GTT trigger_id = ", gttModifyResp.TriggerID)

	gttDeleteResp, err := kc.DeleteGTT(gttResp.TriggerID)
	if err != nil {
		log.Fatalf("Error getting GTTs: %v", err)
	}
	log.Printf("gtt deleted: %v", gttDeleteResp)
}
