package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskExamples/Disk_Update_AddPurchasePlan.json
func ExampleDisksClient_BeginUpdate_updateAManagedDiskToAddPurchasePlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDisksClient().BeginUpdate(ctx, "myResourceGroup", "myDisk", armcompute.DiskUpdate{
		Properties: &armcompute.DiskUpdateProperties{
			PurchasePlan: &armcompute.DiskPurchasePlan{
				Name:          to.Ptr("myPurchasePlanName"),
				Product:       to.Ptr("myPurchasePlanProduct"),
				PromotionCode: to.Ptr("myPurchasePlanPromotionCode"),
				Publisher:     to.Ptr("myPurchasePlanPublisher"),
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
	// res.Disk = armcompute.Disk{
	// 	Name: to.Ptr("myDisk"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcompute.DiskProperties{
	// 		CreationData: &armcompute.CreationData{
	// 			CreateOption: to.Ptr(armcompute.DiskCreateOptionFromImage),
	// 			ImageReference: &armcompute.ImageDiskReference{
	// 				ID: to.Ptr("/Subscriptions/{subscription-id}/Providers/Microsoft.Compute/Locations/westus/Publishers/test_test_pmc2pc1/ArtifactTypes/VMImage/Offers/marketplace_vm_test/Skus/test_sku/Versions/1.0.0"),
	// 			},
	// 		},
	// 		DiskSizeGB: to.Ptr[int32](127),
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 		OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PurchasePlan: &armcompute.DiskPurchasePlan{
	// 			Name: to.Ptr("myPurchasePlanName"),
	// 			Product: to.Ptr("myPurchasePlanProduct"),
	// 			PromotionCode: to.Ptr("myPurchasePlanPromotionCode"),
	// 			Publisher: to.Ptr("myPurchasePlanPublisher"),
	// 		},
	// 	},
	// 	SKU: &armcompute.DiskSKU{
	// 		Name: to.Ptr(armcompute.DiskStorageAccountTypesStandardLRS),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// }
}
