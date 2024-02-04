trackingmore-sdk-go
=================

The Go SDK of TrackingMore API

Contact: <manage@trackingmore.org>

## Official document

[Document](https://www.trackingmore.com/docs/trackingmore/d5ac362fc3cda-api-quick-start)

## Index
1. [Installation](https://github.com/TrackingMore-API/trackingmore-sdk-go#installation)
2. [Testing](https://github.com/TrackingMore-API/trackingmore-sdk-go#testing)
3. [Error Handling](https://github.com/TrackingMore-API/trackingmore-sdk-go#error-handling)
4. SDK
    1. [Couriers](https://github.com/TrackingMore-API/trackingmore-sdk-go#couriers)
    2. [Trackings](https://github.com/TrackingMore-API/trackingmore-sdk-go#trackings)
    3. [Air Waybill](https://github.com/TrackingMore-API/trackingmore-sdk-go#air-waybill)


## Installation

trackingmore-sdk-go requires a Go version with [Modules](https://github.com/golang/go/wiki/Modules) support and uses import versioning. So please make sure to initialize a Go module before installing trackingmore-sdk-go:

```
go mod init github.com/my/repo
go get github.com/trackingmore-api/trackingmore-sdk-go
```

Import:

``` go
import "github.com/trackingmore-api/trackingmore-sdk-go"
```

## Quick Start

```go
package main

import (
"context"
"fmt"
"github.com/trackingmore-api/trackingmore-sdk-go"
)

func main() {
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

```

## Testing
```
go test
```

## Error handling

**Throw** by the new SDK client

```go
cli, err := trackingmore.NewClient("")

if err != nil {
   fmt.Println(err)
   return
}

/*
API Key is missing
*/
```

**Throw** by the parameter validation in function

```go
cli, err := trackingmore.NewClient("your api key")

if err != nil {
   fmt.Println(err)
   return
}

params := trackingmore.DetectParams{
    TrackingNumber: "",
}
result, err := cli.Detect(context.Background(), params)
if err != nil {
   fmt.Println(err)
   return
}

/*
Tracking number cannot be empty
*/
```
## Examples

## Couriers
##### Return a list of all supported couriers.
https://api.trackingmore.com/v4/couriers/all
```go
result, err := cli.GetAllCouriers(context.Background())
if err != nil {
  fmt.Println(err)
  return
}

fmt.Println(result)
```

##### Return a list of matched couriers based on submitted tracking number.
https://api.trackingmore.com/v4/couriers/detect
```go
params := trackingmore.DetectParams{
  TrackingNumber: "92612903029511573030094531",
}

result, err := cli.Detect(context.Background(), params)
if err != nil {
  fmt.Println(err)
  return
}

fmt.Println(result)
```

## Trackings
##### Create a tracking.
https://api.trackingmore.com/v4/trackings/create
```go
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
```

##### Get tracking results of multiple trackings.
https://api.trackingmore.com/v4/trackings/get
```go
// Perform queries based on various conditions
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
```

##### Create multiple trackings (Max. 40 tracking numbers create in one call).
https://api.trackingmore.com/v4/trackings/batch
```go
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
```

##### Update a tracking by ID.
https://api.trackingmore.com/v4/trackings/update/{id}
```go
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
```

##### Delete a tracking by ID.
https://api.trackingmore.com/v4/trackings/delete/{id}
```go
idString := "9a1d3844a50f3851e76e3ee347881588"
result, err := cli.DeleteTrackingByID(context.Background(), idString)
if err != nil {
   fmt.Println(err)
   return
}

fmt.Println(result)
```

##### Retrack expired tracking by ID.
https://api.trackingmore.com/v4/trackings/retrack/{id}
```go
idString := "99ff2ce10105aeb8627ec0c03e1773bd"
result, err := cli.RetrackTrackingByID(context.Background(), idString)
if err != nil {
   fmt.Println(err)
   return
}

fmt.Println(result)
```
## Air Waybill
##### Create an air waybill.
https://api.trackingmore.com/v4/awb
```go
params := trackingmore.AirWaybillParams{
   AwbNumber: "235-69030430",
}

result, err := cli.CreateAnAirWayBill(context.Background(), params)
if err != nil {
   fmt.Println(err)
   return
}

fmt.Println(result)
```

## Response Code

Trackingmore uses conventional HTTP response codes to indicate success or failure of an API request. In general, codes in the 2xx range indicate success, codes in the 4xx range indicate an error that resulted from the provided information (e.g. a required parameter was missing, a charge failed, etc.), and codes in the 5xx range indicate an TrackingMore's server error.


Http CODE|META CODE|TYPE | MESSAGE
----|-----|--------------|-------------------------------
200    |200     | <code>Success</code>        |    Request response is successful
400    |400     | <code>BadRequest</code>     |    Request type error. Please check the API documentation for the request type of this API.
400    |4101    | <code>BadRequest</code>     |    Tracking No. already exists.
400    |4102    | <code>BadRequest</code>     |    Tracking No. no exists. Please use 「Create a tracking」 API first to create shipment.
400    |4103    | <code>BadRequest</code>     |    You have exceeded the shipment quantity of API call. The maximum quantity is 40 shipments per call.
400    |4110    | <code>BadRequest</code>     |    The value of tracking_number is invalid.
400    |4111    | <code>BadRequest</code>     |    Tracking_number is required.
400    |4112    | <code>BadRequest</code>     |    Invalid Tracking ID.
400    |4113    | <code>BadRequest</code>     |    Retrack is not allowed. You can only retrack an expired tracking.
400    |4120    | <code>BadRequest</code>     |    The value of courier_code is invalid.
400    |4121    | <code>BadRequest</code>     |    Cannot detect courier.
400    |4122    | <code>BadRequest</code>     |    Missing or invalid value of the special required fields for this courier.
400    |4130    | <code>BadRequest</code>     |    The format of Field name is invalid.
400    |4160    | <code>BadRequest</code>     |    The awb_number is required or invaild format.
400    |4161    | <code>BadRequest</code>     |    The awb airline does not support yet.
400    |4190    | <code>BadRequest</code>     |    You are reaching the maximum quota limitation, please upgrade your current plan.
401    |401     | <code>Unauthorized</code>   |    Authentication failed or has no permission. Please check and ensure your API Key is correct.
403    |403     | <code>Forbidden</code>      |    Access prohibited. The request has been refused or access is not allowed.
404    |404     | <code>NotFound</code>       |    Page does not exist. Please check and ensure your link is correct.
429    |429     | <code>TooManyRequests</code>|    Exceeded API request limits, please try again later. Please check the API documentation for the limit of this API.
500    |511     | <code>ServerError</code>    |    Server error. Please contact us: service@trackingmore.org.
500    |512     | <code>ServerError</code>    |    Server error. Please contact us: service@trackingmore.org.
500    |513     | <code>ServerError</code>    |    Server error. Please contact us: service@trackingmore.org.