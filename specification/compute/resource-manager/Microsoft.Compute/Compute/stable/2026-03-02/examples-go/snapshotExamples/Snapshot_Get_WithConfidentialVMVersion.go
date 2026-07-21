package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-02/snapshotExamples/Snapshot_Get_WithConfidentialVMVersion.json
func ExampleSnapshotsClient_Get_getInformationAboutAConfidentialVMSnapshotWithConfidentialVMVersion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSnapshotsClient().Get(ctx, "myResourceGroup", "myConfidentialSnapshot", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.SnapshotsClientGetResponse{
	// 	Snapshot: armcompute.Snapshot{
	// 		Properties: &armcompute.SnapshotProperties{
	// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			SecurityProfile: &armcompute.DiskSecurityProfile{
	// 				SecurityType: to.Ptr(armcompute.DiskSecurityTypesConfidentialVMDiskEncryptedWithPlatformKey),
	// 				ConfidentialVMVersion: to.Ptr(armcompute.ConfidentialVMVersionV2),
	// 			},
	// 			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV2),
	// 			PurchasePlan: &armcompute.DiskPurchasePlan{
	// 				Name: to.Ptr("test_sku"),
	// 				Publisher: to.Ptr("test_test_pmc2pc1"),
	// 				Product: to.Ptr("marketplace_vm_test"),
	// 			},
	// 			SupportedCapabilities: &armcompute.SupportedCapabilities{
	// 				AcceleratedNetwork: to.Ptr(true),
	// 			},
	// 			SupportsHibernation: to.Ptr(true),
	// 			CreationData: &armcompute.CreationData{
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionCopy),
	// 				SourceResourceID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myConfidentialDisk"),
	// 				SourceUniqueID: to.Ptr("d633885d-d102-4481-901e-5b2413d1a7be"),
	// 			},
	// 			DiskSizeGB: to.Ptr[int32](10),
	// 			Encryption: &armcompute.Encryption{
	// 				Type: to.Ptr(armcompute.EncryptionTypeEncryptionAtRestWithPlatformKey),
	// 			},
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-20T04:41:35.079872+00:00"); return t}()),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 		Type: to.Ptr("Microsoft.Compute/snapshots"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("Development"),
	// 			"project": to.Ptr("Snapshots"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/myConfidentialSnapshot"),
	// 		Name: to.Ptr("myConfidentialSnapshot"),
	// 	},
	// }
}
