package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListFilesUnderFileWorkspace.json
func ExampleFilesNoSubscriptionClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFilesNoSubscriptionClient().NewListPager("testworkspace", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.FilesListResult = armsupport.FilesListResult{
		// 	Value: []*armsupport.FileDetails{
		// 		{
		// 			Name: to.Ptr("test1.txt"),
		// 			Type: to.Ptr("Microsoft.Support/files"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/fileWorkspaces/testworkspace/files/test1.txt"),
		// 			Properties: &armsupport.FileDetailsProperties{
		// 				ChunkSize: to.Ptr[float32](41423),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-24T20:18:19.000Z"); return t}()),
		// 				FileSize: to.Ptr[float32](41423),
		// 				NumberOfChunks: to.Ptr[float32](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test1.txt"),
		// 			Type: to.Ptr("Microsoft.Support/files"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/fileWorkspaces/testworkspace/files/test1.txt"),
		// 			Properties: &armsupport.FileDetailsProperties{
		// 				ChunkSize: to.Ptr[float32](41423),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-24T20:18:19.000Z"); return t}()),
		// 				FileSize: to.Ptr[float32](41423),
		// 				NumberOfChunks: to.Ptr[float32](1),
		// 			},
		// 	}},
		// }
	}
}
