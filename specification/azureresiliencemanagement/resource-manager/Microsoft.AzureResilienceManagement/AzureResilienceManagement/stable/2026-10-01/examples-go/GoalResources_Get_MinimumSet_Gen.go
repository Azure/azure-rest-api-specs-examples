package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-10-01/GoalResources_Get_MinimumSet_Gen.json
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
	res, err := clientFactory.NewGoalResourcesClient().Get(ctx, "production-sg", "zonal-resiliency-goal", "primary-vm", nil)
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
	// 			ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
	// 			ZonalResiliency: &armresiliencemanagement.ResiliencyProperties{
	// 				GoalParticipation: to.Ptr(armresiliencemanagement.ExclusionStateIncluded),
	// 				AttestationStatus: to.Ptr(armresiliencemanagement.AttestationStateNotAttested),
	// 			},
	// 		},
	// 		ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/production-sg/providers/Microsoft.AzureResilienceManagement/goalAssignments/zonal-resiliency-goal/goalResources/primary-vm"),
	// 		Name: to.Ptr("primary-vm"),
	// 		Type: to.Ptr("Microsoft.AzureResilienceManagement/goalAssignments/goalResources"),
	// 		SystemData: &armresiliencemanagement.SystemData{
	// 			CreatedBy: to.Ptr("admin@contoso.com"),
	// 			CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2025, time.February, 6, 15, 3, 42, 796000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("admin@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2025, time.February, 6, 15, 3, 42, 797000000, time.UTC)),
	// 		},
	// 	},
	// }
}
