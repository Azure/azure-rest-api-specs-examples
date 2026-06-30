package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: 2024-12-30/LongRunningBackupGet.json
func ExampleLongRunningBackupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
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
	// res = armmysqlflexibleservers.LongRunningBackupsClientGetResponse{
	// 	ServerBackupV2: armmysqlflexibleservers.ServerBackupV2{
	// 		Name: to.Ptr("daily_20210615T160516"),
	// 		Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/backupsV2"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/backupsV2/daily_20210615T160516"),
	// 		Properties: &armmysqlflexibleservers.ServerBackupPropertiesV2{
	// 			BackupType: to.Ptr(armmysqlflexibleservers.BackupTypeFULL),
	// 			CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-15T16:05:19.902522+00:00"); return t}()),
	// 			ProvisioningState: to.Ptr(armmysqlflexibleservers.ProvisioningState("Ready")),
	// 			Source: to.Ptr("Automatic"),
	// 		},
	// 	},
	// }
}
