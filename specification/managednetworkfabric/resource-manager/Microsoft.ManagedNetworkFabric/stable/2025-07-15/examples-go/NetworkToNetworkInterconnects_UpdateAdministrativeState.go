package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkToNetworkInterconnects_UpdateAdministrativeState.json
func ExampleNetworkToNetworkInterconnectsClient_BeginUpdateAdministrativeState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkToNetworkInterconnectsClient().BeginUpdateAdministrativeState(ctx, "example-rg", "example-nf", "example-nni", armmanagednetworkfabric.UpdateAdministrativeState{
		ResourceIDs: []*string{
			to.Ptr("/subscriptions/0000ABCD-0A0B-0000-0000-000000ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-nf/networkToNetworkInterconnects/example-nni"),
		},
		State: to.Ptr(armmanagednetworkfabric.EnableDisableStateEnable),
	}, nil)
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
	// res = armmanagednetworkfabric.NetworkToNetworkInterconnectsClientUpdateAdministrativeStateResponse{
	// 	UpdateAdministrativeStateResponse: &armmanagednetworkfabric.UpdateAdministrativeStateResponse{
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
