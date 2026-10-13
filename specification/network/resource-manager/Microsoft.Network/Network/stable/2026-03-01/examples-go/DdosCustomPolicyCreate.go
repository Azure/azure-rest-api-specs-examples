package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-03-01/DdosCustomPolicyCreate.json
func ExampleDdosCustomPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDdosCustomPoliciesClient().BeginCreateOrUpdate(ctx, "rg1", "test-ddos-custom-policy", armnetwork.DdosCustomPolicy{
		Location: to.Ptr("centraluseuap"),
		Properties: &armnetwork.DdosCustomPolicyPropertiesFormat{
			DetectionRules: []*armnetwork.DdosDetectionRule{
				{
					Name: to.Ptr("detectionRuleTcp"),
					Properties: &armnetwork.DdosDetectionRulePropertiesFormat{
						DetectionMode: to.Ptr(armnetwork.DdosDetectionModeTrafficThreshold),
						TrafficDetectionRule: &armnetwork.TrafficDetectionRule{
							PacketsPerSecond: to.Ptr[int32](1000000),
							TrafficType:      to.Ptr(armnetwork.DdosTrafficTypeTCP),
						},
					},
				},
			},
			MitigationRules: []*armnetwork.DdosMitigationRule{
				{
					Name: to.Ptr("mitigationRuleTcp"),
					Properties: &armnetwork.DdosMitigationRulePropertiesFormat{
						TrafficScope: to.Ptr(armnetwork.DdosMitigationTrafficScopeTCP),
						TCPDefaultMitigations: &armnetwork.DdosTCPDefaultMitigations{
							PerSourceRateLimiting: &armnetwork.DdosTCPPerSourceRateLimitPolicy{
								PacketsPerSecond: to.Ptr[int32](100000),
							},
							PerSourceConnectionRateLimiting: &armnetwork.DdosTCPPerSourceConnectionRateLimitPolicy{
								ConnectionsPerSecond: to.Ptr[int32](1000),
							},
						},
						SourcePolicyOverrides: []*armnetwork.DdosSourcePolicyOverride{
							{
								PolicyAction: &armnetwork.DdosSourcePolicyAction{
									ActionType: to.Ptr(armnetwork.DdosSourcePolicyActionTypeDeny),
								},
								Conditions: &armnetwork.DdosSourceMatchConditions{
									IPPrefixes: []*string{
										to.Ptr("198.51.100.0/24"),
									},
									GeoMatches: []*armnetwork.DdosGeoMatch{
										{
											CountryCode: to.Ptr("CA"),
										},
									},
								},
							},
							{
								PolicyAction: &armnetwork.DdosSourcePolicyAction{
									ActionType: to.Ptr(armnetwork.DdosSourcePolicyActionTypePermit),
								},
								Conditions: &armnetwork.DdosSourceMatchConditions{
									IPPrefixes: []*string{
										to.Ptr("203.0.113.0/24"),
									},
									GeoMatches: []*armnetwork.DdosGeoMatch{
										{
											Continent: to.Ptr(armnetwork.DdosContinentNorthAmerica),
										},
									},
								},
							},
						},
					},
				},
				{
					Name: to.Ptr("mitigationRuleUdp"),
					Properties: &armnetwork.DdosMitigationRulePropertiesFormat{
						TrafficScope: to.Ptr(armnetwork.DdosMitigationTrafficScopeUDP),
						UDPDefaultMitigations: &armnetwork.DdosUDPDefaultMitigations{
							PerSourceRateLimiting: &armnetwork.DdosUDPPerSourceRateLimitPolicy{
								PacketsPerSecond: to.Ptr[int32](50000),
							},
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.DdosCustomPoliciesClientCreateOrUpdateResponse{
	// 	DdosCustomPolicy: armnetwork.DdosCustomPolicy{
	// 		Name: to.Ptr("test-ddos-custom-policy"),
	// 		Type: to.Ptr("Microsoft.Network/ddosCustomPolicies"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/ddosCustomPolicies/test-ddos-custom-policy"),
	// 		Location: to.Ptr("centraluseuap"),
	// 		Properties: &armnetwork.DdosCustomPolicyPropertiesFormat{
	// 			DetectionRules: []*armnetwork.DdosDetectionRule{
	// 				{
	// 					Name: to.Ptr("detectionRuleTcp"),
	// 					Type: to.Ptr("Microsoft.Network/ddosCustomPolicies/ddosDetectionRules"),
	// 					Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/ddosCustomPolicies/test-ddos-custom-policy/ddosDetectionRules/detectionRuleTcp"),
	// 					Properties: &armnetwork.DdosDetectionRulePropertiesFormat{
	// 						DetectionMode: to.Ptr(armnetwork.DdosDetectionModeTrafficThreshold),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						TrafficDetectionRule: &armnetwork.TrafficDetectionRule{
	// 							PacketsPerSecond: to.Ptr[int32](1000000),
	// 							TrafficType: to.Ptr(armnetwork.DdosTrafficTypeTCP),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			MitigationRules: []*armnetwork.DdosMitigationRule{
	// 				{
	// 					Name: to.Ptr("mitigationRuleTcp"),
	// 					Type: to.Ptr("Microsoft.Network/ddosCustomPolicies/ddosMitigationRules"),
	// 					Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/ddosCustomPolicies/test-ddos-custom-policy/ddosMitigationRules/mitigationRuleTcp"),
	// 					Properties: &armnetwork.DdosMitigationRulePropertiesFormat{
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						TrafficScope: to.Ptr(armnetwork.DdosMitigationTrafficScopeTCP),
	// 						TCPDefaultMitigations: &armnetwork.DdosTCPDefaultMitigations{
	// 							PerSourceRateLimiting: &armnetwork.DdosTCPPerSourceRateLimitPolicy{
	// 								PacketsPerSecond: to.Ptr[int32](100000),
	// 							},
	// 							PerSourceConnectionRateLimiting: &armnetwork.DdosTCPPerSourceConnectionRateLimitPolicy{
	// 								ConnectionsPerSecond: to.Ptr[int32](1000),
	// 							},
	// 						},
	// 						SourcePolicyOverrides: []*armnetwork.DdosSourcePolicyOverride{
	// 							{
	// 								PolicyAction: &armnetwork.DdosSourcePolicyAction{
	// 									ActionType: to.Ptr(armnetwork.DdosSourcePolicyActionTypeDeny),
	// 								},
	// 								Conditions: &armnetwork.DdosSourceMatchConditions{
	// 									IPPrefixes: []*string{
	// 										to.Ptr("198.51.100.0/24"),
	// 									},
	// 									GeoMatches: []*armnetwork.DdosGeoMatch{
	// 										{
	// 											CountryCode: to.Ptr("CA"),
	// 										},
	// 									},
	// 								},
	// 							},
	// 							{
	// 								PolicyAction: &armnetwork.DdosSourcePolicyAction{
	// 									ActionType: to.Ptr(armnetwork.DdosSourcePolicyActionTypePermit),
	// 								},
	// 								Conditions: &armnetwork.DdosSourceMatchConditions{
	// 									IPPrefixes: []*string{
	// 										to.Ptr("203.0.113.0/24"),
	// 									},
	// 									GeoMatches: []*armnetwork.DdosGeoMatch{
	// 										{
	// 											Continent: to.Ptr(armnetwork.DdosContinentNorthAmerica),
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("mitigationRuleUdp"),
	// 					Type: to.Ptr("Microsoft.Network/ddosCustomPolicies/ddosMitigationRules"),
	// 					Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/ddosCustomPolicies/test-ddos-custom-policy/ddosMitigationRules/mitigationRuleUdp"),
	// 					Properties: &armnetwork.DdosMitigationRulePropertiesFormat{
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						TrafficScope: to.Ptr(armnetwork.DdosMitigationTrafficScopeUDP),
	// 						UDPDefaultMitigations: &armnetwork.DdosUDPDefaultMitigations{
	// 							PerSourceRateLimiting: &armnetwork.DdosUDPPerSourceRateLimitPolicy{
	// 								PacketsPerSecond: to.Ptr[int32](50000),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 	},
	// }
}
