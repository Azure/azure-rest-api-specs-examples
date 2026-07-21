package armenclave_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/enclave/armenclave"
)

// Generated from example definition: 2026-03-01-preview/TransitHub_ListBySubscription.json
func ExampleTransitHubClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armenclave.NewClientFactory("CA1CB369-DD26-4DB2-9D43-9AFEF0F22093", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTransitHubClient().NewListBySubscriptionPager("TestMyCommunity", nil)
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
		// page = armenclave.TransitHubClientListBySubscriptionResponse{
		// 	TransitHubResourceListResult: armenclave.TransitHubResourceListResult{
		// 		Value: []*armenclave.TransitHubResource{
		// 			{
		// 				Properties: &armenclave.TransitHubProperties{
		// 					ProvisioningState: to.Ptr(armenclave.ProvisioningStateSucceeded),
		// 					State: to.Ptr(armenclave.TransitHubStatePendingApproval),
		// 					TransitOption: &armenclave.TransitOption{
		// 						Type: to.Ptr(armenclave.TransitOptionTypeExpressRoute),
		// 						Params: &armenclave.TransitOptionParams{
		// 							ScaleUnits: to.Ptr[int64](1),
		// 						},
		// 					},
		// 					ResourceCollection: []*string{
		// 						to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg"),
		// 					},
		// 					SecurityProvider: to.Ptr(armenclave.SecurityProviderAzureFirewall),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Tag1": to.Ptr("Value1"),
		// 				},
		// 				Location: to.Ptr("westcentralus"),
		// 				ID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/testrg/providers/Microsoft.Mission/communities/TestMyCommunity/transitHubs/TestThName"),
		// 				Name: to.Ptr("TestThName"),
		// 				Type: to.Ptr("microsoft.mission/communities/transithubs"),
		// 				SystemData: &armenclave.SystemData{
		// 					CreatedBy: to.Ptr("myAlias"),
		// 					CreatedByType: to.Ptr(armenclave.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("myAlias"),
		// 					LastModifiedByType: to.Ptr(armenclave.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
