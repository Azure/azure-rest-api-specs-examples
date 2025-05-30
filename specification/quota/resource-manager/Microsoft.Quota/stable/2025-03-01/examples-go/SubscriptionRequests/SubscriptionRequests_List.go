package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8691e5081766c7ad602a9e55de841d07bed5196a/specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/SubscriptionRequests/SubscriptionRequests_List.json
func ExampleGroupQuotaSubscriptionRequestsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGroupQuotaSubscriptionRequestsClient().NewListPager("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1", nil)
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
		// page.GroupQuotaSubscriptionRequestStatusList = armquota.GroupQuotaSubscriptionRequestStatusList{
		// 	Value: []*armquota.GroupQuotaSubscriptionRequestStatus{
		// 		{
		// 			Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 			Type: to.Ptr("Microsoft.Quota/groupQuotas/SubscriptionRequest"),
		// 			ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111"),
		// 			Properties: &armquota.GroupQuotaSubscriptionRequestStatusProperties{
		// 				ProvisioningState: to.Ptr(armquota.RequestStateSucceeded),
		// 				RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-17T00:56:45.009Z"); return t}()),
		// 				SubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 			},
		// 	}},
		// }
	}
}
