package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/46c51b03d99b113ecc3b38883e3cb2d395fe94a4/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/StorageAppliances_Create.json
func ExampleStorageAppliancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageAppliancesClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "storageApplianceName", armnetworkcloud.StorageAppliance{
		Location: to.Ptr("location"),
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
		},
		ExtendedLocation: &armnetworkcloud.ExtendedLocation{
			Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
			Type: to.Ptr("CustomLocation"),
		},
		Properties: &armnetworkcloud.StorageApplianceProperties{
			AdministratorCredentials: &armnetworkcloud.AdministrativeCredentials{
				Password: to.Ptr("{password}"),
				Username: to.Ptr("adminUser"),
			},
			RackID:                to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName"),
			RackSlot:              to.Ptr[int64](1),
			SerialNumber:          to.Ptr("BM1219XXX"),
			StorageApplianceSKUID: to.Ptr("684E-3B16-399E"),
		},
	}, &armnetworkcloud.StorageAppliancesClientBeginCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
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
	// res.StorageAppliance = armnetworkcloud.StorageAppliance{
	// 	Name: to.Ptr("storageApplianceName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/storageAppliances"),
	// 	ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/storageAppliances/storageApplianceName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.StorageApplianceProperties{
	// 		AdministratorCredentials: &armnetworkcloud.AdministrativeCredentials{
	// 			Username: to.Ptr("adminUser"),
	// 		},
	// 		CaCertificate: &armnetworkcloud.CertificateInfo{
	// 			Hash: to.Ptr("dea698309efd2830a1d440a807650d9aa6d954b3243ab8cb556ac98c1f3faa60"),
	// 			Value: to.Ptr("-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWgAwIBAgIJAL4a5b1d8f2wM...A0GCSqGSIb3DQEBCwUAMEUxCzAJB==\n-----END CERTIFICATE-----"),
	// 		},
	// 		Capacity: to.Ptr[int64](893),
	// 		CapacityUsed: to.Ptr[int64](500),
	// 		ClusterID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.StorageApplianceDetailedStatusAvailable),
	// 		DetailedStatusMessage: to.Ptr("Storage appliance is up and running"),
	// 		ManagementIPv4Address: to.Ptr("192.0.2.2"),
	// 		Manufacturer: to.Ptr("Contoso Storage"),
	// 		Model: to.Ptr("ArrayStore-Flash70"),
	// 		ProvisioningState: to.Ptr(armnetworkcloud.StorageApplianceProvisioningStateSucceeded),
	// 		RackID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName"),
	// 		RackSlot: to.Ptr[int64](1),
	// 		RemoteVendorManagementFeature: to.Ptr(armnetworkcloud.RemoteVendorManagementFeatureSupported),
	// 		RemoteVendorManagementStatus: to.Ptr(armnetworkcloud.RemoteVendorManagementStatusEnabled),
	// 		SecretRotationStatus: []*armnetworkcloud.SecretRotationStatus{
	// 			{
	// 				ExpirePeriodDays: to.Ptr[int64](90),
	// 				LastRotationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-30T13:27:23.103Z"); return t}()),
	// 				RotationPeriodDays: to.Ptr[int64](60),
	// 				SecretArchiveReference: &armnetworkcloud.SecretArchiveReference{
	// 					KeyVaultID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.KeyVault/vaults/keyVaultName"),
	// 					SecretName: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff-resource-group-cluster-1679871-storage-appliance-credential-manager-ffffffff"),
	// 					SecretVersion: to.Ptr("02ab6c1f9c0f4982b0632b0d5d74a33b"),
	// 				},
	// 				SecretType: to.Ptr("Storage Appliance User"),
	// 		}},
	// 		SerialNumber: to.Ptr("BM1219XXX"),
	// 		StorageApplianceSKUID: to.Ptr("684E-3B16-399E"),
	// 		Version: to.Ptr("9.9.9"),
	// 	},
	// }
}
