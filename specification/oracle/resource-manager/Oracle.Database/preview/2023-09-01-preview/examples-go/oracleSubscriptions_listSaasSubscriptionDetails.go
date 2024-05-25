package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/520e274d7d95fc6d1002dd3c1fcaf8d55d27f63e/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/oracleSubscriptions_listSaasSubscriptionDetails.json
func ExampleOracleSubscriptionsClient_BeginListSaasSubscriptionDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.SaasSubscriptionDetails = armoracledatabase.SaasSubscriptionDetails{
	// 	ID: to.Ptr("1234567890"),
	// 	IsAutoRenew: to.Ptr(true),
	// 	IsFreeTrial: to.Ptr(false),
	// 	OfferID: to.Ptr("offer1"),
	// 	PlanID: to.Ptr("silver"),
	// 	PublisherID: to.Ptr("Oracle"),
	// 	PurchaserEmailID: to.Ptr("test@test.com"),
	// 	PurchaserTenantID: to.Ptr("1234567890"),
	// 	SaasSubscriptionStatus: to.Ptr("PendingFulfillmentStart"),
	// 	SubscriptionName: to.Ptr("Oracle"),
	// 	TermUnit: to.Ptr("P1M"),
	// 	TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-07T00:00:00.000Z"); return t}()),
	// }
}
