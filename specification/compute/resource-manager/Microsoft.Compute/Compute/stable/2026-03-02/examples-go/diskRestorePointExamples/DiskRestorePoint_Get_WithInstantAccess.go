package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-02/diskRestorePointExamples/DiskRestorePoint_Get_WithInstantAccess.json
func ExampleDiskRestorePointClient_Get_getADiskRestorePointResourceWithInstantAccessSnapshotAccessState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiskRestorePointClient().Get(ctx, "myResourceGroup", "rpc", "vmrp", "TestDisk45ceb03433006d1baee_5c1528-43e2-4c77-9c55-a78bf5a5fc88", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.DiskRestorePointClientGetResponse{
	// 	DiskRestorePoint: armcompute.DiskRestorePoint{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpc/restorePoints/vmrp/diskRestorePoints/TestDisk45ceb03433006d1baee_5c1528-43e2-4c77-9c55-a78bf5a5fc88"),
	// 		Name: to.Ptr("TestDisk45ceb03433006d1baee_5c1528-43e2-4c77-9c55-a78bf5a5fc88"),
	// 		Properties: &armcompute.DiskRestorePointProperties{
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-02T04:41:35.079872+00:00"); return t}()),
	// 			SourceResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/TestDisk45ceb03433006d1baee"),
	// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV2),
	// 			FamilyID: to.Ptr("f22e934b-c768-447f-8c44-ada4b40224ec"),
	// 			SourceUniqueID: to.Ptr("d633885d-d102-4481-901e-5b2413d1a7be"),
	// 			NetworkAccessPolicy: to.Ptr(armcompute.NetworkAccessPolicyAllowAll),
	// 			PublicNetworkAccess: to.Ptr(armcompute.PublicNetworkAccessDisabled),
	// 			CompletionPercent: to.Ptr[float32](100),
	// 			LogicalSectorSize: to.Ptr[int32](4096),
	// 			SnapshotAccessState: to.Ptr(armcompute.SnapshotAccessStateInstantAccess),
	// 		},
	// 	},
	// }
}
