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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyCreate.json
func ExampleSimPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<mobile-network-name>",
		"<sim-policy-name>",
		armmobilenetwork.SimPolicy{
			Location: to.Ptr("<location>"),
			Properties: &armmobilenetwork.SimPolicyPropertiesFormat{
				DefaultSlice: &armmobilenetwork.SliceResourceID{
					ID: to.Ptr("<id>"),
				},
				RegistrationTimer: to.Ptr[int32](3240),
				SliceConfigurations: []*armmobilenetwork.SliceConfiguration{
					{
						DataNetworkConfigurations: []*armmobilenetwork.DataNetworkConfiguration{
							{
								FiveQi:                              to.Ptr[int32](9),
								AdditionalAllowedSessionTypes:       []*armmobilenetwork.PduSessionType{},
								AllocationAndRetentionPriorityLevel: to.Ptr[int32](9),
								AllowedServices: []*armmobilenetwork.ServiceResourceID{
									{
										ID: to.Ptr("<id>"),
									}},
								DataNetwork: &armmobilenetwork.DataNetworkResourceID{
									ID: to.Ptr("<id>"),
								},
								DefaultSessionType:      to.Ptr(armmobilenetwork.PduSessionTypeIPv4),
								PreemptionCapability:    to.Ptr(armmobilenetwork.PreemptionCapabilityNotPreempt),
								PreemptionVulnerability: to.Ptr(armmobilenetwork.PreemptionVulnerabilityPreemptable),
								SessionAmbr: &armmobilenetwork.Ambr{
									Downlink: to.Ptr("<downlink>"),
									Uplink:   to.Ptr("<uplink>"),
								},
							}},
						DefaultDataNetwork: &armmobilenetwork.DataNetworkResourceID{
							ID: to.Ptr("<id>"),
						},
						Slice: &armmobilenetwork.SliceResourceID{
							ID: to.Ptr("<id>"),
						},
					}},
				UeAmbr: &armmobilenetwork.Ambr{
					Downlink: to.Ptr("<downlink>"),
					Uplink:   to.Ptr("<uplink>"),
				},
			},
		},
		&armmobilenetwork.SimPoliciesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
