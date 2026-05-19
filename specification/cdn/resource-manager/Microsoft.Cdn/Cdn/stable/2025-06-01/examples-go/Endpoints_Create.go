package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/Endpoints_Create.json
func ExampleEndpointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEndpointsClient().BeginCreate(ctx, "RG", "profile1", "endpoint1", armcdn.Endpoint{
		Location: to.Ptr("WestUs"),
		Properties: &armcdn.EndpointProperties{
			ContentTypesToCompress: []*string{
				to.Ptr("text/html"),
				to.Ptr("application/octet-stream"),
			},
			DefaultOriginGroup: &armcdn.ResourceReference{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/originGroups/originGroup1"),
			},
			DeliveryPolicy: &armcdn.EndpointPropertiesUpdateParametersDeliveryPolicy{
				Description: to.Ptr("Test description for a policy."),
				Rules: []*armcdn.DeliveryRule{
					{
						Name: to.Ptr("rule1"),
						Actions: []armcdn.DeliveryRuleActionClassification{
							&armcdn.DeliveryRuleCacheExpirationAction{
								Name: to.Ptr(armcdn.DeliveryRuleActionNameCacheExpiration),
								Parameters: &armcdn.CacheExpirationActionParameters{
									CacheBehavior: to.Ptr(armcdn.CacheBehaviorOverride),
									CacheDuration: to.Ptr("10:10:09"),
									CacheType:     to.Ptr(armcdn.CacheTypeAll),
									TypeName:      to.Ptr(armcdn.DeliveryRuleActionParametersTypeDeliveryRuleCacheExpirationActionParameters),
								},
							},
							&armcdn.DeliveryRuleResponseHeaderAction{
								Name: to.Ptr(armcdn.DeliveryRuleActionNameModifyResponseHeader),
								Parameters: &armcdn.HeaderActionParameters{
									HeaderAction: to.Ptr(armcdn.HeaderActionOverwrite),
									HeaderName:   to.Ptr("Access-Control-Allow-Origin"),
									TypeName:     to.Ptr(armcdn.DeliveryRuleActionParametersTypeDeliveryRuleHeaderActionParameters),
									Value:        to.Ptr("*"),
								},
							},
							&armcdn.DeliveryRuleRequestHeaderAction{
								Name: to.Ptr(armcdn.DeliveryRuleActionNameModifyRequestHeader),
								Parameters: &armcdn.HeaderActionParameters{
									HeaderAction: to.Ptr(armcdn.HeaderActionOverwrite),
									HeaderName:   to.Ptr("Accept-Encoding"),
									TypeName:     to.Ptr(armcdn.DeliveryRuleActionParametersTypeDeliveryRuleHeaderActionParameters),
									Value:        to.Ptr("gzip"),
								},
							},
						},
						Conditions: []armcdn.DeliveryRuleConditionClassification{
							&armcdn.DeliveryRuleRemoteAddressCondition{
								Name: to.Ptr(armcdn.MatchVariableRemoteAddress),
								Parameters: &armcdn.RemoteAddressMatchConditionParameters{
									MatchValues: []*string{
										to.Ptr("192.168.1.0/24"),
										to.Ptr("10.0.0.0/24"),
									},
									NegateCondition: to.Ptr(true),
									Operator:        to.Ptr(armcdn.RemoteAddressOperatorIPMatch),
									TypeName:        to.Ptr(armcdn.DeliveryRuleConditionParametersTypeDeliveryRuleRemoteAddressConditionParameters),
								},
							},
						},
						Order: to.Ptr[int32](1),
					},
				},
			},
			IsCompressionEnabled: to.Ptr(true),
			IsHTTPAllowed:        to.Ptr(true),
			IsHTTPSAllowed:       to.Ptr(true),
			OriginGroups: []*armcdn.DeepCreatedOriginGroup{
				{
					Name: to.Ptr("originGroup1"),
					Properties: &armcdn.DeepCreatedOriginGroupProperties{
						HealthProbeSettings: &armcdn.HealthProbeParameters{
							ProbeIntervalInSeconds: to.Ptr[int32](120),
							ProbePath:              to.Ptr("/health.aspx"),
							ProbeProtocol:          to.Ptr(armcdn.ProbeProtocolHTTP),
							ProbeRequestType:       to.Ptr(armcdn.HealthProbeRequestTypeGET),
						},
						Origins: []*armcdn.ResourceReference{
							{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin1"),
							},
							{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin2"),
							},
						},
						ResponseBasedOriginErrorDetectionSettings: &armcdn.ResponseBasedOriginErrorDetectionParameters{
							ResponseBasedDetectedErrorTypes:          to.Ptr(armcdn.ResponseBasedDetectedErrorTypesTCPErrorsOnly),
							ResponseBasedFailoverThresholdPercentage: to.Ptr[int32](10),
						},
					},
				},
			},
			OriginHostHeader: to.Ptr("www.bing.com"),
			OriginPath:       to.Ptr("/photos"),
			Origins: []*armcdn.DeepCreatedOrigin{
				{
					Name: to.Ptr("origin1"),
					Properties: &armcdn.DeepCreatedOriginProperties{
						Enabled:          to.Ptr(true),
						HostName:         to.Ptr("www.someDomain1.net"),
						HTTPPort:         to.Ptr[int32](80),
						HTTPSPort:        to.Ptr[int32](443),
						OriginHostHeader: to.Ptr("www.someDomain1.net"),
						Priority:         to.Ptr[int32](1),
						Weight:           to.Ptr[int32](50),
					},
				},
				{
					Name: to.Ptr("origin2"),
					Properties: &armcdn.DeepCreatedOriginProperties{
						Enabled:          to.Ptr(true),
						HostName:         to.Ptr("www.someDomain2.net"),
						HTTPPort:         to.Ptr[int32](80),
						HTTPSPort:        to.Ptr[int32](443),
						OriginHostHeader: to.Ptr("www.someDomain2.net"),
						Priority:         to.Ptr[int32](2),
						Weight:           to.Ptr[int32](50),
					},
				},
			},
			QueryStringCachingBehavior: to.Ptr(armcdn.QueryStringCachingBehaviorBypassCaching),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
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
	// res = armcdn.EndpointsClientCreateResponse{
	// 	Endpoint: armcdn.Endpoint{
	// 		Name: to.Ptr("endpoint4899"),
	// 		Type: to.Ptr("Microsoft.Cdn/profiles/endpoints"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1"),
	// 		Location: to.Ptr("WestUs"),
	// 		Properties: &armcdn.EndpointProperties{
	// 			ContentTypesToCompress: []*string{
	// 				to.Ptr("text/html"),
	// 				to.Ptr("application/octet-stream"),
	// 			},
	// 			DefaultOriginGroup: &armcdn.ResourceReference{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/originGroups/originGroup1"),
	// 			},
	// 			DeliveryPolicy: &armcdn.EndpointPropertiesUpdateParametersDeliveryPolicy{
	// 				Description: to.Ptr("Test description for a policy."),
	// 				Rules: []*armcdn.DeliveryRule{
	// 					{
	// 						Name: to.Ptr("rule1"),
	// 						Actions: []armcdn.DeliveryRuleActionClassification{
	// 							&armcdn.DeliveryRuleCacheExpirationAction{
	// 								Name: to.Ptr(armcdn.DeliveryRuleActionNameCacheExpiration),
	// 								Parameters: &armcdn.CacheExpirationActionParameters{
	// 									CacheBehavior: to.Ptr(armcdn.CacheBehaviorOverride),
	// 									CacheDuration: to.Ptr("10:10:09"),
	// 									CacheType: to.Ptr(armcdn.CacheTypeAll),
	// 									TypeName: to.Ptr(armcdn.DeliveryRuleActionParametersTypeDeliveryRuleCacheExpirationActionParameters),
	// 								},
	// 							},
	// 							&armcdn.DeliveryRuleResponseHeaderAction{
	// 								Name: to.Ptr(armcdn.DeliveryRuleActionNameModifyResponseHeader),
	// 								Parameters: &armcdn.HeaderActionParameters{
	// 									HeaderAction: to.Ptr(armcdn.HeaderActionOverwrite),
	// 									HeaderName: to.Ptr("Access-Control-Allow-Origin"),
	// 									TypeName: to.Ptr(armcdn.DeliveryRuleActionParametersTypeDeliveryRuleHeaderActionParameters),
	// 									Value: to.Ptr("*"),
	// 								},
	// 							},
	// 							&armcdn.DeliveryRuleRequestHeaderAction{
	// 								Name: to.Ptr(armcdn.DeliveryRuleActionNameModifyRequestHeader),
	// 								Parameters: &armcdn.HeaderActionParameters{
	// 									HeaderAction: to.Ptr(armcdn.HeaderActionOverwrite),
	// 									HeaderName: to.Ptr("Accept-Encoding"),
	// 									TypeName: to.Ptr(armcdn.DeliveryRuleActionParametersTypeDeliveryRuleHeaderActionParameters),
	// 									Value: to.Ptr("gzip"),
	// 								},
	// 							},
	// 						},
	// 						Conditions: []armcdn.DeliveryRuleConditionClassification{
	// 							&armcdn.DeliveryRuleRemoteAddressCondition{
	// 								Name: to.Ptr(armcdn.MatchVariableRemoteAddress),
	// 								Parameters: &armcdn.RemoteAddressMatchConditionParameters{
	// 									MatchValues: []*string{
	// 										to.Ptr("192.168.1.0/24"),
	// 										to.Ptr("10.0.0.0/24"),
	// 									},
	// 									NegateCondition: to.Ptr(true),
	// 									Operator: to.Ptr(armcdn.RemoteAddressOperatorIPMatch),
	// 									Transforms: []*armcdn.Transform{
	// 									},
	// 									TypeName: to.Ptr(armcdn.DeliveryRuleConditionParametersTypeDeliveryRuleRemoteAddressConditionParameters),
	// 								},
	// 							},
	// 						},
	// 						Order: to.Ptr[int32](1),
	// 					},
	// 				},
	// 			},
	// 			GeoFilters: []*armcdn.GeoFilter{
	// 			},
	// 			HostName: to.Ptr("endpoint4899.azureedge-test.net"),
	// 			IsCompressionEnabled: to.Ptr(true),
	// 			IsHTTPAllowed: to.Ptr(true),
	// 			IsHTTPSAllowed: to.Ptr(true),
	// 			OriginGroups: []*armcdn.DeepCreatedOriginGroup{
	// 				{
	// 					Name: to.Ptr("originGroup1"),
	// 					Properties: &armcdn.DeepCreatedOriginGroupProperties{
	// 						HealthProbeSettings: &armcdn.HealthProbeParameters{
	// 							ProbeIntervalInSeconds: to.Ptr[int32](120),
	// 							ProbePath: to.Ptr("/health.aspx"),
	// 							ProbeProtocol: to.Ptr(armcdn.ProbeProtocolHTTP),
	// 							ProbeRequestType: to.Ptr(armcdn.HealthProbeRequestTypeGET),
	// 						},
	// 						Origins: []*armcdn.ResourceReference{
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin1"),
	// 							},
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin2"),
	// 							},
	// 						},
	// 						ResponseBasedOriginErrorDetectionSettings: &armcdn.ResponseBasedOriginErrorDetectionParameters{
	// 							ResponseBasedDetectedErrorTypes: to.Ptr(armcdn.ResponseBasedDetectedErrorTypesTCPErrorsOnly),
	// 							ResponseBasedFailoverThresholdPercentage: to.Ptr[int32](10),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			OriginHostHeader: to.Ptr("www.bing.com"),
	// 			OriginPath: to.Ptr("/photos"),
	// 			Origins: []*armcdn.DeepCreatedOrigin{
	// 				{
	// 					Name: to.Ptr("origin1"),
	// 					Properties: &armcdn.DeepCreatedOriginProperties{
	// 						Enabled: to.Ptr(true),
	// 						HostName: to.Ptr("www.someDomain1.net"),
	// 						HTTPPort: to.Ptr[int32](80),
	// 						HTTPSPort: to.Ptr[int32](443),
	// 						OriginHostHeader: to.Ptr("www.someDomain1.net"),
	// 						Priority: to.Ptr[int32](1),
	// 						Weight: to.Ptr[int32](50),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("origin2"),
	// 					Properties: &armcdn.DeepCreatedOriginProperties{
	// 						Enabled: to.Ptr(true),
	// 						HostName: to.Ptr("www.someDomain2.net"),
	// 						HTTPPort: to.Ptr[int32](80),
	// 						HTTPSPort: to.Ptr[int32](443),
	// 						OriginHostHeader: to.Ptr("www.someDomain2.net"),
	// 						Priority: to.Ptr[int32](2),
	// 						Weight: to.Ptr[int32](50),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armcdn.EndpointProvisioningStateCreating),
	// 			QueryStringCachingBehavior: to.Ptr(armcdn.QueryStringCachingBehaviorBypassCaching),
	// 			ResourceState: to.Ptr(armcdn.EndpointResourceStateCreating),
	// 		},
	// 		Tags: map[string]*string{
	// 			"kay1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
