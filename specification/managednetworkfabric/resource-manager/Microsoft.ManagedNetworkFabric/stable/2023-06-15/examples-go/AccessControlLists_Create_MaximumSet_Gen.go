package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/AccessControlLists_Create_MaximumSet_Gen.json
func ExampleAccessControlListsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccessControlListsClient().BeginCreate(ctx, "example-rg", "example-acl", armmanagednetworkfabric.AccessControlList{
		Location: to.Ptr("eastUs"),
		Tags: map[string]*string{
			"keyID": to.Ptr("KeyValue"),
		},
		Properties: &armmanagednetworkfabric.AccessControlListProperties{
			ACLsURL:           to.Ptr("https://ACL-Storage-URL"),
			ConfigurationType: to.Ptr(armmanagednetworkfabric.ConfigurationTypeFile),
			DynamicMatchConfigurations: []*armmanagednetworkfabric.CommonDynamicMatchConfiguration{
				{
					IPGroups: []*armmanagednetworkfabric.IPGroupProperties{
						{
							Name:          to.Ptr("example-ipGroup"),
							IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
							IPPrefixes: []*string{
								to.Ptr("10.20.3.1/20")},
						}},
					PortGroups: []*armmanagednetworkfabric.PortGroupProperties{
						{
							Name: to.Ptr("example-portGroup"),
							Ports: []*string{
								to.Ptr("100-200")},
						}},
					VlanGroups: []*armmanagednetworkfabric.VlanGroupProperties{
						{
							Name: to.Ptr("example-vlanGroup"),
							Vlans: []*string{
								to.Ptr("20-30")},
						}},
				}},
			MatchConfigurations: []*armmanagednetworkfabric.AccessControlListMatchConfiguration{
				{
					Actions: []*armmanagednetworkfabric.AccessControlListAction{
						{
							Type:        to.Ptr(armmanagednetworkfabric.ACLActionTypeCount),
							CounterName: to.Ptr("example-counter"),
						}},
					IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
					MatchConditions: []*armmanagednetworkfabric.AccessControlListMatchCondition{
						{
							IPCondition: &armmanagednetworkfabric.IPMatchCondition{
								Type: to.Ptr(armmanagednetworkfabric.SourceDestinationTypeSourceIP),
								IPGroupNames: []*string{
									to.Ptr("example-ipGroup")},
								IPPrefixValues: []*string{
									to.Ptr("10.20.20.20/12")},
								PrefixType: to.Ptr(armmanagednetworkfabric.PrefixTypePrefix),
							},
							ProtocolTypes: []*string{
								to.Ptr("TCP")},
							VlanMatchCondition: &armmanagednetworkfabric.VlanMatchCondition{
								InnerVlans: []*string{
									to.Ptr("30")},
								VlanGroupNames: []*string{
									to.Ptr("example-vlanGroup")},
								Vlans: []*string{
									to.Ptr("20-30")},
							},
							DscpMarkings: []*string{
								to.Ptr("32")},
							EtherTypes: []*string{
								to.Ptr("0x1")},
							Fragments: []*string{
								to.Ptr("0xff00-0xffff")},
							IPLengths: []*string{
								to.Ptr("4094-9214")},
							PortCondition: &armmanagednetworkfabric.AccessControlListPortCondition{
								Layer4Protocol: to.Ptr(armmanagednetworkfabric.Layer4ProtocolTCP),
								PortGroupNames: []*string{
									to.Ptr("example-portGroup")},
								PortType: to.Ptr(armmanagednetworkfabric.PortTypeSourcePort),
								Ports: []*string{
									to.Ptr("1-20")},
								Flags: []*string{
									to.Ptr("established")},
							},
							TTLValues: []*string{
								to.Ptr("23")},
						}},
					MatchConfigurationName: to.Ptr("example-match"),
					SequenceNumber:         to.Ptr[int64](123),
				}},
			Annotation: to.Ptr("annotation"),
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
	// res.AccessControlList = armmanagednetworkfabric.AccessControlList{
	// 	Name: to.Ptr("example-acl"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/accessControlLists"),
	// 	ID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 		CreatedBy: to.Ptr("email@address.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("UserId"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastUs"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("KeyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.AccessControlListProperties{
	// 		ACLsURL: to.Ptr("https://ACL-Storage-URL"),
	// 		ConfigurationType: to.Ptr(armmanagednetworkfabric.ConfigurationTypeFile),
	// 		DynamicMatchConfigurations: []*armmanagednetworkfabric.CommonDynamicMatchConfiguration{
	// 			{
	// 				IPGroups: []*armmanagednetworkfabric.IPGroupProperties{
	// 					{
	// 						Name: to.Ptr("example-ipGroup"),
	// 						IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
	// 						IPPrefixes: []*string{
	// 							to.Ptr("10.20.3.1/20")},
	// 					}},
	// 					PortGroups: []*armmanagednetworkfabric.PortGroupProperties{
	// 						{
	// 							Name: to.Ptr("example-portGroup"),
	// 							Ports: []*string{
	// 								to.Ptr("100-200")},
	// 						}},
	// 						VlanGroups: []*armmanagednetworkfabric.VlanGroupProperties{
	// 							{
	// 								Name: to.Ptr("example-vlanGroup"),
	// 								Vlans: []*string{
	// 									to.Ptr("20-30")},
	// 							}},
	// 					}},
	// 					MatchConfigurations: []*armmanagednetworkfabric.AccessControlListMatchConfiguration{
	// 						{
	// 							Actions: []*armmanagednetworkfabric.AccessControlListAction{
	// 								{
	// 									Type: to.Ptr(armmanagednetworkfabric.ACLActionTypeCount),
	// 									CounterName: to.Ptr("example-counter"),
	// 							}},
	// 							IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
	// 							MatchConditions: []*armmanagednetworkfabric.AccessControlListMatchCondition{
	// 								{
	// 									IPCondition: &armmanagednetworkfabric.IPMatchCondition{
	// 										Type: to.Ptr(armmanagednetworkfabric.SourceDestinationTypeSourceIP),
	// 										IPGroupNames: []*string{
	// 											to.Ptr("example-ipGroup")},
	// 											IPPrefixValues: []*string{
	// 												to.Ptr("10.20.20.20/12")},
	// 												PrefixType: to.Ptr(armmanagednetworkfabric.PrefixTypePrefix),
	// 											},
	// 											ProtocolTypes: []*string{
	// 												to.Ptr("TCP")},
	// 												VlanMatchCondition: &armmanagednetworkfabric.VlanMatchCondition{
	// 													InnerVlans: []*string{
	// 														to.Ptr("30")},
	// 														VlanGroupNames: []*string{
	// 															to.Ptr("example-vlanGroup")},
	// 															Vlans: []*string{
	// 																to.Ptr("20-30")},
	// 															},
	// 															DscpMarkings: []*string{
	// 																to.Ptr("32")},
	// 																EtherTypes: []*string{
	// 																	to.Ptr("0x1")},
	// 																	Fragments: []*string{
	// 																		to.Ptr("0xff00-0xffff")},
	// 																		IPLengths: []*string{
	// 																			to.Ptr("4094-9214")},
	// 																			PortCondition: &armmanagednetworkfabric.AccessControlListPortCondition{
	// 																				Layer4Protocol: to.Ptr(armmanagednetworkfabric.Layer4ProtocolTCP),
	// 																				PortGroupNames: []*string{
	// 																					to.Ptr("example-portGroup")},
	// 																					PortType: to.Ptr(armmanagednetworkfabric.PortTypeSourcePort),
	// 																					Ports: []*string{
	// 																						to.Ptr("1-20")},
	// 																						Flags: []*string{
	// 																							to.Ptr("established")},
	// 																						},
	// 																						TTLValues: []*string{
	// 																							to.Ptr("23")},
	// 																					}},
	// 																					MatchConfigurationName: to.Ptr("example-match"),
	// 																					SequenceNumber: to.Ptr[int64](123),
	// 																			}},
	// 																			Annotation: to.Ptr("annotation"),
	// 																			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 																			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 																			LastSyncedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-17T08:56:23.203Z"); return t}()),
	// 																			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 																		},
	// 																	}
}
