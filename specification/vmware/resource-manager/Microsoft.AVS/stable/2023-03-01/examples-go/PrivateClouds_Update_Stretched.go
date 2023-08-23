package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a5e7ff51c8af3781e7f6dd3b82a3fc922e2f6f93/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/PrivateClouds_Update_Stretched.json
func ExamplePrivateCloudsClient_BeginUpdate_privateCloudsUpdateStretched() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateCloudsClient().BeginUpdate(ctx, "group1", "cloud1", armavs.PrivateCloudUpdate{
		Properties: &armavs.PrivateCloudUpdateProperties{
			ManagementCluster: &armavs.ManagementCluster{
				ClusterSize: to.Ptr[int32](4),
			},
		},
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
	// res.PrivateCloud = armavs.PrivateCloud{
	// 	Name: to.Ptr("cloud1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1"),
	// 	Location: to.Ptr("eastus2"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armavs.PrivateCloudProperties{
	// 		Availability: &armavs.AvailabilityProperties{
	// 			SecondaryZone: to.Ptr[int32](2),
	// 			Strategy: to.Ptr(armavs.AvailabilityStrategyDualZone),
	// 			Zone: to.Ptr[int32](1),
	// 		},
	// 		IdentitySources: []*armavs.IdentitySource{
	// 			{
	// 				Name: to.Ptr("group1"),
	// 				Alias: to.Ptr("groupAlias"),
	// 				BaseGroupDN: to.Ptr("ou=baseGroup"),
	// 				BaseUserDN: to.Ptr("ou=baseUser"),
	// 				Domain: to.Ptr("domain1"),
	// 				PrimaryServer: to.Ptr("ldaps://1.1.1.1:636/"),
	// 				SecondaryServer: to.Ptr("ldaps://1.1.1.2:636/"),
	// 				SSL: to.Ptr(armavs.SSLEnumEnabled),
	// 		}},
	// 		Internet: to.Ptr(armavs.InternetEnumDisabled),
	// 		ManagementCluster: &armavs.ManagementCluster{
	// 			ClusterID: to.Ptr[int32](1),
	// 			ClusterSize: to.Ptr[int32](4),
	// 			Hosts: []*string{
	// 				to.Ptr("fakehost18.nyc1.kubernetes.center"),
	// 				to.Ptr("fakehost19.nyc1.kubernetes.center"),
	// 				to.Ptr("fakehost20.nyc1.kubernetes.center"),
	// 				to.Ptr("fakehost21.nyc1.kubernetes.center")},
	// 			},
	// 			Circuit: &armavs.Circuit{
	// 				ExpressRouteID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2/providers/Microsoft.Network/expressroutecircuits/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2-xconnect"),
	// 				ExpressRoutePrivatePeeringID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tnt42-cust-p01-dmo01/providers/Microsoft.Network/expressroutecircuits/tnt42-cust-p01-dmo01-er/peerings/AzurePrivatePeering"),
	// 				PrimarySubnet: to.Ptr("192.168.53.0/30"),
	// 				SecondarySubnet: to.Ptr("192.168.53.4/30"),
	// 			},
	// 			Endpoints: &armavs.Endpoints{
	// 				HcxCloudManager: to.Ptr("https://192.168.50.4/"),
	// 				NsxtManager: to.Ptr("https://192.168.50.3/"),
	// 				Vcsa: to.Ptr("https://192.168.50.2/"),
	// 			},
	// 			ExternalCloudLinks: []*string{
	// 				to.Ptr("/subscriptions/12341234-1234-1234-1234-123412341234/resourceGroups/mygroup/providers/Microsoft.AVS/privateClouds/cloud2")},
	// 				NetworkBlock: to.Ptr("192.168.48.0/22"),
	// 				ProvisioningState: to.Ptr(armavs.PrivateCloudProvisioningStateSucceeded),
	// 				SecondaryCircuit: &armavs.Circuit{
	// 					ExpressRouteID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2/providers/Microsoft.Network/expressroutecircuits/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2-xconnect2"),
	// 					ExpressRoutePrivatePeeringID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tnt42-cust-p01-dmo01/providers/Microsoft.Network/expressroutecircuits/tnt42-cust-p01-dmo01-er2/peerings/AzurePrivatePeering"),
	// 					PrimarySubnet: to.Ptr("192.168.53.0/30"),
	// 					SecondarySubnet: to.Ptr("192.168.53.4/30"),
	// 				},
	// 			},
	// 			SKU: &armavs.SKU{
	// 				Name: to.Ptr("AV36"),
	// 			},
	// 		}
}
