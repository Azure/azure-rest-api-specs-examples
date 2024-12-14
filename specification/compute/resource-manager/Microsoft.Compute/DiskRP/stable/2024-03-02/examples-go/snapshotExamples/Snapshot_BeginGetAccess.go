package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/snapshotExamples/Snapshot_BeginGetAccess.json
func ExampleSnapshotsClient_BeginGrantAccess() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSnapshotsClient().BeginGrantAccess(ctx, "myResourceGroup", "mySnapshot", armcompute.GrantAccessData{
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
