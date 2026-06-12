package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkTapRules_ListBySubscription.json
func ExampleNetworkTapRulesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkTapRulesClient().NewListBySubscriptionPager(nil)
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
		// page = armmanagednetworkfabric.NetworkTapRulesClientListBySubscriptionResponse{
		// 	NetworkTapRulesListResult: armmanagednetworkfabric.NetworkTapRulesListResult{
		// 		Value: []*armmanagednetworkfabric.NetworkTapRule{
		// 			{
		// 				Properties: &armmanagednetworkfabric.NetworkTapRuleProperties{
		// 					Annotation: to.Ptr("annotation"),
		// 					ConfigurationType: to.Ptr(armmanagednetworkfabric.ConfigurationTypeFile),
		// 					TapRulesURL: to.Ptr("https://microsoft.com/a"),
		// 					IdentitySelector: &armmanagednetworkfabric.IdentitySelector{
		// 						IdentityType: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentitySelectorTypeUserAssignedIdentity),
		// 						UserAssignedIdentityResourceID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/example-identity"),
		// 					},
		// 					MatchConfigurations: []*armmanagednetworkfabric.NetworkTapRuleMatchConfiguration{
		// 						{
		// 							MatchConfigurationName: to.Ptr("config1"),
		// 							SequenceNumber: to.Ptr[int64](10),
		// 							IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
		// 							MatchConditions: []*armmanagednetworkfabric.NetworkTapRuleMatchCondition{
		// 								{
		// 									ProtocolTypes: []*string{
		// 										to.Ptr("TCP"),
		// 									},
		// 									VlanMatchCondition: &armmanagednetworkfabric.VlanMatchCondition{
		// 										Vlans: []*string{
		// 											to.Ptr("20-30"),
		// 										},
		// 										InnerVlans: []*string{
		// 											to.Ptr("30"),
		// 										},
		// 										VlanGroupNames: []*string{
		// 											to.Ptr("example-vlanGroup"),
		// 										},
		// 									},
		// 									IPCondition: &armmanagednetworkfabric.IPMatchCondition{
		// 										Type: to.Ptr(armmanagednetworkfabric.SourceDestinationTypeSourceIP),
		// 										PrefixType: to.Ptr(armmanagednetworkfabric.PrefixTypePrefix),
		// 										IPPrefixValues: []*string{
		// 											to.Ptr("10.20.20.20/12"),
		// 										},
		// 										IPGroupNames: []*string{
		// 											to.Ptr("example-ipGroup"),
		// 										},
		// 									},
		// 									EncapsulationType: to.Ptr(armmanagednetworkfabric.EncapsulationTypeNone),
		// 									PortCondition: &armmanagednetworkfabric.PortCondition{
		// 										PortType: to.Ptr(armmanagednetworkfabric.PortTypeSourcePort),
		// 										Layer4Protocol: to.Ptr(armmanagednetworkfabric.Layer4ProtocolTCP),
		// 										Ports: []*string{
		// 											to.Ptr("100"),
		// 										},
		// 										PortGroupNames: []*string{
		// 											to.Ptr("example-portGroup1"),
		// 										},
		// 									},
		// 								},
		// 							},
		// 							Actions: []*armmanagednetworkfabric.NetworkTapRuleAction{
		// 								{
		// 									Type: to.Ptr(armmanagednetworkfabric.TapRuleActionTypeDrop),
		// 									Truncate: to.Ptr("100"),
		// 									IsTimestampEnabled: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
		// 									DestinationID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/neighborGroups/example-neighborGroup"),
		// 									MatchConfigurationName: to.Ptr("match1"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					DynamicMatchConfigurations: []*armmanagednetworkfabric.CommonDynamicMatchConfiguration{
		// 						{
		// 							IPGroups: []*armmanagednetworkfabric.IPGroupProperties{
		// 								{
		// 									Name: to.Ptr("example-ipGroup"),
		// 									IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
		// 									IPPrefixes: []*string{
		// 										to.Ptr("10.20.3.1/20"),
		// 									},
		// 								},
		// 							},
		// 							VlanGroups: []*armmanagednetworkfabric.VlanGroupProperties{
		// 								{
		// 									Name: to.Ptr("example-vlanGroup"),
		// 									Vlans: []*string{
		// 										to.Ptr("20-30"),
		// 									},
		// 								},
		// 							},
		// 							PortGroups: []*armmanagednetworkfabric.PortGroupProperties{
		// 								{
		// 									Name: to.Ptr("example-portGroup"),
		// 									Ports: []*string{
		// 										to.Ptr("100-200"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					NetworkTapIDs: []*string{
		// 						to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTaps/example-tap"),
		// 					},
		// 					PollingIntervalInSeconds: to.Ptr[int32](30),
		// 					LastSyncedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-12T07:11:22.485Z"); return t}()),
		// 					NetworkFabricIDs: []*string{
		// 						to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
		// 					},
		// 					GlobalNetworkTapRuleActions: &armmanagednetworkfabric.GlobalNetworkTapRuleActionProperties{
		// 						EnableCount: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
		// 						Truncate: to.Ptr("truncate-name"),
		// 					},
		// 					LastOperation: &armmanagednetworkfabric.LastOperationProperties{
		// 						Details: to.Ptr("Pending administrative update"),
		// 					},
		// 					ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStatePendingAdministrativeUpdate),
		// 					ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
		// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabledDegraded),
		// 				},
		// 				Identity: &armmanagednetworkfabric.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
		// 					TenantID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
		// 					Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
		// 						"key872": &armmanagednetworkfabric.UserAssignedIdentity{
		// 							PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
		// 							ClientID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"keyId": to.Ptr("keyValue"),
		// 				},
		// 				Location: to.Ptr("eastuseuap"),
		// 				ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTapRules/example-tapRule"),
		// 				Name: to.Ptr("example-tapRule"),
		// 				Type: to.Ptr("microsoft.managednetworkfabric/networkTapRules"),
		// 				SystemData: &armmanagednetworkfabric.SystemData{
		// 					CreatedBy: to.Ptr("email@address.com"),
		// 					CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("UserId"),
		// 					LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/aukskqxh"),
		// 	},
		// }
	}
}
