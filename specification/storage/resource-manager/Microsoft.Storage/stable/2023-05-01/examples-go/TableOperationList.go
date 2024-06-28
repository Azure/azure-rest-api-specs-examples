package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/TableOperationList.json
func ExampleTableClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTableClient().NewListPager("res9290", "sto328", nil)
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
		// page.ListTableResource = armstorage.ListTableResource{
		// 	Value: []*armstorage.Table{
		// 		{
		// 			Name: to.Ptr("table6185"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/tables"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables/table6185"),
		// 			TableProperties: &armstorage.TableProperties{
		// 				TableName: to.Ptr("table6185"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("table6186"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/tables"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables/table6186"),
		// 			TableProperties: &armstorage.TableProperties{
		// 				TableName: to.Ptr("table6186"),
		// 			},
		// 	}},
		// }
	}
}
