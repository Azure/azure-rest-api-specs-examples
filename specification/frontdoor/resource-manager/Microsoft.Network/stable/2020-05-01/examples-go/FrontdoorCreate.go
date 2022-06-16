package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorCreate.json
func ExampleFrontDoorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewFrontDoorsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"frontDoor1",
		armfrontdoor.FrontDoor{
			Location: to.Ptr("westus"),
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
			Properties: &armfrontdoor.Properties{
				BackendPools: []*armfrontdoor.BackendPool{
					{
						Name: to.Ptr("backendPool1"),
						Properties: &armfrontdoor.BackendPoolProperties{
							Backends: []*armfrontdoor.Backend{
								{
									Address:   to.Ptr("w3.contoso.com"),
									HTTPPort:  to.Ptr[int32](80),
									HTTPSPort: to.Ptr[int32](443),
									Priority:  to.Ptr[int32](2),
									Weight:    to.Ptr[int32](1),
								},
								{
									Address:                    to.Ptr("contoso.com.website-us-west-2.othercloud.net"),
									HTTPPort:                   to.Ptr[int32](80),
									HTTPSPort:                  to.Ptr[int32](443),
									Priority:                   to.Ptr[int32](1),
									PrivateLinkApprovalMessage: to.Ptr("Please approve the connection request for this Private Link"),
									PrivateLinkLocation:        to.Ptr("eastus"),
									PrivateLinkResourceID:      to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Network/privateLinkServices/pls1"),
									Weight:                     to.Ptr[int32](2),
								},
								{
									Address:                    to.Ptr("10.0.1.5"),
									HTTPPort:                   to.Ptr[int32](80),
									HTTPSPort:                  to.Ptr[int32](443),
									Priority:                   to.Ptr[int32](1),
									PrivateLinkAlias:           to.Ptr("APPSERVER.d84e61f0-0870-4d24-9746-7438fa0019d1.westus2.azure.privatelinkservice"),
									PrivateLinkApprovalMessage: to.Ptr("Please approve this request to connect to the Private Link"),
									Weight:                     to.Ptr[int32](1),
								}},
							HealthProbeSettings: &armfrontdoor.SubResource{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/healthProbeSettings/healthProbeSettings1"),
							},
							LoadBalancingSettings: &armfrontdoor.SubResource{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/loadBalancingSettings/loadBalancingSettings1"),
							},
						},
					}},
				BackendPoolsSettings: &armfrontdoor.BackendPoolsSettings{
					EnforceCertificateNameCheck: to.Ptr(armfrontdoor.EnforceCertificateNameCheckEnabledStateEnabled),
					SendRecvTimeoutSeconds:      to.Ptr[int32](60),
				},
				EnabledState: to.Ptr(armfrontdoor.FrontDoorEnabledStateEnabled),
				FrontendEndpoints: []*armfrontdoor.FrontendEndpoint{
					{
						Name: to.Ptr("frontendEndpoint1"),
						Properties: &armfrontdoor.FrontendEndpointProperties{
							HostName:                    to.Ptr("www.contoso.com"),
							SessionAffinityEnabledState: to.Ptr(armfrontdoor.SessionAffinityEnabledStateEnabled),
							SessionAffinityTTLSeconds:   to.Ptr[int32](60),
							WebApplicationFirewallPolicyLink: &armfrontdoor.FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/policy1"),
							},
						},
					},
					{
						Name: to.Ptr("default"),
						Properties: &armfrontdoor.FrontendEndpointProperties{
							HostName: to.Ptr("frontDoor1.azurefd.net"),
						},
					}},
				HealthProbeSettings: []*armfrontdoor.HealthProbeSettingsModel{
					{
						Name: to.Ptr("healthProbeSettings1"),
						Properties: &armfrontdoor.HealthProbeSettingsProperties{
							Path:              to.Ptr("/"),
							EnabledState:      to.Ptr(armfrontdoor.HealthProbeEnabledEnabled),
							HealthProbeMethod: to.Ptr(armfrontdoor.FrontDoorHealthProbeMethodHEAD),
							IntervalInSeconds: to.Ptr[int32](120),
							Protocol:          to.Ptr(armfrontdoor.FrontDoorProtocolHTTP),
						},
					}},
				LoadBalancingSettings: []*armfrontdoor.LoadBalancingSettingsModel{
					{
						Name: to.Ptr("loadBalancingSettings1"),
						Properties: &armfrontdoor.LoadBalancingSettingsProperties{
							SampleSize:                to.Ptr[int32](4),
							SuccessfulSamplesRequired: to.Ptr[int32](2),
						},
					}},
				RoutingRules: []*armfrontdoor.RoutingRule{
					{
						Name: to.Ptr("routingRule1"),
						Properties: &armfrontdoor.RoutingRuleProperties{
							AcceptedProtocols: []*armfrontdoor.FrontDoorProtocol{
								to.Ptr(armfrontdoor.FrontDoorProtocolHTTP)},
							EnabledState: to.Ptr(armfrontdoor.RoutingRuleEnabledStateEnabled),
							FrontendEndpoints: []*armfrontdoor.SubResource{
								{
									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/frontendEndpoints/frontendEndpoint1"),
								},
								{
									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/frontendEndpoints/default"),
								}},
							PatternsToMatch: []*string{
								to.Ptr("/*")},
							RouteConfiguration: &armfrontdoor.ForwardingConfiguration{
								ODataType: to.Ptr("#Microsoft.Azure.FrontDoor.Models.FrontdoorForwardingConfiguration"),
								BackendPool: &armfrontdoor.SubResource{
									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/backendPools/backendPool1"),
								},
							},
							RulesEngine: &armfrontdoor.SubResource{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/rulesEngines/rulesEngine1"),
							},
							WebApplicationFirewallPolicyLink: &armfrontdoor.RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/policy1"),
							},
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
