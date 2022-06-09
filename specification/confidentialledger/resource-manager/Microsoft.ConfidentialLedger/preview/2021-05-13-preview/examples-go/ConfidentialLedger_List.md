```go
package armconfidentialledger_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confidentialledger/armconfidentialledger"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2021-05-13-preview/examples/ConfidentialLedger_List.json
func ExampleLedgerClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconfidentialledger.NewLedgerClient("0000000-0000-0000-0000-000000000001", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("DummyResourceGroupName",
		&armconfidentialledger.LedgerClientListByResourceGroupOptions{Filter: nil})
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fconfidentialledger%2Farmconfidentialledger%2Fv0.5.0/sdk/resourcemanager/confidentialledger/armconfidentialledger/README.md) on how to add the SDK to your project and authenticate.
