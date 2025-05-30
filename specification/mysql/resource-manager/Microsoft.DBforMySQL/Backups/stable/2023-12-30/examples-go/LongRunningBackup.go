package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/stable/2023-12-30/examples/LongRunningBackup.json
func ExampleLongRunningBackupClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLongRunningBackupClient().BeginCreate(ctx, "TestGroup", "mysqltestserver", "testback", &armmysqlflexibleservers.LongRunningBackupClientBeginCreateOptions{Parameters: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerBackupV2 = armmysqlflexibleservers.ServerBackupV2{
	// 	Name: to.Ptr("customer_20220507t073755_bb392c3b-17c6-4d3f-9742-8479ca87b3ac_mybackup"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/backups"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/backupsV2/customer_20220507t073755_bb392c3b-17c6-4d3f-9742-8479ca87b3ac_mybackup"),
	// 	Properties: &armmysqlflexibleservers.ServerBackupPropertiesV2{
	// 		BackupType: to.Ptr(armmysqlflexibleservers.BackupTypeFULL),
	// 		CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-07T07:38:01.149Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armmysqlflexibleservers.ProvisioningStateSucceeded),
	// 		Source: to.Ptr("Automatic"),
	// 	},
	// }
}
