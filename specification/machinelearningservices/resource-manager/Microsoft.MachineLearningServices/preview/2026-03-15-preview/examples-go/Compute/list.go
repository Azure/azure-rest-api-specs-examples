package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Compute/list.json
func ExampleComputeClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewComputeClient().NewListPager("testrg123", "workspaces123", nil)
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
		// page = armmachinelearning.ComputeClientListResponse{
		// 	PaginatedComputeResourcesList: armmachinelearning.PaginatedComputeResourcesList{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspaces123/computes?api-version=2025-07-01-preview&$skip=2"),
		// 		Value: []*armmachinelearning.ComputeResource{
		// 			{
		// 				Name: to.Ptr("compute123"),
		// 				Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/computes"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspaces123/computes/compute123"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armmachinelearning.AKS{
		// 					Description: to.Ptr("some compute"),
		// 					ComputeType: to.Ptr(armmachinelearning.ComputeTypeAKS),
		// 					CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T22:00:00.0000000+00:00"); return t}()),
		// 					ModifiedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T22:00:00.0000000+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armmachinelearning.ProvisioningStateSucceeded),
		// 					ResourceID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/Microsoft.ContainerService/managedClusters/compute123-56826-c9b00420020b2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("compute1234"),
		// 				Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/computes"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspaces123/computes/compute1234"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armmachinelearning.AKS{
		// 					Description: to.Ptr("some compute"),
		// 					ComputeType: to.Ptr(armmachinelearning.ComputeTypeAKS),
		// 					CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T22:00:00.0000000+00:00"); return t}()),
		// 					ModifiedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T22:00:00.0000000+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armmachinelearning.ProvisioningStateSucceeded),
		// 					ResourceID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/Microsoft.ContainerService/managedClusters/compute1234-56826-c9b00420020b2"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
