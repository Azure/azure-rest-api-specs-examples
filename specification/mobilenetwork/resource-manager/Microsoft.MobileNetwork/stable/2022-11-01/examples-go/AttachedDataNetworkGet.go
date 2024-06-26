package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/340d577969b7bff5ad0488d79543314bc17daa50/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/AttachedDataNetworkGet.json
func ExampleAttachedDataNetworksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAttachedDataNetworksClient().Get(ctx, "rg1", "TestPacketCoreCP", "TestPacketCoreDP", "TestAttachedDataNetwork", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AttachedDataNetwork = armmobilenetwork.AttachedDataNetwork{
	// 	Name: to.Ptr("TestAttachedDataNetwork"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/attachedDataNetwork"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
	// 	SystemData: &armmobilenetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmobilenetwork.AttachedDataNetworkPropertiesFormat{
	// 		DNSAddresses: []*string{
	// 			to.Ptr("1.1.1.1")},
	// 			NaptConfiguration: &armmobilenetwork.NaptConfiguration{
	// 				Enabled: to.Ptr(armmobilenetwork.NaptEnabledEnabled),
	// 				PinholeLimits: to.Ptr[int32](65536),
	// 				PinholeTimeouts: &armmobilenetwork.PinholeTimeouts{
	// 					Icmp: to.Ptr[int32](30),
	// 					TCP: to.Ptr[int32](180),
	// 					UDP: to.Ptr[int32](30),
	// 				},
	// 				PortRange: &armmobilenetwork.PortRange{
	// 					MaxPort: to.Ptr[int32](49999),
	// 					MinPort: to.Ptr[int32](1024),
	// 				},
	// 				PortReuseHoldTime: &armmobilenetwork.PortReuseHoldTimes{
	// 					TCP: to.Ptr[int32](120),
	// 					UDP: to.Ptr[int32](60),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
	// 			UserEquipmentAddressPoolPrefix: []*string{
	// 				to.Ptr("2.2.0.0/16")},
	// 				UserEquipmentStaticAddressPoolPrefix: []*string{
	// 					to.Ptr("2.4.0.0/16")},
	// 					UserPlaneDataInterface: &armmobilenetwork.InterfaceProperties{
	// 						Name: to.Ptr("N6"),
	// 					},
	// 				},
	// 			}
}
