package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/BatchAccountGet.json
func ExampleAccountClient_Get_batchAccountGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountClient().Get(ctx, "default-azurebatch-japaneast", "sampleacct", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armbatch.Account{
	// 	Name: to.Ptr("sampleacct"),
	// 	Type: to.Ptr("Microsoft.Batch/batchAccounts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct"),
	// 	Location: to.Ptr("japaneast"),
	// 	Identity: &armbatch.AccountIdentity{
	// 		Type: to.Ptr(armbatch.ResourceIdentityTypeNone),
	// 	},
	// 	Properties: &armbatch.AccountProperties{
	// 		AccountEndpoint: to.Ptr("sampleacct.japaneast.batch.azure.com"),
	// 		ActiveJobAndJobScheduleQuota: to.Ptr[int32](20),
	// 		AutoStorage: &armbatch.AutoStorageProperties{
	// 			StorageAccountID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage"),
	// 			LastKeySync: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-10T23:48:38.987Z"); return t}()),
	// 		},
	// 		DedicatedCoreQuota: to.Ptr[int32](20),
	// 		LowPriorityCoreQuota: to.Ptr[int32](20),
	// 		PoolAllocationMode: to.Ptr(armbatch.PoolAllocationModeBatchService),
	// 		PoolQuota: to.Ptr[int32](20),
	// 		ProvisioningState: to.Ptr(armbatch.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armbatch.PublicNetworkAccessTypeEnabled),
	// 	},
	// }
}
