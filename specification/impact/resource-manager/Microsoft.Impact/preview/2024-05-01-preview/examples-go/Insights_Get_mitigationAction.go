package armimpactreporting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
)

// Generated from example definition: 2024-05-01-preview/Insights_Get_mitigationAction.json
func ExampleInsightsClient_Get_getInsightSampleForMitigationActionCategory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInsightsClient().Get(ctx, "impactId", "HPCUASucceeded", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armimpactreporting.InsightsClientGetResponse{
	// 	Insight: &armimpactreporting.Insight{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Impact/workloadimpacts/impactId/insights/HPCUASucceeded"),
	// 		Name: to.Ptr("HPCUASucceeded"),
	// 		Type: to.Ptr("Microsoft.impact/workloadimpacts/insights"),
	// 		Properties: &armimpactreporting.InsightProperties{
	// 			Category: to.Ptr("MitigationAction"),
	// 			EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T04:00:00.009223Z"); return t}()),
	// 			InsightUniqueID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			Impact: &armimpactreporting.ImpactDetails{
	// 				ImpactID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Impact/workloadimpacts/impactId"),
	// 				ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/virtualMachine/vm"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
	// 			},
	// 			Content: &armimpactreporting.Content{
	// 				Title: to.Ptr("Node was flagged for inspection"),
	// 				Description: to.Ptr("At 2024-02-16 21:05:07 (UTC) an impact was reported on west-eur-VM indicating a potential disruption from Azure platform. <strong>Action Taken</strong> The hardware your VM was running on was flagged for inspection. We will conduct our investigation and take corrective action to prevent further impact on your workloads. We apologize for any disruption. Microsoft Azure Team"),
	// 			},
	// 			AdditionalDetails: &armimpactreporting.InsightPropertiesAdditionalDetails{
	// 			},
	// 		},
	// 	},
	// }
}
