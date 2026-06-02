package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-04-01/StorageDataShareCRUD/StorageDataShares_ListByStorageAccount.json
func ExampleDataSharesClient_NewListByStorageAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataSharesClient().NewListByStorageAccountPager("testrg", "teststorageaccount", nil)
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
		// page = armstorage.DataSharesClientListByStorageAccountResponse{
		// 	DataShareListResult: armstorage.DataShareListResult{
		// 		Value: []*armstorage.DataShare{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount/dataShares/testdatashare1"),
		// 				Name: to.Ptr("testdatashare1"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/dataShares"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armstorage.DataShareProperties{
		// 					DataShareIdentifier: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Description: to.Ptr("Dummy data share"),
		// 					DataShareURI: to.Ptr("azds://eastus:testdatashare1:00000000-0000-0000-0000-000000000000"),
		// 					AccessPolicies: []*armstorage.DataShareAccessPolicy{
		// 						{
		// 							PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							Permission: to.Ptr(armstorage.StorageDataShareAccessPolicyPermissionRead),
		// 						},
		// 					},
		// 					Assets: []*armstorage.DataShareAsset{
		// 						{
		// 							AssetPath: to.Ptr("/container/folder/foo"),
		// 							DisplayName: to.Ptr("virtualFoo"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armstorage.NativeDataSharingProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount/dataShares/testdatashare2"),
		// 				Name: to.Ptr("testdatashare2"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/dataShares"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armstorage.DataShareProperties{
		// 					DataShareIdentifier: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Description: to.Ptr("Dummy data share"),
		// 					DataShareURI: to.Ptr("azds://eastus:testdatashare2:00000000-0000-0000-0000-000000000000"),
		// 					AccessPolicies: []*armstorage.DataShareAccessPolicy{
		// 						{
		// 							PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							Permission: to.Ptr(armstorage.StorageDataShareAccessPolicyPermissionRead),
		// 						},
		// 					},
		// 					Assets: []*armstorage.DataShareAsset{
		// 						{
		// 							AssetPath: to.Ptr("/container/folder/foo"),
		// 							DisplayName: to.Ptr("virtualFoo"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armstorage.NativeDataSharingProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
