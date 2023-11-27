package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Disks_Get.json
func ExampleDisksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDisksClient().Get(ctx, "resourceGroupName", "{labName}", "@me", "{diskName}", &armdevtestlabs.DisksClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Disk = armdevtestlabs.Disk{
	// 	Name: to.Ptr("{diskName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/users/disks"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/users/{userId}/disks/{diskName}"),
	// 	Properties: &armdevtestlabs.DiskProperties{
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T20:27:32.144Z"); return t}()),
	// 		DiskSizeGiB: to.Ptr[int32](1023),
	// 		DiskType: to.Ptr(armdevtestlabs.StorageTypeStandard),
	// 		DiskURI: to.Ptr(""),
	// 		HostCaching: to.Ptr("None"),
	// 		LeasedByLabVMID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualmachines/vmName"),
	// 		ManagedDiskID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.compute/disks/{diskName}"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		UniqueIdentifier: to.Ptr("9bf098d1-1b64-41a5-aa05-286767074a0b"),
	// 	},
	// }
}
