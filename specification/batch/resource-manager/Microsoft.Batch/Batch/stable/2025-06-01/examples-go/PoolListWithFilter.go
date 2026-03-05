package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v4"
)

// Generated from example definition: 2025-06-01/PoolListWithFilter.json
func ExamplePoolClient_NewListByBatchAccountPager_listPoolWithFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPoolClient().NewListByBatchAccountPager("default-azurebatch-japaneast", "sampleacct", &armbatch.PoolClientListByBatchAccountOptions{
		Filter: to.Ptr("startswith(name, 'po') or (properties/allocationState eq 'Steady' and properties/provisioningStateTransitionTime lt datetime'2017-02-02')"),
		Select: to.Ptr("properties/allocationState,properties/provisioningStateTransitionTime,properties/currentDedicatedNodes,properties/currentLowPriorityNodes")})
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
		// page = armbatch.PoolClientListByBatchAccountResponse{
		// 	ListPoolsResult: armbatch.ListPoolsResult{
		// 		Value: []*armbatch.Pool{
		// 			{
		// 				Name: to.Ptr("testpool"),
		// 				Type: to.Ptr("Microsoft.Batch/batchAccounts/pools"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool"),
		// 				Properties: &armbatch.PoolProperties{
		// 					AllocationState: to.Ptr(armbatch.AllocationStateSteady),
		// 					CurrentDedicatedNodes: to.Ptr[int32](0),
		// 					CurrentLowPriorityNodes: to.Ptr[int32](2),
		// 					ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.9407275Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("pooltest"),
		// 				Type: to.Ptr("Microsoft.Batch/batchAccounts/pools"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/pooltest"),
		// 				Properties: &armbatch.PoolProperties{
		// 					AllocationState: to.Ptr(armbatch.AllocationStateResizing),
		// 					CurrentDedicatedNodes: to.Ptr[int32](4),
		// 					CurrentLowPriorityNodes: to.Ptr[int32](0),
		// 					ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-26T10:22:55.9407275Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
