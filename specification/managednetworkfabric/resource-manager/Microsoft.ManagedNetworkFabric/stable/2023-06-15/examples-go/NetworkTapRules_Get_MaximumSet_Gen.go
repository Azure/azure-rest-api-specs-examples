package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkTapRules_Get_MaximumSet_Gen.json
func ExampleNetworkTapRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkTapRulesClient().Get(ctx, "example-rg", "example-tapRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkTapRule = armmanagednetworkfabric.NetworkTapRule{
	// 	Name: to.Ptr("example-tapRule"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/networkTapRules"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTapRules/example-tapRule"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-12T07:11:22.488Z"); return t}()),
	// 		CreatedBy: to.Ptr("email@address.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-12T07:11:22.488Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@mail.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastuseuap"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("keyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NetworkTapRuleProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		ConfigurationType: to.Ptr(armmanagednetworkfabric.ConfigurationTypeFile),
	// 		DynamicMatchConfigurations: []*armmanagednetworkfabric.CommonDynamicMatchConfiguration{
	// 			{
	// 				IPGroups: []*armmanagednetworkfabric.IPGroupProperties{
	// 					{
	// 						Name: to.Ptr("example-ipGroup1"),
	// 						IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
	// 						IPPrefixes: []*string{
	// 							to.Ptr("10.10.10.10/30")},
	// 					}},
	// 					PortGroups: []*armmanagednetworkfabric.PortGroupProperties{
	// 						{
	// 							Name: to.Ptr("example-portGroup1"),
	// 							Ports: []*string{
	// 								to.Ptr("100-200")},
	// 							},
	// 							{
	// 								Name: to.Ptr("example-portGroup2"),
	// 								Ports: []*string{
	// 									to.Ptr("900"),
	// 									to.Ptr("1000-2000")},
	// 							}},
	// 							VlanGroups: []*armmanagednetworkfabric.VlanGroupProperties{
	// 								{
	// 									Name: to.Ptr("exmaple-vlanGroup"),
	// 									Vlans: []*string{
	// 										to.Ptr("10"),
	// 										to.Ptr("100-200")},
	// 								}},
	// 						}},
	// 						MatchConfigurations: []*armmanagednetworkfabric.NetworkTapRuleMatchConfiguration{
	// 							{
	// 								Actions: []*armmanagednetworkfabric.NetworkTapRuleAction{
	// 									{
	// 										Type: to.Ptr(armmanagednetworkfabric.TapRuleActionTypeDrop),
	// 										DestinationID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/neighborGroups/example-neighborGroup"),
	// 										IsTimestampEnabled: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
	// 										MatchConfigurationName: to.Ptr("match1"),
	// 										Truncate: to.Ptr("100"),
	// 								}},
	// 								IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
	// 								MatchConditions: []*armmanagednetworkfabric.NetworkTapRuleMatchCondition{
	// 									{
	// 										IPCondition: &armmanagednetworkfabric.IPMatchCondition{
	// 											Type: to.Ptr(armmanagednetworkfabric.SourceDestinationTypeSourceIP),
	// 											IPGroupNames: []*string{
	// 												to.Ptr("example-ipGroup")},
	// 												IPPrefixValues: []*string{
	// 													to.Ptr("10.10.10.10/20")},
	// 													PrefixType: to.Ptr(armmanagednetworkfabric.PrefixTypePrefix),
	// 												},
	// 												ProtocolTypes: []*string{
	// 													to.Ptr("TCP")},
	// 													VlanMatchCondition: &armmanagednetworkfabric.VlanMatchCondition{
	// 														InnerVlans: []*string{
	// 															to.Ptr("11-20")},
	// 															VlanGroupNames: []*string{
	// 																to.Ptr("exmaple-vlanGroup")},
	// 																Vlans: []*string{
	// 																	to.Ptr("10")},
	// 																},
	// 																EncapsulationType: to.Ptr(armmanagednetworkfabric.EncapsulationTypeNone),
	// 																PortCondition: &armmanagednetworkfabric.PortCondition{
	// 																	Layer4Protocol: to.Ptr(armmanagednetworkfabric.Layer4ProtocolTCP),
	// 																	PortGroupNames: []*string{
	// 																		to.Ptr("example-portGroup1")},
	// 																		PortType: to.Ptr(armmanagednetworkfabric.PortTypeSourcePort),
	// 																		Ports: []*string{
	// 																			to.Ptr("100")},
	// 																		},
	// 																}},
	// 																MatchConfigurationName: to.Ptr("config1"),
	// 																SequenceNumber: to.Ptr[int64](10),
	// 														}},
	// 														TapRulesURL: to.Ptr("https://microsoft.com/a"),
	// 														AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 														ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 														LastSyncedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-12T07:11:22.485Z"); return t}()),
	// 														NetworkTapID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTaps/example-tap"),
	// 														PollingIntervalInSeconds: to.Ptr(armmanagednetworkfabric.PollingIntervalInSeconds(30)),
	// 														ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 													},
	// 												}
}
