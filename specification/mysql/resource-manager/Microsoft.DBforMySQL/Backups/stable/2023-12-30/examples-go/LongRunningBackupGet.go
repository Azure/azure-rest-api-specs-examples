package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/stable/2023-12-30/examples/LongRunningBackupGet.json
func ExampleLongRunningBackupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLongRunningBackupsClient().Get(ctx, "TestGroup", "mysqltestserver", "daily_20210615T160516", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerBackupV2 = armmysqlflexibleservers.ServerBackupV2{
	// 	Name: to.Ptr("daily_20210615T160516"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/backupsV2"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/backupsV2/daily_20210615T160516"),
	// 	Properties: &armmysqlflexibleservers.ServerBackupPropertiesV2{
	// 		BackupType: to.Ptr(armmysqlflexibleservers.BackupTypeFULL),
	// 		CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-15T16:05:19.902Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armmysqlflexibleservers.ProvisioningStateSucceeded),
	// 		Source: to.Ptr("Automatic"),
	// 	},
	// }
}
