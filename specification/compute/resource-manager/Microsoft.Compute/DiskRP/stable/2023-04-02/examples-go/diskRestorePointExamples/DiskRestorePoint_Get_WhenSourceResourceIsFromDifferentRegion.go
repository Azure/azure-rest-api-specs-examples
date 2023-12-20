package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/60679ee3db06e93eb73faa0587fed93ed843d6dc/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/diskRestorePointExamples/DiskRestorePoint_Get_WhenSourceResourceIsFromDifferentRegion.json
func ExampleDiskRestorePointClient_Get_getAnIncrementalDiskRestorePointWhenSourceResourceIsFromADifferentRegion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiskRestorePointClient().Get(ctx, "myResourceGroup", "rpc", "vmrp", "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiskRestorePoint = armcompute.DiskRestorePoint{
	// 	Name: to.Ptr("TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpc/restorePoints/vmrp/diskRestorePoints/TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745"),
	// 	Properties: &armcompute.DiskRestorePointProperties{
	// 		CompletionPercent: to.Ptr[float32](100),
	// 		FamilyID: to.Ptr("996bf3ce-b6ff-4e86-9db6-dc27ea06cea5"),
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 		NetworkAccessPolicy: to.Ptr(armcompute.NetworkAccessPolicyAllowAll),
	// 		OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 		PublicNetworkAccess: to.Ptr(armcompute.PublicNetworkAccessDisabled),
	// 		ReplicationState: to.Ptr("Succeeded"),
	// 		SourceResourceID: to.Ptr("/subscriptions/d2260d06-e00d-422f-8b63-93df551a59ae/resourceGroups/rg0680fb0c-89f1-41b4-96c0-35733a181558/providers/Microsoft.Compute/disks/TestDisk45ceb03433006d1baee0"),
	// 		SourceResourceLocation: to.Ptr("eastus2"),
	// 		SourceUniqueID: to.Ptr("48e058b1-7eea-4968-b532-10a8a1130c13"),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-16T04:41:35.079Z"); return t}()),
	// 	},
	// }
}
