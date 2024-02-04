package trackingmore_test

import (
	"context"
	"fmt"
	"github.com/trackingmore-api/trackingmore-sdk-go"
)

func ExampleClient_GetCouriers() {
	key := "your api key"
	cli, err := trackingmore.NewClient(key)

	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := cli.GetAllCouriers(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	var couriers, ok = result.Data.(*[]trackingmore.Courier)
	if !ok {
		fmt.Println("Structure type conversion failed")
		return
	}
	for _, item := range *couriers {
		fmt.Printf("courier_name:%s courier_code:%s\n", item.CourierName, item.CourierCode)
	}
}

func ExampleClient_Detect() {
	key := "your api key"
	cli, err := trackingmore.NewClient(key)

	if err != nil {
		fmt.Println(err)
		return
	}

	params := trackingmore.DetectParams{
		TrackingNumber: "92612903029511573030094531",
	}
	result, err := cli.Detect(context.Background(), params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	var couriers, ok = result.Data.(*[]trackingmore.Courier)

	if !ok {
		fmt.Println("Structure type conversion failed")
		return
	}

	for _, item := range *couriers {
		fmt.Printf("courier_name:%s courier_code:%s\n", item.CourierName, item.CourierCode)
	}

}
