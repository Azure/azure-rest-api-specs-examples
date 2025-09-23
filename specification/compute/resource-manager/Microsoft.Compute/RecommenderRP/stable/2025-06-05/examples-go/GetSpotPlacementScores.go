package armrecommender_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armrecommender"
)

// Generated from example definition: 2025-06-05/GetSpotPlacementScores.json
func ExampleSpotPlacementScoresClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecommender.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSpotPlacementScoresClient().Get(ctx, "eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecommender.SpotPlacementScoresClientGetResponse{
	// 	ComputeDiagnosticBase: &armrecommender.ComputeDiagnosticBase{
	// 		Name: to.Ptr("spotPlacementScores"),
	// 		Type: to.Ptr("Microsoft.Compute/placementScores"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/placementScores/spot?api-version=2025-06-05"),
	// 		Properties: &armrecommender.DiagnosticProperties{
	// 			SupportedResourceTypes: []*string{
	// 				to.Ptr("Microsoft.Compute/virtualMachines"),
	// 			},
	// 		},
	// 	},
	// }
}
