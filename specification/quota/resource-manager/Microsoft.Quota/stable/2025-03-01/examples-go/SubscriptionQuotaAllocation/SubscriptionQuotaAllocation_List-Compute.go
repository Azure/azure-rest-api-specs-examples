package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8691e5081766c7ad602a9e55de841d07bed5196a/specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/SubscriptionQuotaAllocation/SubscriptionQuotaAllocation_List-Compute.json
func ExampleGroupQuotaSubscriptionAllocationClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGroupQuotaSubscriptionAllocationClient("<subscription-id>").List(ctx, "E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1", "Microsoft.Compute", "westus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubscriptionQuotaAllocationsList = armquota.SubscriptionQuotaAllocationsList{
	// 	Name: to.Ptr("westus"),
	// 	Type: to.Ptr("Microsoft.Quota/groupQuotas/quotaAllocations"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Quota/groupQuotas/groupquota1/resourceProviders/Microsoft.Compute/quotaAllocations/westus"),
	// 	Properties: &armquota.SubscriptionQuotaAllocationsListProperties{
	// 		ProvisioningState: to.Ptr(armquota.RequestStateSucceeded),
	// 		Value: []*armquota.SubscriptionQuotaAllocations{
	// 			{
	// 				Properties: &armquota.SubscriptionQuotaAllocationsProperties{
	// 					Name: &armquota.SubscriptionQuotaDetailsName{
	// 						LocalizedValue: to.Ptr("standard DDv4 Family vCPUs"),
	// 						Value: to.Ptr("standardddv4family"),
	// 					},
	// 					Limit: to.Ptr[int64](25),
	// 					ResourceName: to.Ptr("standardddv4family"),
	// 					ShareableQuota: to.Ptr[int64](15),
	// 				},
	// 			},
	// 			{
	// 				Properties: &armquota.SubscriptionQuotaAllocationsProperties{
	// 					Name: &armquota.SubscriptionQuotaDetailsName{
	// 						LocalizedValue: to.Ptr("standard Av2 Family vCPUs"),
	// 						Value: to.Ptr("standardav2family"),
	// 					},
	// 					Limit: to.Ptr[int64](30),
	// 					ResourceName: to.Ptr("standardav2family"),
	// 					ShareableQuota: to.Ptr[int64](0),
	// 				},
	// 		}},
	// 	},
	// }
}
