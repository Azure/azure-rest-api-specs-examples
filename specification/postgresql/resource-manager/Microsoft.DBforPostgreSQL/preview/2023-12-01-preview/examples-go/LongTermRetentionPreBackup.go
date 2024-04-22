package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-12-01-preview/examples/LongTermRetentionPreBackup.json
func ExampleFlexibleServerClient_TriggerLtrPreBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFlexibleServerClient().TriggerLtrPreBackup(ctx, "rgLongTermRetention", "pgsqlltrtestserver", armpostgresqlflexibleservers.LtrPreBackupRequest{
		BackupSettings: &armpostgresqlflexibleservers.BackupSettings{
			BackupName: to.Ptr("backup1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LtrPreBackupResponse = armpostgresqlflexibleservers.LtrPreBackupResponse{
	// 	Properties: &armpostgresqlflexibleservers.LtrPreBackupResponseProperties{
	// 		NumberOfContainers: to.Ptr[int32](1),
	// 	},
	// }
}
