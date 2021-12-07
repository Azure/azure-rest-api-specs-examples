Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fchangeanalysis%2Farmchangeanalysis%2Fv0.1.0/sdk/resourcemanager/changeanalysis/armchangeanalysis/README.md) on how to add the SDK to your project and authenticate.

```go
package armchangeanalysis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/changeanalysis/armchangeanalysis"
)

// x-ms-original-file: specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/ChangesListChangesBySubscription.json
func ExampleChangesClient_ListChangesBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armchangeanalysis.NewChangesClient("<subscription-id>", cred, nil)
	pager := client.ListChangesBySubscription(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-25T12:09:03.141Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-26T12:09:03.141Z"); return t }(),
		&armchangeanalysis.ChangesListChangesBySubscriptionOptions{SkipToken: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Change.ID: %s\n", *v.ID)
		}
	}
}
```
