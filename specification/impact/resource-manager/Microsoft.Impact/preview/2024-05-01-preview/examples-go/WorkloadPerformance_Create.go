package armimpactreporting_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
)

// Generated from example definition: 2024-05-01-preview/WorkloadPerformance_Create.json
func ExampleWorkloadImpactsClient_BeginCreate_reportingPerformanceRelatedImpact() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadImpactsClient().BeginCreate(ctx, "impact-002", armimpactreporting.WorkloadImpact{
		Properties: &armimpactreporting.WorkloadImpactProperties{
			ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservercontext"),
			StartDateTime:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-15T05:59:46.6517821Z"); return t }()),
			ImpactDescription:  to.Ptr("high cpu utilization"),
			ImpactCategory:     to.Ptr("Resource.Performance"),
			Workload: &armimpactreporting.Workload{
				Context: to.Ptr("webapp/scenario1"),
				Toolset: to.Ptr(armimpactreporting.ToolsetOther),
			},
			Performance: []*armimpactreporting.Performance{
				{
					MetricName: to.Ptr("CPU"),
					Actual:     to.Ptr[float64](90),
					Expected:   to.Ptr[float64](60),
				},
			},
			ClientIncidentDetails: &armimpactreporting.ClientIncidentDetails{
				ClientIncidentID:     to.Ptr("AA123"),
				ClientIncidentSource: to.Ptr(armimpactreporting.IncidentSourceJira),
			},
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
	// res = armimpactreporting.WorkloadImpactsClientCreateResponse{
	// 	WorkloadImpact: &armimpactreporting.WorkloadImpact{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Impact/PerformanceImpacts/impact-001"),
	// 		Name: to.Ptr("impact-001"),
	// 		Type: to.Ptr("Microsoft.Impact/PerformanceImpacts"),
	// 		Properties: &armimpactreporting.WorkloadImpactProperties{
	// 			ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservername"),
	// 			StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-15T05:59:46.6517821Z"); return t}()),
	// 			ImpactDescription: to.Ptr("high cup utilization"),
	// 			ImpactCategory: to.Ptr("Resource.Performance"),
	// 			Workload: &armimpactreporting.Workload{
	// 				Context: to.Ptr("webapp/scenario1"),
	// 				Toolset: to.Ptr(armimpactreporting.ToolsetOther),
	// 			},
	// 			Performance: []*armimpactreporting.Performance{
	// 				{
	// 					MetricName: to.Ptr("CPU"),
	// 					Actual: to.Ptr[float64](90),
	// 					Expected: to.Ptr[float64](60),
	// 				},
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
