package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/ParameterGroups_ListConnections.json
func ExampleParameterGroupsClient_NewListConnectionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewParameterGroupsClient().NewListConnectionsPager("exampleresourcegroup", "exampleparametergroup", nil)
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
		// page = armhorizondb.ParameterGroupsClientListConnectionsResponse{
		// 	ParameterGroupConnectionPropertiesListResult: armhorizondb.ParameterGroupConnectionPropertiesListResult{
		// 		Value: []*armhorizondb.ParameterGroupConnectionProperties{
		// 			{
		// 				Name: to.Ptr("examplecluster1"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster1"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters"),
		// 			},
		// 			{
		// 				Name: to.Ptr("examplecluster2"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster2"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters"),
		// 			},
		// 			{
		// 				Name: to.Ptr("examplecluster3"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster3"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
