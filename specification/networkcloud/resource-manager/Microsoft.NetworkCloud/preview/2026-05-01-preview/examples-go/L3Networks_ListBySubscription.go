package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: 2026-05-01-preview/L3Networks_ListBySubscription.json
func ExampleL3NetworksClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("123e4567-e89b-12d3-a456-426655440000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewL3NetworksClient().NewListBySubscriptionPager(nil)
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
		// page = armnetworkcloud.L3NetworksClientListBySubscriptionResponse{
		// 	L3NetworkList: armnetworkcloud.L3NetworkList{
		// 		NextLink: to.Ptr("https://fully.qualified.hyperlink"),
		// 		Value: []*armnetworkcloud.L3Network{
		// 			{
		// 				ExtendedLocation: &armnetworkcloud.ExtendedLocation{
		// 					Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
		// 					Type: to.Ptr("CustomLocation"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 				Location: to.Ptr("location"),
		// 				Name: to.Ptr("l3NetworkName"),
		// 				Properties: &armnetworkcloud.L3NetworkProperties{
		// 					AssociatedResourceIDs: []*string{
		// 						to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/virtualMachines/virtualMachineName"),
		// 					},
		// 					ClusterID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
		// 					DetailedStatus: to.Ptr(armnetworkcloud.L3NetworkDetailedStatusAvailable),
		// 					DetailedStatusMessage: to.Ptr("L3 network is up"),
		// 					InterfaceName: to.Ptr("eth0"),
		// 					IPAllocationType: to.Ptr(armnetworkcloud.IPAllocationTypeDualStack),
		// 					IPv4ConnectedPrefix: to.Ptr("198.51.100.0/24"),
		// 					IPv6ConnectedPrefix: to.Ptr("2001:db8::/64"),
		// 					L3IsolationDomainID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/l3IsolationDomainName"),
		// 					ProvisioningState: to.Ptr(armnetworkcloud.L3NetworkProvisioningStateSucceeded),
		// 					Vlan: to.Ptr[int64](12),
		// 				},
		// 				SystemData: &armnetworkcloud.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
		// 					CreatedBy: to.Ptr("identityA"),
		// 					CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("identityB"),
		// 					LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("myvalue1"),
		// 					"key2": to.Ptr("myvalue2"),
		// 				},
		// 				Type: to.Ptr("Microsoft.NetworkCloud/l3Networks"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
