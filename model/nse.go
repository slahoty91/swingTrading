package model

type NSE struct {
	ID               string  `json:"id" bson:"_id"`
	Instrument_token int64   `json:"instrument_token" bson:"instrument_token"`
	Tradingsymbol    string  `json:"tradingsymbol" bson:"tradingsymbol"`
	Name             string  `json:"name" bson:"name"`
	Last_price       float64 `json:"last_price" bson:"last_price"`
}
