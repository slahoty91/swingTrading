package model

type Orders struct {
	ID            string  `json:"id" bson:"id"`
	EntryPrice    int64   `json:"entryPrice" bson:"entryPrice"`
	StopLoss      int64   `json:"stoploss" bson:"stoploss"`
	ExecutedPrice float64 `json:"executedPrice" bson:"executedPrice"`
	Target        int64   `json:"target" bson:"target"`
	Token         int64   `json:"token" bson:"token"`
	Tradingsymbol string  `json:"tradingsymbol" bson:"tradingsymbol"`
}
