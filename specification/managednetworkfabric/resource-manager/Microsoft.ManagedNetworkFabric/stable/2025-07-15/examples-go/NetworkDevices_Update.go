package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkDevices_Update.json
func ExampleNetworkDevicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkDevicesClient().BeginUpdate(ctx, "example-rg", "example-device", armmanagednetworkfabric.NetworkDevicePatchParameters{
		Tags: map[string]*string{
			"KeyId": to.Ptr("KeyValue"),
		},
		Identity: &armmanagednetworkfabric.ManagedServiceIdentityPatch{
			Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
				"/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/example-identity": {},
			},
		},
		Properties: &armmanagednetworkfabric.NetworkDevicePatchParametersProperties{
			Annotation: to.Ptr("annotation"),
			HostName:   to.Ptr("NFA-Device"),
			IdentitySelector: &armmanagednetworkfabric.IdentitySelectorPatch{
				IdentityType:                   to.Ptr(armmanagednetworkfabric.ManagedServiceIdentitySelectorTypeUserAssignedIdentity),
				UserAssignedIdentityResourceID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/example-identity"),
			},
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
	// res = armmanagednetworkfabric.NetworkDevicesClientUpdateResponse{
	// 	NetworkDevice: &armmanagednetworkfabric.NetworkDevice{
	// 		Properties: &armmanagednetworkfabric.NetworkDeviceProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			HostName: to.Ptr("NFA-Device"),
	// 			SerialNumber: to.Ptr("Vendor;DCS-7280XXX-24;12.05;JPE2111XXXX"),
	// 			Version: to.Ptr("1.0"),
	// 			NetworkDeviceSKU: to.Ptr("DeviceSku"),
	// 			NetworkDeviceRole: to.Ptr(armmanagednetworkfabric.NetworkDeviceRoleCE),
	// 			NetworkRackID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkRacks/example-rack"),
	// 			ManagementIPv4Address: to.Ptr("10.10.10.10"),
	// 			ManagementIPv6Address: to.Ptr("2F::1/100"),
	// 			RwDeviceConfig: to.Ptr("hostname access-switch1\n enable secret somestrongpass"),
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			IdentitySelector: &armmanagednetworkfabric.IdentitySelector{
	// 				IdentityType: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentitySelectorTypeUserAssignedIdentity),
	// 				UserAssignedIdentityResourceID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/example-identity"),
	// 			},
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		},
	// 		Tags: map[string]*string{
	// 			"KeyId": to.Ptr("KeyValue"),
	// 		},
	// 		Identity: &armmanagednetworkfabric.ManagedServiceIdentity{
	// 			Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
	// 				"/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/example-identity": &armmanagednetworkfabric.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 					ClientID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("eastuseuap"),
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-device"),
	// 		Name: to.Ptr("example-device"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/networkdevices"),
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
