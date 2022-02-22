Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.4.1/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionRuleAssociationsCreate.json
func ExampleDataCollectionRuleAssociationsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionRuleAssociationsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-uri>",
		"<association-name>",
		&armmonitor.DataCollectionRuleAssociationsClientCreateOptions{Body: &armmonitor.DataCollectionRuleAssociationProxyOnlyResource{
			Properties: &armmonitor.DataCollectionRuleAssociationProxyOnlyResourceProperties{
				DataCollectionRuleID: to.StringPtr("<data-collection-rule-id>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataCollectionRuleAssociationsClientCreateResult)
}
```
