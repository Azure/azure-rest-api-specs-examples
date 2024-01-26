package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6f8faf5da91b5b9af5f3512fe609e22e99383d41/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-11-15-preview/examples/CosmosDBSqlDedicatedGatewayServiceGet.json
func ExampleServiceClient_Get_sqlDedicatedGatewayServiceGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceClient().Get(ctx, "rg1", "ddb1", "SqlDedicatedGateway", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResource = armcosmos.ServiceResource{
	// 	Name: to.Ptr("SqlDedicatedGateway"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/SqlDedicatedGateway"),
	// 	Properties: &armcosmos.SQLDedicatedGatewayServiceResourceProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.462Z"); return t}()),
	// 		InstanceCount: to.Ptr[int32](1),
	// 		InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 		ServiceType: to.Ptr(armcosmos.ServiceTypeSQLDedicatedGateway),
	// 		Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		Locations: []*armcosmos.SQLDedicatedGatewayRegionalServiceResource{
	// 			{
	// 				Name: to.Ptr("SqlDedicatedGateway-westus2"),
	// 				Location: to.Ptr("West US 2"),
	// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 				SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway-westus.sqlx.cosmos.windows-int.net/"),
	// 		}},
	// 		SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway.sqlx.cosmos.windows-int.net/"),
	// 	},
	// }
}
