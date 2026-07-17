package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/services/sqldedicatedgateway/CosmosDBSqlDedicatedGatewayServiceCreate.json
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
			InstanceSize:         to.Ptr(armcosmos.ServiceSizeCosmosD4S),
			InstanceCount:        to.Ptr[int32](1),
			ServiceType:          to.Ptr(armcosmos.ServiceTypeSQLDedicatedGateway),
			DedicatedGatewayType: to.Ptr(armcosmos.DedicatedGatewayTypeIntegratedCache),
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
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/SqlDedicatedGateway"),
	// 		Name: to.Ptr("SqlDedicatedGateway"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 		Properties: &armcosmos.SQLDedicatedGatewayServiceResourceProperties{
	// 			Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.4622517Z"); return t}()),
	// 			InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 			InstanceCount: to.Ptr[int32](1),
	// 			ServiceType: to.Ptr(armcosmos.ServiceTypeSQLDedicatedGateway),
	// 			DedicatedGatewayType: to.Ptr(armcosmos.DedicatedGatewayTypeIntegratedCache),
	// 			SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway.sqlx.cosmos.windows-int.net/"),
	// 			Locations: []*armcosmos.SQLDedicatedGatewayRegionalServiceResource{
	// 				{
	// 					Name: to.Ptr("SqlDedicatedGateway-westus2"),
	// 					Location: to.Ptr("West US 2"),
	// 					Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 					SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway-westus.sqlx.cosmos.windows-int.net/"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
