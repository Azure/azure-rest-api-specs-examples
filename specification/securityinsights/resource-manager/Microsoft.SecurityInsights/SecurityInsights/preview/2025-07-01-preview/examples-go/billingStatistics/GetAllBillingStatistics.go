package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/billingStatistics/GetAllBillingStatistics.json
func ExampleBillingStatisticsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBillingStatisticsClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page = armsecurityinsights.BillingStatisticsClientListResponse{
		// 	BillingStatisticList: armsecurityinsights.BillingStatisticList{
		// 		Value: []armsecurityinsights.BillingStatisticClassification{
		// 			&armsecurityinsights.SapSolutionUsageStatistic{
		// 				Name: to.Ptr("sapSolutionUsage"),
		// 				Type: to.Ptr("Microsoft.SecurityInsights/billingStatistics"),
		// 				Etag: to.Ptr("\"3f6451dd-1b58-4bef-bce7-72eba6b354d7\""),
		// 				ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/billingStatistics/sapUsage"),
		// 				Kind: to.Ptr(armsecurityinsights.BillingStatisticKindSapSolutionUsage),
		// 				Properties: &armsecurityinsights.SapSolutionUsageStatisticProperties{
		// 					ActiveSystemIDCount: to.Ptr[int64](5),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
