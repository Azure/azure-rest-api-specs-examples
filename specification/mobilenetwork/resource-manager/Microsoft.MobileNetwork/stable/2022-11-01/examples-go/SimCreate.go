package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/SimCreate.json
func ExampleSimsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "testSimGroup", "testSim", armmobilenetwork.Sim{
		Properties: &armmobilenetwork.SimPropertiesFormat{
			DeviceType:                            to.Ptr("Video camera"),
			IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000000"),
			InternationalMobileSubscriberIdentity: to.Ptr("00000"),
			SimPolicy: &armmobilenetwork.SimPolicyResourceID{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
			},
			StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
				{
					AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
					},
					Slice: &armmobilenetwork.SliceResourceID{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
					},
					StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
						IPv4Address: to.Ptr("2.4.0.1"),
					},
				}},
			AuthenticationKey: to.Ptr("00000000000000000000000000000000"),
			OperatorKeyCode:   to.Ptr("00000000000000000000000000000000"),
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
