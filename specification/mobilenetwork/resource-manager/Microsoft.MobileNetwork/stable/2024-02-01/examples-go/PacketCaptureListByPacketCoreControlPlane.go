package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/PacketCaptureListByPacketCoreControlPlane.json
func ExamplePacketCapturesClient_NewListByPacketCoreControlPlanePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPacketCapturesClient().NewListByPacketCoreControlPlanePager("rg1", "TestPacketCoreCP", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PacketCaptureListResult = armmobilenetwork.PacketCaptureListResult{
		// 	Value: []*armmobilenetwork.PacketCapture{
		// 		{
		// 			Name: to.Ptr("pc1"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCaptures/pc1"),
		// 			Properties: &armmobilenetwork.PacketCapturePropertiesFormat{
		// 				BytesToCapturePerPacket: to.Ptr[int64](10000),
		// 				NetworkInterfaces: []*string{
		// 					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP"),
		// 					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP"),
		// 					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestADN")},
		// 					ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
		// 					Status: to.Ptr(armmobilenetwork.PacketCaptureStatusStopped),
		// 					TimeLimitInSeconds: to.Ptr[int32](100),
		// 					TotalBytesPerSession: to.Ptr[int64](100000),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("pc2"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCaptures/pc2"),
		// 				Properties: &armmobilenetwork.PacketCapturePropertiesFormat{
		// 					BytesToCapturePerPacket: to.Ptr[int64](10000),
		// 					NetworkInterfaces: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestADN")},
		// 						ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
		// 						Status: to.Ptr(armmobilenetwork.PacketCaptureStatusStopped),
		// 						TimeLimitInSeconds: to.Ptr[int32](100),
		// 						TotalBytesPerSession: to.Ptr[int64](100000),
		// 					},
		// 			}},
		// 		}
	}
}
