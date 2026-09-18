package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-04-01-preview/services/sqldedicatedgateway/CosmosDBSqlDedicatedGatewayServiceCreate.json
func ExampleServiceClient_BeginCreate_sqlDedicatedGatewayServiceCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceClient().BeginCreate(ctx, "rg1", "ddb1", "SqlDedicatedGateway", armcosmos.ServiceResourceCreateUpdateParameters{
		Properties: &armcosmos.SQLDedicatedGatewayServiceResourceCreateUpdateProperties{
			DedicatedGatewayType: to.Ptr(armcosmos.DedicatedGatewayTypeIntegratedCache),
			InstanceCount:        to.Ptr[int32](1),
			InstanceSize:         to.Ptr(armcosmos.ServiceSizeCosmosD4S),
			ServiceType:          to.Ptr(armcosmos.ServiceTypeSQLDedicatedGateway),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.ServiceClientCreateResponse{
	// 	ServiceResource: armcosmos.ServiceResource{
	// 		Name: to.Ptr("SqlDedicatedGateway"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/SqlDedicatedGateway"),
	// 		Properties: &armcosmos.SQLDedicatedGatewayServiceResourceProperties{
	// 			CreationTime: to.Ptr(time.Date(2021, time.January, 25, 12, 56, 5, 462251700, time.UTC)),
	// 			DedicatedGatewayType: to.Ptr(armcosmos.DedicatedGatewayTypeIntegratedCache),
	// 			InstanceCount: to.Ptr[int32](1),
	// 			InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 			Locations: []*armcosmos.SQLDedicatedGatewayRegionalServiceResource{
	// 				{
	// 					Name: to.Ptr("SqlDedicatedGateway-westus2"),
	// 					Location: to.Ptr("West US 2"),
	// 					SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway-westus.sqlx.cosmos.windows-int.net/"),
	// 					Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 				},
	// 			},
	// 			ServiceType: to.Ptr(armcosmos.ServiceTypeSQLDedicatedGateway),
	// 			SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway.sqlx.cosmos.windows-int.net/"),
	// 			Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		},
	// 	},
	// }
}
