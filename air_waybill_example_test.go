package trackingmore_test

import (
	"context"
	"fmt"
	"trackingmore"
)

func ExampleClient_CreateAnAirWayBill() {
	key := "you api key"
	cli, err := trackingmore.NewClient(key)

	if err != nil {
		fmt.Println(err)
		return
	}

	params := trackingmore.AirWaybillParams{
		AwbNumber: "235-69030430",
	}
	result, err := cli.CreateAnAirWayBill(context.Background(), params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	var airWaybillItem, ok = result.Data.(*trackingmore.AirWaybillItem)
	if !ok {
		fmt.Println("Structure type conversion failed")
		return
	}
	fmt.Printf("awb_number:%s destination:%s\n", airWaybillItem.AwbNumber, airWaybillItem.Destination)

}
