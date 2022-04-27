Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmobilenetwork%2Farmmobilenetwork%2Fv0.4.0/sdk/resourcemanager/mobilenetwork/armmobilenetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armmobilenetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/ServiceCreate.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewServicesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<mobile-network-name>",
		"<service-name>",
		armmobilenetwork.Service{
			Location: to.Ptr("<location>"),
			Properties: &armmobilenetwork.ServicePropertiesFormat{
				PccRules: []*armmobilenetwork.PccRuleConfiguration{
					{
						RuleName:       to.Ptr("<rule-name>"),
						RulePrecedence: to.Ptr[int32](255),
						RuleQosPolicy: &armmobilenetwork.PccRuleQosPolicy{
							FiveQi:                              to.Ptr[int32](9),
							AllocationAndRetentionPriorityLevel: to.Ptr[int32](9),
							MaximumBitRate: &armmobilenetwork.Ambr{
								Downlink: to.Ptr("<downlink>"),
								Uplink:   to.Ptr("<uplink>"),
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
								TemplateName: to.Ptr("<template-name>"),
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
						Downlink: to.Ptr("<downlink>"),
						Uplink:   to.Ptr("<uplink>"),
					},
					PreemptionCapability:    to.Ptr(armmobilenetwork.PreemptionCapabilityNotPreempt),
					PreemptionVulnerability: to.Ptr(armmobilenetwork.PreemptionVulnerabilityPreemptable),
				},
			},
		},
		&armmobilenetwork.ServicesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
