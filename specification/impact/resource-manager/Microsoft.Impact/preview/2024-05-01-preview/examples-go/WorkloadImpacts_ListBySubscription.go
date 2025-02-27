package armimpactreporting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
)

// Generated from example definition: 2024-05-01-preview/WorkloadImpacts_ListBySubscription.json
func ExampleWorkloadImpactsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkloadImpactsClient().NewListBySubscriptionPager(nil)
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
		// page = armimpactreporting.WorkloadImpactsClientListBySubscriptionResponse{
		// 	WorkloadImpactListResult: armimpactreporting.WorkloadImpactListResult{
		// 		Value: []*armimpactreporting.WorkloadImpact{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/Impacts-rg/providers/Microsoft.Impact/workloadImpacts/impact2"),
		// 				Name: to.Ptr("impact2"),
		// 				Type: to.Ptr("Microsoft.Impact/workloadImpacts"),
		// 				Properties: &armimpactreporting.WorkloadImpactProperties{
		// 					ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.compute/virtualmachines/vm1"),
		// 					ImpactUniqueID: to.Ptr("d7f24d04-e7f0-48bf-b09c-9d36ca9e1777"),
		// 					ReportedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-15T06:01:46.6517821Z"); return t}()),
		// 					StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-15T05:59:46.6517821Z"); return t}()),
		// 					ImpactDescription: to.Ptr(""),
		// 					ImpactCategory: to.Ptr("Resource.Other"),
		// 					AdditionalProperties: &armimpactreporting.WorkloadImpactPropertiesAdditionalProperties{
		// 					},
		// 					Workload: &armimpactreporting.Workload{
		// 						Context: to.Ptr("webapp/scenario1"),
		// 						Toolset: to.Ptr(armimpactreporting.ToolsetOther),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
