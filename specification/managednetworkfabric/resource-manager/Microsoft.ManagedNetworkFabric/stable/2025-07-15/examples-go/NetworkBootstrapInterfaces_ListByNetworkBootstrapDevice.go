package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkBootstrapInterfaces_ListByNetworkBootstrapDevice.json
func ExampleNetworkBootstrapInterfacesClient_NewListByNetworkBootstrapDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("1234ABCD-0A1B-1234-5678-123456ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkBootstrapInterfacesClient().NewListByNetworkBootstrapDevicePager("example-rg", "example-device", nil)
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
		// page = armmanagednetworkfabric.NetworkBootstrapInterfacesClientListByNetworkBootstrapDeviceResponse{
		// 	NetworkBootstrapInterfaceListResult: armmanagednetworkfabric.NetworkBootstrapInterfaceListResult{
		// 		Value: []*armmanagednetworkfabric.NetworkBootstrapInterface{
		// 			{
		// 				Properties: &armmanagednetworkfabric.NetworkBootstrapInterfaceProperties{
		// 					Annotation: to.Ptr("annotation"),
		// 					ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
		// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		// 					PhysicalIdentifier: to.Ptr("physicalIdentifier"),
		// 					ConnectedTo: to.Ptr("connectedTo"),
		// 					InterfaceType: to.Ptr(armmanagednetworkfabric.InterfaceTypeManagement),
		// 					AdditionalDescription: to.Ptr("additionalDescription"),
		// 					IPv4Address: to.Ptr("10.10.10.10"),
		// 					IPv6Address: to.Ptr("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
		// 					SerialNumber: to.Ptr("Make;Model;HardwareRevisionId;SerialNumber"),
		// 					Description: to.Ptr("description"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkBootstrapDevices/example-device/networkBootstrapInterfaces/example-interface"),
		// 				Name: to.Ptr("example-interface"),
		// 				Type: to.Ptr("microsoft.managednetworkfabric/networkBootstrapDevices/networkBootstrapInterfaces"),
		// 				SystemData: &armmanagednetworkfabric.SystemData{
		// 					CreatedBy: to.Ptr("email@address.com"),
		// 					CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("UserId"),
		// 					LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/abqlveiv"),
		// 	},
		// }
	}
}
