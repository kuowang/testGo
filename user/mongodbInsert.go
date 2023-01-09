package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {

}

type Student struct {
	Name string
	Age  int
}

func insertOne(s Student) {
	initDB()
	collection := client.Database("go_db").Collection("student") //打开数据库对应的表

	insertResult, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult.InsertedID)
}

var client *mongo.Client

func initDB() {

	clientOption := options.Client().ApplyURI("mongodb://182.92.166.65:27017")

	var err error
	client, err = mongo.Connect(context.TODO(), clientOption) //数据库链接
	if err != nil {
		log.Fatal(err)
	}
	err2 := client.Ping(context.TODO(), nil) //检查链接是否通
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("链接mongodb 成功")

}
