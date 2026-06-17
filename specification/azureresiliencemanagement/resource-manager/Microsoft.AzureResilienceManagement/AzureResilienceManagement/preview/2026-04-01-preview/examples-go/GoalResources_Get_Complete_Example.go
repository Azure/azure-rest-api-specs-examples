package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/GoalResources_Get_Complete_Example.json
func ExampleGoalResourcesClient_Get_goalResourcesGetCompleteExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGoalResourcesClient().Get(ctx, "production-sg", "resiliencyGoalAssignment", "web-app-resource", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresiliencemanagement.GoalResourcesClientGetResponse{
	// 	GoalResource: armresiliencemanagement.GoalResource{
	// 		Properties: &armresiliencemanagement.GoalResourceProperties{
	// 			ResourceArmID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/ProductionRG/providers/Microsoft.Web/sites/MyWebApp"),
	// 			HighAvailabilityGoalParticipation: to.Ptr(armresiliencemanagement.ExclusionStateIncluded),
	// 			HighAvailabilityAttestationStatus: to.Ptr(armresiliencemanagement.AttestationStateManuallyAttested),
	// 			DisasterRecoveryGoalParticipation: to.Ptr(armresiliencemanagement.ExclusionStateIncluded),
	// 			DisasterRecoveryAttestationStatus: to.Ptr(armresiliencemanagement.AttestationStateNotAttested),
	// 			UserConfirmationForHighAvailability: []*armresiliencemanagement.UserConfirmationForHighAvailabilityItem{
	// 				{
	// 					SolutionDisplayName: to.Ptr(armresiliencemanagement.SolutionDisplayNameZonePinnedVMWithZrsDisk),
	// 					ConfirmationStatus: to.Ptr(armresiliencemanagement.ConfirmationStatusApprovedByUser),
	// 					ReasonForRequestingConfirmation: to.Ptr(armresiliencemanagement.ReasonForRequestingConfirmationZonePinnedZrsDataDisksConditional),
	// 				},
	// 			},
	// 			ServiceGroupMemberships: []*armresiliencemanagement.ServiceGroupMembership{
	// 				{
	// 					ServiceGroupID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/ProductionRG/providers/Microsoft.AzureResilienceManagement/serviceGroups/production-sg"),
	// 					MembershipType: to.Ptr(armresiliencemanagement.MembershipTypeDirect),
	// 				},
	// 				{
	// 					ServiceGroupID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/ProductionRG/providers/Microsoft.AzureResilienceManagement/serviceGroups/subscription-level-sg"),
	// 					MembershipType: to.Ptr(armresiliencemanagement.MembershipTypeThroughSubscription),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/ProductionRG/providers/Microsoft.AzureResilienceManagement/goalAssignments/resiliency-goal-assignment/goalResources/web-app-resource"),
	// 		Name: to.Ptr("web-app-resource"),
	// 		Type: to.Ptr("Microsoft.AzureResilienceManagement/goalResources"),
	// 		SystemData: &armresiliencemanagement.SystemData{
	// 			CreatedBy: to.Ptr("admin@contoso.com"),
	// 			CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-30T10:15:30.123Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("admin@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-30T14:22:45.789Z"); return t}()),
	// 		},
	// 	},
	// }
}
