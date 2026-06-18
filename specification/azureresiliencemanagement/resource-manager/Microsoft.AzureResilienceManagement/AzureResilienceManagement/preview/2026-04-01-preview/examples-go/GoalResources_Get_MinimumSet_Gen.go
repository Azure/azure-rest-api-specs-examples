package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/GoalResources_Get_MinimumSet_Gen.json
func ExampleGoalResourcesClient_Get_goalResourcesGetMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGoalResourcesClient().Get(ctx, "sg1", "ga1", "gr1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresiliencemanagement.GoalResourcesClientGetResponse{
	// 	GoalResource: armresiliencemanagement.GoalResource{
	// 		Properties: &armresiliencemanagement.GoalResourceProperties{
	// 			ResourceArmID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/MyResourceGroup/providers/Microsoft.Compute/virtualMachines/MyVirtualMachine"),
	// 			HighAvailabilityGoalParticipation: to.Ptr(armresiliencemanagement.ExclusionStateIncluded),
	// 			HighAvailabilityAttestationStatus: to.Ptr(armresiliencemanagement.AttestationStateNotAttested),
	// 			ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/providers/Microsoft.AzureResilienceManagement/goalAssignments/ga1/goalResources/gr1"),
	// 		Name: to.Ptr("gr1"),
	// 		Type: to.Ptr("Microsoft.AzureResilienceManagement/goalResources"),
	// 		SystemData: &armresiliencemanagement.SystemData{
	// 			CreatedBy: to.Ptr("admin@contoso.com"),
	// 			CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-06T15:03:42.796Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("admin@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-06T15:03:42.797Z"); return t}()),
	// 		},
	// 	},
	// }
}
