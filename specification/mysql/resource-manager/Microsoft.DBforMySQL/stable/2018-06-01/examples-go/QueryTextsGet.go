package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/QueryTextsGet.json
func ExampleQueryTextsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQueryTextsClient().Get(ctx, "testResourceGroupName", "testServerName", "1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryText = armmysql.QueryText{
	// 	Name: to.Ptr("1"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers/queryTexts"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMySQL/servers/testServerName/queryTexts/1"),
	// 	Properties: &armmysql.QueryTextProperties{
	// 		QueryID: to.Ptr("1"),
	// 		QueryText: to.Ptr("UPDATE `performance_schema`.`setup_instruments` SET `ENABLED` = ? , `TIMED` = ? WHERE NAME = ?"),
	// 	},
	// }
}
