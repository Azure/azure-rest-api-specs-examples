package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace/v2"
)

// Generated from example definition: 2025-01-01/QueryApprovedPlans.json
func ExamplePrivateStoreClient_QueryApprovedPlans() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreClient().QueryApprovedPlans(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", &armmarketplace.PrivateStoreClientQueryApprovedPlansOptions{
		Payload: &armmarketplace.QueryApprovedPlansPayload{
			Properties: &armmarketplace.QueryApprovedPlans{
				OfferID: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
				PlanIDs: []*string{
					to.Ptr("testPlanA"),
					to.Ptr("testPlanB"),
					to.Ptr("testPlanC"),
				},
				SubscriptionIDs: []*string{
					to.Ptr("85e3e079-c718-4e4c-abbe-f72fceba8305"),
					to.Ptr("7752d461-4bf1-4185-8b56-8a3f11486ac6"),
				},
			},
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmarketplace.PrivateStoreClientQueryApprovedPlansResponse{
	// 	QueryApprovedPlansResponse: &armmarketplace.QueryApprovedPlansResponse{
	// 		Details: []*armmarketplace.QueryApprovedPlansDetails{
	// 			{
	// 				AllSubscriptions: to.Ptr(false),
	// 				PlanID: to.Ptr("testPlanA"),
	// 				SubscriptionIDs: []*string{
	// 					to.Ptr("85e3e079-c718-4e4c-abbe-f72fceba8305"),
	// 					to.Ptr("7752d461-4bf1-4185-8b56-8a3f11486ac6"),
	// 				},
	// 			},
	// 			{
	// 				AllSubscriptions: to.Ptr(true),
	// 				PlanID: to.Ptr("testPlanB"),
	// 				SubscriptionIDs: []*string{
	// 				},
	// 			},
	// 			{
	// 				AllSubscriptions: to.Ptr(false),
	// 				PlanID: to.Ptr("testPlanC"),
	// 				SubscriptionIDs: []*string{
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
