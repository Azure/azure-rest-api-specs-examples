package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkBootstrapDevices_Update.json
func ExampleNetworkBootstrapDevicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("1234ABCD-0A1B-1234-5678-123456ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkBootstrapDevicesClient().BeginUpdate(ctx, "example-rg", "example-device", armmanagednetworkfabric.NetworkBootstrapDevicePatch{
		Tags: map[string]*string{},
		Identity: &armmanagednetworkfabric.ManagedServiceIdentityPatch{
			Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
				"key8793": {},
			},
		},
		Properties: &armmanagednetworkfabric.NetworkBootstrapDevicePatchProperties{
			Annotation:   to.Ptr("annotation"),
			HostName:     to.Ptr("NFA-Device"),
			SerialNumber: to.Ptr("Vendor;DCS-7280XXX-24;12.05;JPE2111XXXX"),
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
	// res = armmanagednetworkfabric.NetworkBootstrapDevicesClientUpdateResponse{
	// 	NetworkBootstrapDevice: &armmanagednetworkfabric.NetworkBootstrapDevice{
	// 		Properties: &armmanagednetworkfabric.NetworkBootstrapDeviceProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			HostName: to.Ptr("NFA-Device"),
	// 			SerialNumber: to.Ptr("Vendor;DCS-7280XXX-24;12.05;JPE2111XXXX"),
	// 			NetworkDeviceSKU: to.Ptr("DeviceSku"),
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 			Version: to.Ptr("1.0"),
	// 			NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
	// 			SecondaryManagementIPv4Address: to.Ptr("10.10.10.10"),
	// 			DhcpV4ServerIPAddress: to.Ptr("10.10.10.10"),
	// 			PrimaryManagementIPv6Address: to.Ptr("2F::1/100"),
	// 			SecondaryManagementIPv6Address: to.Ptr("2F::1/100"),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			PrimaryManagementIPv4Address: to.Ptr("10.10.10.10"),
	// 		},
	// 		Identity: &armmanagednetworkfabric.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 			TenantID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 			Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
	// 				"key4876": &armmanagednetworkfabric.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 					ClientID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("eastuseuap"),
	// 		ID: to.Ptr("/Subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkBootstrapDevices/example-device"),
	// 		Name: to.Ptr("example-device"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/networkbootstrapdevices"),
	// 		SystemData: &armmanagednetworkfabric.SystemData{
	// 			CreatedBy: to.Ptr("email@address.com"),
	// 			CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("UserId"),
	// 			LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 		},
	// 	},
	// }
}
