package armimpactreporting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
)

// Generated from example definition: 2024-05-01-preview/WorkloadImpact_Get.json
func ExampleWorkloadImpactsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkloadImpactsClient().Get(ctx, "impact-001", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armimpactreporting.WorkloadImpactsClientGetResponse{
	// 	WorkloadImpact: &armimpactreporting.WorkloadImpact{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Impact/workloadImpacts/impact-001"),
	// 		Name: to.Ptr("impact-001"),
	// 		Type: to.Ptr("Microsoft.Impact/workloadImpacts"),
	// 		Properties: &armimpactreporting.WorkloadImpactProperties{
	// 			ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservername"),
	// 			StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-15T05:59:46.6517821Z"); return t}()),
	// 			ImpactDescription: to.Ptr("high cup utilization"),
	// 			ImpactCategory: to.Ptr("Performance"),
	// 			Workload: &armimpactreporting.Workload{
	// 				Context: to.Ptr("webapp/scenario1"),
	// 				Toolset: to.Ptr(armimpactreporting.ToolsetOther),
	// 			},
	// 			ClientIncidentDetails: &armimpactreporting.ClientIncidentDetails{
	// 				ClientIncidentID: to.Ptr("AA123"),
	// 				ClientIncidentSource: to.Ptr(armimpactreporting.IncidentSourceJira),
	// 			},
	// 			ImpactUniqueID: to.Ptr("d7f24d04-e7f0-48bf-b09c-9d36ca9e1777"),
	// 			ReportedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-15T06:01:46.6517821Z"); return t}()),
	// 		},
	// 	},
	// }
}
