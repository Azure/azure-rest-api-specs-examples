package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabrics_Get_MaximumSet_Gen.json
func ExampleNetworkFabricsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkFabricsClient().Get(ctx, "example-rg", "example-fabric", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkFabric = armmanagednetworkfabric.NetworkFabric{
	// 	Name: to.Ptr("example-fabric"),
	// 	Type: to.Ptr("Microsoft.ManagedNetworkFabric/networkFabrics"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-07T09:53:31.314Z"); return t}()),
	// 		CreatedBy: to.Ptr("email@email.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-07T09:53:31.314Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@mail.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastuseuap"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("KeyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NetworkFabricProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 		FabricASN: to.Ptr[int64](29249),
	// 		FabricVersion: to.Ptr("version1"),
	// 		IPv4Prefix: to.Ptr("10.18.0.0/19"),
	// 		IPv6Prefix: to.Ptr("3FFE:FFFF:0:CD40::/59"),
	// 		L2IsolationDomains: []*string{
	// 			to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/example-l2Domain")},
	// 			L3IsolationDomains: []*string{
	// 				to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3Domain")},
	// 				ManagementNetworkConfiguration: &armmanagednetworkfabric.ManagementNetworkConfigurationProperties{
	// 					InfrastructureVPNConfiguration: &armmanagednetworkfabric.VPNConfigurationProperties{
	// 						AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 						NetworkToNetworkInterconnectID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
	// 						OptionAProperties: &armmanagednetworkfabric.VPNConfigurationPropertiesOptionAProperties{
	// 							PrimaryIPv4Prefix: to.Ptr("10.0.0.12/30"),
	// 							PrimaryIPv6Prefix: to.Ptr("4FFE:FFFF:0:CD30::a8/127"),
	// 							SecondaryIPv4Prefix: to.Ptr("20.0.0.13/30"),
	// 							SecondaryIPv6Prefix: to.Ptr("6FFE:FFFF:0:CD30::ac/127"),
	// 							BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
	// 								AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeStateEnabled),
	// 								IntervalInMilliSeconds: to.Ptr[int32](300),
	// 								Multiplier: to.Ptr[int32](10),
	// 							},
	// 							Mtu: to.Ptr[int32](1501),
	// 							PeerASN: to.Ptr[int64](1235),
	// 							VlanID: to.Ptr[int32](3001),
	// 						},
	// 						OptionBProperties: &armmanagednetworkfabric.OptionBProperties{
	// 							ExportRouteTargets: []*string{
	// 								to.Ptr("65046:10050")},
	// 								ImportRouteTargets: []*string{
	// 									to.Ptr("65046:10050")},
	// 									RouteTargets: &armmanagednetworkfabric.RouteTargetInformation{
	// 										ExportIPv4RouteTargets: []*string{
	// 											to.Ptr("65046:10050")},
	// 											ExportIPv6RouteTargets: []*string{
	// 												to.Ptr("65046:10050")},
	// 												ImportIPv4RouteTargets: []*string{
	// 													to.Ptr("65046:10050")},
	// 													ImportIPv6RouteTargets: []*string{
	// 														to.Ptr("65046:10050")},
	// 													},
	// 												},
	// 												PeeringOption: to.Ptr(armmanagednetworkfabric.PeeringOptionOptionA),
	// 											},
	// 											WorkloadVPNConfiguration: &armmanagednetworkfabric.VPNConfigurationProperties{
	// 												AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 												NetworkToNetworkInterconnectID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
	// 												OptionAProperties: &armmanagednetworkfabric.VPNConfigurationPropertiesOptionAProperties{
	// 													PrimaryIPv4Prefix: to.Ptr("10.0.0.14/30"),
	// 													PrimaryIPv6Prefix: to.Ptr("2FFE:FFFF:0:CD30::a7/126"),
	// 													SecondaryIPv4Prefix: to.Ptr("10.0.0.15/30"),
	// 													SecondaryIPv6Prefix: to.Ptr("2FFE:FFFF:0:CD30::ac/126"),
	// 													BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
	// 														AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeStateEnabled),
	// 														IntervalInMilliSeconds: to.Ptr[int32](300),
	// 														Multiplier: to.Ptr[int32](5),
	// 													},
	// 													Mtu: to.Ptr[int32](1500),
	// 													PeerASN: to.Ptr[int64](61234),
	// 													VlanID: to.Ptr[int32](3000),
	// 												},
	// 												OptionBProperties: &armmanagednetworkfabric.OptionBProperties{
	// 													ExportRouteTargets: []*string{
	// 														to.Ptr("65046:10050")},
	// 														ImportRouteTargets: []*string{
	// 															to.Ptr("65046:10050")},
	// 															RouteTargets: &armmanagednetworkfabric.RouteTargetInformation{
	// 																ExportIPv4RouteTargets: []*string{
	// 																	to.Ptr("65046:10050")},
	// 																	ExportIPv6RouteTargets: []*string{
	// 																		to.Ptr("65046:10050")},
	// 																		ImportIPv4RouteTargets: []*string{
	// 																			to.Ptr("65046:10050")},
	// 																			ImportIPv6RouteTargets: []*string{
	// 																				to.Ptr("65046:10050")},
	// 																			},
	// 																		},
	// 																		PeeringOption: to.Ptr(armmanagednetworkfabric.PeeringOptionOptionA),
	// 																	},
	// 																},
	// 																NetworkFabricControllerID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/example-fabricController"),
	// 																NetworkFabricSKU: to.Ptr("M4-A400-A100-C16-aa"),
	// 																ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 																RackCount: to.Ptr[int32](4),
	// 																Racks: []*string{
	// 																	to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkRacks/example-networkRack")},
	// 																	RouterIDs: []*string{
	// 																		to.Ptr("routerId")},
	// 																		ServerCountPerRack: to.Ptr[int32](8),
	// 																		TerminalServerConfiguration: &armmanagednetworkfabric.TerminalServerConfiguration{
	// 																			PrimaryIPv4Prefix: to.Ptr("10.0.0.12/30"),
	// 																			PrimaryIPv6Prefix: to.Ptr("4FFE:FFFF:0:CD30::a8/127"),
	// 																			SecondaryIPv4Prefix: to.Ptr("40.0.0.14/30"),
	// 																			SecondaryIPv6Prefix: to.Ptr("6FFE:FFFF:0:CD30::ac/127"),
	// 																			SerialNumber: to.Ptr("123456"),
	// 																			Username: to.Ptr("username"),
	// 																			NetworkDeviceID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice"),
	// 																		},
	// 																	},
	// 																}
}
