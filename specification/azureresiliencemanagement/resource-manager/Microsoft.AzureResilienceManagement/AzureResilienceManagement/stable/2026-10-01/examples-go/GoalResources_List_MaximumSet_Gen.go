package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-10-01/GoalResources_List_MaximumSet_Gen.json
func ExampleGoalResourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGoalResourcesClient().NewListPager("production-sg", "zonal-resiliency-goal", &armresiliencemanagement.GoalResourcesClientListOptions{
		SkipToken: to.Ptr("xntbyoswztnmvitj"),
		Top:       to.Ptr[int32](69)})
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
		// page = armresiliencemanagement.GoalResourcesClientListResponse{
		// 	GoalResourceListResult: armresiliencemanagement.GoalResourceListResult{
		// 		Value: []*armresiliencemanagement.GoalResource{
		// 			{
		// 				Properties: &armresiliencemanagement.GoalResourceProperties{
		// 					ResourceArmID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/MyResourceGroup/providers/Microsoft.Compute/virtualMachines/MyVirtualMachine"),
		// 					ZonalResiliency: &armresiliencemanagement.ResiliencyProperties{
		// 						UserConfirmation: []*armresiliencemanagement.UserConfirmationItem{
		// 							{
		// 								SolutionDisplayName: to.Ptr(armresiliencemanagement.SolutionDisplayNameZonePinnedVMWithZrsDisk),
		// 								ConfirmationStatus: to.Ptr(armresiliencemanagement.ConfirmationStatusApprovedByUser),
		// 								ReasonForRequestingConfirmation: to.Ptr(armresiliencemanagement.ReasonForRequestingConfirmationZonePinnedZrsDataDisksConditional),
		// 							},
		// 							{
		// 								SolutionDisplayName: to.Ptr(armresiliencemanagement.SolutionDisplayNameVMInMultiZoneVmss),
		// 								ConfirmationStatus: to.Ptr(armresiliencemanagement.ConfirmationStatusApprovalPending),
		// 								ReasonForRequestingConfirmation: to.Ptr(armresiliencemanagement.ReasonForRequestingConfirmationVMInMultiZoneScaleSetStatelessOnly),
		// 							},
		// 						},
		// 						GoalParticipation: to.Ptr(armresiliencemanagement.ExclusionStateExcluded),
		// 						AttestationStatus: to.Ptr(armresiliencemanagement.AttestationStateManuallyAttested),
		// 						ExclusionReason: to.Ptr(armresiliencemanagement.ExclusionReasonUserSelectedExclusion),
		// 					},
		// 					ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/production-sg/providers/Microsoft.AzureResilienceManagement/goalAssignments/zonal-resiliency-goal/goalResources/primary-vm"),
		// 				Name: to.Ptr("primary-vm"),
		// 				Type: to.Ptr("Microsoft.AzureResilienceManagement/goalAssignments/goalResources"),
		// 				SystemData: &armresiliencemanagement.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2025, time.February, 6, 15, 3, 42, 796000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.February, 6, 15, 3, 42, 797000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/providers/Microsoft.Management/serviceGroups/production-sg/providers/Microsoft.AzureResilienceManagement/goalAssignments/zonal-resiliency-goal/goalResources?api-version=2026-10-01&$skipToken=eyJuZXh0UGFnZSI6Mn0%3D"),
		// 	},
		// }
	}
}
