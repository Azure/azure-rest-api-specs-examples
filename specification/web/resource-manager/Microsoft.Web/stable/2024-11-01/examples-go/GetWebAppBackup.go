package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetWebAppBackup.json
func ExampleWebAppsClient_GetBackupStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebAppsClient().GetBackupStatus(ctx, "testrg123", "sitef6141", "12345", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupItem = armappservice.BackupItem{
	// 	Name: to.Ptr("sitef6141"),
	// 	Type: to.Ptr("Microsoft.Web/sites"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/sitef6141/backups/12345"),
	// 	Properties: &armappservice.BackupItemProperties{
	// 		Name: to.Ptr("sitef6141_2024-11-01"),
	// 		BlobName: to.Ptr("sitef6141_2024-11-01"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-03T17:33:11.641Z"); return t}()),
	// 		Databases: []*armappservice.DatabaseBackupSetting{
	// 			{
	// 				Name: to.Ptr("backenddb"),
	// 				ConnectionString: to.Ptr("DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]"),
	// 				ConnectionStringName: to.Ptr("backend"),
	// 				DatabaseType: to.Ptr(armappservice.DatabaseTypeSQLAzure),
	// 			},
	// 			{
	// 				Name: to.Ptr("statsdb"),
	// 				ConnectionString: to.Ptr("DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]"),
	// 				ConnectionStringName: to.Ptr("stats"),
	// 				DatabaseType: to.Ptr(armappservice.DatabaseTypeSQLAzure),
	// 		}},
	// 		FinishedTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-05T17:33:11.641Z"); return t}()),
	// 		BackupID: to.Ptr[int32](12345),
	// 		LastRestoreTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-04T17:33:11.641Z"); return t}()),
	// 		Log: to.Ptr("Succeeded"),
	// 		Scheduled: to.Ptr(true),
	// 		SizeInBytes: to.Ptr[int64](56091883),
	// 		Status: to.Ptr(armappservice.BackupItemStatusInProgress),
	// 		StorageAccountURL: to.Ptr("DefaultEndpointsProtocol=https;AccountName=storagesample;AccountKey=<account-key>"),
	// 		WebsiteSizeInBytes: to.Ptr[int64](56091883),
	// 	},
	// }
}
