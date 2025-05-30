package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/QuotaUsagesForFlexibleServers.json
func ExampleQuotaUsagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQuotaUsagesClient().NewListPager("westus", nil)
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
		// page.QuotaUsagesListResult = armpostgresqlflexibleservers.QuotaUsagesListResult{
		// 	Value: []*armpostgresqlflexibleservers.QuotaUsage{
		// 		{
		// 			Name: &armpostgresqlflexibleservers.NameProperty{
		// 				LocalizedValue: to.Ptr("standardBSFamily"),
		// 				Value: to.Ptr("standardBSFamily"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr(""),
		// 			Limit: to.Ptr[int64](10000),
		// 			Unit: to.Ptr("Count"),
		// 		},
		// 		{
		// 			Name: &armpostgresqlflexibleservers.NameProperty{
		// 				LocalizedValue: to.Ptr("standardDDSv4Family"),
		// 				Value: to.Ptr("standardDDSv4Family"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr(""),
		// 			Limit: to.Ptr[int64](100),
		// 			Unit: to.Ptr("Count"),
		// 	}},
		// }
	}
}
