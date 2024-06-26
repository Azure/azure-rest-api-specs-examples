package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-11-15-preview/examples/CosmosDBManagedCassandraBackupsList.json
func ExampleCassandraClustersClient_NewListBackupsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCassandraClustersClient().NewListBackupsPager("cassandra-prod-rg", "cassandra-prod", nil)
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
		// page.ListBackups = armcosmos.ListBackups{
		// 	Value: []*armcosmos.BackupResource{
		// 		{
		// 			Name: to.Ptr("1611250348"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters/backups"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters/cassandra-prod/backups/1611250348"),
		// 			Properties: &armcosmos.BackupResourceProperties{
		// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-21T17:32:28Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("1611214220"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters/backups"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters/cassandra-prod/backups/1611214220"),
		// 			Properties: &armcosmos.BackupResourceProperties{
		// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-21T07:30:20Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
