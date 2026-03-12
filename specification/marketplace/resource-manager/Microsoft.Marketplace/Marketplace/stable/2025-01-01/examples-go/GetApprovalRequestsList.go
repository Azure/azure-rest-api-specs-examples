package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace/v2"
)

// Generated from example definition: 2025-01-01/GetApprovalRequestsList.json
func ExamplePrivateStoreClient_GetApprovalRequestsList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreClient().GetApprovalRequestsList(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmarketplace.PrivateStoreClientGetApprovalRequestsListResponse{
	// 	RequestApprovalsList: &armmarketplace.RequestApprovalsList{
	// 		Value: []*armmarketplace.RequestApprovalResource{
	// 			{
	// 				Name: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
	// 				Type: to.Ptr("Microsoft.Marketplace/privateStores/requestApprovals"),
	// 				ID: to.Ptr("/providers/Microsoft.Marketplace/privateStores/9afd3c45-5230-4d58-9469-2cacc00bba68/requestApprovals/marketplacetestthirdparty.md-test-third-party-2"),
	// 				Properties: &armmarketplace.RequestApprovalProperties{
	// 					IsClosed: to.Ptr(false),
	// 					MessageCode: to.Ptr[int64](0),
	// 					OfferDisplayName: to.Ptr("Offer Display name"),
	// 					OfferID: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
	// 					PlansDetails: []*armmarketplace.PlanDetails{
	// 						{
	// 							Justification: to.Ptr("Because I want to...."),
	// 							PlanID: to.Ptr("testPlanA"),
	// 							RequestDate: "2021-02-01T10:23:17.6571572+02:00",
	// 							Status: to.Ptr(armmarketplace.StatusPending),
	// 							SubscriptionID: to.Ptr(""),
	// 							SubscriptionName: to.Ptr(""),
	// 						},
	// 						{
	// 							Justification: to.Ptr("try me :)"),
	// 							PlanID: to.Ptr("*"),
	// 							RequestDate: "2021-02-01T10:23:17.6571572+02:00",
	// 							Status: to.Ptr(armmarketplace.StatusPending),
	// 							SubscriptionID: to.Ptr(""),
	// 							SubscriptionName: to.Ptr(""),
	// 						},
	// 					},
	// 					PublisherID: to.Ptr("marketplacetestthirdparty"),
	// 				},
	// 				SystemData: &armmarketplace.SystemData{
	// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-01T10:23:17.6571572+02:00"); return t}()),
	// 					CreatedBy: to.Ptr("user@somedoamin.com"),
	// 					CreatedByType: to.Ptr(armmarketplace.CreatedByTypeUser),
	// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-01T10:23:17.6571572+02:00"); return t}()),
	// 					LastModifiedBy: to.Ptr("user@somedoamin.com"),
	// 					LastModifiedByType: to.Ptr(armmarketplace.CreatedByTypeUser),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
