Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fconsumption%2Farmconsumption%2Fv1.0.0/sdk/resourcemanager/consumption/armconsumption/README.md) on how to add the SDK to your project and authenticate.

```go
package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesDaily.json
func ExampleReservationsSummariesClient_NewListByReservationOrderPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconsumption.NewReservationsSummariesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByReservationOrderPager("00000000-0000-0000-0000-000000000000",
		armconsumption.DatagrainDailyGrain,
		&armconsumption.ReservationsSummariesClientListByReservationOrderOptions{Filter: to.Ptr("properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-11-20")})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
