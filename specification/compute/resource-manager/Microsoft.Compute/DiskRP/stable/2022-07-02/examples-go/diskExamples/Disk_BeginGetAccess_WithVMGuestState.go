package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/19f98c9f526f8db961f172276dd6d6882a86ed86/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/diskExamples/Disk_BeginGetAccess_WithVMGuestState.json
func ExampleDisksClient_BeginGrantAccess_getSasOnManagedDiskAndVmGuestState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDisksClient().BeginGrantAccess(ctx, "myResourceGroup", "myDisk", armcompute.GrantAccessData{
		Access:                   to.Ptr(armcompute.AccessLevelRead),
		DurationInSeconds:        to.Ptr[int32](300),
		GetSecureVMGuestStateSAS: to.Ptr(true),
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
	// 	SecurityDataAccessSAS: to.Ptr("https://md-gpvmcxzlzxgd.partition.blob.storage.azure.net/xx3cqcx53f0v/b9bf5824-6122-49e0-ba22-042f76ccd8a1_vmgs?sv=2014-02-14&sr=b&sk=key1&sig=XXX&st=2021-05-24T18:02:34Z&se=2021-05-24T18:19:14Z&sp=r"),
	// }
}
