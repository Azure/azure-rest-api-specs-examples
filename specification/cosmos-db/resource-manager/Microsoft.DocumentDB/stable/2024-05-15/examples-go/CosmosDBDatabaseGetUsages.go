package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ec7ee8842bf615c2f0354bf8b5b8725fdac9454a/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/CosmosDBDatabaseGetUsages.json
func ExampleDatabaseClient_NewListUsagesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabaseClient().NewListUsagesPager("rg1", "ddb1", "databaseRid", &armcosmos.DatabaseClientListUsagesOptions{Filter: to.Ptr("$filter=name.value eq 'Storage'")})
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
		// page.UsagesResult = armcosmos.UsagesResult{
		// 	Value: []*armcosmos.Usage{
		// 		{
		// 			Name: &armcosmos.MetricName{
		// 				LocalizedValue: to.Ptr("Storage"),
		// 				Value: to.Ptr("Storage"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](10737418240),
		// 			QuotaPeriod: to.Ptr("P1D"),
		// 			Unit: to.Ptr(armcosmos.UnitTypeBytes),
		// 	}},
		// }
	}
}
