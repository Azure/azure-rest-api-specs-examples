package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/AttachedDataNetworkCreate.json
func ExampleAttachedDataNetworksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAttachedDataNetworksClient().BeginCreateOrUpdate(ctx, "rg1", "TestPacketCoreCP", "TestPacketCoreDP", "TestAttachedDataNetwork", armmobilenetwork.AttachedDataNetwork{
		Location: to.Ptr("eastus"),
		Properties: &armmobilenetwork.AttachedDataNetworkPropertiesFormat{
			DNSAddresses: []*string{
				to.Ptr("1.1.1.1")},
			NaptConfiguration: &armmobilenetwork.NaptConfiguration{
				Enabled:       to.Ptr(armmobilenetwork.NaptEnabledEnabled),
				PinholeLimits: to.Ptr[int32](65536),
				PinholeTimeouts: &armmobilenetwork.PinholeTimeouts{
					Icmp: to.Ptr[int32](30),
					TCP:  to.Ptr[int32](180),
					UDP:  to.Ptr[int32](30),
				},
				PortRange: &armmobilenetwork.PortRange{
					MaxPort: to.Ptr[int32](49999),
					MinPort: to.Ptr[int32](1024),
				},
				PortReuseHoldTime: &armmobilenetwork.PortReuseHoldTimes{
					TCP: to.Ptr[int32](120),
					UDP: to.Ptr[int32](60),
				},
			},
			UserEquipmentAddressPoolPrefix: []*string{
				to.Ptr("2.2.0.0/16")},
			UserEquipmentStaticAddressPoolPrefix: []*string{
				to.Ptr("2.4.0.0/16")},
			UserPlaneDataInterface: &armmobilenetwork.InterfaceProperties{
				Name: to.Ptr("N6"),
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AttachedDataNetwork = armmobilenetwork.AttachedDataNetwork{
	// 	Name: to.Ptr("TestAttachedDataNetwork"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/attachedDataNetwork"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
	// 	SystemData: &armmobilenetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
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
