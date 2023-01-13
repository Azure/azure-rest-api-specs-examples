package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/SimPolicyCreate.json
func ExampleSimPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "testMobileNetwork", "testPolicy", armmobilenetwork.SimPolicy{
		Location: to.Ptr("eastus"),
		Properties: &armmobilenetwork.SimPolicyPropertiesFormat{
			DefaultSlice: &armmobilenetwork.SliceResourceID{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
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
									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/services/testService"),
								}},
							DataNetwork: &armmobilenetwork.DataNetworkResourceID{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/dataNetworks/testdataNetwork"),
							},
							DefaultSessionType:             to.Ptr(armmobilenetwork.PduSessionTypeIPv4),
							MaximumNumberOfBufferedPackets: to.Ptr[int32](200),
							PreemptionCapability:           to.Ptr(armmobilenetwork.PreemptionCapabilityNotPreempt),
							PreemptionVulnerability:        to.Ptr(armmobilenetwork.PreemptionVulnerabilityPreemptable),
							SessionAmbr: &armmobilenetwork.Ambr{
								Downlink: to.Ptr("1 Gbps"),
								Uplink:   to.Ptr("500 Mbps"),
							},
						}},
					DefaultDataNetwork: &armmobilenetwork.DataNetworkResourceID{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/dataNetworks/testdataNetwork"),
					},
					Slice: &armmobilenetwork.SliceResourceID{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
					},
				}},
			UeAmbr: &armmobilenetwork.Ambr{
				Downlink: to.Ptr("1 Gbps"),
				Uplink:   to.Ptr("500 Mbps"),
			},
		},
	}, nil)
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
