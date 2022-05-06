Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatalake-analytics%2Farmdatalakeanalytics%2Fv0.5.0/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatalakeanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/ComputePolicies_Update.json
func ExampleComputePoliciesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatalakeanalytics.NewComputePoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<compute-policy-name>",
		&armdatalakeanalytics.ComputePoliciesClientUpdateOptions{Parameters: &armdatalakeanalytics.UpdateComputePolicyParameters{
			Properties: &armdatalakeanalytics.UpdateComputePolicyProperties{
				MaxDegreeOfParallelismPerJob: to.Ptr[int32](11),
				MinPriorityPerJob:            to.Ptr[int32](31),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
