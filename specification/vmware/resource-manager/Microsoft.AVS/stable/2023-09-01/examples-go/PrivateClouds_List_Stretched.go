package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/PrivateClouds_List_Stretched.json
func ExamplePrivateCloudsClient_NewListPager_privateCloudsListStretched() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateCloudsClient().NewListPager("group1", nil)
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
		// page.PrivateCloudList = armavs.PrivateCloudList{
		// 	Value: []*armavs.PrivateCloud{
		// 		{
		// 			Name: to.Ptr("cloud1"),
		// 			Type: to.Ptr("Microsoft.AVS/privateClouds"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1"),
		// 			Location: to.Ptr("eastus2"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armavs.PrivateCloudProperties{
		// 				Availability: &armavs.AvailabilityProperties{
		// 					SecondaryZone: to.Ptr[int32](2),
		// 					Strategy: to.Ptr(armavs.AvailabilityStrategyDualZone),
		// 					Zone: to.Ptr[int32](1),
		// 				},
		// 				Circuit: &armavs.Circuit{
		// 					ExpressRouteID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2/providers/Microsoft.Network/expressroutecircuits/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2-xconnect"),
		// 					ExpressRoutePrivatePeeringID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tnt42-cust-p01-dmo01/providers/Microsoft.Network/expressroutecircuits/tnt42-cust-p01-dmo01-er/peerings/AzurePrivatePeering"),
		// 					PrimarySubnet: to.Ptr("192.168.53.0/30"),
		// 					SecondarySubnet: to.Ptr("192.168.53.4/30"),
		// 				},
		// 				Endpoints: &armavs.Endpoints{
		// 					HcxCloudManager: to.Ptr("https://hcx.290351365f5b41a19b77af.eastus.avslab.azure.com/"),
		// 					HcxCloudManagerIP: to.Ptr("192.168.50.4"),
		// 					NsxtManager: to.Ptr("https://nsx.290351365f5b41a19b77af.eastus.avslab.azure.com/"),
		// 					NsxtManagerIP: to.Ptr("192.168.50.3"),
		// 					VcenterIP: to.Ptr("192.168.50.2"),
		// 					Vcsa: to.Ptr("https://vc.290351365f5b41a19b77af.eastus.avslab.azure.com/"),
		// 				},
		// 				ExternalCloudLinks: []*string{
		// 					to.Ptr("/subscriptions/12341234-1234-1234-1234-123412341234/resourceGroups/mygroup/providers/Microsoft.AVS/privateClouds/cloud2")},
		// 					IdentitySources: []*armavs.IdentitySource{
		// 						{
		// 							Name: to.Ptr("group1"),
		// 							Alias: to.Ptr("groupAlias"),
		// 							BaseGroupDN: to.Ptr("ou=baseGroup"),
		// 							BaseUserDN: to.Ptr("ou=baseUser"),
		// 							Domain: to.Ptr("domain1"),
		// 							PrimaryServer: to.Ptr("ldaps://1.1.1.1:636/"),
		// 							SecondaryServer: to.Ptr("ldaps://1.1.1.2:636/"),
		// 							SSL: to.Ptr(armavs.SSLEnumEnabled),
		// 					}},
		// 					Internet: to.Ptr(armavs.InternetEnumDisabled),
		// 					ManagementCluster: &armavs.ManagementCluster{
		// 						ClusterID: to.Ptr[int32](1),
		// 						ClusterSize: to.Ptr[int32](4),
		// 						Hosts: []*string{
		// 							to.Ptr("fakehost18.nyc1.kubernetes.center"),
		// 							to.Ptr("fakehost19.nyc1.kubernetes.center"),
		// 							to.Ptr("fakehost20.nyc1.kubernetes.center"),
		// 							to.Ptr("fakehost21.nyc1.kubernetes.center")},
		// 						},
		// 						NetworkBlock: to.Ptr("192.168.48.0/22"),
		// 						ProvisioningState: to.Ptr(armavs.PrivateCloudProvisioningStateSucceeded),
		// 						SecondaryCircuit: &armavs.Circuit{
		// 							ExpressRouteID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2/providers/Microsoft.Network/expressroutecircuits/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2-xconnect2"),
		// 							ExpressRoutePrivatePeeringID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tnt42-cust-p01-dmo01/providers/Microsoft.Network/expressroutecircuits/tnt42-cust-p01-dmo01-er2/peerings/AzurePrivatePeering"),
		// 							PrimarySubnet: to.Ptr("192.168.53.0/30"),
		// 							SecondarySubnet: to.Ptr("192.168.53.4/30"),
		// 						},
		// 					},
		// 					SKU: &armavs.SKU{
		// 						Name: to.Ptr("AV36"),
		// 					},
		// 			}},
		// 		}
	}
}
