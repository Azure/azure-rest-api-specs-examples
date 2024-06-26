package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/preview/2023-10-01-preview/examples/LongRunningBackupsListByServer.json
func ExampleLongRunningBackupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongRunningBackupsClient().NewListPager("TestGroup", "mysqltestserver", nil)
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
		// page.ServerBackupV2ListResult = armmysqlflexibleservers.ServerBackupV2ListResult{
		// 	Value: []*armmysqlflexibleservers.ServerBackupV2{
		// 		{
		// 			Name: to.Ptr("daily_20210615T160516"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/backupsV2"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/backupsV2/daily_20210615T160516"),
		// 			Properties: &armmysqlflexibleservers.ServerBackupPropertiesV2{
		// 				BackupType: to.Ptr(armmysqlflexibleservers.BackupTypeFULL),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-15T16:05:19.902Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armmysqlflexibleservers.ProvisioningStateSucceeded),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 	}},
		// }
	}
}
