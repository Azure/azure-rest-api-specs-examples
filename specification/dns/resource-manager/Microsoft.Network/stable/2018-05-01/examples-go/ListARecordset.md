Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdns%2Farmdns%2Fv0.4.0/sdk/resourcemanager/dns/armdns/README.md) on how to add the SDK to your project and authenticate.

```go
package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListARecordset.json
func ExampleRecordSetsClient_NewListByTypePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdns.NewRecordSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListByTypePager("<resource-group-name>",
		"<zone-name>",
		armdns.RecordTypeA,
		&armdns.RecordSetsClientListByTypeOptions{Top: nil,
			Recordsetnamesuffix: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
