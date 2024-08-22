package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Quota/update.json
func ExampleQuotasClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQuotasClient().Update(ctx, "eastus", armmachinelearning.QuotaUpdateParameters{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UpdateWorkspaceQuotasResult = armmachinelearning.UpdateWorkspaceQuotasResult{
	// 	Value: []*armmachinelearning.UpdateWorkspaceQuotas{
	// 		{
	// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs"),
	// 			Limit: to.Ptr[int64](100),
	// 			Status: to.Ptr(armmachinelearning.StatusSuccess),
	// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
	// 		},
	// 		{
	// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs"),
	// 			Limit: to.Ptr[int64](200),
	// 			Status: to.Ptr(armmachinelearning.StatusSuccess),
	// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
	// 	}},
	// }
}
