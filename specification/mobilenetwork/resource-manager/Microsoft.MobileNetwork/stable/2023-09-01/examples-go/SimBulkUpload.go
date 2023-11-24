package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/SimBulkUpload.json
func ExampleSimsClient_BeginBulkUpload() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSimsClient().BeginBulkUpload(ctx, "rg1", "testSimGroup", armmobilenetwork.SimUploadList{
		Sims: []*armmobilenetwork.SimNameAndProperties{
			{
				Name: to.Ptr("testSim"),
				Properties: &armmobilenetwork.SimPropertiesFormat{
					DeviceType:                            to.Ptr("Video camera"),
					IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000000"),
					InternationalMobileSubscriberIdentity: to.Ptr("00000"),
					SimPolicy: &armmobilenetwork.SimPolicyResourceID{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
					},
					StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
						{
							AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
							},
							Slice: &armmobilenetwork.SliceResourceID{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
							},
							StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
								IPv4Address: to.Ptr("2.4.0.1"),
							},
						}},
					AuthenticationKey: to.Ptr("00000000000000000000000000000000"),
					OperatorKeyCode:   to.Ptr("00000000000000000000000000000000"),
				},
			},
			{
				Name: to.Ptr("testSim2"),
				Properties: &armmobilenetwork.SimPropertiesFormat{
					DeviceType:                            to.Ptr("Video camera"),
					IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000001"),
					InternationalMobileSubscriberIdentity: to.Ptr("00000"),
					SimPolicy: &armmobilenetwork.SimPolicyResourceID{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
					},
					StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
						{
							AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
							},
							Slice: &armmobilenetwork.SliceResourceID{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
							},
							StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
								IPv4Address: to.Ptr("2.4.0.2"),
							},
						}},
					AuthenticationKey: to.Ptr("00000000000000000000000000000000"),
					OperatorKeyCode:   to.Ptr("00000000000000000000000000000000"),
				},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AsyncOperationStatus = armmobilenetwork.AsyncOperationStatus{
	// 	Name: to.Ptr("testOperation"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:38:07.000Z"); return t}()),
	// 	ID: to.Ptr("/providers/Microsoft.MobileNetwork/locations/testLocation/operationStatuses/testOperation"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:36:07.000Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
