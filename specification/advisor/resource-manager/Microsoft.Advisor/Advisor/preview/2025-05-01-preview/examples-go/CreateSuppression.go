package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor/v2"
)

// Generated from example definition: 2025-05-01-preview/CreateSuppression.json
func ExampleSuppressionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSuppressionsClient().Create(ctx, "resourceUri", "recommendationId", "suppressionName1", armadvisor.SuppressionContract{
		Properties: &armadvisor.SuppressionProperties{
			TTL: to.Ptr("07:00:00:00"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armadvisor.SuppressionsClientCreateResponse{
	// 	SuppressionContract: armadvisor.SuppressionContract{
	// 		ID: to.Ptr("/subscriptions/a5481ee1-95df-47d0-85d4-dd3f0dfa19bc/resourceGroups/resourceGroup/providers/Microsoft.Compute/availabilitysets/armavset/providers/Microsoft.Advisor/recommendations/bd27ddc6-1312-4067-b4af-cbb45e32cfd7/suppressions/HardcodedSuppressionName"),
	// 		Name: to.Ptr("suppressionName1"),
	// 		Type: to.Ptr("Microsoft.Advisor/suppressions"),
	// 		Properties: &armadvisor.SuppressionProperties{
	// 			SuppressionID: to.Ptr("suppressionId"),
	// 			TTL: to.Ptr("07:00:00:00"),
	// 		},
	// 	},
	// }
}
