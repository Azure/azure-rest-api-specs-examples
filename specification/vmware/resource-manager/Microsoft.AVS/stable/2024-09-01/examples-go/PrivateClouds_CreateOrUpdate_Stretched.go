package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2024-09-01/PrivateClouds_CreateOrUpdate_Stretched.json
func ExamplePrivateCloudsClient_BeginCreateOrUpdate_privateCloudsCreateOrUpdateStretched() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateCloudsClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", armavs.PrivateCloud{
		Location: to.Ptr("eastus2"),
		SKU: &armavs.SKU{
			Name: to.Ptr("AV36"),
		},
		Properties: &armavs.PrivateCloudProperties{
			Availability: &armavs.AvailabilityProperties{
				Strategy:      to.Ptr(armavs.AvailabilityStrategyDualZone),
				Zone:          to.Ptr[int32](1),
				SecondaryZone: to.Ptr[int32](2),
			},
			NetworkBlock: to.Ptr("192.168.48.0/22"),
			ManagementCluster: &armavs.ManagementCluster{
				ClusterSize: to.Ptr[int32](4),
			},
		},
		Tags: map[string]*string{},
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
	// res = armavs.PrivateCloudsClientCreateOrUpdateResponse{
	// 	PrivateCloud: &armavs.PrivateCloud{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1"),
	// 		Location: to.Ptr("eastus2"),
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 			to.Ptr("2"),
	// 		},
	// 		Name: to.Ptr("cloud1"),
	// 		SKU: &armavs.SKU{
	// 			Name: to.Ptr("AV36"),
	// 		},
	// 		Properties: &armavs.PrivateCloudProperties{
	// 			Availability: &armavs.AvailabilityProperties{
	// 				Strategy: to.Ptr(armavs.AvailabilityStrategyDualZone),
	// 				Zone: to.Ptr[int32](1),
	// 				SecondaryZone: to.Ptr[int32](2),
	// 			},
	// 			NetworkBlock: to.Ptr("192.168.48.0/22"),
	// 			Circuit: &armavs.Circuit{
	// 				ExpressRouteID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2/providers/Microsoft.Network/expressroutecircuits/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2-xconnect"),
	// 				PrimarySubnet: to.Ptr("192.168.53.0/30"),
	// 				SecondarySubnet: to.Ptr("192.168.53.4/30"),
	// 			},
	// 			SecondaryCircuit: &armavs.Circuit{
	// 				ExpressRouteID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2/providers/Microsoft.Network/expressroutecircuits/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2-xconnect2"),
	// 				PrimarySubnet: to.Ptr("192.168.53.0/30"),
	// 				SecondarySubnet: to.Ptr("192.168.53.4/30"),
	// 			},
	// 			ManagementCluster: &armavs.ManagementCluster{
	// 				ClusterID: to.Ptr[int32](1),
	// 				ClusterSize: to.Ptr[int32](4),
	// 				Hosts: []*string{
	// 					to.Ptr("fakehost18.nyc1.kubernetes.center"),
	// 					to.Ptr("fakehost19.nyc1.kubernetes.center"),
	// 					to.Ptr("fakehost20.nyc1.kubernetes.center"),
	// 					to.Ptr("fakehost21.nyc1.kubernetes.center"),
	// 				},
	// 			},
	// 			Endpoints: &armavs.Endpoints{
	// 				NsxtManager: to.Ptr("https://nsx.290351365f5b41a19b77af.eastus.avslab.azure.com/"),
	// 				Vcsa: to.Ptr("https://vc.290351365f5b41a19b77af.eastus.avslab.azure.com/"),
	// 				HcxCloudManager: to.Ptr("https://hcx.290351365f5b41a19b77af.eastus.avslab.azure.com/"),
	// 				NsxtManagerIP: to.Ptr("192.168.50.3"),
	// 				VcenterIP: to.Ptr("192.168.50.2"),
	// 				HcxCloudManagerIP: to.Ptr("192.168.50.4"),
	// 			},
	// 			ExternalCloudLinks: []*string{
	// 				to.Ptr("/subscriptions/12341234-1234-1234-1234-123412341234/resourceGroups/mygroup/providers/Microsoft.AVS/privateClouds/cloud2"),
	// 			},
	// 			IdentitySources: []*armavs.IdentitySource{
	// 				{
	// 					Alias: to.Ptr("groupAlias"),
	// 					BaseGroupDN: to.Ptr("ou=baseGroup"),
	// 					BaseUserDN: to.Ptr("ou=baseUser"),
	// 					Domain: to.Ptr("domain1"),
	// 					Name: to.Ptr("group1"),
	// 					PrimaryServer: to.Ptr("ldaps://1.1.1.1:636/"),
	// 					SecondaryServer: to.Ptr("ldaps://1.1.1.2:636/"),
	// 					SSL: to.Ptr(armavs.SSLEnumEnabled),
	// 				},
	// 			},
	// 			Internet: to.Ptr(armavs.InternetEnumDisabled),
	// 			ProvisioningState: to.Ptr(armavs.PrivateCloudProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds"),
	// 	},
	// }
}
