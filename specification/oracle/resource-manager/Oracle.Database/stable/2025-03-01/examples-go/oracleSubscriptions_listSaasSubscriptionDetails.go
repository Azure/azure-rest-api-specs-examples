package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: 2025-03-01/oracleSubscriptions_listSaasSubscriptionDetails.json
func ExampleOracleSubscriptionsClient_BeginListSaasSubscriptionDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOracleSubscriptionsClient().BeginListSaasSubscriptionDetails(ctx, nil)
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
	// res = armoracledatabase.OracleSubscriptionsClientListSaasSubscriptionDetailsResponse{
	// 	SaasSubscriptionDetails: &armoracledatabase.SaasSubscriptionDetails{
	// 		ID: to.Ptr("1234567890"),
	// 		SubscriptionName: to.Ptr("Oracle"),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-07T00:00:00Z"); return t}()),
	// 		OfferID: to.Ptr("offer1"),
	// 		PlanID: to.Ptr("silver"),
	// 		SaasSubscriptionStatus: to.Ptr("PendingFulfillmentStart"),
	// 		PublisherID: to.Ptr("Oracle"),
	// 		PurchaserEmailID: to.Ptr("test@test.com"),
	// 		PurchaserTenantID: to.Ptr("1234567890"),
	// 		TermUnit: to.Ptr("P1M"),
	// 		IsAutoRenew: to.Ptr(true),
	// 		IsFreeTrial: to.Ptr(false),
	// 	},
	// }
}
