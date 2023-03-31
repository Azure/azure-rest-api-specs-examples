package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/QueryRequestApproval.json
func ExamplePrivateStoreClient_QueryRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreClient().QueryRequestApproval(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", "marketplacetestthirdparty.md-test-third-party-2", &armmarketplace.PrivateStoreClientQueryRequestApprovalOptions{Payload: &armmarketplace.QueryRequestApprovalProperties{
		Properties: &armmarketplace.RequestDetails{
			PlanIDs: []*string{
				to.Ptr("testPlanA"),
				to.Ptr("testPlanB"),
				to.Ptr("*")},
			PublisherID: to.Ptr("marketplacetestthirdparty"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryRequestApproval = armmarketplace.QueryRequestApproval{
	// 	MessageCode: to.Ptr[int64](0),
	// 	PlansDetails: map[string]*armmarketplace.PlanDetails{
	// 		"*": &armmarketplace.PlanDetails{
	// 			PlanID: to.Ptr("*"),
	// 			Status: to.Ptr(armmarketplace.StatusNone),
	// 			SubscriptionID: to.Ptr(""),
	// 			SubscriptionName: to.Ptr(""),
	// 		},
	// 		"byol": &armmarketplace.PlanDetails{
	// 			Justification: to.Ptr(""),
	// 			PlanID: to.Ptr("testPlanA"),
	// 			Status: to.Ptr(armmarketplace.StatusNone),
	// 			SubscriptionID: to.Ptr(""),
	// 			SubscriptionName: to.Ptr(""),
	// 		},
	// 		"hourly": &armmarketplace.PlanDetails{
	// 			Justification: to.Ptr(""),
	// 			PlanID: to.Ptr("testPlanB"),
	// 			Status: to.Ptr(armmarketplace.Status("ApprovedByAdmin")),
	// 			SubscriptionID: to.Ptr("4ca4753c-5a1e-4913-b849-2c68880e03c2"),
	// 			SubscriptionName: to.Ptr("Test subscription 2"),
	// 		},
	// 	},
	// 	UniqueOfferID: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
	// }
}
