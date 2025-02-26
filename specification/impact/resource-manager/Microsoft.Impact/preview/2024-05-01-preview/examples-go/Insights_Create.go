package armimpactreporting_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
)

// Generated from example definition: 2024-05-01-preview/Insights_Create.json
func ExampleInsightsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInsightsClient().Create(ctx, "impactid22", "insightId12", armimpactreporting.Insight{
		Properties: &armimpactreporting.InsightProperties{
			Content: &armimpactreporting.Content{
				Title:       to.Ptr("Impact Has been correlated to an outage"),
				Description: to.Ptr("At 2018-11-08T00:00:00Z UTC, your services dependent on these resources <link href=”…”>VM1</link> may have experienced an issue. <br/><div>We have identified an outage that affected these resources(s). You can look at outage information on <link href=\"https:// portal.azure.com/#view/Microsoft_Azure_Health/AzureHealthBrowseBlade/~/serviceIssues/trackingId/NL2W-VCZ\">NL2W-VCZ</link> link.<div>"),
			},
			Category:        to.Ptr("repair"),
			Status:          to.Ptr("resolved"),
			EventTime:       to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T04:00:00.009223Z"); return t }()),
			InsightUniqueID: to.Ptr("00000000-0000-0000-0000-000000000000"),
			Impact: &armimpactreporting.ImpactDetails{
				ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservername"),
				StartTime:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T01:00:00.009223Z"); return t }()),
				ImpactID:           to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactid22"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armimpactreporting.InsightsClientCreateResponse{
	// 	Insight: &armimpactreporting.Insight{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactid22/insights/insightId12"),
	// 		Name: to.Ptr("insightId12"),
	// 		Type: to.Ptr("Microsoft.Impact/insights"),
	// 		Properties: &armimpactreporting.InsightProperties{
	// 			EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T04:00:00.009223Z"); return t}()),
	// 			Content: &armimpactreporting.Content{
	// 				Title: to.Ptr("Impact Has been correlated to an outage"),
	// 				Description: to.Ptr("At 2018-11-08T00:00:00Z UTC, your services dependent on these resources <link href=”…”>VM1</link> may have experienced an issue. <br/><div>We have identified an outage that affected these resources(s). You can look at outage information on <link href=\"https:// portal.azure.com/#view/Microsoft_Azure_Health/AzureHealthBrowseBlade/~/serviceIssues/trackingId/NL2W-VCZ\">NL2W-VCZ</link> link.<div>"),
	// 			},
	// 			Category: to.Ptr("repair"),
	// 			Status: to.Ptr("resolved"),
	// 			InsightUniqueID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			Impact: &armimpactreporting.ImpactDetails{
	// 				ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservername"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T01:00:00.009223Z"); return t}()),
	// 				ImpactID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactid22"),
	// 			},
	// 		},
	// 	},
	// }
}
