Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fchangeanalysis%2Farmchangeanalysis%2Fv1.0.0/sdk/resourcemanager/changeanalysis/armchangeanalysis/README.md) on how to add the SDK to your project and authenticate.

```go
package armchangeanalysis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/changeanalysis/armchangeanalysis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/ChangesListChangesByResourceGroup.json
func ExampleChangesClient_NewListChangesByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armchangeanalysis.NewChangesClient("4d962866-1e3f-47f2-bd18-450c08f914c1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListChangesByResourceGroupPager("MyResourceGroup",
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-25T12:09:03.141Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-26T12:09:03.141Z"); return t }(),
		&armchangeanalysis.ChangesClientListChangesByResourceGroupOptions{SkipToken: nil})
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
