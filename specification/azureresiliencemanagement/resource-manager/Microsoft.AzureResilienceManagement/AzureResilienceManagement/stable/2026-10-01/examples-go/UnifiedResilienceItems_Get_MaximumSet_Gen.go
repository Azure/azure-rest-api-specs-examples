package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-10-01/UnifiedResilienceItems_Get_MaximumSet_Gen.json
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
	// 				AssignmentID: to.Ptr("/providers/Microsoft.AzureResilienceManagement/goalAssignments/ga1"),
	// 				ZonalResiliency: &armresiliencemanagement.UnifiedResilienceItemGoalRequirement{
	// 					Required: to.Ptr(true),
	// 				},
	// 			},
	// 			ResiliencyPosture: &armresiliencemanagement.UnifiedResilienceItemResiliencyPosture{
	// 				ZonalResiliency: &armresiliencemanagement.UnifiedResilienceItemZonalResiliencyPosture{
	// 					EnabledResourceCount: to.Ptr[int64](5),
	// 					NotEnabledResourceCount: to.Ptr[int64](2),
	// 					NotEvaluatedResourceCount: to.Ptr[int64](1),
	// 					UserConfirmationNeededCount: to.Ptr[int64](3),
	// 					EvaluationDateTime: to.Ptr(time.Date(2025, time.May, 1, 8, 0, 0, 0, time.UTC)),
	// 				},
	// 			},
	// 			BillingInfo: &armresiliencemanagement.UnifiedResilienceItemBillingInfo{
	// 				UsagePlanArmID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/MyResourceGroup/providers/Microsoft.AzureResilienceManagement/usagePlans/myUsagePlan"),
	// 				UsagePlanEnrollmentArmID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/MyResourceGroup/providers/Microsoft.AzureResilienceManagement/usagePlans/myUsagePlan/enrollments/sg1-enrollment"),
	// 				UsagePlanEnrollmentCreatedOn: to.Ptr(time.Date(2025, time.June, 1, 10, 0, 0, 0, time.UTC)),
	// 				UsagePlanEnrollmentLastUpdatedOn: to.Ptr(time.Date(2025, time.June, 15, 10, 0, 0, 0, time.UTC)),
	// 			},
	// 			LastModifiedTime: to.Ptr(time.Date(2025, time.May, 1, 8, 0, 0, 0, time.UTC)),
	// 		},
	// 		ID: to.Ptr("/providers/Microsoft.AzureResilienceManagement/unifiedResilienceItems/uri1"),
	// 		Name: to.Ptr("uri1"),
	// 		Type: to.Ptr("Microsoft.AzureResilienceManagement/unifiedResilienceItems"),
	// 		SystemData: &armresiliencemanagement.SystemData{
	// 			CreatedBy: to.Ptr("dvnfxbuyqhvivfjddjccdtlwajfht"),
	// 			CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2025, time.February, 6, 15, 3, 42, 796000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("lndhhaimomorael"),
	// 			LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2025, time.February, 6, 15, 3, 42, 797000000, time.UTC)),
	// 		},
	// 	},
	// }
}
