package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/RecoverableServersGet.json
func ExampleRecoverableServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecoverableServersClient().Get(ctx, "testrg", "testsvc4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecoverableServerResource = armmysql.RecoverableServerResource{
	// 	Name: to.Ptr("recoverableServers"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers/recoverableServers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/testsvc4/recoverableServers"),
	// 	Properties: &armmysql.RecoverableServerProperties{
	// 		Edition: to.Ptr("GeneralPurpose"),
	// 		HardwareGeneration: to.Ptr("Gen5"),
	// 		LastAvailableBackupDateTime: to.Ptr("2020-11-20T01:06:29.78Z"),
	// 		ServiceLevelObjective: to.Ptr("GP_Gen5_2"),
	// 		VCore: to.Ptr[int32](2),
	// 		Version: to.Ptr("5.7"),
	// 	},
	// }
}
