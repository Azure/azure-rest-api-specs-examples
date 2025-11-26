package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListWebAppBackups.json
func ExampleWebAppsClient_NewListBackupsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebAppsClient().NewListBackupsPager("testrg123", "sitef6141", nil)
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
		// 			Name: to.Ptr("sitef6141"),
		// 			Type: to.Ptr("Microsoft.Web/sites"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/sitef6141/backups/12345"),
		// 			Properties: &armappservice.BackupItemProperties{
		// 				Name: to.Ptr("sitef6141_2025-03-01"),
		// 				BlobName: to.Ptr("sitef6141_2025-03-01"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-03T17:33:11.641Z"); return t}()),
		// 				Databases: []*armappservice.DatabaseBackupSetting{
		// 					{
		// 						Name: to.Ptr("backenddb"),
		// 						ConnectionString: to.Ptr("DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]"),
		// 						ConnectionStringName: to.Ptr("backend"),
		// 						DatabaseType: to.Ptr(armappservice.DatabaseTypeSQLAzure),
		// 					},
		// 					{
		// 						Name: to.Ptr("statsdb"),
		// 						ConnectionString: to.Ptr("DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]"),
		// 						ConnectionStringName: to.Ptr("stats"),
		// 						DatabaseType: to.Ptr(armappservice.DatabaseTypeSQLAzure),
		// 				}},
		// 				FinishedTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-05T17:33:11.641Z"); return t}()),
		// 				BackupID: to.Ptr[int32](12345),
		// 				LastRestoreTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-04T17:33:11.641Z"); return t}()),
		// 				Log: to.Ptr("Succeeded"),
		// 				Scheduled: to.Ptr(true),
		// 				SizeInBytes: to.Ptr[int64](56091883),
		// 				Status: to.Ptr(armappservice.BackupItemStatusInProgress),
		// 				StorageAccountURL: to.Ptr("DefaultEndpointsProtocol=https;AccountName=storagesample;AccountKey=<account-key>"),
		// 				WebsiteSizeInBytes: to.Ptr[int64](56091883),
		// 			},
		// 	}},
		// }
	}
}
