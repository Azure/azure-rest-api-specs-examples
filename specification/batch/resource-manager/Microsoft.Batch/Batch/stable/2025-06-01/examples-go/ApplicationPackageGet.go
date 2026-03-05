package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v4"
)

// Generated from example definition: 2025-06-01/ApplicationPackageGet.json
func ExampleApplicationPackageClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationPackageClient().Get(ctx, "default-azurebatch-japaneast", "sampleacct", "app1", "1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbatch.ApplicationPackageClientGetResponse{
	// 	ApplicationPackage: &armbatch.ApplicationPackage{
	// 		Name: to.Ptr("1"),
	// 		Type: to.Ptr("Microsoft.Batch/batchAccounts/applications/versions"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/applications/app1/versions/1"),
	// 		Properties: &armbatch.ApplicationPackageProperties{
	// 			Format: to.Ptr("zip"),
	// 			LastActivationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-27T18:48:09.9330991Z"); return t}()),
	// 			State: to.Ptr(armbatch.PackageStateActive),
	// 		},
	// 	},
	// }
}
