Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcostmanagement%2Farmcostmanagement%2Fv0.4.0/sdk/resourcemanager/costmanagement/armcostmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExternalBillingAccountsDimensions.json
func ExampleDimensionsClient_NewByExternalCloudProviderTypePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewDimensionsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewByExternalCloudProviderTypePager(armcostmanagement.ExternalCloudProviderTypeExternalBillingAccounts,
		"<external-cloud-provider-id>",
		&armcostmanagement.DimensionsClientByExternalCloudProviderTypeOptions{Filter: nil,
			Expand:    nil,
			Skiptoken: nil,
			Top:       nil,
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
