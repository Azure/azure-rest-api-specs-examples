package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/ApplicationPackageCreate.json
func ExampleApplicationPackageClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationPackageClient().Create(ctx, "default-azurebatch-japaneast", "sampleacct", "app1", "1", &armbatch.ApplicationPackageClientCreateOptions{Parameters: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationPackage = armbatch.ApplicationPackage{
	// 	Name: to.Ptr("1"),
	// 	Type: to.Ptr("Microsoft.Batch/batchAccounts/applications/versions"),
	// 	Etag: to.Ptr("W/\"0x8D64FEC83A3B436\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/applications/app1/versions/1"),
	// 	Properties: &armbatch.ApplicationPackageProperties{
	// 		State: to.Ptr(armbatch.PackageStatePending),
	// 		StorageURL: to.Ptr("http://mystorage1.blob.core.windows.net/myapp?mysas"),
	// 		StorageURLExpiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-27T18:48:09.933Z"); return t}()),
	// 	},
	// }
}
