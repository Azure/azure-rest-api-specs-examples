package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkDevices_Get.json
func ExampleNetworkDevicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkDevicesClient().Get(ctx, "example-rg", "example-device", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmanagednetworkfabric.NetworkDevicesClientGetResponse{
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
	// 			NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
	// 			LastOperation: &armmanagednetworkfabric.LastOperationProperties{
	// 				Details: to.Ptr("Succeeded"),
	// 			},
	// 			SecretRotationStatus: []*armmanagednetworkfabric.SecretRotationStatus{
	// 				{
	// 					LastRotationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-09T04:51:41.251Z"); return t}()),
	// 					SynchronizationStatus: to.Ptr(armmanagednetworkfabric.SynchronizationStatusInSync),
	// 					SecretArchiveReference: &armmanagednetworkfabric.SecretArchiveReference{
	// 						KeyVaultURI: to.Ptr("https://example-kv.vault.azure.net/secrets/example-secret-1/7e61b8efbcdd4e28963560dba3021df7"),
	// 						KeyVaultID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.KeyVault/vaults/example-kv"),
	// 						SecretName: to.Ptr("example-secret-1"),
	// 						SecretVersion: to.Ptr("7e61b8efbcdd4e28963560dba3021df7"),
	// 					},
	// 					SecretType: to.Ptr("Admin user password"),
	// 				},
	// 				{
	// 					LastRotationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-09T02:32:12.904Z"); return t}()),
	// 					SynchronizationStatus: to.Ptr(armmanagednetworkfabric.SynchronizationStatusOutOfSync),
	// 					SecretArchiveReference: &armmanagednetworkfabric.SecretArchiveReference{
	// 						KeyVaultURI: to.Ptr("https://example-kv.vault.azure.net/secrets/example-secret-2/11a536561da34d6b8b452d880df58f3a"),
	// 						KeyVaultID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.KeyVault/vaults/example-kv"),
	// 						SecretName: to.Ptr("example-secret-2"),
	// 						SecretVersion: to.Ptr("11a536561da34d6b8b452d880df58f3a"),
	// 					},
	// 					SecretType: to.Ptr("AzureOperatorRW user password"),
	// 				},
	// 				{
	// 					LastRotationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-09T02:32:12.904Z"); return t}()),
	// 					SynchronizationStatus: to.Ptr(armmanagednetworkfabric.SynchronizationStatusSynchronizing),
	// 					SecretArchiveReference: &armmanagednetworkfabric.SecretArchiveReference{
	// 						KeyVaultURI: to.Ptr("https://example-kv.vault.azure.net/secrets/example-secret-3/22b646561da34d6b8b452d880df69a4b"),
	// 						KeyVaultID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.KeyVault/vaults/example-kv"),
	// 						SecretName: to.Ptr("example-secret-3"),
	// 						SecretVersion: to.Ptr("22b646561da34d6b8b452d880df69a4b"),
	// 					},
	// 					SecretType: to.Ptr("AzureOperatorRO user password"),
	// 				},
	// 			},
	// 			CertificateRotationStatus: []*armmanagednetworkfabric.CertificateRotationStatus{
	// 				{
	// 					LastRotationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-09T04:51:41.251Z"); return t}()),
	// 					SynchronizationStatus: to.Ptr(armmanagednetworkfabric.SynchronizationStatusInSync),
	// 					CertificateArchiveReference: &armmanagednetworkfabric.CertificateArchiveReference{
	// 						KeyVaultURI: to.Ptr("https://example-kv.vault.azure.net/certificates/example-certificate-1/9f3c2a7d4e8b4c1a9a7f2e6d5c3b1a2f"),
	// 						KeyVaultID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.KeyVault/vaults/example-kv"),
	// 						CertificateName: to.Ptr("example-certificate-1"),
	// 						CertificateVersion: to.Ptr("9f3c2a7d4e8b4c1a9a7f2e6d5c3b1a2f"),
	// 					},
	// 					CertificateType: to.Ptr("gNMI server certificate"),
	// 				},
	// 			},
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		},
	// 		Tags: map[string]*string{
	// 			"KeyId": to.Ptr("KeyValue"),
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
