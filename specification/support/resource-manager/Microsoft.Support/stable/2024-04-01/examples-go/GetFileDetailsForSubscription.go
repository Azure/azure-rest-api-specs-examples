package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/GetFileDetailsForSubscription.json
func ExampleFilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFilesClient().Get(ctx, "testworkspace", "test.txt", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileDetails = armsupport.FileDetails{
	// 	Name: to.Ptr("test.txt"),
	// 	Type: to.Ptr("Microsoft.Support/files"),
	// 	ID: to.Ptr("/subscriptions/132d901f-189d-4381-9214-fe68e27e05a1/providers/Microsoft.Support/fileWorkspaces/testworkspace/files/test.txt"),
	// 	Properties: &armsupport.FileDetailsProperties{
	// 		ChunkSize: to.Ptr[int32](41423),
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-24T20:18:19.000Z"); return t}()),
	// 		FileSize: to.Ptr[int32](41423),
	// 		NumberOfChunks: to.Ptr[int32](1),
	// 	},
	// }
}
