package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Order struct {
	ID          int    `json:"id,omitempty" bson:"id,omitempty"`
	Order_Name  string `json:"order_name,omitempty" bson:"order_name,omitempty"`
	Customer_Id string `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	Created_at  string `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type Customer struct {
	User_Id    string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Copmany_Id int    `json:"company_id,omitempty" bson:"company_id,omitempty"`
}

type Company struct {
	Company_Id int    `json:"company_id,omitempty" bson:"company_id,omitempty"`
	Name       string `json:"company_name,omitempty" bson:"company_name,omitempty"`
}

type Order_Item struct {
	ID       int    `json:"id" bson:"id"`
	Order_Id int    `json:"order_id" bson:"order_id"`
	Quantity int    `json:"quantity" bson:"quantity"`
	Price    string `json:"price_per_unit" bson:"price_per_unit"`
}

type Delivery struct {
	Order_Item_Id int32 `json:"order_item_id,omitempty" bson:"order_item_id,omitempty"`
	Qty           int   `json:"delivered_quantity" bson:"delivered_quantity"`
}

var client *mongo.Client

func fetchOrdersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	var orders []Order
	collection := client.Database("Dataset").Collection("orders")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var order Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(orders)
}

func fetchCustomersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	var customers []Customer
	collection := client.Database("Dataset").Collection("customers")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var customer Customer
		cursor.Decode(&customer)
		customers = append(customers, customer)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(customers)
}

func fetchCompanyEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	var companies []Company
	collection := client.Database("Dataset").Collection("customer_companies")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var company Company
		cursor.Decode(&company)
		companies = append(companies, company)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(companies)
}

func fetchOrderItemsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	var items []Order_Item
	collection := client.Database("Dataset").Collection("order_items")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var item Order_Item
		cursor.Decode(&item)
		items = append(items, item)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(items)
}

func fetchDeliveriesEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	var deliveries []Delivery
	collection := client.Database("Dataset").Collection("deliveries")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var delivery Delivery
		cursor.Decode(&delivery)
		deliveries = append(deliveries, delivery)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(deliveries)
}

func main() {
	fmt.Println("Server Running...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/orders", fetchOrdersEndpoint).Methods("GET")
	router.HandleFunc("/customers", fetchCustomersEndpoint).Methods("GET")
	router.HandleFunc("/companies", fetchCompanyEndpoint).Methods("GET")
	router.HandleFunc("/order_items", fetchOrderItemsEndpoint).Methods("GET")
	router.HandleFunc("/deliveries", fetchDeliveriesEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}
