package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/AccessControlLists_Create_ControlPlaneAcl.json
func ExampleAccessControlListsClient_BeginCreate_accessControlListsCreateMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("1234ABCD-0A1B-1234-5678-123456ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccessControlListsClient().BeginCreate(ctx, "example-resource-group", "example-acl", armmanagednetworkfabric.AccessControlList{
		Properties: &armmanagednetworkfabric.AccessControlListProperties{
			Annotation:        to.Ptr("annotation"),
			ConfigurationType: to.Ptr(armmanagednetworkfabric.ConfigurationTypeInline),
			ACLsURL:           to.Ptr("https://microsoft.com/a"),
			DefaultAction:     to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
			ACLType:           to.Ptr(armmanagednetworkfabric.ACLTypeCp),
			DeviceRole:        to.Ptr(armmanagednetworkfabric.DeviceRoleCE),
			ControlPlaneACLConfiguration: []*armmanagednetworkfabric.ControlPlaneACLProperties{
				{
					IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
					MatchConfigurations: []*armmanagednetworkfabric.ControlPlaneACLMatchConfigurationProperties{
						{
							MatchConfigurationName: to.Ptr("example-match-config"),
							SequenceNumber:         to.Ptr[int64](3779271459),
							MatchCondition: &armmanagednetworkfabric.ControlPlaneACLMatchCondition{
								ProtocolTypes: to.Ptr("tcp"),
								IPCondition: &armmanagednetworkfabric.ControlPlanACLIPMatchCondition{
									SourceIPPrefix:      to.Ptr("10.0.0.0/24"),
									DestinationIPPrefix: to.Ptr("10.0.0.0/24"),
								},
								TTLMatchCondition: &armmanagednetworkfabric.ControlPlaneACLTTLMatchCondition{
									TTLValue:     to.Ptr("1"),
									TTLMatchType: to.Ptr(armmanagednetworkfabric.ControlPlaneACLTTLMatchTypeEquals),
								},
								PortCondition: &armmanagednetworkfabric.ControlPlaneACLPortMatchCondition{
									SourcePorts: &armmanagednetworkfabric.ControlPlaneACLPortCondition{
										Ports: []*string{
											to.Ptr("100"),
										},
										PortMatchType: to.Ptr(armmanagednetworkfabric.ControlPlaneACLPortMatchTypeEquals),
									},
									DestinationPorts: &armmanagednetworkfabric.ControlPlaneACLPortCondition{
										Ports: []*string{
											to.Ptr("200"),
										},
										PortMatchType: to.Ptr(armmanagednetworkfabric.ControlPlaneACLPortMatchTypeEquals),
									},
								},
								Flags: []*string{
									to.Ptr("established"),
								},
								IcmpConfiguration: &armmanagednetworkfabric.IcmpConfigurationProperties{
									IcmpTypes: []*string{
										to.Ptr("icmp"),
									},
								},
							},
							Action: &armmanagednetworkfabric.ControlPlaneACLAction{
								Type:          to.Ptr(armmanagednetworkfabric.ControlPlaneACLActionType("Drop")),
								RemarkComment: to.Ptr("remark"),
							},
						},
					},
				},
			},
			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		},
		Tags: map[string]*string{
			"key5032": to.Ptr("example-tag-value"),
		},
		Location: to.Ptr("eastus"),
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
	// 			Annotation: to.Ptr("ttbh"),
	// 			ConfigurationType: to.Ptr(armmanagednetworkfabric.ConfigurationTypeInline),
	// 			ACLsURL: to.Ptr("https://microsoft.com/a"),
	// 			DefaultAction: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
	// 			LastSyncedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-27T08:49:55.426Z"); return t}()),
	// 			ACLType: to.Ptr(armmanagednetworkfabric.ACLTypeCp),
	// 			DeviceRole: to.Ptr(armmanagednetworkfabric.DeviceRoleCE),
	// 			ControlPlaneACLConfiguration: []*armmanagednetworkfabric.ControlPlaneACLProperties{
	// 				{
	// 					IPAddressType: to.Ptr(armmanagednetworkfabric.IPAddressTypeIPv4),
	// 					MatchConfigurations: []*armmanagednetworkfabric.ControlPlaneACLMatchConfigurationProperties{
	// 						{
	// 							MatchConfigurationName: to.Ptr("example-match-config"),
	// 							SequenceNumber: to.Ptr[int64](3779271459),
	// 							MatchCondition: &armmanagednetworkfabric.ControlPlaneACLMatchCondition{
	// 								ProtocolTypes: to.Ptr("tcp"),
	// 								IPCondition: &armmanagednetworkfabric.ControlPlanACLIPMatchCondition{
	// 									SourceIPPrefix: to.Ptr("10.0.0.0/24"),
	// 									DestinationIPPrefix: to.Ptr("10.0.0.0/24"),
	// 								},
	// 								TTLMatchCondition: &armmanagednetworkfabric.ControlPlaneACLTTLMatchCondition{
	// 									TTLValue: to.Ptr("1"),
	// 									TTLMatchType: to.Ptr(armmanagednetworkfabric.ControlPlaneACLTTLMatchTypeEquals),
	// 								},
	// 								PortCondition: &armmanagednetworkfabric.ControlPlaneACLPortMatchCondition{
	// 									SourcePorts: &armmanagednetworkfabric.ControlPlaneACLPortCondition{
	// 										Ports: []*string{
	// 											to.Ptr("100"),
	// 										},
	// 										PortMatchType: to.Ptr(armmanagednetworkfabric.ControlPlaneACLPortMatchTypeEquals),
	// 									},
	// 									DestinationPorts: &armmanagednetworkfabric.ControlPlaneACLPortCondition{
	// 										Ports: []*string{
	// 											to.Ptr("200"),
	// 										},
	// 										PortMatchType: to.Ptr(armmanagednetworkfabric.ControlPlaneACLPortMatchTypeEquals),
	// 									},
	// 								},
	// 								Flags: []*string{
	// 									to.Ptr("established"),
	// 								},
	// 								IcmpConfiguration: &armmanagednetworkfabric.IcmpConfigurationProperties{
	// 									IcmpTypes: []*string{
	// 										to.Ptr("icmp"),
	// 									},
	// 								},
	// 							},
	// 							Action: &armmanagednetworkfabric.ControlPlaneACLAction{
	// 								Type: to.Ptr(armmanagednetworkfabric.ControlPlaneACLActionType("Drop")),
	// 								RemarkComment: to.Ptr("remark"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key5032": to.Ptr("tags"),
	// 		},
	// 		Location: to.Ptr("eastus"),
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
