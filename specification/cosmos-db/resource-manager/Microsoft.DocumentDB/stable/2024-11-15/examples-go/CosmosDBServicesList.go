package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ded6306d00ae294c24211e5069c1f56b15ba8ef5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBServicesList.json
func ExampleServiceClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceClient().NewListPager("rg1", "ddb1", nil)
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
		// page.ServiceResourceListResult = armcosmos.ServiceResourceListResult{
		// 	Value: []*armcosmos.ServiceResource{
		// 		{
		// 			Name: to.Ptr("sqlDedicatedGateway"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/sqlDedicatedGateway"),
		// 			Properties: &armcosmos.SQLDedicatedGatewayServiceResourceProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.462Z"); return t}()),
		// 				InstanceCount: to.Ptr[int32](1),
		// 				InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
		// 				ServiceType: to.Ptr(armcosmos.ServiceTypeSQLDedicatedGateway),
		// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
		// 				Locations: []*armcosmos.SQLDedicatedGatewayRegionalServiceResource{
		// 					{
		// 						Name: to.Ptr("sqlDedicatedGateway-westus2"),
		// 						Location: to.Ptr("West US 2"),
		// 						Status: to.Ptr(armcosmos.ServiceStatusRunning),
		// 						SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway-westus.sqlx.cosmos.windows-int.net/"),
		// 				}},
		// 				SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway.sqlx.cosmos.windows-int.net/"),
		// 			},
		// 	}},
		// }
	}
}
