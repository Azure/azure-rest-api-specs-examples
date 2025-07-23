package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard/v2"
)

// Generated from example definition: 2024-11-01-preview/EnterpriseDetails_Post.json
func ExampleGrafanaClient_CheckEnterpriseDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGrafanaClient().CheckEnterpriseDetails(ctx, "myResourceGroup", "myWorkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdashboard.GrafanaClientCheckEnterpriseDetailsResponse{
	// 	EnterpriseDetails: &armdashboard.EnterpriseDetails{
	// 		MarketplaceTrialQuota: &armdashboard.MarketplaceTrialQuota{
	// 			AvailablePromotion: to.Ptr(armdashboard.AvailablePromotionNone),
	// 			GrafanaResourceID: to.Ptr("/subscriptions/e1e3b30d-e7ec-4e25-8587-db037bcb9a4d/resourcegroups/amg-local-script-test-rg/providers/microsoft.dashboard/grafana/eus2-enterprise-1001-07"),
	// 			TrialEndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-03T01:06:00.447Z"); return t}()),
	// 			TrialStartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-04T01:06:00.447Z"); return t}()),
	// 		},
	// 		SaasSubscriptionDetails: &armdashboard.SaasSubscriptionDetails{
	// 			OfferID: to.Ptr("amg_test"),
	// 			PlanID: to.Ptr("amg_globalplan"),
	// 			PublisherID: to.Ptr("isvtestuklegacy"),
	// 			Term: &armdashboard.SubscriptionTerm{
	// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-03T00:00:00Z"); return t}()),
	// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-04T00:00:00Z"); return t}()),
	// 				TermUnit: to.Ptr("P1M"),
	// 			},
	// 		},
	// 	},
	// }
}
