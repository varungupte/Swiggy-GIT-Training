package main

import (
	"./types"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main(){
	sess:= session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region: aws.String("us-east-1"),
	}))
	db:= dynamodb.New(sess)

	//query parameters
	artist:= "Sonu"
	songTitle:="Hello Again"

	params:=&dynamodb.GetItemInput{
		TableName: aws.String("Music1"),
		Key: map[string]*dynamodb.AttributeValue{
			"Artist":{
				S:aws.String(artist),
			},
			"SongTitle":{
				S:aws.String(songTitle),
			},
		},
	}

	resp,err:=db.GetItem(params)
	if err!=nil{
		panic("Item not found.")
	}

	// result in marshalled form
	fmt.Println(resp)

	// result in unmarshalled form
	music := types.Music{}

	err = dynamodbattribute.UnmarshalMap(resp.Item,&music)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}
	fmt.Println("Unmarshalled Item %+v",music)


}

