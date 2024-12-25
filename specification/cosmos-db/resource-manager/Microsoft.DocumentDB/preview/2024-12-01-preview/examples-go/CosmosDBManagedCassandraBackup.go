package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBManagedCassandraBackup.json
func ExampleCassandraClustersClient_GetBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCassandraClustersClient().GetBackup(ctx, "cassandra-prod-rg", "cassandra-prod", "1611250348", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupResource = armcosmos.BackupResource{
	// 	BackupExpiryTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-28T20:05:22.384Z"); return t}()),
	// 	BackupID: to.Ptr("2517222704776158383"),
	// 	BackupStartTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-28T19:05:22.384Z"); return t}()),
	// 	BackupState: to.Ptr(armcosmos.BackupStateSucceeded),
	// 	BackupStopTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-28T19:05:22.384Z"); return t}()),
	// }
}
