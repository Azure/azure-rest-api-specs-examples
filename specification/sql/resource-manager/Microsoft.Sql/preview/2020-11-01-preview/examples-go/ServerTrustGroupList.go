package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerTrustGroupList.json
func ExampleServerTrustGroupsClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerTrustGroupsClient().NewListByLocationPager("Default", "Japan East", nil)
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
		// page.ServerTrustGroupListResult = armsql.ServerTrustGroupListResult{
		// 	Value: []*armsql.ServerTrustGroup{
		// 		{
		// 			Name: to.Ptr("server-trust-group-test"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/serverTrustGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/locations/Japan East/serverTrustGroups/server-trust-group-test"),
		// 			Properties: &armsql.ServerTrustGroupProperties{
		// 				GroupMembers: []*armsql.ServerInfo{
		// 					{
		// 						ServerID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/managedInstance-1"),
		// 					},
		// 					{
		// 						ServerID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/managedInstance-2"),
		// 				}},
		// 				TrustScopes: []*armsql.ServerTrustGroupPropertiesTrustScopesItem{
		// 					to.Ptr(armsql.ServerTrustGroupPropertiesTrustScopesItemGlobalTransactions),
		// 					to.Ptr(armsql.ServerTrustGroupPropertiesTrustScopesItemServiceBroker)},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("server-trust-group-test-2"),
		// 				Type: to.Ptr("Microsoft.Sql/locations/serverTrustGroups"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/locations/Japan East/serverTrustGroups/server-trust-group-test-2"),
		// 				Properties: &armsql.ServerTrustGroupProperties{
		// 					GroupMembers: []*armsql.ServerInfo{
		// 						{
		// 							ServerID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/managedInstance-1"),
		// 						},
		// 						{
		// 							ServerID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/managedInstance-3"),
		// 					}},
		// 					TrustScopes: []*armsql.ServerTrustGroupPropertiesTrustScopesItem{
		// 						to.Ptr(armsql.ServerTrustGroupPropertiesTrustScopesItemGlobalTransactions),
		// 						to.Ptr(armsql.ServerTrustGroupPropertiesTrustScopesItemServiceBroker)},
		// 					},
		// 			}},
		// 		}
	}
}
