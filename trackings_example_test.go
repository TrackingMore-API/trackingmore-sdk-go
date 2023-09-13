package trackingmore_test

import (
	"context"
	"fmt"
	"github.com/trackingmores/trackingmore-sdk-go"
	"time"
)

func ExampleClient_CreateTracking() {
	key := "you api key"
	cli, err := trackingmore.NewClient(key)

	if err != nil {
		fmt.Println(err)
		return
	}

	params := trackingmore.CreateTrackingParams{
		TrackingNumber: "9400111899562537683144",
		CourierCode:    "usps",
	}
	result, err := cli.CreateTracking(context.Background(), params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	var tracking, ok = result.Data.(*trackingmore.Tracking)
	if !ok {
		fmt.Println("Structure type conversion failed")
		return
	}
	fmt.Printf("tracking_number:%s courier_code:%s\n", tracking.TrackingNumber, tracking.CourierCode)

}

func ExampleClient_GetTrackingResults() {
	key := "you api key"
	cli, err := trackingmore.NewClient(key)

	if err != nil {
		fmt.Println(err)
		return
	}

	//params := trackingmore.GetTrackingResultsParams{
	//	TrackingNumbers: "92612903029511573030094532",
	//	CourierCode:     "usps",
	//}
	//
	//params := trackingmore.GetTrackingResultsParams{
	//	TrackingNumbers: "92612903029511573030094531,9400111899562539126562",
	//	CourierCode:     "usps",
	//}

	currentTime := time.Now()
	zeroTime := currentTime.UTC()
	layout := "2006-01-02T15:04:05-07:00"
	formattedTime := zeroTime.Format(layout)
	params := trackingmore.GetTrackingResultsParams{
		CreatedDateMin: "2023-08-23T06:00:00+00:00",
		CreatedDateMax: formattedTime,
	}

	result, err := cli.GetTrackingResults(context.Background(), params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	var trackings, ok = result.Data.(*[]trackingmore.Tracking)
	if !ok {
		fmt.Println("Structure type conversion failed")
		return
	}
	for _, item := range *trackings {
		fmt.Printf("id:%s tracking_number: %s\n", item.Id, item.TrackingNumber)
	}

}

func ExampleClient_BatchCreateTrackings() {
	key := "you api key"
	cli, err := trackingmore.NewClient(key)

	if err != nil {
		fmt.Println(err)
		return
	}

	params := []trackingmore.CreateTrackingParams{
		{
			TrackingNumber: "92632903279511573030094832",
			CourierCode:    "usps",
		},
		{
			TrackingNumber: "92642903289511563030094932",
			CourierCode:    "usps",
		},
	}
	result, err := cli.BatchCreateTrackings(context.Background(), params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	var batchResults, ok = result.Data.(*trackingmore.BatchResults)

	if !ok {
		fmt.Println("Structure type conversion failed")
		return
	}

	for _, item := range batchResults.Error {
		fmt.Printf("error_code: %d error_message: %s courier_code: %s tracking_number: %s\n", item.ErrorCode, item.ErrorMessage, item.CourierCode, item.TrackingNumber)
	}

	for _, item := range batchResults.Success {
		fmt.Printf("courier_code: %s tracking_number: %s\n", item.CourierCode, item.TrackingNumber)
	}

}

func ExampleClient_UpdateTrackingByID() {
	key := "you api key"
	cli, err := trackingmore.NewClient(key)

	if err != nil {
		fmt.Println(err)
		return
	}

	params := trackingmore.UpdateTrackingParams{
		CustomerName: "New name",
		Note:         "New tests order note",
	}
	idString := "9a1d3844a50f3851e76e3ee347881588"
	result, err := cli.UpdateTrackingByID(context.Background(), idString, params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	var updateAfterResult, ok = result.Data.(*trackingmore.UpdateAfterResult)
	if !ok {
		fmt.Println("Structure type conversion failed")
		return
	}
	fmt.Printf("customer_name: %s note: %s\n", updateAfterResult.CustomerName, updateAfterResult.Note)

}

func ExampleClient_DeleteTrackingByID() {
	key := "you api key"
	cli, err := trackingmore.NewClient(key)

	if err != nil {
		fmt.Println(err)
		return
	}

	idString := "9a1d3844a50f3851e76e3ee347881588"
	result, err := cli.DeleteTrackingByID(context.Background(), idString)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	var tracking, ok = result.Data.(*trackingmore.Tracking)
	if !ok {
		fmt.Println("Structure type conversion failed")
		return
	}
	fmt.Printf("id:%s courier_code: %s tracking_number: %s\n", tracking.Id, tracking.CourierCode, tracking.TrackingNumber)

}

func ExampleClient_RetrackTrackingByID() {
	key := "you api key"
	cli, err := trackingmore.NewClient(key)

	if err != nil {
		fmt.Println(err)
		return
	}

	idString := "99ff2ce10105aeb8627ec0c03e1773bd"
	result, err := cli.RetrackTrackingByID(context.Background(), idString)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	var tracking, ok = result.Data.(*trackingmore.Tracking)
	if !ok {
		fmt.Println("Structure type conversion failed")
		return
	}
	fmt.Printf("id:%s courier_code: %s tracking_number: %s\n", tracking.Id, tracking.CourierCode, tracking.TrackingNumber)

}
