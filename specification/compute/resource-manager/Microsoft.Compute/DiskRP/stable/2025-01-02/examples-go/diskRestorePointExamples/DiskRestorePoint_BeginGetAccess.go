package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskRestorePointExamples/DiskRestorePoint_BeginGetAccess.json
func ExampleDiskRestorePointClient_BeginGrantAccess() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskRestorePointClient().BeginGrantAccess(ctx, "myResourceGroup", "rpc", "vmrp", "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745", armcompute.GrantAccessData{
		Access:            to.Ptr(armcompute.AccessLevelRead),
		DurationInSeconds: to.Ptr[int32](300),
		FileFormat:        to.Ptr(armcompute.FileFormatVHDX),
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
	// res.AccessURI = armcompute.AccessURI{
	// 	AccessSAS: to.Ptr("https://md-gpvmcxzlzxgd.partition.blob.storage.azure.net/xx3cqcx53f0v/abcd?sv=2014-02-14&sr=b&sk=key1&sig=XXX&st=2021-05-24T18:02:34Z&se=2021-05-24T18:19:14Z&sp=r"),
	// }
}
