package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkTapRules_UpdateAdministrativeState.json
func ExampleNetworkTapRulesClient_BeginUpdateAdministrativeState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkTapRulesClient().BeginUpdateAdministrativeState(ctx, "example-rg", "example-tapRule", armmanagednetworkfabric.UpdateAdministrativeState{
		ResourceIDs: []*string{
			to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTapRules/example-tapRule"),
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
	// res = armmanagednetworkfabric.NetworkTapRulesClientUpdateAdministrativeStateResponse{
	// 	CommonPostActionResponseForStateUpdate: &armmanagednetworkfabric.CommonPostActionResponseForStateUpdate{
	// 		Error: &armmanagednetworkfabric.ErrorDetail{
	// 			Code: to.Ptr("ConfigurationStateError"),
	// 			Message: to.Ptr("The operation resulted in a degraded state"),
	// 			Target: to.Ptr("NetworkTapRule"),
	// 			Details: []*armmanagednetworkfabric.ErrorDetail{
	// 				{
	// 					Code: to.Ptr("EnabledDegraded"),
	// 					Message: to.Ptr("The resource is in a degraded state"),
	// 					Target: to.Ptr("AdministrativeState"),
	// 				},
	// 			},
	// 			AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 				{
	// 					Type: to.Ptr("NetworkTapRuleInfo"),
	// 					Info: map[string]any{
	// 						"state": "EnabledDegraded",
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStatePendingAdministrativeUpdate),
	// 	},
	// }
}
