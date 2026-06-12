package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: 2026-05-01-preview/AccessBridges_Get.json
func ExampleAccessBridgesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("123e4567-e89b-12d3-a456-426655440000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessBridgesClient().Get(ctx, "resourceGroupName", armnetworkcloud.AccessBridgeAllowedNameBastion, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetworkcloud.AccessBridgesClientGetResponse{
	// 	AccessBridge: armnetworkcloud.AccessBridge{
	// 		ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 			Type: to.Ptr("CustomLocation"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/accessBridges/accessBridgeName"),
	// 		Location: to.Ptr("location"),
	// 		Name: to.Ptr("accessBridgeName"),
	// 		Properties: &armnetworkcloud.AccessBridgeProperties{
	// 			DetailedStatus: to.Ptr(armnetworkcloud.AccessBridgeDetailedStatusRunning),
	// 			DetailedStatusMessage: to.Ptr("The access bridge is serving traffic from the VIP and storage appliance endpoints."),
	// 			Endpoints: []*armnetworkcloud.AccessBridgeEndpoint{
	// 				{
	// 					Fqdn: to.Ptr("bastion.contoso.example.com"),
	// 					IPv4Address: to.Ptr("192.51.100.10"),
	// 					IPv6Address: to.Ptr("2001:db8::10"),
	// 					Name: to.Ptr("vip"),
	// 				},
	// 			},
	// 			IPv4ConnectedPrefix: to.Ptr("198.51.100.0/24"),
	// 			IPv6ConnectedPrefix: to.Ptr("2001:db8::/64"),
	// 			NetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/l3IsolationDomainName/internalNetworks/internalNetworkName"),
	// 			Protocol: to.Ptr(armnetworkcloud.TransportProtocolTCP),
	// 			ProvisioningState: to.Ptr(armnetworkcloud.AccessBridgeProvisioningStateSucceeded),
	// 			SecurityRules: []*armnetworkcloud.AccessBridgeSecurityRule{
	// 				{
	// 					Description: to.Ptr("Allow management plane egress"),
	// 					Direction: to.Ptr(armnetworkcloud.SecurityRuleDirectionOutbound),
	// 					IPv4Addresses: []*string{
	// 						to.Ptr("10.10.20.10-10.10.20.20"),
	// 					},
	// 					IPv6Addresses: []*string{
	// 						to.Ptr("2001:db8:abcd:12::1000-2001:db8:abcd:12::1fff"),
	// 					},
	// 					Port: to.Ptr("24562-24570"),
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armnetworkcloud.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 			CreatedBy: to.Ptr("identityA"),
	// 			CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("identityB"),
	// 			LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("myvalue1"),
	// 			"key2": to.Ptr("myvalue2"),
	// 		},
	// 		Type: to.Ptr("Microsoft.NetworkCloud/accessBridges"),
	// 	},
	// }
}
