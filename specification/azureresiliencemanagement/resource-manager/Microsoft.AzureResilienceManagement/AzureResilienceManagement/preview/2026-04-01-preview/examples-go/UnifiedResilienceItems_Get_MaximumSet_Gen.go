package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/UnifiedResilienceItems_Get_MaximumSet_Gen.json
func ExampleUnifiedResilienceItemsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUnifiedResilienceItemsClient().Get(ctx, "sg1", "uri1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresiliencemanagement.UnifiedResilienceItemsClientGetResponse{
	// 	UnifiedResilienceItem: armresiliencemanagement.UnifiedResilienceItem{
	// 		Properties: &armresiliencemanagement.UnifiedResilienceItemProperties{
	// 			ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
	// 			Goals: &armresiliencemanagement.GoalsData{
	// 				TemplateID: to.Ptr("/providers/Microsoft.AzureResilienceManagement/goalTemplates/gt1"),
	// 				AssignmentID: to.Ptr("/providers/Microsoft.AzureResilienceManagement/goalAssignments/ga1"),
	// 				RegionalRecoveryPointObjectiveInMinutes: to.Ptr(armresiliencemanagement.IsoDurationPT15M),
	// 				RegionalRecoveryPointEstimatedInMinutes: to.Ptr(armresiliencemanagement.IsoDuration("PT10M")),
	// 				RegionalRecoveryPointObjectiveStatus: to.Ptr(armresiliencemanagement.ResilienceHealthStatusHealthy),
	// 				RegionalRecoveryTimeObjectiveInMinutes: to.Ptr(armresiliencemanagement.IsoDurationPT1H),
	// 				RegionalRecoveryTimeActualInMinutes: to.Ptr(armresiliencemanagement.IsoDuration("PT45M")),
	// 				RegionalRecoveryTimeObjectiveStatus: to.Ptr(armresiliencemanagement.ResilienceHealthStatusHealthy),
	// 				RequireHighAvailability: to.Ptr(armresiliencemanagement.UnifiedResilienceItemRequirementSelectedRequired),
	// 				RequireDisasterRecovery: to.Ptr(armresiliencemanagement.UnifiedResilienceItemRequirementSelectedNotRequired),
	// 			},
	// 			Recommendations: &armresiliencemanagement.RecommendationsData{
	// 				HighAvailability: &armresiliencemanagement.RecommendationsHighAvailabilityData{
	// 					EnabledResourceCount: to.Ptr[int64](5),
	// 					NotEnabledResourceCount: to.Ptr[int64](2),
	// 					NotEvaluatedResourceCount: to.Ptr[int64](1),
	// 					EvaluationDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-01T08:00:00Z"); return t}()),
	// 				},
	// 			},
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-01T08:00:00Z"); return t}()),
	// 		},
	// 		ID: to.Ptr("/providers/Microsoft.AzureResilienceManagement/unifiedResilienceItems/uri1"),
	// 		Name: to.Ptr("uri1"),
	// 		Type: to.Ptr("Microsoft.AzureResilienceManagement/unifiedResilienceItems"),
	// 		SystemData: &armresiliencemanagement.SystemData{
	// 			CreatedBy: to.Ptr("dvnfxbuyqhvivfjddjccdtlwajfht"),
	// 			CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-06T15:03:42.796Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("lndhhaimomorael"),
	// 			LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-06T15:03:42.797Z"); return t}()),
	// 		},
	// 	},
	// }
}
