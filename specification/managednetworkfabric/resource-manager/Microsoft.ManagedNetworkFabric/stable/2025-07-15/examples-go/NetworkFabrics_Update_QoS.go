package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkFabrics_Update_QoS.json
func ExampleNetworkFabricsClient_BeginUpdate_networkFabricsUpdateQoSEnable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkFabricsClient().BeginUpdate(ctx, "example-rg", "example-fabric", armmanagednetworkfabric.NetworkFabricPatch{
		Properties: &armmanagednetworkfabric.NetworkFabricPatchProperties{
			QosConfiguration: &armmanagednetworkfabric.QosPatchProperties{
				QosConfigurationState: to.Ptr(armmanagednetworkfabric.QosConfigurationStateEnabled),
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
	// res = armmanagednetworkfabric.NetworkFabricsClientUpdateResponse{
	// 	NetworkFabric: &armmanagednetworkfabric.NetworkFabric{
	// 		Properties: &armmanagednetworkfabric.NetworkFabricProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			NetworkFabricSKU: to.Ptr("M4-A400-A100-C16-aa"),
	// 			FabricVersion: to.Ptr("version1"),
	// 			RouterIDs: []*string{
	// 				to.Ptr("routerId"),
	// 			},
	// 			NetworkFabricControllerID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/example-fabricController"),
	// 			RackCount: to.Ptr[int32](4),
	// 			ServerCountPerRack: to.Ptr[int32](8),
	// 			IPv4Prefix: to.Ptr("10.18.0.0/19"),
	// 			IPv6Prefix: to.Ptr("3FFE:FFFF:0:CD40::/59"),
	// 			FabricASN: to.Ptr[int64](29249),
	// 			TerminalServerConfiguration: &armmanagednetworkfabric.TerminalServerConfiguration{
	// 				Username: to.Ptr("username"),
	// 				SerialNumber: to.Ptr("123456"),
	// 				PrimaryIPv4Prefix: to.Ptr("10.0.0.12/30"),
	// 				PrimaryIPv6Prefix: to.Ptr("4FFE:FFFF:0:CD30::a8/127"),
	// 				SecondaryIPv4Prefix: to.Ptr("40.0.0.14/30"),
	// 				SecondaryIPv6Prefix: to.Ptr("6FFE:FFFF:0:CD30::ac/127"),
	// 				NetworkDeviceID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice"),
	// 			},
	// 			ManagementNetworkConfiguration: &armmanagednetworkfabric.ManagementNetworkConfigurationProperties{
	// 				InfrastructureVPNConfiguration: &armmanagednetworkfabric.VPNConfigurationProperties{
	// 					NetworkToNetworkInterconnectID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
	// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 					PeeringOption: to.Ptr(armmanagednetworkfabric.PeeringOptionOptionA),
	// 					OptionAProperties: &armmanagednetworkfabric.VPNOptionAProperties{
	// 						Mtu: to.Ptr[int32](1501),
	// 						VlanID: to.Ptr[int32](3001),
	// 						PeerASN: to.Ptr[int64](1235),
	// 						BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
	// 							AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeStateEnabled),
	// 							IntervalInMilliSeconds: to.Ptr[int32](300),
	// 							Multiplier: to.Ptr[int32](10),
	// 						},
	// 						PrimaryIPv4Prefix: to.Ptr("10.0.0.12/30"),
	// 						PrimaryIPv6Prefix: to.Ptr("4FFE:FFFF:0:CD30::a8/127"),
	// 						SecondaryIPv4Prefix: to.Ptr("20.0.0.13/30"),
	// 						SecondaryIPv6Prefix: to.Ptr("6FFE:FFFF:0:CD30::ac/127"),
	// 					},
	// 				},
	// 				WorkloadVPNConfiguration: &armmanagednetworkfabric.VPNConfigurationProperties{
	// 					NetworkToNetworkInterconnectID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
	// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 					PeeringOption: to.Ptr(armmanagednetworkfabric.PeeringOptionOptionA),
	// 					OptionAProperties: &armmanagednetworkfabric.VPNOptionAProperties{
	// 						Mtu: to.Ptr[int32](1500),
	// 						VlanID: to.Ptr[int32](3000),
	// 						PeerASN: to.Ptr[int64](61234),
	// 						BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
	// 							AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeStateEnabled),
	// 							IntervalInMilliSeconds: to.Ptr[int32](300),
	// 							Multiplier: to.Ptr[int32](5),
	// 						},
	// 						PrimaryIPv4Prefix: to.Ptr("10.0.0.14/30"),
	// 						PrimaryIPv6Prefix: to.Ptr("2FFE:FFFF:0:CD30::a7/126"),
	// 						SecondaryIPv4Prefix: to.Ptr("10.0.0.15/30"),
	// 						SecondaryIPv6Prefix: to.Ptr("2FFE:FFFF:0:CD30::ac/126"),
	// 					},
	// 				},
	// 			},
	// 			Racks: []*string{
	// 				to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkRacks/example-networkRack"),
	// 			},
	// 			L2IsolationDomains: []*string{
	// 				to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/example-l2Domain"),
	// 			},
	// 			L3IsolationDomains: []*string{
	// 				to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3Domain"),
	// 			},
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 			QosConfiguration: &armmanagednetworkfabric.QosProperties{
	// 				QosConfigurationState: to.Ptr(armmanagednetworkfabric.QosConfigurationStateEnabled),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 		},
	// 		Location: to.Ptr("eastuseuap"),
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
	// 		Name: to.Ptr("example-fabric"),
	// 		Type: to.Ptr("Microsoft.ManagedNetworkFabric/networkFabrics"),
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
