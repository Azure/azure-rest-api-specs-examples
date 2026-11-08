package armrecommender_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armrecommender"
)

// Generated from example definition: 2026-10-05-preview/GenerateComputePlacementScores.json
func ExampleComputePlacementScoresClient_Post_returnsComputeVMPlacementScoresForGivenConfigurations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecommender.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComputePlacementScoresClient().Post(ctx, "eastus", armrecommender.ComputePlacementScoresInput{
		AvailabilityZones: to.Ptr(true),
		DesiredCount:      to.Ptr[int32](1),
		DesiredLocations: []*string{
			to.Ptr("eastus"),
			to.Ptr("eastus2"),
		},
		DesiredSizes: []*armrecommender.ResourceSize{
			{
				SKU: to.Ptr("Standard_D2_v2"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecommender.ComputePlacementScoresClientPostResponse{
	// 	ComputePlacementScoresResponse: armrecommender.ComputePlacementScoresResponse{
	// 		AvailabilityZones: to.Ptr(true),
	// 		DesiredCount: to.Ptr[int32](1),
	// 		DesiredLocations: []*string{
	// 			to.Ptr("eastus"),
	// 			to.Ptr("eastus2"),
	// 		},
	// 		DesiredSizes: []*armrecommender.ResourceSize{
	// 			{
	// 				SKU: to.Ptr("Standard_D2_v2"),
	// 			},
	// 		},
	// 		PlacementScores: []*armrecommender.PlacementScore{
	// 			{
	// 				AvailabilityZone: to.Ptr("1"),
	// 				PlacementPreference: to.Ptr[int32](10),
	// 				Region: to.Ptr("eastus"),
	// 				Score: to.Ptr(armrecommender.PlacementScoreValueHigh),
	// 				SKU: to.Ptr("Standard_D2_v2"),
	// 			},
	// 			{
	// 				AvailabilityZone: to.Ptr("2"),
	// 				PlacementPreference: to.Ptr[int32](9),
	// 				Region: to.Ptr("eastus"),
	// 				Score: to.Ptr(armrecommender.PlacementScoreValueHigh),
	// 				SKU: to.Ptr("Standard_D2_v2"),
	// 			},
	// 			{
	// 				AvailabilityZone: to.Ptr("3"),
	// 				PlacementPreference: to.Ptr[int32](8),
	// 				Region: to.Ptr("eastus"),
	// 				Score: to.Ptr(armrecommender.PlacementScoreValueHigh),
	// 				SKU: to.Ptr("Standard_D2_v2"),
	// 			},
	// 			{
	// 				AvailabilityZone: to.Ptr("1"),
	// 				PlacementPreference: to.Ptr[int32](0),
	// 				Region: to.Ptr("eastus2"),
	// 				Score: to.Ptr(armrecommender.PlacementScoreValueDataNotFoundOrStale),
	// 				SKU: to.Ptr("Standard_D2_v2"),
	// 			},
	// 			{
	// 				AvailabilityZone: to.Ptr("2"),
	// 				PlacementPreference: to.Ptr[int32](10),
	// 				Region: to.Ptr("eastus2"),
	// 				Score: to.Ptr(armrecommender.PlacementScoreValueHigh),
	// 				SKU: to.Ptr("Standard_D2_v2"),
	// 			},
	// 			{
	// 				AvailabilityZone: to.Ptr("3"),
	// 				PlacementPreference: to.Ptr[int32](9),
	// 				Region: to.Ptr("eastus2"),
	// 				Score: to.Ptr(armrecommender.PlacementScoreValueHigh),
	// 				SKU: to.Ptr("Standard_D2_v2"),
	// 			},
	// 		},
	// 	},
	// }
}
