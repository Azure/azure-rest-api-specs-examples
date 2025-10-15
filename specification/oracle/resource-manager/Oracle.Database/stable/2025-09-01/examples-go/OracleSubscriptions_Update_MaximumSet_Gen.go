package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
)

// Generated from example definition: 2025-09-01/OracleSubscriptions_Update_MaximumSet_Gen.json
func ExampleOracleSubscriptionsClient_BeginUpdate_patchOracleSubscriptionGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOracleSubscriptionsClient().BeginUpdate(ctx, armoracledatabase.OracleSubscriptionUpdate{
		Plan: &armoracledatabase.PlanUpdate{
			Name:          to.Ptr("klnnbggrxhvvaiajvjx"),
			Publisher:     to.Ptr("xvsarzadrjqergudsohjk"),
			Product:       to.Ptr("hivkczjyrimjilbmqj"),
			PromotionCode: to.Ptr("zhotaxrodldvmwpksvsrwbnc"),
			Version:       to.Ptr("ueudckjmuqpjvsmmenzyflgpa"),
		},
		Properties: &armoracledatabase.OracleSubscriptionUpdateProperties{
			ProductCode: to.Ptr("kbqzsukkjceoplyalyrdayfj"),
			Intent:      to.Ptr(armoracledatabase.IntentRetain),
		},
	}, nil)
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
	// res = armoracledatabase.OracleSubscriptionsClientUpdateResponse{
	// 	OracleSubscription: &armoracledatabase.OracleSubscription{
	// 		Properties: &armoracledatabase.OracleSubscriptionProperties{
	// 			ProvisioningState: to.Ptr(armoracledatabase.OracleSubscriptionProvisioningStateSucceeded),
	// 			SaasSubscriptionID: to.Ptr("saas1"),
	// 			CloudAccountID: to.Ptr("ocid1..aaaaaa"),
	// 			CloudAccountState: to.Ptr(armoracledatabase.CloudAccountProvisioningStatePending),
	// 			TermUnit: to.Ptr("akzq"),
	// 			Intent: to.Ptr(armoracledatabase.IntentRetain),
	// 			AzureSubscriptionIDs: []*string{
	// 				to.Ptr("qrplrbvwarbkqsqdodbtdoyo"),
	// 			},
	// 			AddSubscriptionOperationState: to.Ptr(armoracledatabase.AddSubscriptionOperationStateSucceeded),
	// 			LastOperationStatusDetail: to.Ptr("inkvybccxpzxvkwzhfcztr"),
	// 		},
	// 		Plan: &armoracledatabase.Plan{
	// 			Name: to.Ptr("plan1"),
	// 			Publisher: to.Ptr("publisher1"),
	// 			Product: to.Ptr("product1"),
	// 			PromotionCode: to.Ptr("none"),
	// 			Version: to.Ptr("alpha"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/oracleSubscriptions/default"),
	// 		Name: to.Ptr("sxraf"),
	// 		Type: to.Ptr("Oracle.Database/oracleSubscriptions"),
	// 		SystemData: &armoracledatabase.SystemData{
	// 			CreatedBy: to.Ptr("sqehacivpuim"),
	// 			CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("axrqfdkqylvjv"),
	// 			LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
	// 		},
	// 	},
	// }
}
