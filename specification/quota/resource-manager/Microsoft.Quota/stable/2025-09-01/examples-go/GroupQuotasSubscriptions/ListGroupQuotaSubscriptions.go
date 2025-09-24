package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota/v2"
)

// Generated from example definition: 2025-09-01/GroupQuotasSubscriptions/ListGroupQuotaSubscriptions.json
func ExampleGroupQuotaSubscriptionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGroupQuotaSubscriptionsClient("00000000-0000-0000-0000-000000000000").NewListPager("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1", nil)
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
		// page = armquota.GroupQuotaSubscriptionsClientListResponse{
		// 	GroupQuotaSubscriptionIDList: armquota.GroupQuotaSubscriptionIDList{
		// 		NextLink: to.Ptr("https://yourLinkHere.com"),
		// 		Value: []*armquota.GroupQuotaSubscriptionID{
		// 			{
		// 				Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				Type: to.Ptr("Microsoft.Quota/groupQuotas/subscriptions"),
		// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/providers/Microsoft.Quota/groupQuotas/groupquota1/subscriptions/00000000-0000-0000-0000-000000000000"),
		// 				Properties: &armquota.GroupQuotaSubscriptionIDProperties{
		// 					ProvisioningState: to.Ptr(armquota.RequestStateSucceeded),
		// 					SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				Type: to.Ptr("Microsoft.Quota/groupQuotas/subscriptions"),
		// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/providers/Microsoft.Quota/groupQuotas/groupquota1/subscriptions/11111111-1111-1111-1111-111111111111"),
		// 				Properties: &armquota.GroupQuotaSubscriptionIDProperties{
		// 					ProvisioningState: to.Ptr(armquota.RequestStateSucceeded),
		// 					SubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
