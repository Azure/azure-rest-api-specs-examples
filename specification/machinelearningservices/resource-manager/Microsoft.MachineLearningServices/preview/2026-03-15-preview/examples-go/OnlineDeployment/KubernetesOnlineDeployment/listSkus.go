package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/OnlineDeployment/KubernetesOnlineDeployment/listSkus.json
func ExampleOnlineDeploymentsClient_NewListSKUsPager_listKubernetesOnlineDeploymentSkus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOnlineDeploymentsClient().NewListSKUsPager("test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", &armmachinelearning.OnlineDeploymentsClientListSKUsOptions{
		Count: to.Ptr[int32](1)})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armmachinelearning.OnlineDeploymentsClientListSKUsResponse{
		// 	SKUResourceArmPaginatedResult: armmachinelearning.SKUResourceArmPaginatedResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/my-aml-workspace/onlineEndpoints/testEndpointName/deployments/testDeploymentName/skus?api-version=2025-07-01-preview&$skip=2"),
		// 		Value: []*armmachinelearning.SKUResource{
		// 			{
		// 				Capacity: &armmachinelearning.SKUCapacity{
		// 					Default: to.Ptr[int32](1),
		// 					Maximum: to.Ptr[int32](1),
		// 					Minimum: to.Ptr[int32](1),
		// 					ScaleType: to.Ptr(armmachinelearning.SKUScaleTypeAutomatic),
		// 				},
		// 				ResourceType: to.Ptr("Microsoft.MachineLearning.Services/endpoints/deployments"),
		// 				SKU: &armmachinelearning.SKUSetting{
		// 					Name: to.Ptr("string"),
		// 					Tier: to.Ptr(armmachinelearning.SKUTierFree),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
