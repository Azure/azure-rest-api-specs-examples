Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ffrontdoor%2Farmfrontdoor%2Fv0.4.0/sdk/resourcemanager/frontdoor/armfrontdoor/README.md) on how to add the SDK to your project and authenticate.

```go
package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorCreate.json
func ExampleFrontDoorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		armfrontdoor.FrontDoor{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
			Properties: &armfrontdoor.Properties{
				BackendPools: []*armfrontdoor.BackendPool{
					{
						Name: to.Ptr("<name>"),
						Properties: &armfrontdoor.BackendPoolProperties{
							Backends: []*armfrontdoor.Backend{
								{
									Address:   to.Ptr("<address>"),
									HTTPPort:  to.Ptr[int32](80),
									HTTPSPort: to.Ptr[int32](443),
									Priority:  to.Ptr[int32](2),
									Weight:    to.Ptr[int32](1),
								},
								{
									Address:                    to.Ptr("<address>"),
									HTTPPort:                   to.Ptr[int32](80),
									HTTPSPort:                  to.Ptr[int32](443),
									Priority:                   to.Ptr[int32](1),
									PrivateLinkApprovalMessage: to.Ptr("<private-link-approval-message>"),
									PrivateLinkLocation:        to.Ptr("<private-link-location>"),
									PrivateLinkResourceID:      to.Ptr("<private-link-resource-id>"),
									Weight:                     to.Ptr[int32](2),
								},
								{
									Address:                    to.Ptr("<address>"),
									HTTPPort:                   to.Ptr[int32](80),
									HTTPSPort:                  to.Ptr[int32](443),
									Priority:                   to.Ptr[int32](1),
									PrivateLinkAlias:           to.Ptr("<private-link-alias>"),
									PrivateLinkApprovalMessage: to.Ptr("<private-link-approval-message>"),
									Weight:                     to.Ptr[int32](1),
								}},
							HealthProbeSettings: &armfrontdoor.SubResource{
								ID: to.Ptr("<id>"),
							},
							LoadBalancingSettings: &armfrontdoor.SubResource{
								ID: to.Ptr("<id>"),
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
						Name: to.Ptr("<name>"),
						Properties: &armfrontdoor.FrontendEndpointProperties{
							HostName:                    to.Ptr("<host-name>"),
							SessionAffinityEnabledState: to.Ptr(armfrontdoor.SessionAffinityEnabledStateEnabled),
							SessionAffinityTTLSeconds:   to.Ptr[int32](60),
							WebApplicationFirewallPolicyLink: &armfrontdoor.FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink{
								ID: to.Ptr("<id>"),
							},
						},
					},
					{
						Name: to.Ptr("<name>"),
						Properties: &armfrontdoor.FrontendEndpointProperties{
							HostName: to.Ptr("<host-name>"),
						},
					}},
				HealthProbeSettings: []*armfrontdoor.HealthProbeSettingsModel{
					{
						Name: to.Ptr("<name>"),
						Properties: &armfrontdoor.HealthProbeSettingsProperties{
							Path:              to.Ptr("<path>"),
							EnabledState:      to.Ptr(armfrontdoor.HealthProbeEnabledEnabled),
							HealthProbeMethod: to.Ptr(armfrontdoor.FrontDoorHealthProbeMethodHEAD),
							IntervalInSeconds: to.Ptr[int32](120),
							Protocol:          to.Ptr(armfrontdoor.FrontDoorProtocolHTTP),
						},
					}},
				LoadBalancingSettings: []*armfrontdoor.LoadBalancingSettingsModel{
					{
						Name: to.Ptr("<name>"),
						Properties: &armfrontdoor.LoadBalancingSettingsProperties{
							SampleSize:                to.Ptr[int32](4),
							SuccessfulSamplesRequired: to.Ptr[int32](2),
						},
					}},
				RoutingRules: []*armfrontdoor.RoutingRule{
					{
						Name: to.Ptr("<name>"),
						Properties: &armfrontdoor.RoutingRuleProperties{
							AcceptedProtocols: []*armfrontdoor.FrontDoorProtocol{
								to.Ptr(armfrontdoor.FrontDoorProtocolHTTP)},
							EnabledState: to.Ptr(armfrontdoor.RoutingRuleEnabledStateEnabled),
							FrontendEndpoints: []*armfrontdoor.SubResource{
								{
									ID: to.Ptr("<id>"),
								},
								{
									ID: to.Ptr("<id>"),
								}},
							PatternsToMatch: []*string{
								to.Ptr("/*")},
							RouteConfiguration: &armfrontdoor.ForwardingConfiguration{
								ODataType: to.Ptr("<odata-type>"),
								BackendPool: &armfrontdoor.SubResource{
									ID: to.Ptr("<id>"),
								},
							},
							RulesEngine: &armfrontdoor.SubResource{
								ID: to.Ptr("<id>"),
							},
							WebApplicationFirewallPolicyLink: &armfrontdoor.RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink{
								ID: to.Ptr("<id>"),
							},
						},
					}},
			},
		},
		&armfrontdoor.FrontDoorsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
