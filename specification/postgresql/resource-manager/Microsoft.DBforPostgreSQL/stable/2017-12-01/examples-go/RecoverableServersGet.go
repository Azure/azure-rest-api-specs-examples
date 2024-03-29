package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/RecoverableServersGet.json
func ExampleRecoverableServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecoverableServersClient().Get(ctx, "testrg", "pgtestsvc4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecoverableServerResource = armpostgresql.RecoverableServerResource{
	// 	Name: to.Ptr("recoverableServers"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/recoverableServers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/servers/pgtestsvc4/recoverableServers"),
	// 	Properties: &armpostgresql.RecoverableServerProperties{
	// 		Edition: to.Ptr("GeneralPurpose"),
	// 		HardwareGeneration: to.Ptr("Gen5"),
	// 		LastAvailableBackupDateTime: to.Ptr("2020-11-20T01:06:29.78Z"),
	// 		ServiceLevelObjective: to.Ptr("GP_Gen5_2"),
	// 		VCore: to.Ptr[int32](2),
	// 		Version: to.Ptr("9.6"),
	// 	},
	// }
}
