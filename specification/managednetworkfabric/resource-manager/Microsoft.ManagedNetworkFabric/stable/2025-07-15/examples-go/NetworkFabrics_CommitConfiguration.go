package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkFabrics_CommitConfiguration.json
func ExampleNetworkFabricsClient_BeginCommitConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkFabricsClient().BeginCommitConfiguration(ctx, "example-rg", "example-networkFabric", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmanagednetworkfabric.NetworkFabricsClientCommitConfigurationResponse{
	// 	CommitConfigurationResponse: &armmanagednetworkfabric.CommitConfigurationResponse{
	// 		Status: to.Ptr("Failed"),
	// 		Error: &armmanagednetworkfabric.ErrorDetail{
	// 			Code: to.Ptr(""),
	// 			Message: to.Ptr(""),
	// 			Target: to.Ptr(""),
	// 			Details: []*armmanagednetworkfabric.ErrorDetail{
	// 			},
	// 			AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 				{
	// 					Type: to.Ptr(""),
	// 					Info: map[string]any{
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
