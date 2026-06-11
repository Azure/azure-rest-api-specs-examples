package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/InternalNetworks_UpdateBfdAdministrativeState.json
func ExampleInternalNetworksClient_BeginUpdateBfdAdministrativeState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInternalNetworksClient().BeginUpdateBfdAdministrativeState(ctx, "example-rg", "example-l3isd", "example-internalnetwork", armmanagednetworkfabric.InternalNetworkUpdateBfdAdministrativeStateRequest{
		RouteType:           to.Ptr(armmanagednetworkfabric.InternalNetworkRouteTypeStatic),
		NeighborAddress:     to.Ptr("10.10.1.10"),
		AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeState("Enable")),
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
	// res = armmanagednetworkfabric.InternalNetworksClientUpdateBfdAdministrativeStateResponse{
	// 	InternalNetworkUpdateBfdAdministrativeStateResponse: &armmanagednetworkfabric.InternalNetworkUpdateBfdAdministrativeStateResponse{
	// 		Status: to.Ptr("Failed"),
	// 		Error: &armmanagednetworkfabric.ErrorDetail{
	// 			Code: to.Ptr("bidbvvmy"),
	// 			Message: to.Ptr("ekbarfd"),
	// 			Target: to.Ptr("wpnbisaqlnidiilkhaadmu"),
	// 			Details: []*armmanagednetworkfabric.ErrorDetail{
	// 			},
	// 			AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 				{
	// 					Type: to.Ptr("qakayuauodscmplibrcmvrnvfnia"),
	// 					Info: map[string]any{
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Properties: &armmanagednetworkfabric.InternalNetworkUpdateBfdAdministrativeStateResponseProperties{
	// 			NeighborAddressAdministrativeStatus: []*armmanagednetworkfabric.NeighborAddressBfdAdministrativeStatus{
	// 				{
	// 					NeighborAddress: to.Ptr("10.10.1.10"),
	// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeState("Enable")),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
