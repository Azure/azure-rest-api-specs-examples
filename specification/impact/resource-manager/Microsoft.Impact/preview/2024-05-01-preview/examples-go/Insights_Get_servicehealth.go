package armimpactreporting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
)

// Generated from example definition: 2024-05-01-preview/Insights_Get_servicehealth.json
func ExampleInsightsClient_Get_getInsightSampleForServiceHealthCategory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInsightsClient().Get(ctx, "impactid", "insightname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armimpactreporting.InsightsClientGetResponse{
	// 	Insight: &armimpactreporting.Insight{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactname/insights/insightname"),
	// 		Name: to.Ptr("insightname"),
	// 		Type: to.Ptr("Microsoft.Impact/insights"),
	// 		Properties: &armimpactreporting.InsightProperties{
	// 			Category: to.Ptr("ServiceHealth"),
	// 			EventID: to.Ptr("ABC-123"),
	// 			EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T04:00:00.009223Z"); return t}()),
	// 			Impact: &armimpactreporting.ImpactDetails{
	// 				ImpactID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Impact/workloadimpacts/impactid"),
	// 				ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.compute/virtualmachines/vm1"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
	// 			},
	// 			InsightUniqueID: to.Ptr("a3d91a07-698b-4044-a230-e918252c4c59"),
	// 			Content: &armimpactreporting.Content{
	// 				Title: to.Ptr("Impact Has been correlated to an outage"),
	// 				Description: to.Ptr("On November 8, 2018, at 00:00 UTC, there may have been disruptions to your services linked to the resource <a href='...'>VM1<a>. We have pinpointed a service outage impacting these resources. For details on this outage, please refer to the service health information available at <a href='...'>service health</a>."),
	// 			},
	// 		},
	// 	},
	// }
}
