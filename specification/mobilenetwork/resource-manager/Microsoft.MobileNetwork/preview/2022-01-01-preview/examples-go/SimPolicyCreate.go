package armmobilenetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SimPolicyCreate.json
func ExampleSimPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmobilenetwork.NewSimPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<mobile-network-name>",
		"<sim-policy-name>",
		armmobilenetwork.SimPolicy{
			Location: to.StringPtr("<location>"),
			Properties: &armmobilenetwork.SimPolicyPropertiesFormat{
				DefaultSlice: &armmobilenetwork.SliceResourceID{
					ID: to.StringPtr("<id>"),
				},
				RegistrationTimer: to.Int32Ptr(3240),
				SliceConfigurations: []*armmobilenetwork.SliceConfiguration{
					{
						DataNetworkConfigurations: []*armmobilenetwork.DataNetworkConfiguration{
							{
								FiveQi:                              to.Int32Ptr(9),
								AdditionalAllowedSessionTypes:       []*armmobilenetwork.PduSessionType{},
								AllocationAndRetentionPriorityLevel: to.Int32Ptr(9),
								AllowedServices: []*armmobilenetwork.ServiceResourceID{
									{
										ID: to.StringPtr("<id>"),
									}},
								DataNetwork: &armmobilenetwork.DataNetworkResourceID{
									ID: to.StringPtr("<id>"),
								},
								DefaultSessionType:      armmobilenetwork.PduSessionType("IPv4").ToPtr(),
								PreemptionCapability:    armmobilenetwork.PreemptionCapability("NotPreempt").ToPtr(),
								PreemptionVulnerability: armmobilenetwork.PreemptionVulnerability("Preemptable").ToPtr(),
								SessionAmbr: &armmobilenetwork.Ambr{
									Downlink: to.StringPtr("<downlink>"),
									Uplink:   to.StringPtr("<uplink>"),
								},
							}},
						DefaultDataNetwork: &armmobilenetwork.DataNetworkResourceID{
							ID: to.StringPtr("<id>"),
						},
						Slice: &armmobilenetwork.SliceResourceID{
							ID: to.StringPtr("<id>"),
						},
					}},
				UeAmbr: &armmobilenetwork.Ambr{
					Downlink: to.StringPtr("<downlink>"),
					Uplink:   to.StringPtr("<uplink>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SimPoliciesClientCreateOrUpdateResult)
}
