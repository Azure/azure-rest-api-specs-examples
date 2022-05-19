Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmobilenetwork%2Farmmobilenetwork%2Fv0.5.0/sdk/resourcemanager/mobilenetwork/armmobilenetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/ServiceCreate.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewServicesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"testMobileNetwork",
		"TestService",
		armmobilenetwork.Service{
			Location: to.Ptr("eastus"),
			Properties: &armmobilenetwork.ServicePropertiesFormat{
				PccRules: []*armmobilenetwork.PccRuleConfiguration{
					{
						RuleName:       to.Ptr("default-rule"),
						RulePrecedence: to.Ptr[int32](255),
						RuleQosPolicy: &armmobilenetwork.PccRuleQosPolicy{
							FiveQi:                              to.Ptr[int32](9),
							AllocationAndRetentionPriorityLevel: to.Ptr[int32](9),
							MaximumBitRate: &armmobilenetwork.Ambr{
								Downlink: to.Ptr("1 Gbps"),
								Uplink:   to.Ptr("500 Mbps"),
							},
							PreemptionCapability:    to.Ptr(armmobilenetwork.PreemptionCapabilityNotPreempt),
							PreemptionVulnerability: to.Ptr(armmobilenetwork.PreemptionVulnerabilityPreemptable),
						},
						ServiceDataFlowTemplates: []*armmobilenetwork.ServiceDataFlowTemplate{
							{
								Direction: to.Ptr(armmobilenetwork.SdfDirectionUplink),
								Ports:     []*string{},
								RemoteIPList: []*string{
									to.Ptr("10.3.4.0/24")},
								TemplateName: to.Ptr("IP-to-server"),
								Protocol: []*string{
									to.Ptr("ip")},
							}},
						TrafficControl: to.Ptr(armmobilenetwork.TrafficControlPermissionEnabled),
					}},
				ServicePrecedence: to.Ptr[int32](255),
				ServiceQosPolicy: &armmobilenetwork.QosPolicy{
					FiveQi:                              to.Ptr[int32](9),
					AllocationAndRetentionPriorityLevel: to.Ptr[int32](9),
					MaximumBitRate: &armmobilenetwork.Ambr{
						Downlink: to.Ptr("1 Gbps"),
						Uplink:   to.Ptr("500 Mbps"),
					},
					PreemptionCapability:    to.Ptr(armmobilenetwork.PreemptionCapabilityNotPreempt),
					PreemptionVulnerability: to.Ptr(armmobilenetwork.PreemptionVulnerabilityPreemptable),
				},
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
```
