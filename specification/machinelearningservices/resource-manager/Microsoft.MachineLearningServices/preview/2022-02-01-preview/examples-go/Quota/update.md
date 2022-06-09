```go
package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Quota/update.json
func ExampleQuotasClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewQuotasClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"eastus",
		armmachinelearning.QuotaUpdateParameters{
			Value: []*armmachinelearning.QuotaBaseProperties{
				{
					Type:  to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
					ID:    to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs"),
					Limit: to.Ptr[int64](100),
					Unit:  to.Ptr(armmachinelearning.QuotaUnitCount),
				},
				{
					Type:  to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
					ID:    to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs"),
					Limit: to.Ptr[int64](200),
					Unit:  to.Ptr(armmachinelearning.QuotaUnitCount),
				}},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearning%2Farmmachinelearning%2Fv2.0.0-beta.1/sdk/resourcemanager/machinelearning/armmachinelearning/README.md) on how to add the SDK to your project and authenticate.
