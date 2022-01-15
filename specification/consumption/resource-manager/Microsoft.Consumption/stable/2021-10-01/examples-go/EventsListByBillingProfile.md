Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fconsumption%2Farmconsumption%2Fv0.3.0/sdk/resourcemanager/consumption/armconsumption/README.md) on how to add the SDK to your project and authenticate.

```go
package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/EventsListByBillingProfile.json
func ExampleEventsClient_ListByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconsumption.NewEventsClient(cred, nil)
	pager := client.ListByBillingProfile("<billing-account-id>",
		"<billing-profile-id>",
		"<start-date>",
		"<end-date>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
```
