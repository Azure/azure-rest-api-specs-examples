package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkPacketBrokers_Create.json
func ExampleNetworkPacketBrokersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkPacketBrokersClient().BeginCreate(ctx, "example-rg", "example-networkPacketBroker", armmanagednetworkfabric.NetworkPacketBroker{
		Properties: &armmanagednetworkfabric.NetworkPacketBrokerProperties{
			NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-networkFabric"),
		},
		Identity: &armmanagednetworkfabric.ManagedServiceIdentity{
			Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
				"key3673": {},
			},
		},
		Tags: map[string]*string{
			"keyId": to.Ptr("keyValue"),
		},
		Location: to.Ptr("eastuseuap"),
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
	// res = armmanagednetworkfabric.NetworkPacketBrokersClientCreateResponse{
	// 	NetworkPacketBroker: &armmanagednetworkfabric.NetworkPacketBroker{
	// 		Properties: &armmanagednetworkfabric.NetworkPacketBrokerProperties{
	// 			NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-networkFabric"),
	// 			NetworkDeviceIDs: []*string{
	// 				to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice"),
	// 			},
	// 			SourceInterfaceIDs: []*string{
	// 				to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice/networkInterfaces/example-networkInterface"),
	// 			},
	// 			NetworkTapIDs: []*string{
	// 				to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTaps/example-networkTap"),
	// 			},
	// 			NeighborGroupIDs: []*string{
	// 				to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/neighborGroups/example-neighborGroup"),
	// 			},
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 		},
	// 		Identity: &armmanagednetworkfabric.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 			TenantID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 			Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
	// 				"key3673": &armmanagednetworkfabric.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 					ClientID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"keyId": to.Ptr("keyValue"),
	// 		},
	// 		Location: to.Ptr("eastuseuap"),
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkPacketBrokers/example-networkPacketBroker"),
	// 		Name: to.Ptr("example-networkPacketBroker"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/networkPacketBrokers"),
	// 		SystemData: &armmanagednetworkfabric.SystemData{
	// 			CreatedBy: to.Ptr("email@address.com"),
	// 			CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("UserId"),
	// 			LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 		},
	// 	},
	// }
}
