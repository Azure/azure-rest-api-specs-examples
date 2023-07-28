package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SubscriptionUsageGet.json
func ExampleSubscriptionUsagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubscriptionUsagesClient().Get(ctx, "WestUS", "ServerQuota", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubscriptionUsage = armsql.SubscriptionUsage{
	// 	Name: to.Ptr("ServerQuota"),
	// 	Type: to.Ptr("Microsoft.Sql/locations/usages"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/Onebox/usages/ServerQuota"),
	// 	Properties: &armsql.SubscriptionUsageProperties{
	// 		CurrentValue: to.Ptr[float64](1),
	// 		DisplayName: to.Ptr("Regional Server Quota for West US"),
	// 		Limit: to.Ptr[float64](20),
	// 		Unit: to.Ptr("Count"),
	// 	},
	// }
}
