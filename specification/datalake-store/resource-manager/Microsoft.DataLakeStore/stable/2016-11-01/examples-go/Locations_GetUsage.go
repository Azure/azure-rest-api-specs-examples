package armdatalakestore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Locations_GetUsage.json
func ExampleLocationsClient_NewGetUsagePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakestore.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationsClient().NewGetUsagePager("WestUS", nil)
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
		// page.UsageListResult = armdatalakestore.UsageListResult{
		// 	Value: []*armdatalakestore.Usage{
		// 		{
		// 			Name: &armdatalakestore.UsageName{
		// 				LocalizedValue: to.Ptr("Data Lake Store Accounts"),
		// 				Value: to.Ptr("DataLakeStoreAccounts"),
		// 			},
		// 			CurrentValue: to.Ptr[int32](69),
		// 			ID: to.Ptr("/subscriptions/0f936178-0d48-4777-a7ab-47860b604941/providers/Microsoft.DataLakeStore/locations/WestUS/usages/DataLakeStoreAccounts"),
		// 			Limit: to.Ptr[int32](-1),
		// 			Unit: to.Ptr(armdatalakestore.UsageUnitCount),
		// 	}},
		// }
	}
}
