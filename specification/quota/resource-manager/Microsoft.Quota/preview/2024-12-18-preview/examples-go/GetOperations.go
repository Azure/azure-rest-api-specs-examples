package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/55c5a0cd6da80b2700333c01e9a9c6067de9cef0/specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/GetOperations.json
func ExampleOperationClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationClient().NewListPager(nil)
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
		// page.OperationList = armquota.OperationList{
		// 	Value: []*armquota.OperationResponse{
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/quotas/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get the current Service limit or quota of the specified resource"),
		// 				Operation: to.Ptr("Get resource Quota limit"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("Resource Quota limit"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/quotas/write"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Creates the service limit or quota request for the specified resource"),
		// 				Operation: to.Ptr("Creates resource Quota limit request"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("Resource Quota limit"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/quotaRequests/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get any service limit request for the specified resource"),
		// 				Operation: to.Ptr("Get Quota limit request"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("Resource Quota limit request"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/usages/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get the usages for resource providers"),
		// 				Operation: to.Ptr("Get the usages for providers"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("usages information"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/operations/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get the Operations supported by Microsoft.Quota"),
		// 				Operation: to.Ptr("Get the Operations supported by Microsoft.Quota"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("Read Operation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/register/action"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Register the subscription with Microsoft.Quota Resource Provider"),
		// 				Operation: to.Ptr("Register the subscription with Microsoft.Quota Resource Provider"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("Subscription registration with Resource provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/groupQuotas/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get the GroupQuota"),
		// 				Operation: to.Ptr("Get GroupQuota resource"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("GroupQuota resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/groupQuotas/write"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Creates the GroupQuota resource"),
		// 				Operation: to.Ptr("Creates GroupQuota resource"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("GroupQuota Resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/groupQuotas/subscriptions/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get the GroupQuota subscriptions"),
		// 				Operation: to.Ptr("Get GroupQuota subscriptions"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("Subscriptions added to GroupQuota resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/groupQuotas/subscriptions/write"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Add Subscriptions to GroupQuota resource"),
		// 				Operation: to.Ptr("Adds subscription to GroupQuota resource"),
		// 				Provider: to.Ptr("Creates request to add subscription to GroupQuota resource"),
		// 				Resource: to.Ptr("Subscriptions added to GroupQuota resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/groupQuotas/groupQuotaLimits/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get the current GroupQuota of the specified resource"),
		// 				Operation: to.Ptr("Get GroupQuota resource Quota limit"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("GroupQuota Resource Quota limit"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/groupQuotas/groupQuotaLimits/write"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Creates the GroupQuota request for the specified resource"),
		// 				Operation: to.Ptr("Creates GroupQuota resource Quota limit request"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("GroupQuota Resource Quota limit"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/groupQuotas/groupQuotaRequests/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get the GroupQuota request status for the specific request"),
		// 				Operation: to.Ptr("Get GroupQuota request status"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("GroupQuota request"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/groupQuotas/quotaAllocations/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get the current GroupQuota to Subscription Quota allocation"),
		// 				Operation: to.Ptr("Get GroupQuota to Subscription Quota allocation"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("GroupQuota to Subscription Quota allocation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/groupQuotas/quotaAllocations/write"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Creates the GroupQuota to subscription Quota limit request for the specified resource"),
		// 				Operation: to.Ptr("Creates GroupQuota to subscription Quota limit request"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("GroupQuota to Subscription Quota allocation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/groupQuotas/quotaAllocationRequests/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get the GroupQuota to Subscription Quota allocation request status for the specific request"),
		// 				Operation: to.Ptr("Get GroupQuota to Subscription Quota allocation request status"),
		// 				Provider: to.Ptr("Microsoft.Quota"),
		// 				Resource: to.Ptr("GroupQuota to Subscription Quota allocation request"),
		// 			},
		// 	}},
		// }
	}
}
