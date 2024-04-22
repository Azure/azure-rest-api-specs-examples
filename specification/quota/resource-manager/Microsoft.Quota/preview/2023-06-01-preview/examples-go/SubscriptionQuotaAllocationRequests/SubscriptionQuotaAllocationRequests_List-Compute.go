package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/SubscriptionQuotaAllocationRequests/SubscriptionQuotaAllocationRequests_List-Compute.json
func ExampleGroupQuotaSubscriptionAllocationRequestClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGroupQuotaSubscriptionAllocationRequestClient().NewListPager("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1", "Microsoft.Compute", "location eq westus", nil)
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
		// page.AllocationRequestStatusList = armquota.AllocationRequestStatusList{
		// 	Value: []*armquota.AllocationRequestStatus{
		// 		{
		// 			Name: to.Ptr("AE000000-0000-0000-0000-00000000000A"),
		// 			Type: to.Ptr("Microsoft.Quota/groupQuotas/quotaAllocationRequests"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Quota/groupQuotas/groupquota1/resourceProviders/Microsoft.Compute/quotaAllocationRequests/AE000000-0000-0000-0000-00000000000A"),
		// 			Properties: &armquota.AllocationRequestStatusProperties{
		// 				FaultCode: to.Ptr("ContactSupport"),
		// 				ProvisioningState: to.Ptr(armquota.RequestStateSucceeded),
		// 				RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-20T06:18:59.913Z"); return t}()),
		// 				RequestedResource: &armquota.AllocationRequestBase{
		// 					Properties: &armquota.AllocationRequestBaseProperties{
		// 						Name: &armquota.AllocationRequestBasePropertiesName{
		// 							LocalizedValue: to.Ptr("standard Av2 Family vCPUs"),
		// 							Value: to.Ptr("standardav2family"),
		// 						},
		// 						Limit: to.Ptr[int64](75),
		// 						Region: to.Ptr("westus"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
