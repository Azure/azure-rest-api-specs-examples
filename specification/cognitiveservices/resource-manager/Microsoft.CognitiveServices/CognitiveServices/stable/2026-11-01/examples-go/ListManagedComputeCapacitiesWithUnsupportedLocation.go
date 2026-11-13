package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-11-01/ListManagedComputeCapacitiesWithUnsupportedLocation.json
func ExampleManagedComputeCapacitiesClient_NewListPager_listManagedComputeCapacitiesWithUnsupportedDataZoneLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedComputeCapacitiesClient().NewListPager("MaaP", &armcognitiveservices.ManagedComputeCapacitiesClientListOptions{
		AcceleratorType: to.Ptr("H100_80GB"),
		Location:        to.Ptr("japaneast")})
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
		// page = armcognitiveservices.ManagedComputeCapacitiesClientListResponse{
		// 	ManagedComputeCapacityListResult: armcognitiveservices.ManagedComputeCapacityListResult{
		// 		Value: []*armcognitiveservices.ManagedComputeCapacity{
		// 			{
		// 				ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/providers/Microsoft.CognitiveServices/managedComputeCapacities/H100_80GB.Global"),
		// 				Name: to.Ptr("H100_80GB.Global"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/managedComputeCapacities"),
		// 				Properties: &armcognitiveservices.ManagedComputeCapacityProperties{
		// 					AcceleratorType: to.Ptr("H100_80GB"),
		// 					OfferScope: to.Ptr("Global"),
		// 					ScopeID: to.Ptr(""),
		// 					AvailableAccelerators: to.Ptr[int32](16),
		// 					TotalAvailableAccelerators: to.Ptr[int32](32),
		// 					DeploymentSizeCapacities: []*armcognitiveservices.DeploymentSizeCapacity{
		// 						{
		// 							ModelInstanceAcceleratorCount: to.Ptr[int32](1),
		// 							TotalAvailableCapacity: to.Ptr[int32](16),
		// 							LargestDeploymentCapacity: to.Ptr[int32](8),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
