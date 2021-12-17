Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.3.0/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionRulesUpdate.json
func ExampleDataCollectionRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionRulesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<data-collection-rule-name>",
		&armmonitor.DataCollectionRulesUpdateOptions{Body: &armmonitor.ResourceForUpdate{
			Tags: map[string]*string{
				"tag1": to.StringPtr("A"),
				"tag2": to.StringPtr("B"),
				"tag3": to.StringPtr("C"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataCollectionRuleResource.ID: %s\n", *res.ID)
}
```
