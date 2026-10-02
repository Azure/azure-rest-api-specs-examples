package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-07-01-preview/DBAgentList.json
func ExampleDbAgentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDbAgentsClient().NewListPager("exampleresourcegroup", "exampleserver", nil)
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
		// page = armpostgresqlflexibleservers.DbAgentsClientListResponse{
		// 	DbAgentListResult: armpostgresqlflexibleservers.DbAgentListResult{
		// 		Value: []*armpostgresqlflexibleservers.DbAgent{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/dbAgents/Default"),
		// 				Name: to.Ptr("Default"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/dbAgents"),
		// 				Properties: &armpostgresqlflexibleservers.DbAgentProperties{
		// 					State: to.Ptr(armpostgresqlflexibleservers.DbAgentStateEnabled),
		// 					ProvisioningState: to.Ptr(armpostgresqlflexibleservers.DbAgentProvisioningStateSucceeded),
		// 					LastModifiedTime: to.Ptr(time.Date(2026, time.August, 28, 15, 30, 0, 0, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
