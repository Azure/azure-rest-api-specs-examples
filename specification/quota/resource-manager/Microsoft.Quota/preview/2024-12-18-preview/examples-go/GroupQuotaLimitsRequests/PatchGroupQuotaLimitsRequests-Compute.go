package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/55c5a0cd6da80b2700333c01e9a9c6067de9cef0/specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/GroupQuotaLimitsRequests/PatchGroupQuotaLimitsRequests-Compute.json
func ExampleGroupQuotaLimitsRequestClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGroupQuotaLimitsRequestClient().BeginUpdate(ctx, "E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1", "Microsoft.Compute", "westus", &armquota.GroupQuotaLimitsRequestClientBeginUpdateOptions{GroupQuotaRequest: &armquota.GroupQuotaLimitList{
		Properties: &armquota.GroupQuotaLimitListProperties{
			Value: []*armquota.GroupQuotaLimit{
				{
					Properties: &armquota.GroupQuotaLimitProperties{
						Comment:      to.Ptr("Contoso requires more quota."),
						Limit:        to.Ptr[int64](110),
						ResourceName: to.Ptr("standardddv4family"),
					},
				},
				{
					Properties: &armquota.GroupQuotaLimitProperties{
						Comment:      to.Ptr("Contoso requires more quota."),
						Limit:        to.Ptr[int64](110),
						ResourceName: to.Ptr("standardav2family"),
					},
				}},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GroupQuotaLimitList = armquota.GroupQuotaLimitList{
	// 	Name: to.Ptr("westus"),
	// 	Type: to.Ptr("Microsoft.Quota/groupQuotas/groupQuotaLimits"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/providers/Microsoft.Quota/groupQuotas/groupquota1/resourceProviders/Microsoft.Compute/groupQuotaLimits/westus"),
	// 	Properties: &armquota.GroupQuotaLimitListProperties{
	// 		ProvisioningState: to.Ptr(armquota.RequestStateSucceeded),
	// 		Value: []*armquota.GroupQuotaLimit{
	// 			{
	// 				Properties: &armquota.GroupQuotaLimitProperties{
	// 					Name: &armquota.GroupQuotaDetailsName{
	// 						LocalizedValue: to.Ptr("standard DDv4 Family vCPUs"),
	// 						Value: to.Ptr("standardddv4family"),
	// 					},
	// 					AllocatedToSubscriptions: &armquota.AllocatedQuotaToSubscriptionList{
	// 						Value: []*armquota.AllocatedToSubscription{
	// 							{
	// 								QuotaAllocated: to.Ptr[int64](20),
	// 								SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 							},
	// 							{
	// 								QuotaAllocated: to.Ptr[int64](30),
	// 								SubscriptionID: to.Ptr("A000000-0000-0000-0000-000000000000"),
	// 						}},
	// 					},
	// 					AvailableLimit: to.Ptr[int64](50),
	// 					Limit: to.Ptr[int64](100),
	// 					ResourceName: to.Ptr("standardddv4family"),
	// 					Unit: to.Ptr("count"),
	// 				},
	// 			},
	// 			{
	// 				Properties: &armquota.GroupQuotaLimitProperties{
	// 					Name: &armquota.GroupQuotaDetailsName{
	// 						LocalizedValue: to.Ptr("Standard AV2 Family vCPUs"),
	// 						Value: to.Ptr("standardav2family"),
	// 					},
	// 					AllocatedToSubscriptions: &armquota.AllocatedQuotaToSubscriptionList{
	// 						Value: []*armquota.AllocatedToSubscription{
	// 							{
	// 								QuotaAllocated: to.Ptr[int64](20),
	// 								SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 						}},
	// 					},
	// 					AvailableLimit: to.Ptr[int64](80),
	// 					Limit: to.Ptr[int64](100),
	// 					ResourceName: to.Ptr("standardav2family"),
	// 					Unit: to.Ptr("count"),
	// 				},
	// 		}},
	// 	},
	// }
}
