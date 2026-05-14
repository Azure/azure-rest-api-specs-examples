package armfileshares_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fileshares/armfileshares"
)

// Generated from example definition: 2026-06-01/FileShareSnapshot_List_MinimumSet_Gen.json
func ExampleFileShareSnapshotsClient_NewListByFileSharePager_fileShareSnapshotListMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfileshares.NewClientFactory("0681745E-3F9F-4966-80E6-69624A3B29F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileShareSnapshotsClient().NewListByFileSharePager("rgfileshares", "testfileshare", nil)
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
		// page = armfileshares.FileShareSnapshotsClientListByFileShareResponse{
		// 	FileShareSnapshotListResult: armfileshares.FileShareSnapshotListResult{
		// 		Value: []*armfileshares.FileShareSnapshot{
		// 		},
		// 	},
		// }
	}
}
