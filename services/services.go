package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/slahoty91/tradingBot/database"
	"github.com/slahoty91/tradingBot/model"
	kiteconnect "github.com/zerodha/gokiteconnect/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NSEservice struct {
	nseCollection *database.Collection
}

type Orderservice struct {
	orderCollection *database.Collection
	nseCollection   *database.Collection
	// userCollection  *database.Collection
}

func NewNSEService() *NSEservice {
	return &NSEservice{
		nseCollection: database.NewCollection("instrumentNSE"),
	}
}

func NewOrderService() *Orderservice {
	return &Orderservice{
		orderCollection: database.NewCollection("swingOrders"),
		nseCollection:   database.NewCollection("instrumentNSE"),
	}
}

func (ns *NSEservice) GetName(inst_token int64) (string, error) {
	var nse model.NSE
	err := ns.nseCollection.FindOne(context.Background(), map[string]interface{}{"instrument_token": inst_token}).Decode(&nse)
	if err != nil {
		return "", err
	}

	return nse.Name, nil
}

func (ns *NSEservice) FuzzySearch(query string) ([]model.NSE, error) {

	fmt.Println(query, "queryyyyyy")
	var result []model.NSE
	cursor, err := ns.nseCollection.Find(context.Background(), map[string]interface{}{"tradingsymbol": primitive.Regex{Pattern: query, Options: "i"}})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var nse model.NSE
		if err := cursor.Decode(&nse); err != nil {
			return nil, err
		}
		if fuzzy.Match(query, nse.Tradingsymbol) {
			result = append(result, nse)
		}
	}

	return result, nil

}

func (ord *Orderservice) CreateOrders(data model.Orders) error {

	var ns model.NSE
	// projection := bson.M{"instrument_token": 1}
	if err := ord.nseCollection.FindOne(context.Background(), bson.M{"tradingsymbol": data.Tradingsymbol}).Decode(&ns); err != nil {
		return err
	}
	fmt.Println(ns.Instrument_token, "nsssssss")
	data.Token = ns.Instrument_token
	// _, err := ord.orderCollection.InsertOne(context.Background(), data)
	// if err != nil {
	// 	return err
	// }
	ord.GetCurrentPrice(data.Tradingsymbol)
	return nil
}

func (ord *Orderservice) GetCurrentPrice(trading_inst string) {

	// var user model.User
	fmt.Println("GET quoteee")

	kc := kiteconnect.New("k55bdfkr27eqguv6")

	kc.SetAccessToken("1AvCivSE1Bpkjb4Ji7H7ZgO9e0xouit1")

	quote, err := kc.GetLTP("NSE:" + trading_inst)
	if err != nil {
		fmt.Println(err, "err from quote")
	}

	jsonData, err := json.Marshal(quote)
	if err != nil {
		fmt.Println(err, "err convertingto json")
	}
	fmt.Println(string(jsonData), "json data")

}
