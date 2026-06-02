package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-04-01/AdvancedPlatformMetricsCRUD/AdvancedPlatformMetricsRules_List.json
func ExampleAdvancedPlatformMetricsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAdvancedPlatformMetricsClient().NewListPager("res6977", "sto2527", nil)
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
		// page = armstorage.AdvancedPlatformMetricsClientListResponse{
		// 	AdvancedPlatformMetricsRuleListResult: armstorage.AdvancedPlatformMetricsRuleListResult{
		// 		Value: []*armstorage.AdvancedPlatformMetricsRule{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/advancedPlatformMetrics/ContainerLevelCapacityMetrics"),
		// 				Name: to.Ptr("DefaultAdvancedPlatformMetricsRule"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/advancedPlatformMetrics"),
		// 				Properties: &armstorage.AdvancedPlatformMetricsRuleProperties{
		// 					RuleType: to.Ptr(armstorage.AdvancedPlatformMetricsRuleTypeContainerLevelCapacityMetrics),
		// 					Enabled: to.Ptr(true),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-01T11:00:00.0000000Z"); return t}()),
		// 					MetricsEmitted: []*armstorage.MetricsEmitted{
		// 						to.Ptr(armstorage.MetricsEmittedContainerUsedSize),
		// 						to.Ptr(armstorage.MetricsEmittedContainerBlobCount),
		// 					},
		// 					RuleConfig: &armstorage.AdvancedPlatformMetricsRuleConfig{
		// 						FilterType: to.Ptr(armstorage.AdvancedPlatformMetricsFilterTypeAllContainersFilter),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
