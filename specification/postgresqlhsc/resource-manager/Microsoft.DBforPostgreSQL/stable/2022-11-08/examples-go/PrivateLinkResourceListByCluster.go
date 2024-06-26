package armcosmosforpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosforpostgresql/armcosmosforpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/PrivateLinkResourceListByCluster.json
func ExamplePrivateLinkResourcesClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmosforpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByClusterPager("TestResourceGroup", "testcluster", nil)
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
		// page.PrivateLinkResourceListResult = armcosmosforpostgresql.PrivateLinkResourceListResult{
		// 	Value: []*armcosmosforpostgresql.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("coordinator"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/privateLinkResources/coordinator"),
		// 			SystemData: &armcosmosforpostgresql.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 			},
		// 			Properties: &armcosmosforpostgresql.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("coordinator"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("coordinator")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.testcluster.postgres.database.azure.com")},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("worker-0"),
		// 					Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateLinkResources"),
		// 					ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/privateLinkResources/worker-0"),
		// 					SystemData: &armcosmosforpostgresql.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-01T17:18:19.123Z"); return t}()),
		// 						CreatedBy: to.Ptr("user1"),
		// 						CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-02T17:18:19.123Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("user2"),
		// 						LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 					},
		// 					Properties: &armcosmosforpostgresql.PrivateLinkResourceProperties{
		// 						GroupID: to.Ptr("worker-0"),
		// 						RequiredMembers: []*string{
		// 							to.Ptr("worker-0")},
		// 							RequiredZoneNames: []*string{
		// 								to.Ptr("privatelink.testcluster.postgres.database.azure.com")},
		// 							},
		// 					}},
		// 				}
	}
}
