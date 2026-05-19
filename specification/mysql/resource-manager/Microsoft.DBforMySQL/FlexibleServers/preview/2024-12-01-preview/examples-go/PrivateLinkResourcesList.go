package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: 2024-12-01-preview/PrivateLinkResourcesList.json
func ExamplePrivateLinkResourcesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByServerPager("Default", "test-svr", nil)
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
		// page = armmysqlflexibleservers.PrivateLinkResourcesClientListByServerResponse{
		// 	PrivateLinkResourceListResult: armmysqlflexibleservers.PrivateLinkResourceListResult{
		// 		Value: []*armmysqlflexibleservers.PrivateLinkResource{
		// 			{
		// 				Name: to.Ptr("plr"),
		// 				Type: to.Ptr("Microsoft.DBforMySQL/servers/privateLinkResources"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.DBforMySQL/servers/test-svr/privateLinkResources/plr"),
		// 				Properties: &armmysqlflexibleservers.PrivateLinkResourceProperties{
		// 					GroupID: to.Ptr("mysqlServer"),
		// 					RequiredMembers: []*string{
		// 						to.Ptr("mysqlServer"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
