package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8691e5081766c7ad602a9e55de841d07bed5196a/specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/SubscriptionQuotaAllocationRequests/SubscriptionQuotaAllocationRequests_Get-Compute.json
func ExampleGroupQuotaSubscriptionAllocationRequestClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGroupQuotaSubscriptionAllocationRequestClient("<subscription-id>").Get(ctx, "E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1", "Microsoft.Compute", "AE000000-0000-0000-0000-00000000000A", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AllocationRequestStatus = armquota.AllocationRequestStatus{
	// 	Name: to.Ptr("AE000000-0000-0000-0000-00000000000A"),
	// 	Type: to.Ptr("Microsoft.Quota/groupQuotas/quotaAllocationRequests"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Quota/groupQuotas/groupquota1/quotaAllocationRequests/AE000000-0000-0000-0000-00000000000A"),
	// 	Properties: &armquota.AllocationRequestStatusProperties{
	// 		FaultCode: to.Ptr("ContactSupport"),
	// 		ProvisioningState: to.Ptr(armquota.RequestStateSucceeded),
	// 		RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-17T01:06:02.191Z"); return t}()),
	// 		RequestedResource: &armquota.AllocationRequestBase{
	// 			Properties: &armquota.AllocationRequestBaseProperties{
	// 				Name: &armquota.AllocationRequestBasePropertiesName{
	// 					LocalizedValue: to.Ptr("standard Av2 Family vCPUs"),
	// 					Value: to.Ptr("standardav2family"),
	// 				},
	// 				Limit: to.Ptr[int64](75),
	// 				Region: to.Ptr("westus"),
	// 			},
	// 		},
	// 	},
	// }
}
