package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ee6d9fd7687d4b67117c5a167c191a7e7e70b53/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/oracleSubscriptions_listBySubscription.json
func ExampleOracleSubscriptionsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOracleSubscriptionsClient().NewListBySubscriptionPager(nil)
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
		// page.OracleSubscriptionListResult = armoracledatabase.OracleSubscriptionListResult{
		// 	Value: []*armoracledatabase.OracleSubscription{
		// 		{
		// 			Type: to.Ptr("Oracle.Database/oracleSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/oracleSubscriptions/default"),
		// 			Plan: &armoracledatabase.Plan{
		// 				Name: to.Ptr("plan1"),
		// 				Product: to.Ptr("product1"),
		// 				PromotionCode: to.Ptr("none"),
		// 				Publisher: to.Ptr("publisher1"),
		// 				Version: to.Ptr("alpha"),
		// 			},
		// 			Properties: &armoracledatabase.OracleSubscriptionProperties{
		// 				CloudAccountID: to.Ptr("ocid1..aaaaaa"),
		// 				CloudAccountState: to.Ptr(armoracledatabase.CloudAccountProvisioningStatePending),
		// 				ProvisioningState: to.Ptr(armoracledatabase.OracleSubscriptionProvisioningStateSucceeded),
		// 				SaasSubscriptionID: to.Ptr("saas1"),
		// 			},
		// 	}},
		// }
	}
}
