package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/DefaultCniNetworks_ListBySubscription.json
func ExampleDefaultCniNetworksClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDefaultCniNetworksClient().NewListBySubscriptionPager(nil)
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
		// page.DefaultCniNetworkList = armnetworkcloud.DefaultCniNetworkList{
		// 	Value: []*armnetworkcloud.DefaultCniNetwork{
		// 		{
		// 			Name: to.Ptr("defaultcniNetworkName"),
		// 			Type: to.Ptr("Microsoft.NetworkCloud/defaultcniNetworks"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/defaultcniNetworks/defaultcniNetworkName"),
		// 			SystemData: &armnetworkcloud.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
		// 				CreatedBy: to.Ptr("identityA"),
		// 				CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("identityB"),
		// 				LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("location"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("myvalue1"),
		// 				"key2": to.Ptr("myvalue2"),
		// 			},
		// 			ExtendedLocation: &armnetworkcloud.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armnetworkcloud.DefaultCniNetworkProperties{
		// 				ClusterID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
		// 				CniAsNumber: to.Ptr[int64](64512),
		// 				CniBgpConfiguration: &armnetworkcloud.CniBgpConfiguration{
		// 					BgpPeers: []*armnetworkcloud.BgpPeer{
		// 						{
		// 							AsNumber: to.Ptr[int64](64497),
		// 							PeerIP: to.Ptr("203.0.113.254"),
		// 					}},
		// 					CommunityAdvertisements: []*armnetworkcloud.CommunityAdvertisement{
		// 						{
		// 							Communities: []*string{
		// 								to.Ptr("64512:100")},
		// 								SubnetPrefix: to.Ptr("192.0.2.0/27"),
		// 						}},
		// 						ServiceExternalPrefixes: []*string{
		// 							to.Ptr("192.0.2.0/28")},
		// 							ServiceLoadBalancerPrefixes: []*string{
		// 								to.Ptr("192.0.2.16/28")},
		// 							},
		// 							DetailedStatus: to.Ptr(armnetworkcloud.DefaultCniNetworkDetailedStatusAvailable),
		// 							DetailedStatusMessage: to.Ptr("Default CNI network is up"),
		// 							FabricBgpPeers: []*armnetworkcloud.BgpPeer{
		// 								{
		// 									AsNumber: to.Ptr[int64](64496),
		// 									PeerIP: to.Ptr("203.0.113.2"),
		// 								},
		// 								{
		// 									AsNumber: to.Ptr[int64](64496),
		// 									PeerIP: to.Ptr("203.0.113.3"),
		// 							}},
		// 							HybridAksClustersAssociatedIDs: []*string{
		// 								to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/hybridAksClusters/hybridAksClusterName")},
		// 								InterfaceName: to.Ptr("eth0"),
		// 								IPAllocationType: to.Ptr(armnetworkcloud.IPAllocationTypeDualStack),
		// 								IPv4ConnectedPrefix: to.Ptr("203.0.113.0/24"),
		// 								IPv6ConnectedPrefix: to.Ptr("2001:db8:0:3::/64"),
		// 								L3IsolationDomainID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/l3IsolationDomainName"),
		// 								ProvisioningState: to.Ptr(armnetworkcloud.DefaultCniNetworkProvisioningStateSucceeded),
		// 								Vlan: to.Ptr[int64](12),
		// 							},
		// 					}},
		// 				}
	}
}
