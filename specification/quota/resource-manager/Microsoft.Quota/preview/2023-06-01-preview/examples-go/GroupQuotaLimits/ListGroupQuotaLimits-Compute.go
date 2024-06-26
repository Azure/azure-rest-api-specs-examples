package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotaLimits/ListGroupQuotaLimits-Compute.json
func ExampleGroupQuotaLimitsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGroupQuotaLimitsClient().NewListPager("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1", "Microsoft.Compute", "location eq westus", nil)
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
		// page.GroupQuotaLimitList = armquota.GroupQuotaLimitList{
		// 	Value: []*armquota.GroupQuotaLimit{
		// 		{
		// 			Name: to.Ptr("cores"),
		// 			Type: to.Ptr("Microsoft.Quota/groupQuotas/groupQuotaLimits"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/providers/Microsoft.Quota/groupQuotas/groupquota1/providers/Microsoft.Compute/locations/westus/groupQuotaLimits/cores"),
		// 			Properties: &armquota.GroupQuotaDetails{
		// 				Name: &armquota.GroupQuotaDetailsName{
		// 					LocalizedValue: to.Ptr("standard DDv4 Family vCPUs"),
		// 					Value: to.Ptr("standardddv4family"),
		// 				},
		// 				AllocatedToSubscriptions: &armquota.AllocatedQuotaToSubscriptionList{
		// 					Value: []*armquota.AllocatedToSubscription{
		// 						{
		// 							QuotaAllocated: to.Ptr[int64](20),
		// 							SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						},
		// 						{
		// 							QuotaAllocated: to.Ptr[int64](30),
		// 							SubscriptionID: to.Ptr("A000000-0000-0000-0000-000000000000"),
		// 					}},
		// 				},
		// 				AvailableLimit: to.Ptr[int64](50),
		// 				Limit: to.Ptr[int64](100),
		// 				Region: to.Ptr("westus"),
		// 				Unit: to.Ptr("count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("standardav2family"),
		// 			Type: to.Ptr("Microsoft.Quota/groupQuotas/groupQuotaLimits"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/providers/Microsoft.Quota/groupQuotas/groupquota1/providers/Microsoft.Compute/locations/westus/groupQuotaLimits/standardav2family"),
		// 			Properties: &armquota.GroupQuotaDetails{
		// 				Name: &armquota.GroupQuotaDetailsName{
		// 					LocalizedValue: to.Ptr("Standard AV2 Family vCPUs"),
		// 					Value: to.Ptr("standardav2family"),
		// 				},
		// 				AllocatedToSubscriptions: &armquota.AllocatedQuotaToSubscriptionList{
		// 					Value: []*armquota.AllocatedToSubscription{
		// 						{
		// 							QuotaAllocated: to.Ptr[int64](20),
		// 							SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					}},
		// 				},
		// 				AvailableLimit: to.Ptr[int64](80),
		// 				Limit: to.Ptr[int64](100),
		// 				Region: to.Ptr("westus"),
		// 				Unit: to.Ptr("count"),
		// 			},
		// 	}},
		// }
	}
}
