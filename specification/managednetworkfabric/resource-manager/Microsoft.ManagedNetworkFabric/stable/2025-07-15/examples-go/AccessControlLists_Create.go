package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/AccessControlLists_Create.json
func ExampleAccessControlListsClient_BeginCreate_accessControlListsCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccessControlListsClient().BeginCreate(ctx, "example-rg", "example-acl", armmanagednetworkfabric.AccessControlList{
		Properties: &armmanagednetworkfabric.AccessControlListProperties{
			Annotation:        to.Ptr("annotation"),
			ConfigurationType: to.Ptr(armmanagednetworkfabric.ConfigurationTypeFile),
			ACLsURL:           to.Ptr("https://ACL-Storage-URL"),
			DefaultAction:     to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
			MatchConfigurations: []*armmanagednetworkfabric.AccessControlListMatchConfiguration{
				{
					MatchConfigurationName: to.Ptr("example-match"),
					SequenceNumber:         to.Ptr[int64](123),
					IPAddressType:          to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
					MatchConditions: []*armmanagednetworkfabric.AccessControlListMatchCondition{
						{
							ProtocolTypes: []*string{
								to.Ptr("TCP"),
							},
							VlanMatchCondition: &armmanagednetworkfabric.VlanMatchCondition{
								Vlans: []*string{
									to.Ptr("20-30"),
								},
								InnerVlans: []*string{
									to.Ptr("30"),
								},
								VlanGroupNames: []*string{
									to.Ptr("example-vlanGroup"),
								},
							},
							IPCondition: &armmanagednetworkfabric.IPMatchCondition{
								Type:       to.Ptr(armmanagednetworkfabric.SourceDestinationTypeSourceIP),
								PrefixType: to.Ptr(armmanagednetworkfabric.PrefixTypePrefix),
								IPPrefixValues: []*string{
									to.Ptr("10.20.20.20/12"),
								},
								IPGroupNames: []*string{
									to.Ptr("example-ipGroup"),
								},
							},
							EtherTypes: []*string{
								to.Ptr("0x1"),
							},
							Fragments: []*string{
								to.Ptr("0xff00-0xffff"),
							},
							IPLengths: []*string{
								to.Ptr("4094-9214"),
							},
							TTLValues: []*string{
								to.Ptr("23"),
							},
							DscpMarkings: []*string{
								to.Ptr("32"),
							},
							ProtocolNeighbors: []*string{
								to.Ptr("example-neighbor"),
							},
							PortCondition: &armmanagednetworkfabric.AccessControlListPortCondition{
								PortType:       to.Ptr(armmanagednetworkfabric.PortTypeSourcePort),
								Layer4Protocol: to.Ptr(armmanagednetworkfabric.Layer4ProtocolTCP),
								Ports: []*string{
									to.Ptr("1-20"),
								},
								PortGroupNames: []*string{
									to.Ptr("example-portGroup"),
								},
								Flags: []*string{
									to.Ptr("established"),
								},
							},
							IcmpConfiguration: &armmanagednetworkfabric.IcmpConfigurationProperties{
								IcmpTypes: []*string{
									to.Ptr("echo"),
								},
							},
						},
					},
					Actions: []*armmanagednetworkfabric.AccessControlListAction{
						{
							Type:          to.Ptr(armmanagednetworkfabric.ACLActionTypeCount),
							CounterName:   to.Ptr("example-counter"),
							RemarkComment: to.Ptr("example-remark"),
							PoliceRateConfiguration: &armmanagednetworkfabric.PoliceRateConfigurationProperties{
								BitRate: &armmanagednetworkfabric.BitRate{
									Rate: to.Ptr[int64](15),
									Unit: to.Ptr(armmanagednetworkfabric.BitRateUnitBps),
								},
								BurstSize: &armmanagednetworkfabric.BurstSize{
									Size: to.Ptr[int64](2),
									Unit: to.Ptr(armmanagednetworkfabric.BurstSizeUnitBytes),
								},
							},
						},
					},
				},
			},
			DynamicMatchConfigurations: []*armmanagednetworkfabric.CommonDynamicMatchConfiguration{
				{
					IPGroups: []*armmanagednetworkfabric.IPGroupProperties{
						{
							Name:          to.Ptr("example-ipGroup"),
							IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
							IPPrefixes: []*string{
								to.Ptr("10.20.3.1/20"),
							},
						},
					},
					VlanGroups: []*armmanagednetworkfabric.VlanGroupProperties{
						{
							Name: to.Ptr("example-vlanGroup"),
							Vlans: []*string{
								to.Ptr("20-30"),
							},
						},
					},
					PortGroups: []*armmanagednetworkfabric.PortGroupProperties{
						{
							Name: to.Ptr("example-portGroup"),
							Ports: []*string{
								to.Ptr("100-200"),
							},
						},
					},
				},
			},
			ACLType:    to.Ptr(armmanagednetworkfabric.ACLTypeCp),
			DeviceRole: to.Ptr(armmanagednetworkfabric.DeviceRoleCE),
			GlobalAccessControlListActions: &armmanagednetworkfabric.GlobalAccessControlListActionProperties{
				EnableCount: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
			},
		},
		Tags: map[string]*string{
			"keyID": to.Ptr("keyValue"),
		},
		Location: to.Ptr("eastUs"),
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
	// res = armmanagednetworkfabric.AccessControlListsClientCreateResponse{
	// 	AccessControlList: &armmanagednetworkfabric.AccessControlList{
	// 		Properties: &armmanagednetworkfabric.AccessControlListProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			ConfigurationType: to.Ptr(armmanagednetworkfabric.ConfigurationTypeFile),
	// 			ACLsURL: to.Ptr("https://ACL-Storage-URL"),
	// 			DefaultAction: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
	// 			MatchConfigurations: []*armmanagednetworkfabric.AccessControlListMatchConfiguration{
	// 				{
	// 					MatchConfigurationName: to.Ptr("example-match"),
	// 					SequenceNumber: to.Ptr[int64](123),
	// 					IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
	// 					MatchConditions: []*armmanagednetworkfabric.AccessControlListMatchCondition{
	// 						{
	// 							ProtocolTypes: []*string{
	// 								to.Ptr("TCP"),
	// 							},
	// 							VlanMatchCondition: &armmanagednetworkfabric.VlanMatchCondition{
	// 								Vlans: []*string{
	// 									to.Ptr("20-30"),
	// 								},
	// 								InnerVlans: []*string{
	// 									to.Ptr("30"),
	// 								},
	// 								VlanGroupNames: []*string{
	// 									to.Ptr("example-vlanGroup"),
	// 								},
	// 							},
	// 							IPCondition: &armmanagednetworkfabric.IPMatchCondition{
	// 								Type: to.Ptr(armmanagednetworkfabric.SourceDestinationTypeSourceIP),
	// 								PrefixType: to.Ptr(armmanagednetworkfabric.PrefixTypePrefix),
	// 								IPPrefixValues: []*string{
	// 									to.Ptr("10.20.20.20/12"),
	// 								},
	// 								IPGroupNames: []*string{
	// 									to.Ptr("example-ipGroup"),
	// 								},
	// 							},
	// 							EtherTypes: []*string{
	// 								to.Ptr("0x1"),
	// 							},
	// 							Fragments: []*string{
	// 								to.Ptr("0xff00-0xffff"),
	// 							},
	// 							IPLengths: []*string{
	// 								to.Ptr("4094-9214"),
	// 							},
	// 							TTLValues: []*string{
	// 								to.Ptr("23"),
	// 							},
	// 							DscpMarkings: []*string{
	// 								to.Ptr("32"),
	// 							},
	// 							ProtocolNeighbors: []*string{
	// 								to.Ptr("example-neighbor"),
	// 							},
	// 							PortCondition: &armmanagednetworkfabric.AccessControlListPortCondition{
	// 								PortType: to.Ptr(armmanagednetworkfabric.PortTypeSourcePort),
	// 								Layer4Protocol: to.Ptr(armmanagednetworkfabric.Layer4ProtocolTCP),
	// 								Ports: []*string{
	// 									to.Ptr("1-20"),
	// 								},
	// 								PortGroupNames: []*string{
	// 									to.Ptr("example-portGroup"),
	// 								},
	// 								Flags: []*string{
	// 									to.Ptr("established"),
	// 								},
	// 							},
	// 							IcmpConfiguration: &armmanagednetworkfabric.IcmpConfigurationProperties{
	// 								IcmpTypes: []*string{
	// 									to.Ptr("echo"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 					Actions: []*armmanagednetworkfabric.AccessControlListAction{
	// 						{
	// 							Type: to.Ptr(armmanagednetworkfabric.ACLActionTypeCount),
	// 							CounterName: to.Ptr("example-counter"),
	// 							RemarkComment: to.Ptr("example-remark"),
	// 							PoliceRateConfiguration: &armmanagednetworkfabric.PoliceRateConfigurationProperties{
	// 								BitRate: &armmanagednetworkfabric.BitRate{
	// 									Rate: to.Ptr[int64](15),
	// 									Unit: to.Ptr(armmanagednetworkfabric.BitRateUnitBps),
	// 								},
	// 								BurstSize: &armmanagednetworkfabric.BurstSize{
	// 									Size: to.Ptr[int64](2),
	// 									Unit: to.Ptr(armmanagednetworkfabric.BurstSizeUnitBytes),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			DynamicMatchConfigurations: []*armmanagednetworkfabric.CommonDynamicMatchConfiguration{
	// 				{
	// 					IPGroups: []*armmanagednetworkfabric.IPGroupProperties{
	// 						{
	// 							Name: to.Ptr("example-ipGroup"),
	// 							IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
	// 							IPPrefixes: []*string{
	// 								to.Ptr("10.20.3.1/20"),
	// 							},
	// 						},
	// 					},
	// 					VlanGroups: []*armmanagednetworkfabric.VlanGroupProperties{
	// 						{
	// 							Name: to.Ptr("example-vlanGroup"),
	// 							Vlans: []*string{
	// 								to.Ptr("20-30"),
	// 							},
	// 						},
	// 					},
	// 					PortGroups: []*armmanagednetworkfabric.PortGroupProperties{
	// 						{
	// 							Name: to.Ptr("example-portGroup"),
	// 							Ports: []*string{
	// 								to.Ptr("100-200"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			LastSyncedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-17T08:56:23.203Z"); return t}()),
	// 			ACLType: to.Ptr(armmanagednetworkfabric.ACLTypeCp),
	// 			DeviceRole: to.Ptr(armmanagednetworkfabric.DeviceRoleCE),
	// 			GlobalAccessControlListActions: &armmanagednetworkfabric.GlobalAccessControlListActionProperties{
	// 				EnableCount: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
	// 			},
	// 			LastOperation: &armmanagednetworkfabric.LastOperationProperties{
	// 				Details: to.Ptr("Succeeded"),
	// 			},
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		},
	// 		Tags: map[string]*string{
	// 			"KeyId": to.Ptr("keyValue"),
	// 		},
	// 		Location: to.Ptr("eastUs"),
	// 		ID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
	// 		Name: to.Ptr("example-acl"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/accessControlLists"),
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
