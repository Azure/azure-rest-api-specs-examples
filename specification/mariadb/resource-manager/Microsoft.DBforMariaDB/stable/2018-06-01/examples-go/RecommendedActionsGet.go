package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/RecommendedActionsGet.json
func ExampleRecommendedActionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmariadb.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecommendedActionsClient().Get(ctx, "testResourceGroupName", "testServerName", "Index", "Index-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecommendationAction = armmariadb.RecommendationAction{
	// 	Name: to.Ptr("Index-1"),
	// 	Type: to.Ptr("Microsoft.DBforMariaDB/servers/advisors/recommendedActions"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.Sql/servers/testServerName/advisors/Index/recommendedActions/Index-1"),
	// 	Properties: &armmariadb.RecommendationActionProperties{
	// 		ActionID: to.Ptr[int32](1),
	// 		AdvisorName: to.Ptr("Index"),
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T23:43:24.000Z"); return t}()),
	// 		ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-08T23:43:24.000Z"); return t}()),
	// 		Reason: to.Ptr("Column `movies_genres`.`movie_id` appear in Join On clause(s)."),
	// 		RecommendationType: to.Ptr("Add"),
	// 		SessionID: to.Ptr("c63c2114-e2a4-4c7a-98c1-85577d1a5d50"),
	// 		Details: map[string]*string{
	// 			"engine": to.Ptr("InnoDB"),
	// 			"indexColumns": to.Ptr("`movies_genres`.`movie_id`"),
	// 			"indexName": to.Ptr("idx_movie_id"),
	// 			"indexType": to.Ptr("BTREE"),
	// 			"parentTableName": to.Ptr("movies_genres"),
	// 			"queryIds": to.Ptr("779"),
	// 			"schemaName": to.Ptr("movies"),
	// 			"script": to.Ptr("alter table `movies`.`movies_genres` add index `idx_movie_id` (`movie_id`)"),
	// 			"tableName": to.Ptr("movies_genres"),
	// 		},
	// 	},
	// }
}
