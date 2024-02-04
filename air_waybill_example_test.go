package trackingmore_test

import (
	"context"
	"fmt"
	"github.com/trackingmore-api/trackingmore-sdk-go"
)

func ExampleClient_CreateAnAirWayBill() {
	key := "your api key"
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
