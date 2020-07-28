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

	music:=  types.Music{
		Artist: "Sonu",
		SongTitle: "Hello Again",
		AlbumTitle: "My Album",
		Awards: 2,
	}

	musicmap,err:=dynamodbattribute.MarshalMap(music)
	if err!=nil{
		panic("Cannot map the values given in music struct")
	}

	params:=&dynamodb.PutItemInput{
		TableName: aws.String("Music1"),
		Item: musicmap,
	}

	resp,err:= db.PutItem(params)

	if err!=nil{
		panic("Some problem.")
	}
	fmt.Println("Record inserted.")
	fmt.Println(resp)

}