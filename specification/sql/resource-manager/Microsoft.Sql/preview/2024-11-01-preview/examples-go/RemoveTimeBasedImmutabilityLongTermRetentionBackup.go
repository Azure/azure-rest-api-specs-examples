package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8358c7473dfe057d84a6b6a921225063c040b31a/specification/sql/resource-manager/Microsoft.Sql/preview/2024-11-01-preview/examples/RemoveTimeBasedImmutabilityLongTermRetentionBackup.json
func ExampleLongTermRetentionBackupsClient_BeginRemoveTimeBasedImmutability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLongTermRetentionBackupsClient().BeginRemoveTimeBasedImmutability(ctx, "japaneast", "testserver", "testDatabase", "55555555-6666-7777-8888-999999999999;131637960820000000;Hot", nil)
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
	// res.LongTermRetentionBackup = armsql.LongTermRetentionBackup{
	// 	Name: to.Ptr("01234567-890a-bcde-f012-34567890abcd"),
	// 	Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionBackupOperationResults"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionBackupOperationResults/01234567-890a-bcde-f012-34567890abcd"),
	// 	Properties: &armsql.LongTermRetentionBackupProperties{
	// 	},
	// }
}
