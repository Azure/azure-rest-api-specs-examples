package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2025-08-01/TableOperationList.json
func ExampleTableClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
		// page = armstorage.TableClientListResponse{
		// 	ListTableResource: armstorage.ListTableResource{
		// 		NextLink: to.Ptr("https://sto1590endpoint/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables?api-version=2022-09-01&NextTableName=1!40!bXl0YWJsZXNoYzU0OAEwMWQ2MTI5ZTJmYjVmODFh"),
		// 		Value: []*armstorage.Table{
		// 			{
		// 				Name: to.Ptr("table6185"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/tables"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables/table6185"),
		// 				TableProperties: &armstorage.TableProperties{
		// 					TableName: to.Ptr("table6185"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("table6186"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/tables"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables/table6186"),
		// 				TableProperties: &armstorage.TableProperties{
		// 					TableName: to.Ptr("table6186"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
