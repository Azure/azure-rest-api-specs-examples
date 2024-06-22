package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de1f3772629b6f4d3ac01548a5f6d719bfb97c9e/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ListSlotBackups.json
func ExampleWebAppsClient_NewListSiteBackupsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebAppsClient().NewListSiteBackupsPager("testrg123", "tests346", nil)
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
		// page.BackupItemCollection = armappservice.BackupItemCollection{
		// 	Value: []*armappservice.BackupItem{
		// 		{
		// 			Name: to.Ptr("tests346/staging"),
		// 			Type: to.Ptr("Microsoft.Web/sites/slots"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/tests346/slot/staging"),
		// 			Properties: &armappservice.BackupItemProperties{
		// 				BlobName: to.Ptr("blob1"),
		// 				Status: to.Ptr(armappservice.BackupItemStatusInProgress),
		// 				StorageAccountURL: to.Ptr("https://blobstorage.windows.net"),
		// 			},
		// 	}},
		// }
	}
}
