package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/011ecc5633300a5eefe43dde748f269d39e96458/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBRestorableSqlResourceList.json
func ExampleRestorableSQLResourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorableSQLResourcesClient().NewListPager("WestUS", "d9b26648-2f53-4541-b3d8-3044f4f9810d", &armcosmos.RestorableSQLResourcesClientListOptions{RestoreLocation: to.Ptr("WestUS"),
		RestoreTimestampInUTC: to.Ptr("06/01/2022 4:56"),
	})
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
		// page.RestorableSQLResourcesListResult = armcosmos.RestorableSQLResourcesListResult{
		// 	Value: []*armcosmos.RestorableSQLResourcesGetResult{
		// 		{
		// 			Name: to.Ptr("Database1"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/locations/restorableDatabaseAccounts/restorablesqlresources"),
		// 			CollectionNames: []*string{
		// 				to.Ptr("Container1")},
		// 				DatabaseName: to.Ptr("Database1"),
		// 				ID: to.Ptr("/subscriptions/2296c272-5d55-40d9-bc05-4d56dc2d7588/providers/Microsoft.DocumentDB/locations/westus/restorableDatabaseAccounts/d9b26648-2f53-4541-b3d8-3044f4f9810d/restorablesqlresources/Database1"),
		// 			},
		// 			{
		// 				Name: to.Ptr("Database2"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/locations/restorableDatabaseAccounts/restorablesqlresources"),
		// 				CollectionNames: []*string{
		// 					to.Ptr("Container1"),
		// 					to.Ptr("Container2")},
		// 					DatabaseName: to.Ptr("Database2"),
		// 					ID: to.Ptr("/subscriptions/2296c272-5d55-40d9-bc05-4d56dc2d7588/providers/Microsoft.DocumentDB/locations/westus/restorableDatabaseAccounts/d9b26648-2f53-4541-b3d8-3044f4f9810d/restorablesqlresources/Database2"),
		// 				},
		// 				{
		// 					Name: to.Ptr("Database3"),
		// 					Type: to.Ptr("Microsoft.DocumentDB/locations/restorableDatabaseAccounts/restorablesqlresources"),
		// 					CollectionNames: []*string{
		// 					},
		// 					DatabaseName: to.Ptr("Database3"),
		// 					ID: to.Ptr("/subscriptions/2296c272-5d55-40d9-bc05-4d56dc2d7588/providers/Microsoft.DocumentDB/locations/westus/restorableDatabaseAccounts/d9b26648-2f53-4541-b3d8-3044f4f9810d/restorablesqlresources/Database3"),
		// 			}},
		// 		}
	}
}
