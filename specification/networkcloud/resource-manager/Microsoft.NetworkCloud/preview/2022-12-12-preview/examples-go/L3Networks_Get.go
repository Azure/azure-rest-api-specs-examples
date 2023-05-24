package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/L3Networks_Get.json
func ExampleL3NetworksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewL3NetworksClient().Get(ctx, "resourceGroupName", "l3NetworkName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.L3Network = armnetworkcloud.L3Network{
	// 	Name: to.Ptr("l3NetworkName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/l3Networks"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.L3NetworkProperties{
	// 		ClusterID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.L3NetworkDetailedStatusAvailable),
	// 		DetailedStatusMessage: to.Ptr("L3 network is up"),
	// 		HybridAksClustersAssociatedIDs: []*string{
	// 			to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/hybridAksClusters/hybridAksClusterName")},
	// 			HybridAksIpamEnabled: to.Ptr(armnetworkcloud.HybridAksIpamEnabledTrue),
	// 			HybridAksPluginType: to.Ptr(armnetworkcloud.HybridAksPluginTypeDPDK),
	// 			InterfaceName: to.Ptr("eth0"),
	// 			IPAllocationType: to.Ptr(armnetworkcloud.IPAllocationTypeDualStack),
	// 			IPv4ConnectedPrefix: to.Ptr("198.51.100.0/24"),
	// 			IPv6ConnectedPrefix: to.Ptr("2001:db8::/64"),
	// 			L3IsolationDomainID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/l3IsolationDomainName"),
	// 			ProvisioningState: to.Ptr(armnetworkcloud.L3NetworkProvisioningStateSucceeded),
	// 			VirtualMachinesAssociatedIDs: []*string{
	// 				to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/virtualMachines/virtualMachineName")},
	// 				Vlan: to.Ptr[int64](12),
	// 			},
	// 		}
}
