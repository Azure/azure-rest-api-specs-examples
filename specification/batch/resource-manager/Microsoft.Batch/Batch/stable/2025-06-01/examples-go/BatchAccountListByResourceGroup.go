package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v4"
)

// Generated from example definition: 2025-06-01/BatchAccountListByResourceGroup.json
func ExampleAccountClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountClient().NewListByResourceGroupPager("default-azurebatch-japaneast", nil)
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
		// page = armbatch.AccountClientListByResourceGroupResponse{
		// 	AccountListResult: armbatch.AccountListResult{
		// 		Value: []*armbatch.Account{
		// 			{
		// 				Name: to.Ptr("sampleacct"),
		// 				Type: to.Ptr("Microsoft.Batch/batchAccounts"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct"),
		// 				Identity: &armbatch.AccountIdentity{
		// 					Type: to.Ptr(armbatch.ResourceIdentityTypeNone),
		// 				},
		// 				Location: to.Ptr("japaneast"),
		// 				Properties: &armbatch.AccountProperties{
		// 					AccountEndpoint: to.Ptr("sampleacct.japaneast.batch.azure.com"),
		// 					ActiveJobAndJobScheduleQuota: to.Ptr[int32](20),
		// 					AutoStorage: &armbatch.AutoStorageProperties{
		// 						LastKeySync: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-10T23:48:38.9878479Z"); return t}()),
		// 						StorageAccountID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage"),
		// 					},
		// 					DedicatedCoreQuota: to.Ptr[int32](20),
		// 					LowPriorityCoreQuota: to.Ptr[int32](20),
		// 					PoolAllocationMode: to.Ptr(armbatch.PoolAllocationModeBatchService),
		// 					PoolQuota: to.Ptr[int32](20),
		// 					ProvisioningState: to.Ptr(armbatch.ProvisioningStateSucceeded),
		// 					PublicNetworkAccess: to.Ptr(armbatch.PublicNetworkAccessTypeEnabled),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
